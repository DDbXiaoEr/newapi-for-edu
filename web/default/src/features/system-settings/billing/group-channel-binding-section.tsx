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
import { Combobox } from '@/components/ui/combobox'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Skeleton } from '@/components/ui/skeleton'
import { api } from '@/lib/api'
import { handleServerError } from '@/lib/handle-server-error'
import { requireServerSuccess } from '@/lib/server-error-message'

import { SettingsSection } from '../components/settings-section'

const CHANNEL_STATUS_ENABLED = 1
const WILDCARD_MODEL = '*'

type ChannelSummary = {
  id: number
  name: string
  status: number
  models: string[]
}

type BindingPayload = {
  group_channels: Record<string, Record<string, number[]>>
  channels: ChannelSummary[]
  stale_channels: Record<string, Record<string, number[]>>
  groups: string[]
  group_models: Record<string, string[]>
}

type Envelope<T> = {
  success: boolean
  message?: string
  data: T
}

function channelSupportsModel(channel: ChannelSummary, model: string) {
  if (model === WILDCARD_MODEL) return true
  return channel.models.includes(model)
}

function pinsForGroup(
  payload: BindingPayload | null,
  group: string
): Record<string, number[]> {
  return payload?.group_channels[group] ?? {}
}

export function GroupChannelBindingSection() {
  const { t } = useTranslation()
  const [loading, setLoading] = useState(true)
  const [saving, setSaving] = useState(false)
  const [payload, setPayload] = useState<BindingPayload | null>(null)
  const [group, setGroup] = useState('')
  const [model, setModel] = useState('')
  const [selected, setSelected] = useState<Set<number>>(new Set())

  const fetchPayload = useCallback(async () => {
    const res = await api.get<Envelope<BindingPayload>>('/api/group/channels')
    return requireServerSuccess(res.data).data
  }, [])

  const applyGroupState = useCallback(
    (data: BindingPayload, nextGroup: string, keepModel?: string) => {
      const models = modelOptionsForGroup(data, nextGroup, t)
      const nextModel =
        keepModel && models.some((item) => item.value === keepModel)
          ? keepModel
          : preferredModel(models, pinsForGroup(data, nextGroup))
      setGroup(nextGroup)
      setModel(nextModel)
      setSelected(
        new Set(pinsForGroup(data, nextGroup)[nextModel] ?? [])
      )
    },
    [t]
  )

  useEffect(() => {
    let active = true
    setLoading(true)
    void fetchPayload()
      .then((data) => {
        if (!active || !data) return
        setPayload(data)
        applyGroupState(data, data.groups[0] ?? '')
      })
      .catch((error) => {
        if (active) {
          handleServerError(error, t('Failed to load group channel bindings'))
        }
      })
      .finally(() => {
        if (active) setLoading(false)
      })
    return () => {
      active = false
    }
  }, [applyGroupState, fetchPayload, t])

  const channelById = useMemo(
    () => new Map((payload?.channels ?? []).map((channel) => [channel.id, channel])),
    [payload]
  )

  const modelOptions = useMemo(
    () => modelOptionsForGroup(payload, group, t),
    [payload, group, t]
  )

  const matchingChannels = useMemo(() => {
    if (!model) return []
    return (payload?.channels ?? []).filter((channel) =>
      channelSupportsModel(channel, model)
    )
  }, [payload, model])

  const currentPins = useMemo(() => {
    const pins = pinsForGroup(payload, group)
    return Object.keys(pins)
      .sort((left, right) => {
        if (left === WILDCARD_MODEL) return -1
        if (right === WILDCARD_MODEL) return 1
        return left.localeCompare(right)
      })
      .map((name) => ({
        model: name,
        channelIds: pins[name] ?? [],
      }))
  }, [payload, group])

  const handleGroupChange = (value: string | null) => {
    if (!value || !payload) return
    applyGroupState(payload, value)
  }

  const handleModelChange = (value: string | null) => {
    if (!value) return
    setModel(value)
    setSelected(new Set(pinsForGroup(payload, group)[value] ?? []))
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
    if (!group || !model) return
    setSaving(true)
    try {
      const res = await api.put<Envelope<BindingPayload>>(
        '/api/group/channels',
        {
          group,
          model,
          channel_ids: channelIds,
        }
      )
      requireServerSuccess(res.data)
      const data = await fetchPayload()
      if (data) {
        setPayload(data)
        applyGroupState(data, group, model)
      }
      toast.success(
        channelIds.length === 0
          ? t('Binding cleared for this model')
          : t('Group channel binding saved')
      )
    } catch (error) {
      handleServerError(error, t('Failed to save'))
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
          'Pin a model in a user group to specific channels. Requests for that model only use the pinned channels; other models keep the default channel selection.'
        )}
      </p>

      <div className='flex flex-wrap items-center gap-3'>
        <span className='text-sm font-medium'>{t('User Group')}</span>
        <Select value={group} onValueChange={handleGroupChange}>
          <SelectTrigger className='w-56' aria-label={t('Select group')}>
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
        <span className='text-sm font-medium'>{t('Model')}</span>
        <Combobox
          options={modelOptions}
          value={model}
          onValueChange={handleModelChange}
          placeholder={t('Select model')}
          searchPlaceholder={t('Select model')}
          className='w-64'
          aria-label={t('Select model')}
        />
      </div>

      {currentPins.length > 0 && (
        <div className='flex flex-col gap-2'>
          <span className='text-sm font-medium'>{t('Current pins')}</span>
          <div className='flex flex-wrap gap-1'>
            {currentPins.map((pin) => (
              <Badge
                key={pin.model}
                variant={pin.model === model ? 'default' : 'outline'}
              >
                {pin.model === WILDCARD_MODEL
                  ? t('All models')
                  : pin.model}{' '}
                →{' '}
                {pin.channelIds
                  .map((id) => {
                    const channel = channelById.get(id)
                    return channel ? `${channel.name} #${id}` : `#${id}`
                  })
                  .join(', ')}
              </Badge>
            ))}
          </div>
        </div>
      )}

      <div className='flex flex-col gap-2'>
        <span className='text-sm font-medium'>
          {t('Channels that serve this model')}
        </span>
        <div className='border-border h-72 overflow-y-auto rounded-md border p-2'>
          <div className='flex flex-col gap-1'>
            {matchingChannels.map((channel) => (
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
              </label>
            ))}
            {matchingChannels.length === 0 && (
              <span className='text-muted-foreground px-2 py-1.5 text-sm'>
                {t('No channel serves this model')}
              </span>
            )}
          </div>
        </div>
      </div>

      <div className='flex flex-wrap items-center gap-2'>
        <Button
          onClick={() => void persist([...selected])}
          disabled={saving || !group || !model}
        >
          {t('Save binding')}
        </Button>
        <Button
          variant='outline'
          onClick={() => void persist([])}
          disabled={saving || !group || !model}
        >
          {t('Clear binding')}
        </Button>
        {selected.size === 0 && (
          <span className='text-muted-foreground text-xs'>
            {t(
              "No channel selected. Saving will clear this model's pin."
            )}
          </span>
        )}
      </div>
    </SettingsSection>
  )
}

