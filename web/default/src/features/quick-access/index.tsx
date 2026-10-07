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
import { useQuery } from '@tanstack/react-query'
import { useMemo, useState } from 'react'
import { useTranslation } from 'react-i18next'

import { SectionPageLayout } from '@/components/layout'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import { useActiveChatKey } from '@/features/chat/hooks/use-active-chat-key'
import { useChatPresets } from '@/features/chat/hooks/use-chat-presets'
import { getUserModels } from '@/lib/api'
import { requireServerSuccess } from '@/lib/server-error-message'

import { ClientDetail } from './components/client-detail'
import { getClientIcon } from './components/client-icon'
import {
  BUILTIN_CLIENTS,
  type ClientOption,
  type QuickAccessContext,
} from './lib/clients'

export function QuickAccess() {
  const { t } = useTranslation()
  const { chatPresets, serverAddress } = useChatPresets()
  const { data: activeKey } = useActiveChatKey(true, true)
  const { data: modelsData } = useQuery({
    queryKey: ['user-models-quick-access'],
    queryFn: async () => requireServerSuccess(await getUserModels()),
    staleTime: 5 * 60 * 1000,
  })

  const clients = useMemo<ClientOption[]>(() => {
    const builtins: ClientOption[] = BUILTIN_CLIENTS.map((builtin) => ({
      id: `builtin:${builtin.id}`,
      name: builtin.name,
      kind: 'builtin',
      builtin,
    }))
    const presets: ClientOption[] = chatPresets
      .filter((preset) => preset.type !== 'fluent')
      .map((preset) => ({
        id: `preset:${preset.id}`,
        name: preset.name,
        kind: 'preset',
        preset,
      }))
    return [...builtins, ...presets]
  }, [chatPresets])

  const [selectedId, setSelectedId] = useState('builtin:opencode')
  const selected =
    clients.find((client) => client.id === selectedId) ?? clients[0] ?? null

  const apiKey = activeKey?.trim() ?? ''
  const ctx: QuickAccessContext = {
    serverAddress,
    apiKey,
    model: modelsData?.data?.[0],
  }

  return (
    <SectionPageLayout>
      <SectionPageLayout.Title>{t('Quick Access')}</SectionPageLayout.Title>
      <SectionPageLayout.Content>
        <div className='mx-auto flex w-full max-w-5xl flex-col gap-6'>
          <p className='text-muted-foreground text-sm'>
            {t(
              'Connect developer tools and chat clients to this service. Pick a client to see how to set it up.'
            )}
          </p>

          <div>
            <h2 className='text-muted-foreground mb-2 text-xs font-semibold tracking-wide uppercase'>
              {t('Clients')}
            </h2>
            <div className='flex flex-wrap gap-2'>
              {clients.map((client) => {
                const Icon = getClientIcon(client)
                const active = client.id === selected?.id
                return (
                  <Button
                    key={client.id}
                    variant={active ? 'default' : 'outline'}
                    size='sm'
                    aria-pressed={active}
                    onClick={() => setSelectedId(client.id)}
                  >
                    <Icon className='size-4' aria-hidden='true' />
                    {client.name}
                  </Button>
                )
              })}
            </div>
          </div>

          <Separator />

          {selected ? (
            <ClientDetail
              client={selected}
              ctx={ctx}
              hasApiKey={Boolean(apiKey)}
            />
          ) : null}
        </div>
      </SectionPageLayout.Content>
    </SectionPageLayout>
  )
}
