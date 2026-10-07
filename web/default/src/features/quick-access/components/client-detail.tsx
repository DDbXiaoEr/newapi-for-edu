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
import { Link } from '@tanstack/react-router'
import { ArrowUpRight, ExternalLink } from 'lucide-react'
import { useTranslation } from 'react-i18next'

import { CopyButton } from '@/components/copy-button'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import {
  chatLinkRequiresApiKey,
  resolveChatUrl,
} from '@/features/chat/lib/chat-links'
import { buildCCSwitchImportUrl } from '@/features/keys/lib/cc-switch'

import type { ClientOption, QuickAccessContext } from '../lib/clients'
import { getClientIcon } from './client-icon'

function SnippetBlock(props: {
  title: string
  filePath?: string
  value: string
}) {
  return (
    <div className='border-border/60 overflow-hidden rounded-lg border'>
      <div className='bg-muted/40 flex items-center justify-between gap-2 px-3 py-1.5'>
        <span className='truncate text-xs font-medium'>
          {props.title}
          {props.filePath ? (
            <span className='text-muted-foreground font-normal'>
              {' · '}
              {props.filePath}
            </span>
          ) : null}
        </span>
        <CopyButton value={props.value} />
      </div>
      <pre className='overflow-x-auto px-3 py-2 text-xs'>
        <code>{props.value}</code>
      </pre>
    </div>
  )
}

function BuiltinDetail(props: {
  client: ClientOption & { kind: 'builtin' }
  ctx: QuickAccessContext
  hasApiKey: boolean
}) {
  const { t } = useTranslation()
  const client = props.client.builtin
  const Icon = getClientIcon(props.client)

  const quickImportUrl = client.quickImport
    ? buildCCSwitchImportUrl({
        app: client.quickImport.app,
        name: client.quickImport.name,
        serverAddress: props.ctx.serverAddress,
        apiKey: props.ctx.apiKey,
        models: props.ctx.model ? { model: props.ctx.model } : {},
      })
    : null

  return (
    <div className='space-y-5'>
      <div className='space-y-2'>
        <div className='flex flex-wrap items-center gap-x-3 gap-y-1'>
          <Icon className='size-5 shrink-0' aria-hidden='true' />
          <h2 className='text-lg font-semibold'>{client.name}</h2>
          <Button
            role='link'
            variant='link'
            size='sm'
            className='h-auto px-0'
            render={
              <a
                href={client.docsUrl}
                target='_blank'
                rel='noopener noreferrer'
              />
            }
          >
            {t('Documentation')}
            <ArrowUpRight className='size-3.5' aria-hidden='true' />
          </Button>
        </div>
        <p className='text-muted-foreground text-sm'>
          {t(client.descriptionKey)}
        </p>
      </div>

      {quickImportUrl && client.quickImport ? (
        <>
          <Separator />
          <div className='space-y-2'>
            <h3 className='text-sm font-semibold'>{t('Quick import')}</h3>
            <p className='text-muted-foreground text-xs'>
              {t(
                'Import this service into CC Switch with your active API key in one click.'
              )}
            </p>
            <div className='flex flex-wrap items-center gap-2'>
              <Button
                size='sm'
                disabled={!props.hasApiKey}
                onClick={() =>
                  window.open(quickImportUrl, '_blank', 'noopener')
                }
              >
                <ExternalLink className='size-4' aria-hidden='true' />
                {t('Open CC Switch')}
              </Button>
              <CopyButton
                value={quickImportUrl}
                variant='outline'
                size='sm'
                tooltip={t('Copy import link')}
              />
            </div>
            {!props.hasApiKey ? (
              <p className='text-destructive text-xs'>
                {t(
                  'No enabled API key found. Create or enable one before importing.'
                )}
              </p>
            ) : null}
          </div>
        </>
      ) : null}

      {client.installCommand ? (
        <>
          <Separator />
          <div className='space-y-2'>
            <h3 className='text-sm font-semibold'>{t('Install')}</h3>
            <SnippetBlock
              title={t('Install command')}
              value={client.installCommand}
            />
          </div>
        </>
      ) : null}

      <Separator />
      <div className='space-y-2'>
        <h3 className='text-sm font-semibold'>{t('Configuration')}</h3>
        <div className='space-y-3'>
          {client.snippets(props.ctx).map((snippet) => (
            <SnippetBlock
              key={`${snippet.labelKey}:${snippet.filePath ?? ''}`}
              title={t(snippet.labelKey)}
              filePath={snippet.filePath}
              value={snippet.value}
            />
          ))}
        </div>
      </div>
    </div>
  )
}

function PresetDetail(props: {
  client: ClientOption & { kind: 'preset' }
  ctx: QuickAccessContext
  hasApiKey: boolean
}) {
  const { t } = useTranslation()
  const preset = props.client.preset
  const Icon = getClientIcon(props.client)
  const needsKey = chatLinkRequiresApiKey(preset.url)
  const disabled = needsKey && !props.hasApiKey

  return (
    <div className='space-y-4'>
      <div className='flex flex-wrap items-center gap-x-3 gap-y-1'>
        <Icon className='size-5 shrink-0' aria-hidden='true' />
        <h2 className='text-lg font-semibold'>{preset.name}</h2>
      </div>
      <p className='text-muted-foreground text-sm'>
        {t(
          'This chat client was configured by your administrator. Open it below.'
        )}
      </p>

      {preset.type === 'web' ? (
        <Button
          size='sm'
          render={<Link to='/chat/$chatId' params={{ chatId: preset.id }} />}
        >
          <ExternalLink className='size-4' aria-hidden='true' />
          {t('Open in chat')}
        </Button>
      ) : (
        <div className='space-y-2'>
          <Button
            size='sm'
            disabled={disabled}
            onClick={() => {
              const url = resolveChatUrl({
                template: preset.url,
                apiKey: needsKey ? props.ctx.apiKey : undefined,
                serverAddress: props.ctx.serverAddress,
              })
              if (url) window.open(url, '_blank', 'noopener')
            }}
          >
            <ExternalLink className='size-4' aria-hidden='true' />
            {t('Open')}
          </Button>
          {disabled ? (
            <p className='text-destructive text-xs'>
              {t(
                'No enabled API key found. Create or enable one before opening this client.'
              )}
            </p>
          ) : null}
        </div>
      )}
    </div>
  )
}

export function ClientDetail(props: {
  client: ClientOption
  ctx: QuickAccessContext
  hasApiKey: boolean
}) {
  if (props.client.kind === 'builtin') {
    return (
      <BuiltinDetail
        client={props.client}
        ctx={props.ctx}
        hasApiKey={props.hasApiKey}
      />
    )
  }
  return (
    <PresetDetail
      client={props.client}
      ctx={props.ctx}
      hasApiKey={props.hasApiKey}
    />
  )
}
