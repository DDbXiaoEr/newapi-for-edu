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
import { useCallback, useEffect, useMemo, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { toast } from 'sonner'

import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Checkbox } from '@/components/ui/checkbox'
import { ScrollArea } from '@/components/ui/scroll-area'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Skeleton } from '@/components/ui/skeleton'
import { api } from '@/lib/api'

import { SettingsSection } from '../components/settings-section'

type ChannelSummary = {
  id: number
  name: string
  status: number
  models: string[]
}

type BindingPayload = {
  group_channels: Record<string, number[]>
  channels: ChannelSummary[]
  stale_channels: Record<string, number[]>
  groups: string[]
  group_models: Record<string, string[]>
}

type Envelope<T> = {
  success: boolean
  message?: string
  data: T
}

const CHANNEL_STATUS_ENABLED = 1

export function GroupChannelBindingSection() {
  const { t } = useTranslation()
  const [loading, setLoading] = useState(true)
  const [saving, setSaving] = useState(false)
  const [payload, setPayload] = useState<BindingPayload | null>(null)
  const [group, setGroup] = useState('')
  const [selected, setSelected] = useState<Set<number>>(new Set())

  const fetchPayload = useCallback(async () => {
    const res = await api.get<Envelope<BindingPayload>>('/api/group/channels')
    return res.data?.data
  }, [])

  useEffect(() => {
    let active = true
    setLoading(true)
    void fetchPayload()
      .then((data) => {
        if (!active || !data) return
        setPayload(data)
        const first = data.groups[0] ?? ''
        setGroup(first)
        setSelected(new Set(data.group_channels[first] ?? []))
      })
      .catch(() => {
        if (active) toast.error(t('Failed to load group channel bindings'))
      })
      .finally(() => {
        if (active) setLoading(false)
      })
    return () => {
      active = false
    }
  }, [fetchPayload, t])

  const channelById = useMemo(
    () => new Map((payload?.channels ?? []).map((c) => [c.id, c])),
    [payload]
  )

  const previewModels = useMemo(() => {
    const models = new Set<string>()
    for (const id of selected) {
      const channel = channelById.get(id)
      if (!channel || channel.status !== CHANNEL_STATUS_ENABLED) continue
      for (const model of channel.models) models.add(model)
    }
    return [...models].sort()
  }, [selected, channelById])

  const currentModels = payload?.group_models[group]
  const lostModels = useMemo(
    () =>
      (currentModels ?? []).filter(
        (model) => !previewModels.includes(model)
      ),
    [currentModels, previewModels]
  )

  const handleGroupChange = (value: string | null) => {
    if (!value) return
    setGroup(value)
    setSelected(new Set(payload?.group_channels[value] ?? []))
  }

  const toggleChannel = (id: number) => {
    setSelected((prev) => {
      const next = new Set(prev)
      if (next.has(id)) {
        next.delete(id)
      } else {
        next.add(id)
      }
      return next
    })
  }

  const persist = async (channelIds: number[]) => {
    if (!group) return
    setSaving(true)
    try {
      await api.put('/api/group/channels', {
        group,
        channel_ids: channelIds,
      })
      const data = await fetchPayload()
      if (data) setPayload(data)
      setSelected(new Set(channelIds))
      toast.success(
        channelIds.length === 0
          ? t('Binding cleared for this group')
          : t('Group channel binding saved')
      )
    } catch {
      toast.error(t('Failed to save'))
    } finally {
      setSaving(false)
    }
  }

  if (loading) {
    return (
      <SettingsSection title={t('Group Channel Binding')}>
        <Skeleton className='h-40 w-full' />
      </SettingsSection>
    )
  }

  return (
    <SettingsSection title={t('Group Channel Binding')}>
      <p className='text-muted-foreground text-sm'>
        {t(
          "Bind a user group to specific channels. Keys in a bound group only route through the bound channels; unbound groups keep each channel's own group settings."
        )}
      </p>

      <div className='flex flex-wrap items-center gap-3'>
        <span className='text-sm font-medium'>{t('User Group')}</span>
        <Select value={group} onValueChange={handleGroupChange}>
          <SelectTrigger className='w-56'>
            <SelectValue placeholder={t('Select group')} />
          </SelectTrigger>
          <SelectContent>
            {(payload?.groups ?? []).map((name) => (
              <SelectItem key={name} value={name}>
                {name}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>

      <div className='grid gap-4 lg:grid-cols-2'>
        <div className='flex flex-col gap-2'>
          <span className='text-sm font-medium'>{t('Bindable Channels')}</span>
          <ScrollArea className='border-border h-72 rounded-md border p-2'>
            <div className='flex flex-col gap-1'>
              {(payload?.channels ?? []).map((channel) => (
                <label
                  key={channel.id}
                  className='hover:bg-muted/60 flex cursor-pointer items-center gap-2 rounded px-2 py-1.5 text-sm'
                >
                  <Checkbox
                    checked={selected.has(channel.id)}
                    onCheckedChange={() => toggleChannel(channel.id)}
                    disabled={channel.status !== CHANNEL_STATUS_ENABLED}
                  />
                  <span className='truncate'>
                    {channel.name} #{channel.id}
                  </span>
                  {channel.status !== CHANNEL_STATUS_ENABLED && (
                    <Badge variant='secondary'>{t('Disabled')}</Badge>
                  )}
                  <span className='text-muted-foreground ml-auto shrink-0 text-xs'>
                    {channel.models.length} {t('models')}
                  </span>
                </label>
              ))}
            </div>
          </ScrollArea>
        </div>

        <div className='flex flex-col gap-2'>
          <span className='text-sm font-medium'>
            {t('Available models after binding')}（{previewModels.length}）
          </span>
          <ScrollArea className='border-border h-36 rounded-md border p-2'>
            <div className='flex flex-wrap gap-1'>
              {previewModels.map((model) => (
                <Badge key={model} variant='outline'>
                  {model}
                </Badge>
              ))}
            </div>
          </ScrollArea>
          <span className='text-sm font-medium'>
            {t('Models lost after binding')}（{lostModels.length}）
          </span>
          <ScrollArea className='border-border h-20 rounded-md border p-2'>
            <div className='flex flex-wrap gap-1'>
              {lostModels.map((model) => (
                <Badge key={model} variant='destructive'>
                  {model}
                </Badge>
              ))}
            </div>
          </ScrollArea>
        </div>
      </div>

      <div className='flex flex-wrap items-center gap-2'>
        <Button
          onClick={() => void persist([...selected])}
          disabled={saving || !group}
        >
          {t('Save binding')}
        </Button>
        <Button
          variant='outline'
          onClick={() => void persist([])}
          disabled={saving || !group}
        >
          {t('Clear binding')}
        </Button>
        {selected.size === 0 && (
          <span className='text-muted-foreground text-xs'>
            {t("No channel selected. Saving will clear this group's binding.")}
          </span>
        )}
      </div>
    </SettingsSection>
  )
}
