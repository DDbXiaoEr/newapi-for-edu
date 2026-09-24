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
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { useMemo, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { toast } from 'sonner'

import { ErrorState } from '@/components/error-state'
import { Button } from '@/components/ui/button'
import { Combobox } from '@/components/ui/combobox'
import { getGroups, searchUsers } from '@/features/users/api'
import { handleServerError } from '@/lib/handle-server-error'
import { requireServerSuccess } from '@/lib/server-error-message'

import { assignUsersGroup } from '../api'
import {
  buildAssignPayloads,
  mergeTransferLists,
  resolveFallbackGroup,
} from '../lib/group-users'
import type { GroupUser } from '../types'
import { GroupUserTransfer } from './group-user-transfer'

const PAGE_SIZE = 100

async function loadUsers(params: {
  group?: string
  exclude_group?: string
  keyword?: string
}): Promise<GroupUser[]> {
  const collected: GroupUser[] = []
  let page = 1
  let total = Infinity
  while (collected.length < total) {
    const result = await searchUsers({
      keyword: params.keyword ?? '',
      group: params.group ?? '',
      exclude_group: params.exclude_group ?? '',
      exclude_deleted: true,
      p: page,
      page_size: PAGE_SIZE,
      sort_by: 'id',
      sort_order: 'asc',
    })
    const payload = requireServerSuccess(result)
    const items = payload.data?.items ?? []
    total = payload.data?.total ?? items.length
    collected.push(...items)
    if (items.length === 0 || collected.length >= total) break
    page += 1
  }
  return collected
}

export function GroupUsersPanel() {
  const { t } = useTranslation()
  const queryClient = useQueryClient()
  const [selectedGroup, setSelectedGroup] = useState<string | null>(null)
  const [availableKeyword, setAvailableKeyword] = useState('')
  const [membersKeyword, setMembersKeyword] = useState('')
  const [pendingAdd, setPendingAdd] = useState<GroupUser[]>([])
  const [pendingRemove, setPendingRemove] = useState<GroupUser[]>([])

  const groupsQuery = useQuery({
    queryKey: ['groups'],
    queryFn: async () => requireServerSuccess(await getGroups()),
    staleTime: 5 * 60 * 1000,
  })
  const groups = groupsQuery.data?.data ?? []
  const activeGroup = selectedGroup ?? groups[0] ?? ''

  const membersQuery = useQuery({
    queryKey: ['group-users', 'members', activeGroup],
    queryFn: () => loadUsers({ group: activeGroup }),
    enabled: Boolean(activeGroup),
  })
  const availableQuery = useQuery({
    queryKey: ['group-users', 'available', activeGroup],
    queryFn: () => loadUsers({ exclude_group: activeGroup }),
    enabled: Boolean(activeGroup),
  })

  const transfer = useMemo(
    () =>
      mergeTransferLists({
        available: availableQuery.data ?? [],
        members: membersQuery.data ?? [],
        pendingAdd,
        pendingRemove,
      }),
    [availableQuery.data, membersQuery.data, pendingAdd, pendingRemove]
  )

  const saveMutation = useMutation({
    mutationFn: async () => {
      const payloads = buildAssignPayloads({
        group: activeGroup,
        pendingAdd,
        pendingRemove,
        fallbackGroup: resolveFallbackGroup(activeGroup, groups),
      })
      for (const payload of payloads) {
        requireServerSuccess(await assignUsersGroup(payload))
      }
    },
    onSuccess: async () => {
      setPendingAdd([])
      setPendingRemove([])
      await queryClient.invalidateQueries({ queryKey: ['group-users'] })
      await queryClient.invalidateQueries({ queryKey: ['users'] })
      toast.success(t('Group members updated'))
    },
    onError: (error) =>
      handleServerError(error, t('Failed to update group members')),
  })

  const handleAdd = (ids: number[]) => {
    const idSet = new Set(ids)
    setPendingRemove((prev) => prev.filter((user) => !idSet.has(user.id)))
    const extra = transfer.available.filter((user) => idSet.has(user.id))
    setPendingAdd((prev) => {
      const existing = new Set(prev.map((user) => user.id))
      return [...prev, ...extra.filter((user) => !existing.has(user.id))]
    })
  }

  const handleRemove = (ids: number[]) => {
    const idSet = new Set(ids)
    setPendingAdd((prev) => prev.filter((user) => !idSet.has(user.id)))
    const extra = transfer.members.filter((user) => idSet.has(user.id))
    setPendingRemove((prev) => {
      const existing = new Set(prev.map((user) => user.id))
      return [...prev, ...extra.filter((user) => !existing.has(user.id))]
    })
  }

  if (groupsQuery.isError && !groupsQuery.data) {
    return <ErrorState onRetry={() => void groupsQuery.refetch()} />
  }

  const dirty = pendingAdd.length + pendingRemove.length > 0
  const fallbackGroup = resolveFallbackGroup(activeGroup, groups)
  const cannotRemoveFromDefault =
    activeGroup === fallbackGroup && pendingRemove.length > 0

  return (
    <div className='flex min-h-0 flex-1 flex-col gap-4'>
      <div className='flex flex-wrap items-end justify-between gap-3'>
        <div className='min-w-60 flex-1'>
          <p className='mb-1.5 text-sm font-medium'>{t('Select a group')}</p>
          <Combobox
            options={groups.map((group) => ({ value: group, label: group }))}
            value={activeGroup}
            onValueChange={(value) => {
              if (!value) return
              setSelectedGroup(value)
              setPendingAdd([])
              setPendingRemove([])
              setAvailableKeyword('')
              setMembersKeyword('')
            }}
            placeholder={t('Select a group')}
          />
        </div>
        <div className='flex items-center gap-2'>
          <Button
            type='button'
            variant='outline'
            onClick={() => {
              setPendingAdd([])
              setPendingRemove([])
            }}
            disabled={!dirty || saveMutation.isPending}
          >
            {t('Reset')}
          </Button>
          <Button
            type='button'
            onClick={() => saveMutation.mutate()}
            disabled={
              !dirty ||
              saveMutation.isPending ||
              !activeGroup ||
              cannotRemoveFromDefault
            }
          >
            {saveMutation.isPending ? t('Saving...') : t('Save group members')}
          </Button>
        </div>
      </div>
      {cannotRemoveFromDefault && (
        <p className='text-destructive text-sm'>
          {t(
            'Create another group before moving users out of the default group.'
          )}
        </p>
      )}
      <GroupUserTransfer
        available={transfer.available}
        members={transfer.members}
        availableKeyword={availableKeyword}
        membersKeyword={membersKeyword}
        onAvailableKeywordChange={setAvailableKeyword}
        onMembersKeywordChange={setMembersKeyword}
        onAdd={handleAdd}
        onRemove={handleRemove}
        isLoading={membersQuery.isLoading || availableQuery.isLoading}
        disabled={!activeGroup || saveMutation.isPending}
      />
    </div>
  )
}
