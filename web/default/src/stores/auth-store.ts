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
import { create } from 'zustand'

import type { LoginChallenge } from '@/features/auth/secure-verification/types'
import type { AdminCapabilities } from '@/lib/admin-permissions'

export type UserPermissions = {
  sidebar_settings?: boolean
  sidebar_modules?: Record<string, unknown>
  admin_permissions?: AdminCapabilities
}

export interface AuthUser {
  has_password?: boolean
  id: number
  username: string
  display_name?: string
  email?: string
  role: number
  status?: number
  group?: string
  quota?: number
  used_quota?: number
  request_count?: number
  aff_code?: string
  aff_count?: number
  aff_quota?: number
  aff_history_quota?: number
  inviter_id?: number
  github_id?: string
  discord_id?: string
  oidc_id?: string
  wechat_id?: string
  telegram_id?: string
  linux_do_id?: string
  language?: string
  setting?: Record<string, unknown> | string
  stripe_customer?: string
  sidebar_modules?: string
  permissions?: UserPermissions
}

export interface LoginSession {
  sid: string
  current: boolean
  login_method: string
  ip: string
  user_agent: string
  created_at: number
  last_active_at: number
  expires_at: number
}

export interface AuthBundle {
  access_token: string
  token_type: 'Bearer' | string
  access_expires_at: number
  user: AuthUser
  session: LoginSession
}

export type AuthBootstrapState = 'idle' | 'checking' | 'complete'

export interface PendingLoginVerification {
  challenge: LoginChallenge
  redirectTo?: string
}

/**
 * The refresh flow was removed, so the access token is the only credential a
 * reload has. Persist the whole bundle (token + user + session) per browser so
 * a cold start can restore the dashboard without a network round trip.
 *
 * This is localStorage, not a cookie: it carries no ambient authority, so it is
 * never attached automatically to requests and cannot be abused cross-site.
 * It is however readable by injected script, which is the accepted trade-off of
 * dropping HttpOnly refresh cookies.
 */
export const AUTH_PERSIST_KEY = 'new-api:auth'

function isPersistedBundle(value: unknown): value is AuthBundle {
  if (!value || typeof value !== 'object') return false
  const bundle = value as Partial<AuthBundle>
  return (
    typeof bundle.access_token === 'string' &&
    bundle.access_token.length > 0 &&
    typeof bundle.access_expires_at === 'number' &&
    Number.isFinite(bundle.access_expires_at) &&
    Boolean(bundle.user) &&
    typeof bundle.user === 'object' &&
    Boolean(bundle.session) &&
    typeof bundle.session === 'object'
  )
}

export function loadPersistedAuthBundle(): AuthBundle | null {
  if (typeof window === 'undefined') return null
  try {
    const raw = window.localStorage.getItem(AUTH_PERSIST_KEY)
    if (!raw) return null
    const parsed: unknown = JSON.parse(raw)
    return isPersistedBundle(parsed) ? parsed : null
  } catch {
    return null
  }
}

function persistAuthBundle(bundle: AuthBundle | null): void {
  if (typeof window === 'undefined') return
  try {
    if (bundle) {
      window.localStorage.setItem(AUTH_PERSIST_KEY, JSON.stringify(bundle))
    } else {
      window.localStorage.removeItem(AUTH_PERSIST_KEY)
    }
  } catch {
    // Persistence is best-effort when storage is unavailable or full.
  }
}

interface AuthState {
  auth: {
    user: AuthUser | null
    accessToken: string | null
    accessExpiresAt: number | null
    session: LoginSession | null
    pendingLoginVerification: PendingLoginVerification | null
    bootstrapState: AuthBootstrapState
    setBundle: (bundle: AuthBundle) => void
    setUser: (user: AuthUser | null) => void
    setPendingLoginVerification: (
      pending: PendingLoginVerification | null
    ) => void
    setBootstrapState: (bootstrapState: AuthBootstrapState) => void
    reset: (bootstrapState?: AuthBootstrapState) => void
  }
}

const persistedAuth = loadPersistedAuthBundle()

export const useAuthStore = create<AuthState>()((set, get) => ({
  auth: {
    user: persistedAuth?.user ?? null,
    accessToken: persistedAuth?.access_token ?? null,
    accessExpiresAt: persistedAuth?.access_expires_at ?? null,
    session: persistedAuth?.session ?? null,
    pendingLoginVerification: null,
    bootstrapState: 'idle',
    setBundle: (bundle) => {
      persistAuthBundle(bundle)
      set((state) => ({
        ...state,
        auth: {
          ...state.auth,
          user: bundle.user,
          accessToken: bundle.access_token,
          accessExpiresAt: bundle.access_expires_at,
          session: bundle.session,
          pendingLoginVerification: null,
          bootstrapState: 'complete',
        },
      }))
    },
    setUser: (user) => {
      const current = get().auth
      if (current.accessToken && current.session && user) {
        persistAuthBundle({
          access_token: current.accessToken,
          token_type: 'Bearer',
          access_expires_at: current.accessExpiresAt ?? 0,
          user,
          session: current.session,
        })
      }
      set((state) => ({
        ...state,
        auth: {
          ...state.auth,
          user,
          pendingLoginVerification:
            state.auth.user?.id === user?.id
              ? state.auth.pendingLoginVerification
              : null,
        },
      }))
    },
    setPendingLoginVerification: (pendingLoginVerification) =>
      set((state) => ({
        ...state,
        auth: { ...state.auth, pendingLoginVerification },
      })),
    setBootstrapState: (bootstrapState) =>
      set((state) => ({
        ...state,
        auth: { ...state.auth, bootstrapState },
      })),
    reset: (bootstrapState = 'complete') => {
      persistAuthBundle(null)
      set((state) => ({
        ...state,
        auth: {
          ...state.auth,
          user: null,
          accessToken: null,
          accessExpiresAt: null,
          session: null,
          pendingLoginVerification: null,
          bootstrapState,
        },
      }))
    },
  },
}))
