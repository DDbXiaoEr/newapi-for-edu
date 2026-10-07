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
import { useState } from 'react'
import { I18nextProvider } from 'react-i18next'
import { beforeEach, describe, expect, it, vi } from 'vitest'

import en from '@/i18n/locales/en.json'
import { api } from '@/lib/api'

import { SettingsPageProvider } from '../../components/settings-page-context'
import { RateLimitSection } from '../rate-limit-section'

const defaults = {
  ModelRequestRateLimitEnabled: true,
  ModelRequestRateLimitCount: 0,
  ModelRequestRateLimitSuccessCount: 1000,
  ModelRequestRateLimitDurationMinutes: 1,
  ModelRequestRateLimitGroup: '',
  ModelRequestRateLimitMode: 'api_key' as const,
}

let testI18n: i18n

function Fixture(props: {
  defaultValues?: typeof defaults
  onSave?: () => void
}) {
  const [container, setContainer] = useState<HTMLDivElement | null>(null)
  const [client] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: { retry: false },
          mutations: { retry: false },
        },
      })
  )
  return (
    <I18nextProvider i18n={testI18n}>
      <QueryClientProvider client={client}>
        <div ref={setContainer} />
        <SettingsPageProvider actionsContainer={container}>
          <RateLimitSection defaultValues={props.defaultValues ?? defaults} />
        </SettingsPageProvider>
      </QueryClientProvider>
    </I18nextProvider>
  )
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

describe('Model request rate-limit dimension', () => {
  it('defaults to counting the limit per API key', async () => {
    render(<Fixture />)

    const selector = await screen.findByRole('combobox', {
      name: 'Rate limit by',
    })
    expect(selector).toHaveTextContent('API Key')
  })

  it('saves per-client-IP counting when Client IP is selected', async () => {
    const put = vi
      .spyOn(api, 'put')
      .mockResolvedValue({ data: { success: true } })
    const user = userEvent.setup()
    render(<Fixture />)

    const selector = await screen.findByRole('combobox', {
      name: 'Rate limit by',
    })
    await user.click(selector)
    await user.click(await screen.findByRole('option', { name: 'Client IP' }))
    await user.click(screen.getByRole('button', { name: 'Save rate limits' }))

    await waitFor(() =>
      expect(put).toHaveBeenCalledWith('/api/option/', {
        key: 'ModelRequestRateLimitMode',
        value: 'ip',
      })
    )
  })
})
