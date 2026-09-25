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
import { cleanup, render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { afterEach, expect, it, vi } from 'vitest'

import { GroupUserTransfer } from '../components/group-user-transfer'
import type { GroupUserTransferItem } from '../types'

const alice: GroupUserTransferItem = {
  id: 1,
  username: 'alice',
  display_name: 'Alice',
  group: 'default',
  role: 1,
  status: 1,
}
const bob: GroupUserTransferItem = {
  id: 2,
  username: 'bob',
  display_name: 'Bob',
  group: 'vip',
  role: 1,
  status: 1,
}

afterEach(() => {
  cleanup()
})

it('moves a selected available user into the group when the add button is pressed', async () => {
  const onAdd = vi.fn()
  render(
    <GroupUserTransfer
      available={[alice]}
      members={[bob]}
      availableKeyword=''
      membersKeyword=''
      onAvailableKeywordChange={() => undefined}
      onMembersKeywordChange={() => undefined}
      onAdd={onAdd}
      onRemove={() => undefined}
    />
  )
  await userEvent.click(screen.getByRole('checkbox', { name: /Select alice/i }))
  await userEvent.click(
    screen.getByRole('button', { name: 'Add selected users' })
  )
  expect(onAdd).toHaveBeenCalledWith([1])
})

it('renders custom panel labels when provided', () => {
  render(
    <GroupUserTransfer
      available={[alice]}
      members={[]}
      availableKeyword=''
      membersKeyword=''
      onAvailableKeywordChange={() => undefined}
      onMembersKeywordChange={() => undefined}
      onAdd={() => undefined}
      onRemove={() => undefined}
      availableTitle='Available users'
      availableDescription='Search users to assign'
      membersTitle='Selected users'
      membersDescription='Users that will be assigned to the chosen group'
      membersEmptyTitle='No selected users'
    />
  )
  expect(screen.getByText('Search users to assign')).toBeVisible()
  expect(screen.getByText('Selected users')).toBeVisible()
  expect(
    screen.getByText('Users that will be assigned to the chosen group')
  ).toBeVisible()
  expect(screen.getByText('No selected users')).toBeVisible()
})

it('moves a selected member out of the group when the remove button is pressed', async () => {
  const onRemove = vi.fn()
  render(
    <GroupUserTransfer
      available={[alice]}
      members={[bob]}
      availableKeyword=''
      membersKeyword=''
      onAvailableKeywordChange={() => undefined}
      onMembersKeywordChange={() => undefined}
      onAdd={() => undefined}
      onRemove={onRemove}
    />
  )
  await userEvent.click(screen.getByRole('checkbox', { name: /Select bob/i }))
  await userEvent.click(
    screen.getByRole('button', { name: 'Remove selected users' })
  )
  expect(onRemove).toHaveBeenCalledWith([2])
})
