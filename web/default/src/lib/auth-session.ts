/*
Copyright (C) 2023-2026 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing, please contact support@quantumnous.com
*/
import type { QueryClient } from '@tanstack/react-query'
import { t } from 'i18next'

import { publishAuthSessionEvent } from '@/lib/auth-session-sync'
import {
  useAuthStore,
  type AuthBootstrapState,
  type AuthBundle,
  type AuthUser,
  type LoginSession,
} from '@/stores/auth-store'

export type RefreshOutcome =
  | { kind: 'authenticated'; bundle: AuthBundle }
  | { kind: 'anonymous' }
  | { kind: 'transient_error'; error: unknown }
  | { kind: 'out_of_sync'; code?: string }

export interface AuthRefreshHTTPResponse {
  status: number
  data?: unknown
  error?: unknown
}

export interface AuthRefreshRuntime {
  request: (expectedSID?: string) => Promise<AuthRefreshHTTPResponse>
  getExpectedSID: () => string | undefined
  parseBundle: (value: unknown) => AuthBundle | null
  acceptBundle: (bundle: AuthBundle) => void
  clear: (synchronizeTabs: boolean, bootstrapState?: AuthBootstrapState) => void
  markTransient: () => void
  wait: (delay: number) => Promise<void>
  isCurrent?: () => boolean
}

export interface AuthTokenRotation {
  access_token: string
  token_type: string
  access_expires_at: number
  session: LoginSession
}

export class AuthRotationError extends Error {
  constructor(message: string) {
    super(message)
    this.name = 'AuthRotationError'
  }
}

const refreshRaceDelays = [80, 200, 500] as const
let authEpoch = 0

class AuthRefreshSupersededError extends Error {
  constructor() {
    super('Authentication refresh was superseded')
    this.name = 'AuthRefreshSupersededError'
  }
}

function isRecord(value: unknown): value is Record<string, unknown> {
  return Boolean(value) && typeof value === 'object'
}

function isAuthUser(value: unknown): value is AuthUser {
  if (!isRecord(value)) return false
  return (
    Number.isInteger(value.id) &&
    Number(value.id) > 0 &&
    typeof value.username === 'string' &&
    typeof value.role === 'number'
  )
}

function isLoginSession(value: unknown): value is LoginSession {
  if (!isRecord(value)) return false
  return (
    typeof value.sid === 'string' &&
    value.sid.length > 0 &&
    typeof value.current === 'boolean' &&
    typeof value.login_method === 'string' &&
    typeof value.ip === 'string' &&
    typeof value.user_agent === 'string' &&
    typeof value.created_at === 'number' &&
    typeof value.last_active_at === 'number' &&
    typeof value.expires_at === 'number'
  )
}

function hasValidTokenFields(value: Record<string, unknown>): boolean {
  return (
    typeof value.access_token === 'string' &&
    value.access_token.length > 0 &&
    typeof value.token_type === 'string' &&
    value.token_type.length > 0 &&
    typeof value.access_expires_at === 'number' &&
    Number.isFinite(value.access_expires_at) &&
    value.access_expires_at > 0
  )
}

export function isAuthBundle(value: unknown): value is AuthBundle {
  if (!isRecord(value)) return false
  return (
    hasValidTokenFields(value) &&
    isAuthUser(value.user) &&
    isLoginSession(value.session)
  )
}

function isAuthTokenRotation(value: unknown): value is AuthTokenRotation {
  return (
    isRecord(value) &&
    hasValidTokenFields(value) &&
    value.token_type === 'Bearer' &&
    isLoginSession(value.session) &&
    value.session.current
  )
}

export function applyAuthBundle(
  bundle: AuthBundle,
  synchronizeTabs = true
): void {
  const previousSID = useAuthStore.getState().auth.session?.sid
  authEpoch += 1
  useAuthStore.getState().auth.setBundle(bundle)
  if (synchronizeTabs && previousSID !== bundle.session.sid) {
    publishAuthSessionEvent('authenticated', bundle.session.sid)
  }
}

export function applyAuthRotation(value: unknown): void {
  if (!isAuthTokenRotation(value)) {
    throw new AuthRotationError('Invalid authentication rotation response')
  }

  const auth = useAuthStore.getState().auth
  if (!auth.user || !auth.session) {
    throw new AuthRotationError('Authentication rotation has no active session')
  }
  if (value.session.sid !== auth.session.sid) {
    throw new AuthRotationError('Authentication rotation session mismatch')
  }

  applyAuthBundle(
    {
      access_token: value.access_token,
      token_type: value.token_type,
      access_expires_at: value.access_expires_at,
      session: value.session,
      user: auth.user,
    },
    false
  )
}

export function clearAuthentication(
  synchronizeTabs = true,
  bootstrapState: AuthBootstrapState = 'complete'
): void {
  const sid = useAuthStore.getState().auth.session?.sid
  authEpoch += 1
  useAuthStore.getState().auth.reset(bootstrapState)
  if (synchronizeTabs && sid) {
    publishAuthSessionEvent('signed_out', sid)
  }
}

