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
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { createInstance, type i18n } from 'i18next'
import { I18nextProvider } from 'react-i18next'
import { beforeEach, describe, expect, it, vi } from 'vitest'

import { buildCCSwitchImportUrl } from '@/features/keys/lib/cc-switch'
import en from '@/i18n/locales/en.json'

import { ClientDetail } from '../components/client-detail'
import { QuickAccess } from '../index'
import { BUILTIN_CLIENTS, type ClientOption } from '../lib/clients'

vi.mock('@/features/chat/hooks/use-chat-presets', () => ({
  useChatPresets: () => ({
    chatPresets: [
      {
        id: '0',
        name: 'Cherry Studio',
        url: 'https://chat.example.com/?key={key}',
        type: 'custom-protocol',
      },
    ],
    serverAddress: 'https://api.example.com',
  }),
}))

vi.mock('@/features/chat/hooks/use-active-chat-key', () => ({
  useActiveChatKey: () => ({ data: 'sk-active' }),
}))

vi.mock('@/lib/api', async (importOriginal) => {
  const actual = await importOriginal<typeof import('@/lib/api')>()
  return {
    ...actual,
    getUserModels: vi.fn(async () => ({ success: true, data: ['gpt-4o'] })),
  }
})

let testI18n: i18n

function builtinOption(id: string): ClientOption & { kind: 'builtin' } {
  const builtin = BUILTIN_CLIENTS.find((client) => client.id === id)
  if (!builtin) throw new Error(`missing builtin ${id}`)
  return { id: `builtin:${id}`, name: builtin.name, kind: 'builtin', builtin }
}

beforeEach(async () => {
  testI18n = createInstance()
  await testI18n.init({
    lng: 'en',
    fallbackLng: 'en',
    resources: { en },
    interpolation: { escapeValue: false },
  })
})

describe('buildCCSwitchImportUrl', () => {
  it('appends /v1 for the codex endpoint and carries the api key', () => {
    const url = buildCCSwitchImportUrl({
      app: 'codex',
      name: 'New API (Codex)',
      serverAddress: 'https://api.example.com',
      apiKey: 'sk-test',
      models: { model: 'gpt-5' },
    })
    const parsed = new URL(url)
    expect(parsed.protocol).toBe('ccswitch:')
    expect(parsed.searchParams.get('app')).toBe('codex')
    expect(parsed.searchParams.get('endpoint')).toBe(
      'https://api.example.com/v1'
    )
    expect(parsed.searchParams.get('apiKey')).toBe('sk-test')
    expect(parsed.searchParams.get('model')).toBe('gpt-5')
  })

  it('uses the service root for claude', () => {
    const url = buildCCSwitchImportUrl({
      app: 'claude',
      name: 'New API (Claude Code)',
      serverAddress: 'https://api.example.com',
      apiKey: 'sk-test',
    })
    expect(new URL(url).searchParams.get('endpoint')).toBe(
      'https://api.example.com'
    )
  })
})

describe('ClientDetail', () => {
  const ctx = {
    serverAddress: 'https://api.example.com',
    apiKey: 'sk-active',
    model: 'gpt-4o',
  }

  it('shows only manual configuration for OpenCode', () => {
    render(
      <I18nextProvider i18n={testI18n}>
        <ClientDetail client={builtinOption('opencode')} ctx={ctx} hasApiKey />
      </I18nextProvider>
    )

    expect(screen.getByRole('heading', { name: 'OpenCode' })).toBeVisible()
    expect(screen.queryByText('Quick import')).not.toBeInTheDocument()
    expect(
      screen.getByText(/https:\/\/api\.example\.com\/v1/)
    ).toBeInTheDocument()
    expect(screen.getByText(/sk-active/)).toBeInTheDocument()
  })

  it('offers one-click CC Switch import for Codex with the active key', () => {
    render(
      <I18nextProvider i18n={testI18n}>
        <ClientDetail client={builtinOption('codex')} ctx={ctx} hasApiKey />
      </I18nextProvider>
    )

    expect(screen.getByText('Quick import')).toBeVisible()
    expect(screen.getByRole('button', { name: 'Open CC Switch' })).toBeEnabled()
    expect(screen.getByText(/~\/\.codex\/config\.toml/)).toBeInTheDocument()
  })

  it('disables quick import when no API key is available', () => {
    render(
      <I18nextProvider i18n={testI18n}>
        <ClientDetail
          client={builtinOption('claude-code')}
          ctx={{ ...ctx, apiKey: '' }}
          hasApiKey={false}
        />
      </I18nextProvider>
    )

    expect(
      screen.getByRole('button', { name: 'Open CC Switch' })
    ).toBeDisabled()
    expect(
      screen.getByText(
        'No enabled API key found. Create or enable one before importing.'
      )
    ).toBeVisible()
  })
})

describe('QuickAccess page', () => {
  function renderPage() {
    const client = new QueryClient({
      defaultOptions: { queries: { retry: false } },
    })
    return render(
      <I18nextProvider i18n={testI18n}>
        <QueryClientProvider client={client}>
          <QuickAccess />
        </QueryClientProvider>
      </I18nextProvider>
    )
  }

  it('lists built-in clients and configured presets as selectable buttons', async () => {
    renderPage()

    expect(screen.getByRole('heading', { name: 'Quick Access' })).toBeVisible()
    expect(screen.getByRole('button', { name: 'OpenCode' })).toBeVisible()
    expect(screen.getByRole('button', { name: 'Claude Code' })).toBeVisible()
    expect(screen.getByRole('button', { name: 'Codex' })).toBeVisible()
    expect(screen.getByRole('button', { name: 'Cherry Studio' })).toBeVisible()

    await userEvent.click(screen.getByRole('button', { name: 'Codex' }))
    await waitFor(() =>
      expect(screen.getByRole('heading', { name: 'Codex' })).toBeVisible()
    )
    expect(screen.getByText('Quick import')).toBeVisible()
  })
})
