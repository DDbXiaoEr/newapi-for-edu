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

import { Dialog } from '@/components/dialog'
import { ErrorState } from '@/components/error-state'
import { Button } from '@/components/ui/button'
import { Combobox } from '@/components/ui/combobox'
import { Label } from '@/components/ui/label'
import { assignUsersGroup } from '@/features/groups/api'
import { GroupUserTransfer } from '@/features/groups/components/group-user-transfer'
import { uniqueUsers } from '@/features/groups/lib/group-users'
import type { GroupUser } from '@/features/groups/types'
import { handleServerError } from '@/lib/handle-server-error'
import { requireServerSuccess } from '@/lib/server-error-message'

import { getGroups, searchUsers } from '../../api'
import {
  ERROR_MESSAGES,
  MAX_ASSIGN_USERS_GROUP,
  SUCCESS_MESSAGES,
} from '../../constants'
import { useUsers } from '../users-provider'

const PAGE_SIZE = 100

async function loadUsers(): Promise<GroupUser[]> {
  const collected: GroupUser[] = []
  let page = 1
  let total = Infinity
  while (collected.length < total) {
    const result = await searchUsers({
      keyword: '',
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
  return uniqueUsers(collected)
}

type UsersAssignGroupDialogProps = {
  open: boolean
  onOpenChange: (open: boolean) => void
}

export function UsersAssignGroupDialog(props: UsersAssignGroupDialogProps) {
  const { t } = useTranslation()
  const queryClient = useQueryClient()
  const { triggerRefresh } = useUsers()
  const [selectedGroup, setSelectedGroup] = useState<string | null>(null)
  const [selectedUsers, setSelectedUsers] = useState<GroupUser[]>([])
  const [availableKeyword, setAvailableKeyword] = useState('')
  const [selectedKeyword, setSelectedKeyword] = useState('')

  const groupsQuery = useQuery({
    queryKey: ['groups'],
    queryFn: async () => requireServerSuccess(await getGroups()),
    enabled: props.open,
    staleTime: 5 * 60 * 1000,
  })
  const usersQuery = useQuery({
    queryKey: ['users', 'assign-group'],
    queryFn: loadUsers,
    enabled: props.open,
  })

  const groups = groupsQuery.data?.data ?? []
  const selectedIds = useMemo(
    () => new Set(selectedUsers.map((user) => user.id)),
    [selectedUsers]
  )
  const availableUsers = useMemo(
    () => (usersQuery.data ?? []).filter((user) => !selectedIds.has(user.id)),
    [usersQuery.data, selectedIds]
  )
  const tooManySelected = selectedUsers.length > MAX_ASSIGN_USERS_GROUP

  const resetState = () => {
    setSelectedGroup(null)
    setSelectedUsers([])
    setAvailableKeyword('')
    setSelectedKeyword('')
  }

  const saveMutation = useMutation({
    mutationFn: async () => {
      if (!selectedGroup || selectedUsers.length === 0) {
        throw new Error(t('Please select users and a group'))
      }
      if (tooManySelected) {
        throw new Error(
          t('You can assign up to {{count}} users at a time.', {
            count: MAX_ASSIGN_USERS_GROUP,
          })
        )
      }
      return requireServerSuccess(
        await assignUsersGroup({
          ids: selectedUsers.map((user) => user.id),
          group: selectedGroup,
        })
      )
    },
    onSuccess: async (result) => {
      const updated = result.data?.updated ?? 0
      const skipped = result.data?.skipped ?? 0
      if (updated === 0) {
        toast.success(t(SUCCESS_MESSAGES.USERS_NONE_UPDATED))
      } else if (skipped > 0) {
        toast.success(
          t(SUCCESS_MESSAGES.USERS_ASSIGNED_WITH_SKIPPED, {
            updated,
            group: selectedGroup,
            skipped,
          })
        )
      } else {
        toast.success(
          t(SUCCESS_MESSAGES.USERS_ASSIGNED, {
            updated,
            group: selectedGroup,
          })
        )
      }
      resetState()
      props.onOpenChange(false)
      triggerRefresh()
      await queryClient.invalidateQueries({ queryKey: ['users'] })
      await queryClient.invalidateQueries({ queryKey: ['group-users'] })
    },
    onError: (error) =>
      handleServerError(error, t(ERROR_MESSAGES.ASSIGN_GROUP_FAILED)),
  })

  const handleOpenChange = (open: boolean) => {
    if (!open) resetState()
    props.onOpenChange(open)
  }

  const handleAdd = (ids: number[]) => {
    const idSet = new Set(ids)
    const extra = availableUsers.filter((user) => idSet.has(user.id))
    setSelectedUsers((prev) => uniqueUsers([...prev, ...extra]))
  }

  const handleRemove = (ids: number[]) => {
    const idSet = new Set(ids)
    setSelectedUsers((prev) => prev.filter((user) => !idSet.has(user.id)))
  }

  const canSubmit =
    Boolean(selectedGroup) &&
    selectedUsers.length > 0 &&
    !tooManySelected &&
    !saveMutation.isPending

  let dialogBody = (
    <div className='flex h-full min-h-0 flex-col gap-4'>
      <div className='space-y-1.5'>
        <Label htmlFor='assign-group-select'>{t('Group')}</Label>
        <Combobox
          id='assign-group-select'
          options={groups.map((group) => ({ value: group, label: group }))}
          value={selectedGroup}
          onValueChange={setSelectedGroup}
          placeholder={t('Select a group')}
          aria-label={t('Select a group')}
          disabled={saveMutation.isPending}
        />
      </div>
      {tooManySelected ? (
        <p className='text-destructive text-sm'>
          {t('You can assign up to {{count}} users at a time.', {
            count: MAX_ASSIGN_USERS_GROUP,
          })}
        </p>
      ) : null}
      <div className='flex min-h-0 flex-1 flex-col'>
        <GroupUserTransfer
          available={availableUsers}
          members={selectedUsers}
          availableKeyword={availableKeyword}
          membersKeyword={selectedKeyword}
          onAvailableKeywordChange={setAvailableKeyword}
          onMembersKeywordChange={setSelectedKeyword}
          onAdd={handleAdd}
          onRemove={handleRemove}
          isLoading={usersQuery.isLoading}
          disabled={saveMutation.isPending}
          availableTitle={t('Available users')}
          availableDescription={t('Search users to assign')}
          availableEmptyTitle={t('No available users')}
          membersTitle={t('Selected users')}
          membersDescription={t(
            'Users that will be assigned to the chosen group'
          )}
          membersEmptyTitle={t('No selected users')}
          emptyDescription={t('Try another search.')}
          searchAvailablePlaceholder={t('Search users')}
          searchMembersPlaceholder={t('Search selected users')}
        />
      </div>
    </div>
  )
  if (groupsQuery.isError && !groupsQuery.data) {
    dialogBody = <ErrorState onRetry={() => void groupsQuery.refetch()} />
  } else if (usersQuery.isError && !usersQuery.data) {
    dialogBody = <ErrorState onRetry={() => void usersQuery.refetch()} />
  }

  return (
    <Dialog
      open={props.open}
      onOpenChange={handleOpenChange}
      title={t('Batch assign group')}
      description={t(
        'Search and select users, then choose a group. Users who already belong to a group will be moved to the new group.'
      )}
      contentClassName='sm:max-w-5xl'
      contentHeight='min(68vh, 640px)'
      bodyClassName='flex h-full min-h-0 flex-col gap-4 overflow-hidden'
      footer={
        <>
          <Button
            type='button'
            variant='outline'
            onClick={() => handleOpenChange(false)}
            disabled={saveMutation.isPending}
          >
            {t('Cancel')}
          </Button>
          <Button
            type='button'
            onClick={() => saveMutation.mutate()}
            disabled={!canSubmit}
          >
            {saveMutation.isPending ? t('Saving...') : t('Confirm assignment')}
          </Button>
        </>
      }
    >
      {dialogBody}
    </Dialog>
  )
}
