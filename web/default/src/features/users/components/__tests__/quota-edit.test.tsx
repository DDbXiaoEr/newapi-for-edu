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
import { afterEach, expect, it, vi } from 'vitest'

import { api } from '@/lib/api'
import { useAuthStore } from '@/stores/auth-store'

import type { User } from '../../types'
import { UsersMutateDrawer } from '../users-mutate-drawer'
import { UsersProvider } from '../users-provider'

const target: User = {
  id: 2,
  username: 'managed-user',
  display_name: 'Managed user',
  role: 1,
  status: 1,
  quota: 0,
  used_quota: 0,
  request_count: 0,
  group: 'default',
}

function renderDrawer(currentRow?: User) {
  useAuthStore
    .getState()
    .auth.setUser({ id: 1, username: 'operator', role: 100 })
  vi.spyOn(api, 'get').mockImplementation(async (url) => {
    if (url === '/api/authz/catalog') {
      return { data: { success: true, data: { resources: [], roles: [] } } }
    }
    if (url === '/api/group/') {
      return { data: { success: true, data: ['default'] } }
    }
    return { data: { success: true, data: target } }
  })
  const client = new QueryClient({
    defaultOptions: { queries: { retry: false } },
  })
  return render(
    <QueryClientProvider client={client}>
      <UsersProvider>
        <UsersMutateDrawer
          open
          onOpenChange={() => undefined}
          currentRow={currentRow}
        />
      </UsersProvider>
    </QueryClientProvider>
  )
}

afterEach(() => {
  cleanup()
  vi.restoreAllMocks()
  useAuthStore.getState().auth.reset()
})

it('lets admins type remaining quota when creating a user', async () => {
  renderDrawer()
  const quotaInput = document.querySelector(
    'input[name="quota_dollars"]'
  ) as HTMLInputElement | null
  expect(quotaInput).not.toBeNull()
  expect(quotaInput).not.toHaveAttribute('readonly')
  expect(quotaInput).toHaveAttribute('type', 'number')
  expect(
    screen.queryByRole('button', { name: 'Adjust Quota' })
  ).not.toBeInTheDocument()
})

it('lets admins type remaining quota when editing a user', async () => {
  renderDrawer(target)
  const quotaInput = await screen.findByDisplayValue('0')
  expect(quotaInput).not.toHaveAttribute('readonly')
  expect(quotaInput).toHaveAttribute('type', 'number')
  expect(screen.getByRole('button', { name: 'Adjust Quota' })).toBeVisible()
})

it('overrides remaining quota when saving an edited user', async () => {
  vi.spyOn(api, 'put').mockResolvedValue({ data: { success: true } })
  const post = vi.spyOn(api, 'post').mockResolvedValue({ data: { success: true } })
  renderDrawer(target)
  const quotaInput = await screen.findByDisplayValue('0')
  await userEvent.clear(quotaInput)
  await userEvent.type(quotaInput, '1.5')
  await userEvent.click(screen.getByRole('button', { name: 'Save changes' }))
  await waitFor(() =>
    expect(post).toHaveBeenCalledWith(
      '/api/user/manage',
      expect.objectContaining({
        id: 2,
        action: 'add_quota',
        mode: 'override',
      })
    )
  )
})
