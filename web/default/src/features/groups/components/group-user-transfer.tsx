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
import { ChevronLeft, ChevronRight, Search } from 'lucide-react'
import { useMemo, useState } from 'react'
import { useTranslation } from 'react-i18next'

import { EmptyState } from '@/components/empty-state'
import { LoadingState } from '@/components/loading-state'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Checkbox } from '@/components/ui/checkbox'
import { Input } from '@/components/ui/input'
import { cn } from '@/lib/utils'

import type { GroupUserTransferItem } from '../types'

type GroupUserTransferProps = {
  available: GroupUserTransferItem[]
  members: GroupUserTransferItem[]
  availableKeyword: string
  membersKeyword: string
  onAvailableKeywordChange: (value: string) => void
  onMembersKeywordChange: (value: string) => void
  onAdd: (ids: number[]) => void
  onRemove: (ids: number[]) => void
  isLoading?: boolean
  disabled?: boolean
}

function matchesKeyword(user: GroupUserTransferItem, keyword: string) {
  const query = keyword.trim().toLowerCase()
  if (!query) return true
  return (
    user.username.toLowerCase().includes(query) ||
    (user.display_name || '').toLowerCase().includes(query) ||
    String(user.id).includes(query)
  )
}

function TransferPanel(props: {
  title: string
  description: string
  users: GroupUserTransferItem[]
  keyword: string
  onKeywordChange: (value: string) => void
  selectedIds: Set<number>
  onToggle: (id: number, checked: boolean) => void
  onToggleAll: (checked: boolean) => void
  emptyTitle: string
  searchPlaceholder: string
}) {
  const { t } = useTranslation()
  const filtered = props.users.filter((user) =>
    matchesKeyword(user, props.keyword)
  )
  const allSelected =
    filtered.length > 0 &&
    filtered.every((user) => props.selectedIds.has(user.id))
  const someSelected = filtered.some((user) => props.selectedIds.has(user.id))

  return (
    <div className='flex min-h-0 flex-1 flex-col rounded-xl border'>
      <div className='flex items-start justify-between gap-3 border-b p-3'>
        <div className='min-w-0'>
          <p className='text-sm font-medium'>{props.title}</p>
          <p className='text-muted-foreground text-xs'>{props.description}</p>
        </div>
        <Badge variant='secondary'>{filtered.length}</Badge>
      </div>
      <div className='border-b p-3'>
        <div className='relative'>
          <Search className='text-muted-foreground pointer-events-none absolute top-1/2 left-2.5 size-3.5 -translate-y-1/2' />
          <Input
            value={props.keyword}
            onChange={(event) => props.onKeywordChange(event.target.value)}
            placeholder={props.searchPlaceholder}
            className='pl-8'
            aria-label={props.searchPlaceholder}
          />
        </div>
      </div>
      <div className='flex items-center gap-2 border-b px-3 py-2'>
        <Checkbox
          checked={allSelected}
          indeterminate={someSelected && !allSelected}
          onCheckedChange={(value) => props.onToggleAll(!!value)}
          aria-label={t('Select all')}
          disabled={filtered.length === 0}
        />
        <span className='text-muted-foreground text-xs'>{t('Select all')}</span>
      </div>
      <div className='min-h-0 flex-1 overflow-auto'>
        {filtered.length === 0 ? (
          <EmptyState
            className='min-h-40'
            title={props.emptyTitle}
            description={t('Try another search or choose a different group.')}
          />
        ) : (
          <ul className='divide-y'>
            {filtered.map((user) => {
              const selected = props.selectedIds.has(user.id)
              return (
                <li key={user.id}>
                  <div
                    className={cn(
                      'flex items-start gap-2 px-3 py-2.5',
                      selected && 'bg-muted/50'
                    )}
                  >
                    <Checkbox
                      checked={selected}
                      onCheckedChange={(value) =>
                        props.onToggle(user.id, !!value)
                      }
                      aria-label={t('Select {{name}}', { name: user.username })}
                      className='mt-0.5'
                    />
                    <span className='min-w-0 flex-1'>
                      <span className='flex items-center gap-2'>
                        <span className='truncate text-sm font-medium'>
                          {user.display_name || user.username}
                        </span>
                        {user.pending === 'add' && (
                          <Badge variant='secondary'>{t('Pending add')}</Badge>
                        )}
                        {user.pending === 'remove' && (
                          <Badge variant='outline'>{t('Pending remove')}</Badge>
                        )}
                      </span>
                      <span className='text-muted-foreground block truncate text-xs'>
                        @{user.username} · ID {user.id}
                        {user.group ? ` · ${user.group}` : ''}
                      </span>
                    </span>
                  </div>
                </li>
              )
            })}
          </ul>
        )}
      </div>
    </div>
  )
}

