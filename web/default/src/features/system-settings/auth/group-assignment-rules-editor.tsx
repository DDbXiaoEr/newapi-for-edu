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
import { Plus, Trash2 } from 'lucide-react'
import { useQuery } from '@tanstack/react-query'
import { useTranslation } from 'react-i18next'

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { getGroups } from '@/features/users/api'
import { requireServerSuccess } from '@/lib/server-error-message'

import {
  GROUP_FILTER_ATTRIBUTES,
  type GroupAssignmentEditorValue,
} from './group-assignment-rules'

type GroupAssignmentRulesEditorProps = {
  value: GroupAssignmentEditorValue
  onChange: (value: GroupAssignmentEditorValue) => void
}

export function GroupAssignmentRulesEditor(
  props: GroupAssignmentRulesEditorProps
) {
  const { t } = useTranslation()
  const groupsQuery = useQuery({
    queryKey: ['system-groups'],
    queryFn: async () => requireServerSuccess(await getGroups()),
  })
  const groups = (groupsQuery.data?.data || []).filter(
    (group) => group && group !== 'auto'
  )

  const updateRule = (
    index: number,
    patch: Partial<GroupAssignmentEditorValue['rules'][number]>
  ) => {
    props.onChange({
      ...props.value,
      rules: props.value.rules.map((rule, ruleIndex) =>
        ruleIndex === index ? { ...rule, ...patch } : rule
      ),
    })
  }

  return (
    <div className='space-y-3'>
      <div className='space-y-1.5'>
        <p className='text-sm font-medium'>{t('Filter attribute')}</p>
        <Select
          value={props.value.attribute || 'uid'}
          onValueChange={(attribute) => {
            if (!attribute) return
            props.onChange({ ...props.value, attribute })
          }}
        >
          <SelectTrigger>
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            {GROUP_FILTER_ATTRIBUTES.map((attribute) => (
              <SelectItem key={attribute.value} value={attribute.value}>
                {t(attribute.labelKey)} ({attribute.value})
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
        <p className='text-muted-foreground text-xs'>
          {t(
            'All conditions share this attribute. The first matching regex assigns the group.'
          )}
        </p>
      </div>

      {props.value.rules.map((rule, index) => (
        <div
          key={rule.id}
          className='flex flex-col gap-2 sm:flex-row sm:items-center'
        >
          <span className='text-muted-foreground w-20 shrink-0 text-sm'>
            {t('Condition {{n}}', { n: index + 1 })}
          </span>
          <Input
            value={rule.pattern}
            placeholder={t('Regex, e.g. ^stu-')}
            onChange={(event) =>
              updateRule(index, { pattern: event.target.value })
            }
          />
          <Select
            value={rule.group || undefined}
            onValueChange={(group) => {
              if (!group) return
              updateRule(index, { group })
            }}
          >
            <SelectTrigger className='sm:w-48'>
              <SelectValue placeholder={t('Assign group')} />
            </SelectTrigger>
            <SelectContent>
              {groups.map((group) => (
                <SelectItem key={group} value={group}>
                  {group}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
          <Button
            type='button'
            variant='ghost'
            size='icon'
            onClick={() =>
              props.onChange({
                ...props.value,
                rules: props.value.rules.filter(
                  (_rule, ruleIndex) => ruleIndex !== index
                ),
              })
            }
            aria-label={t('Remove condition')}
          >
            <Trash2 className='h-4 w-4' />
          </Button>
        </div>
      ))}

      <Button
        type='button'
        variant='outline'
        size='sm'
        onClick={() =>
          props.onChange({
            ...props.value,
            rules: [
              ...props.value.rules,
              {
                id: `empty-${props.value.rules.length}-${Date.now()}`,
                pattern: '',
                group: '',
              },
            ],
          })
        }
      >
        <Plus className='mr-1.5 h-4 w-4' />
        {t('Add condition')}
      </Button>
    </div>
  )
}
