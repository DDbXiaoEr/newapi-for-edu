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
import { DEFAULT_GROUP } from '@/features/users/constants'

import type { GroupUser, GroupUserTransferItem } from '../types'

export const FALLBACK_GROUP = DEFAULT_GROUP

export function uniqueUsers(users: GroupUser[]): GroupUser[] {
  const seen = new Set<number>()
  const result: GroupUser[] = []
  for (const user of users) {
    if (seen.has(user.id)) continue
    seen.add(user.id)
    result.push(user)
  }
  return result
}

export function mergeTransferLists(params: {
  available: GroupUser[]
  members: GroupUser[]
  pendingAdd: GroupUser[]
  pendingRemove: GroupUser[]
}): {
  available: GroupUserTransferItem[]
  members: GroupUserTransferItem[]
} {
  const pendingAddIds = new Set(params.pendingAdd.map((user) => user.id))
  const pendingRemoveIds = new Set(params.pendingRemove.map((user) => user.id))

  const available = uniqueUsers([
    ...params.available.filter((user) => !pendingAddIds.has(user.id)),
    ...params.pendingRemove,
  ]).map((user) => ({
    ...user,
    pending: pendingRemoveIds.has(user.id) ? ('remove' as const) : undefined,
  }))

  const members = uniqueUsers([
    ...params.members.filter((user) => !pendingRemoveIds.has(user.id)),
    ...params.pendingAdd,
  ]).map((user) => ({
    ...user,
    pending: pendingAddIds.has(user.id) ? ('add' as const) : undefined,
  }))

  return { available, members }
}

export function buildAssignPayloads(params: {
  group: string
  pendingAdd: GroupUser[]
  pendingRemove: GroupUser[]
  fallbackGroup: string
}): { group: string; ids: number[] }[] {
  const payloads: { group: string; ids: number[] }[] = []
  const addIds = params.pendingAdd.map((user) => user.id)
  if (addIds.length > 0) {
    payloads.push({ group: params.group, ids: addIds })
  }
  const removeIds = params.pendingRemove.map((user) => user.id)
  if (removeIds.length > 0) {
    payloads.push({ group: params.fallbackGroup, ids: removeIds })
  }
  return payloads
}

export function resolveFallbackGroup(currentGroup: string, groups: string[]) {
  if (currentGroup !== FALLBACK_GROUP) return FALLBACK_GROUP
  return groups.find((group) => group !== currentGroup) ?? FALLBACK_GROUP
}
