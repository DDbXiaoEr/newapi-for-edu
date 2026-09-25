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
import { cleanup, render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { Toaster, toast } from 'sonner'
import { afterEach, expect, it, vi } from 'vitest'

import { api } from '@/lib/api'

import { UsersProvider } from '../../users-provider'
import { UsersAssignGroupDialog } from '../users-assign-group-dialog'

const alice = {
  id: 1,
  username: 'alice',
  display_name: 'Alice',
  group: 'default',
  role: 1,
  status: 1,
  quota: 0,
  used_quota: 0,
  request_count: 0,
}
const bob = {
  id: 2,
  username: 'bob',
  display_name: 'Bob',
  group: 'vip',
  role: 1,
  status: 1,
  quota: 0,
  used_quota: 0,
  request_count: 0,
}

const queryClients: QueryClient[] = []

function renderDialog(onOpenChange = vi.fn()) {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false, gcTime: 0 } },
  })
  queryClients.push(queryClient)
  return {
    onOpenChange,
    ...render(
      <QueryClientProvider client={queryClient}>
        <UsersProvider>
          <UsersAssignGroupDialog open onOpenChange={onOpenChange} />
        </UsersProvider>
        <Toaster />
      </QueryClientProvider>
    ),
  }
}

afterEach(() => {
  toast.dismiss()
  cleanup()
  for (const client of queryClients) client.clear()
  queryClients.length = 0
  vi.restoreAllMocks()
})

it('keeps confirm disabled until users and a group are selected', async () => {
  vi.spyOn(api, 'get').mockImplementation(async (url) => {
    if (url === '/api/group/') {
      return { data: { success: true, data: ['default', 'vip'] } }
    }
    return {
      data: {
        success: true,
        data: { items: [alice, bob], total: 2, page: 1, page_size: 100 },
      },
    }
  })
  renderDialog()
  const confirm = await screen.findByRole('button', {
    name: 'Confirm assignment',
  })
  expect(confirm).toBeDisabled()
  const user = userEvent.setup()
  await user.click(
    await screen.findByRole('checkbox', { name: /Select alice/i })
  )
  await user.click(screen.getByRole('button', { name: 'Add selected users' }))
  expect(confirm).toBeDisabled()
  await user.click(screen.getByRole('combobox'))
  await user.click(await screen.findByRole('option', { name: 'vip' }))
  expect(confirm).toBeEnabled()
})

it('assigns selected users to the chosen group, including users who already have a group', async () => {
  vi.spyOn(api, 'get').mockImplementation(async (url) => {
    if (url === '/api/group/') {
      return { data: { success: true, data: ['default', 'vip'] } }
    }
    return {
      data: {
        success: true,
        data: { items: [alice, bob], total: 2, page: 1, page_size: 100 },
      },
    }
  })
  const post = vi.spyOn(api, 'post').mockResolvedValue({
    data: {
      success: true,
      data: { updated: 2, skipped: 0, ids: [1, 2] },
    },
  })
  const { onOpenChange } = renderDialog()
  const user = userEvent.setup()
  await user.click(
    await screen.findByRole('checkbox', { name: /Select alice/i })
  )
  await user.click(screen.getByRole('checkbox', { name: /Select bob/i }))
  await user.click(screen.getByRole('button', { name: 'Add selected users' }))
  expect(screen.getByText('Selected users')).toBeVisible()
  expect(
    screen.getByText('Users that will be assigned to the chosen group')
  ).toBeVisible()
  await user.click(screen.getByRole('combobox'))
  await user.click(await screen.findByRole('option', { name: 'vip' }))
  await user.click(screen.getByRole('button', { name: 'Confirm assignment' }))
  await waitFor(() =>
    expect(post).toHaveBeenCalledWith('/api/user/group', {
      ids: [1, 2],
      group: 'vip',
    })
  )
  expect(
    await screen.findByText('Assigned 2 users to vip')
  ).toBeVisible()
  expect(onOpenChange).toHaveBeenCalledWith(false)
})

it('filters available users when searching in the transfer panel', async () => {
  vi.spyOn(api, 'get').mockImplementation(async (url) => {
    if (url === '/api/group/') {
      return { data: { success: true, data: ['default', 'vip'] } }
    }
    return {
      data: {
        success: true,
        data: { items: [alice, bob], total: 2, page: 1, page_size: 100 },
      },
    }
  })
  renderDialog()
  const user = userEvent.setup()
  await screen.findByRole('checkbox', { name: /Select alice/i })
  await user.type(screen.getByRole('textbox', { name: 'Search users' }), 'bob')
  expect(screen.getByText('Bob')).toBeVisible()
  expect(screen.queryByText('Alice')).not.toBeInTheDocument()
})
