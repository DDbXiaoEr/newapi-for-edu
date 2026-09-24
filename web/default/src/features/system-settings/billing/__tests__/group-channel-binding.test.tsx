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
*/
import { cleanup, render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { afterEach, expect, it, vi } from 'vitest'

import { api } from '@/lib/api'

import { GroupChannelBindingSection } from '../group-channel-binding-section'

const payload = {
  group_channels: {
    default: { 'gpt-4': [1] },
  },
  channels: [
    {
      id: 1,
      name: 'openai',
      status: 1,
      models: ['gpt-4', 'gpt-3.5'],
    },
    {
      id: 2,
      name: 'azure',
      status: 1,
      models: ['gpt-4'],
    },
  ],
  stale_channels: {},
  groups: ['default', 'vip'],
  group_models: {
    default: ['gpt-4', 'gpt-3.5'],
    vip: ['gpt-4'],
  },
}

afterEach(() => {
  cleanup()
  vi.restoreAllMocks()
})

it('saves a per-model channel pin for the selected group', async () => {
  vi.spyOn(api, 'get').mockResolvedValue({
    data: { success: true, data: payload },
  })
  const put = vi.spyOn(api, 'put').mockResolvedValue({
    data: { success: true, data: { group_channels: payload.group_channels } },
  })

  render(<GroupChannelBindingSection />)

  await screen.findByText('openai #1')
  expect(screen.getByText('gpt-4 → openai #1')).toBeVisible()

  await userEvent.click(screen.getByText('azure #2'))
  await userEvent.click(
    screen.getByRole('button', { name: 'Save binding' })
  )

  await waitFor(() => {
    expect(put).toHaveBeenCalledWith('/api/group/channels', {
      group: 'default',
      model: 'gpt-4',
      channel_ids: [1, 2],
    })
  })
})

it('lists only channels that serve the selected model', async () => {
  vi.spyOn(api, 'get').mockResolvedValue({
    data: { success: true, data: payload },
  })

  render(<GroupChannelBindingSection />)

  await screen.findByText('openai #1')
  expect(screen.getByText('azure #2')).toBeVisible()

  await userEvent.click(screen.getByRole('button', { name: 'Select model' }))
  await userEvent.click(await screen.findByRole('option', { name: 'gpt-3.5' }))

  await waitFor(() => {
    expect(screen.queryByText('azure #2')).not.toBeInTheDocument()
  })
  expect(screen.getByText('openai #1')).toBeVisible()
})
