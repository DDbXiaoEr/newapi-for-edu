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
export type CCSwitchApp = 'claude' | 'codex' | 'gemini'

/**
 * Reads the server address the console is configured to talk to, falling back
 * to the current origin when no status snapshot is available.
 */
export function getCCSwitchServerAddress(): string {
  try {
    const raw = localStorage.getItem('status')
    if (raw) {
      const status = JSON.parse(raw)
      if (status.server_address) return status.server_address
    }
  } catch {
    /* empty */
  }
  return typeof window === 'undefined' ? '' : window.location.origin
}

export type BuildCCSwitchImportUrlParams = {
  app: string
  name: string
  serverAddress: string
  apiKey: string
  models?: Record<string, string>
}

/**
 * Builds a `ccswitch://` deeplink that imports the given provider into CC
 * Switch. Codex expects the OpenAI-compatible base URL with a `/v1` suffix,
 * while Claude and Gemini point at the service root.
 */
export function buildCCSwitchImportUrl(
  params: BuildCCSwitchImportUrlParams
): string {
  const serverAddress = params.serverAddress || ''
  const endpoint =
    params.app === 'codex' ? `${serverAddress}/v1` : serverAddress
  const searchParams = new URLSearchParams()
  searchParams.set('resource', 'provider')
  searchParams.set('app', params.app)
  searchParams.set('name', params.name)
  searchParams.set('endpoint', endpoint)
  searchParams.set('apiKey', params.apiKey)
  for (const [key, value] of Object.entries(params.models ?? {})) {
    if (value) searchParams.set(key, value)
  }
  searchParams.set('homepage', serverAddress)
  searchParams.set('enabled', 'true')
  return `ccswitch://v1/import?${searchParams.toString()}`
}