export function GroupUserTransfer(props: GroupUserTransferProps) {
  const { t } = useTranslation()
  const [availableSelected, setAvailableSelected] = useState<Set<number>>(
    new Set()
  )
  const [memberSelected, setMemberSelected] = useState<Set<number>>(new Set())

  const visibleAvailable = useMemo(
    () =>
      props.available.filter((user) =>
        matchesKeyword(user, props.availableKeyword)
      ),
    [props.available, props.availableKeyword]
  )
  const visibleMembers = useMemo(
    () =>
      props.members.filter((user) =>
        matchesKeyword(user, props.membersKeyword)
      ),
    [props.members, props.membersKeyword]
  )

  const moveRight = () => {
    const ids = visibleAvailable
      .filter((user) => availableSelected.has(user.id))
      .map((user) => user.id)
    if (ids.length === 0) return
    props.onAdd(ids)
    setAvailableSelected(new Set())
  }

  const moveLeft = () => {
    const ids = visibleMembers
      .filter((user) => memberSelected.has(user.id))
      .map((user) => user.id)
    if (ids.length === 0) return
    props.onRemove(ids)
    setMemberSelected(new Set())
  }

  if (props.isLoading) {
    return <LoadingState />
  }

  return (
    <div className='flex min-h-0 flex-1 flex-col gap-3 lg:flex-row'>
      <TransferPanel
        title={t('Available users')}
        description={t('Users not in the selected group')}
        users={props.available}
        keyword={props.availableKeyword}
        onKeywordChange={props.onAvailableKeywordChange}
        selectedIds={availableSelected}
        onToggle={(id, checked) => {
          setAvailableSelected((prev) => {
            const next = new Set(prev)
            if (checked) next.add(id)
            else next.delete(id)
            return next
          })
        }}
        onToggleAll={(checked) => {
          setAvailableSelected(
            checked
              ? new Set(visibleAvailable.map((user) => user.id))
              : new Set()
          )
        }}
        emptyTitle={t('No available users')}
        searchPlaceholder={t('Search users')}
      />

      <div className='flex shrink-0 items-center justify-center gap-2 lg:flex-col'>
        <Button
          type='button'
          variant='outline'
          size='icon'
          onClick={moveRight}
          disabled={props.disabled || availableSelected.size === 0}
          aria-label={t('Add selected users')}
        >
          <ChevronRight className='rotate-90 lg:rotate-0' />
        </Button>
        <Button
          type='button'
          variant='outline'
          size='icon'
          onClick={moveLeft}
          disabled={props.disabled || memberSelected.size === 0}
          aria-label={t('Remove selected users')}
        >
          <ChevronLeft className='-rotate-90 lg:rotate-0' />
        </Button>
      </div>

      <TransferPanel
        title={t('Group members')}
        description={t('Users currently assigned to this group')}
        users={props.members}
        keyword={props.membersKeyword}
        onKeywordChange={props.onMembersKeywordChange}
        selectedIds={memberSelected}
        onToggle={(id, checked) => {
          setMemberSelected((prev) => {
            const next = new Set(prev)
            if (checked) next.add(id)
            else next.delete(id)
            return next
          })
        }}
        onToggleAll={(checked) => {
          setMemberSelected(
            checked ? new Set(visibleMembers.map((user) => user.id)) : new Set()
          )
        }}
        emptyTitle={t('No group members')}
        searchPlaceholder={t('Search members')}
      />
    </div>
  )
}
