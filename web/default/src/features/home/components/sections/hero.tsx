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
import { CherryStudio } from '@lobehub/icons'
import { Link } from '@tanstack/react-router'
import { ArrowRight, BookOpen } from 'lucide-react'
import { useTranslation } from 'react-i18next'

import { Button } from '@/components/ui/button'
import { useStatus } from '@/hooks/use-status'

import { PROTOCOL_ROUTES } from '../../lib/protocol-routes'
import { HeroTerminalDemo } from '../hero-terminal-demo'

interface HeroProps {
  isAuthenticated?: boolean
}

export function Hero(props: HeroProps) {
  const { t } = useTranslation()
  const { status } = useStatus()
  const docsUrl =
    (status?.docs_link as string | undefined) || 'https://docs.newapi.pro'

  const renderDocsButton = () => {
    const isExternal = docsUrl.startsWith('http')
    if (isExternal) {
      return (
        <Button
          variant='outline'
          size='lg'
          render={
            <a href={docsUrl} target='_blank' rel='noopener noreferrer' />
          }
        >
          <BookOpen data-icon='inline-start' />
          {t('Docs')}
        </Button>
      )
    }
    return (
      <Button variant='outline' size='lg' render={<Link to={docsUrl} />}>
        <BookOpen data-icon='inline-start' />
        {t('Docs')}
      </Button>
    )
  }

  return (
    <section className='relative z-10 px-6 pt-24 pb-16 md:pt-32 md:pb-20 lg:pt-36'>
      <div className='mx-auto grid max-w-6xl grid-cols-1 items-start gap-12 lg:grid-cols-12 lg:gap-16'>
        <div className='flex flex-col items-start text-left lg:col-span-5'>
          <p className='text-muted-foreground text-sm'>
            {t('AI Application Infrastructure Foundation')}
          </p>

          <h1 className='mt-4 text-[clamp(2.25rem,4vw,3.5rem)] leading-[1.12] font-semibold tracking-tight'>
            {t('Unified API Gateway for')}
            <br />
            {t('Vast Range of AI Models')}
          </h1>
          <p className='text-muted-foreground mt-5 max-w-md text-base leading-relaxed'>
            {t(
              'Access a vast selection of models via a standard, unified API protocol. Power AI applications, manage digital assets, and connect the Future.'
            )}
          </p>

          <ul className='mt-6 w-full max-w-md space-y-1.5 font-mono text-[13px]'>
            {PROTOCOL_ROUTES.map((route) => (
              <li key={route.id} className='flex items-baseline gap-3'>
                <span className='text-primary w-10 shrink-0'>{route.method}</span>
                <span className='text-foreground/80 truncate'>{route.path}</span>
              </li>
            ))}
          </ul>

          <div className='mt-8 flex flex-wrap items-center gap-3'>
            {props.isAuthenticated ? (
              <>
                <Button size='lg' render={<Link to='/dashboard' />}>
                  {t('Go to Dashboard')}
                  <ArrowRight data-icon='inline-end' />
                </Button>
                {renderDocsButton()}
              </>
            ) : (
              <>
                <Button size='lg' render={<Link to='/sign-up' />}>
                  {t('Get Started')}
                  <ArrowRight data-icon='inline-end' />
                </Button>
                <Button
                  variant='outline'
                  size='lg'
                  render={<Link to='/pricing' />}
                >
                  {t('View Pricing')}
                </Button>
                {renderDocsButton()}
              </>
            )}
          </div>

          <div className='mt-10 w-full max-w-md'>
            <p className='text-muted-foreground text-sm'>
              {t('Supported Applications')}
            </p>
            <p className='text-muted-foreground/80 mt-1 text-xs leading-relaxed'>
              {t(
                'Supports one-click configuration and perfectly adapts to NewAPI multi-protocol configuration.'
              )}
            </p>
            <div className='mt-3 flex flex-wrap items-center gap-2'>
              <a
                href='https://cherry-ai.com'
                target='_blank'
                rel='noopener noreferrer'
                className='border-border bg-background text-foreground hover:bg-muted inline-flex items-center gap-2 rounded-md border px-3 py-1.5 text-sm transition-colors'
              >
                <CherryStudio.Color size={18} className='shrink-0' />
                <span>Cherry Studio</span>
              </a>
              <a
                href='https://ccswitch.io'
                target='_blank'
                rel='noopener noreferrer'
                className='border-border bg-background text-foreground hover:bg-muted inline-flex items-center gap-2 rounded-md border px-3 py-1.5 text-sm transition-colors'
              >
                <img
                  src='https://ccswitch.io/favicon.png'
                  alt=''
                  className='size-4 shrink-0 rounded-sm object-contain'
                  onError={(e) => {
                    e.currentTarget.style.display = 'none'
                    const fallback = e.currentTarget.nextSibling as HTMLElement
                    if (fallback) fallback.style.display = 'flex'
                  }}
                />
                <span
                  style={{ display: 'none' }}
                  className='bg-muted text-muted-foreground size-4 shrink-0 items-center justify-center rounded-sm text-[9px] font-medium'
                >
                  CC
                </span>
                <span>CC Switch</span>
              </a>
              <span className='text-muted-foreground px-2 text-sm'>
                {t('More Apps')}
              </span>
            </div>
          </div>
        </div>

        <div className='w-full lg:col-span-7'>
          <HeroTerminalDemo />
        </div>
      </div>
    </section>
  )
}