export function clearAuthenticatedClientState(
  queryClient: QueryClient,
  synchronizeTabs = true
): void {
  queryClient.clear()
  clearAuthentication(synchronizeTabs)
}

export function createRefreshRunner(
  runtime: AuthRefreshRuntime
): () => Promise<RefreshOutcome> {
  const superseded = (): RefreshOutcome => ({
    kind: 'transient_error',
    error: new AuthRefreshSupersededError(),
  })
  const run = async (
    raceAttempt: number,
    allowMismatchRetry: boolean
  ): Promise<RefreshOutcome> => {
    if (runtime.isCurrent && !runtime.isCurrent()) return superseded()
    const response = await runtime.request(runtime.getExpectedSID())
    if (runtime.isCurrent && !runtime.isCurrent()) return superseded()
    const responseData = isRecord(response.data) ? response.data : undefined
    const code =
      typeof responseData?.code === 'string' ? responseData.code : undefined
    const bundle = runtime.parseBundle(responseData?.data)
    if (responseData?.success === true && bundle) {
      runtime.acceptBundle(bundle)
      return { kind: 'authenticated', bundle }
    }

    if (response.status === 409 && code === 'AUTH_REFRESH_RACE') {
      const delay = refreshRaceDelays[raceAttempt]
      if (delay !== undefined) {
        await runtime.wait(delay)
        return run(raceAttempt + 1, allowMismatchRetry)
      }
      runtime.clear(false)
      return { kind: 'out_of_sync', code }
    }

    if (response.status === 409 && code === 'AUTH_SESSION_MISMATCH') {
      if (allowMismatchRetry) {
        runtime.clear(false, 'idle')
        return run(0, false)
      }
      runtime.clear(false)
      return { kind: 'out_of_sync', code }
    }

    if (response.status === 401) {
      runtime.clear(true)
      return { kind: 'anonymous' }
    }

    if (!response.status || response.status >= 500 || response.status === 429) {
      runtime.markTransient()
      return {
        kind: 'transient_error',
        error: response.error ?? response.data,
      }
    }

    runtime.clear(false)
    return {
      kind: 'out_of_sync',
      code: code ?? 'AUTH_INVALID_REFRESH_RESPONSE',
    }
  }

  return () => run(0, true)
}

/**
 * Resolve authentication from the persisted/in-memory bundle.
 *
 * The refresh endpoint and its HttpOnly cookie are gone: the access token is
 * now the long-lived credential itself, so there is nothing left to exchange
 * and no network call to make. The name is kept for callers that still treat
 * "refresh" as "make sure we have a usable token".
 */
export function refreshAuthentication(): Promise<RefreshOutcome> {
  const bundle = currentValidAuthBundle()
  if (bundle) return Promise.resolve({ kind: 'authenticated', bundle })
  return Promise.resolve({ kind: 'anonymous' })
}

function currentValidAuthBundle(): AuthBundle | null {
  const auth = useAuthStore.getState().auth
  if (
    !auth.user ||
    !auth.accessToken ||
    !auth.accessExpiresAt ||
    !auth.session ||
    auth.accessExpiresAt <= Math.floor(Date.now() / 1000)
  ) {
    return null
  }
  return {
    access_token: auth.accessToken,
    token_type: 'Bearer',
    access_expires_at: auth.accessExpiresAt,
    user: auth.user,
    session: auth.session,
  }
}

/**
 * Resolve authentication from the persisted bundle in localStorage.
 *
 * There is no server round trip any more: a valid, unexpired token in local
 * storage is the whole answer. A stale/expired bundle is cleared so the user
 * lands on the sign-in page instead of an authenticated shell that 401s.
 */
export async function resolveAuthentication(): Promise<RefreshOutcome> {
  const bundle = currentValidAuthBundle()
  if (bundle) {
    useAuthStore.getState().auth.setBootstrapState('complete')
    return { kind: 'authenticated', bundle }
  }

  const auth = useAuthStore.getState().auth
  if (auth.accessToken || auth.user || auth.session) {
    clearAuthentication(false)
  } else {
    auth.setBootstrapState('complete')
  }
  return { kind: 'anonymous' }
}

/**
 * Public boot path. With the refresh cookie gone there is nothing to probe, so
 * this is just resolveAuthentication; kept as a separate name for the route
 * boot sequence.
 */
export async function bootstrapAuthentication(): Promise<RefreshOutcome> {
  return resolveAuthentication()
}

export function getCommonHeaders(): Record<string, string> {
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
  }
  const accessToken = useAuthStore.getState().auth.accessToken
  if (accessToken) {
    headers.Authorization = `Bearer ${accessToken}`
  }
  return headers
}

export async function getFreshAuthHeaders(): Promise<Record<string, string>> {
  const auth = useAuthStore.getState().auth
  if (
    auth.accessToken &&
    auth.accessExpiresAt &&
    auth.accessExpiresAt > Math.floor(Date.now() / 1000)
  ) {
    return getCommonHeaders()
  }
  throw new Error(t('Session expired!'))
}
