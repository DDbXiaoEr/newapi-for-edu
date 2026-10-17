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
import { ArtifactStorageSection } from '../artifact-storage-section'

const defaults = {
  'task_artifact_store.mode': 's3',
  'task_artifact_store.endpoint': 'https://objects.example.com',
  'task_artifact_store.bucket': 'task-artifacts',
  'task_artifact_store.region': '',
  'task_artifact_store.access_key': 'access-key',
  'task_artifact_store.prefix': '',
  'task_artifact_store.path_style': true,
}

let testI18n: i18n

function Fixture(props: { defaultValues?: typeof defaults }) {
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
          <ArtifactStorageSection
            defaultValues={props.defaultValues ?? defaults}
          />
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

describe('Artifact storage region', () => {
  it('leaves the region blank when it is not configured', async () => {
    render(<Fixture />)

    const region = await screen.findByRole('textbox', { name: 'Region' })
    expect(region).toHaveValue('')
  })

  it('persists an emptied region instead of forcing us-east-1', async () => {
    const put = vi
      .spyOn(api, 'put')
      .mockResolvedValue({ data: { success: true } })
    const user = userEvent.setup()
    render(
      <Fixture
        defaultValues={{
          ...defaults,
          'task_artifact_store.region': 'us-east-1',
        }}
      />
    )

    const region = await screen.findByRole('textbox', { name: 'Region' })
    await user.clear(region)
    await user.click(
      screen.getByRole('button', { name: 'Save artifact storage settings' })
    )

    await waitFor(() =>
      expect(put).toHaveBeenCalledWith('/api/option/', {
        key: 'task_artifact_store.region',
        value: '',
      })
    )
  })
})
