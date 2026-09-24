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
import { describe, expect, it } from 'vitest'

import {
  buildAssignPayloads,
  mergeTransferLists,
  resolveFallbackGroup,
} from '../lib/group-users'
import type { GroupUser } from '../types'

const alice: GroupUser = {
  id: 1,
  username: 'alice',
  display_name: 'Alice',
  group: 'default',
  role: 1,
  status: 1,
}
const bob: GroupUser = {
  id: 2,
  username: 'bob',
  display_name: 'Bob',
  group: 'vip',
  role: 1,
  status: 1,
}
const cara: GroupUser = {
  id: 3,
  username: 'cara',
  display_name: 'Cara',
  group: 'default',
  role: 1,
  status: 1,
}

describe('mergeTransferLists', () => {
  it('moves pending additions into members and pending removals into available', () => {
    const result = mergeTransferLists({
      available: [alice, cara],
      members: [bob],
      pendingAdd: [alice],
      pendingRemove: [bob],
    })
    expect(result.members.map((user) => user.id)).toEqual([alice.id])
    expect(result.available.map((user) => user.id)).toEqual([cara.id, bob.id])
    expect(result.members[0]?.pending).toBe('add')
    expect(result.available.find((user) => user.id === bob.id)?.pending).toBe(
      'remove'
    )
  })
})

describe('buildAssignPayloads', () => {
  it('creates add and fallback remove payloads', () => {
    expect(
      buildAssignPayloads({
        group: 'vip',
        pendingAdd: [alice],
        pendingRemove: [bob],
        fallbackGroup: 'default',
      })
    ).toEqual([
      { group: 'vip', ids: [1] },
      { group: 'default', ids: [2] },
    ])
  })

  it('omits empty sides', () => {
    expect(
      buildAssignPayloads({
        group: 'vip',
        pendingAdd: [alice],
        pendingRemove: [],
        fallbackGroup: 'default',
      })
    ).toEqual([{ group: 'vip', ids: [1] }])
  })
})

describe('resolveFallbackGroup', () => {
  it('uses default unless the current group is already default', () => {
    expect(resolveFallbackGroup('vip', ['default', 'vip'])).toBe('default')
    expect(resolveFallbackGroup('default', ['default', 'vip'])).toBe('vip')
  })
})