function preferredModel(
  models: { value: string }[],
  pins: Record<string, number[]>
) {
  const pinned = Object.keys(pins).sort((left, right) => {
    if (left === WILDCARD_MODEL) return 1
    if (right === WILDCARD_MODEL) return -1
    return left.localeCompare(right)
  })
  const fromPins = pinned.find((name) =>
    models.some((item) => item.value === name)
  )
  if (fromPins) return fromPins
  return (
    models.find((item) => item.value !== WILDCARD_MODEL)?.value ??
    models[0]?.value ??
    ''
  )
}

function modelOptionsForGroup(
  payload: BindingPayload | null,
  group: string,
  t: (key: string) => string
) {
  const names = new Set<string>([WILDCARD_MODEL])
  for (const name of payload?.group_models[group] ?? []) names.add(name)
  for (const name of Object.keys(pinsForGroup(payload, group))) {
    names.add(name)
  }
  for (const channel of payload?.channels ?? []) {
    for (const name of channel.models) names.add(name)
  }
  return [...names]
    .sort((left, right) => {
      if (left === WILDCARD_MODEL) return -1
      if (right === WILDCARD_MODEL) return 1
      return left.localeCompare(right)
    })
    .map((name) => ({
      value: name,
      label: name === WILDCARD_MODEL ? t('All models') : name,
    }))
}
