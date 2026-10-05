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
import { useTranslation } from 'react-i18next'

import { AnimateInView } from '@/components/animate-in-view'

import { PROTOCOL_ROUTES } from '../../lib/protocol-routes'

export function Features() {
  const { t } = useTranslation()

  const features = [
    {
      id: 'fast',
      title: t('Lightning Fast'),
      desc: t(
        'Optimized network architecture ensures millisecond response times'
      ),
    },
    {
      id: 'secure',
      title: t('Secure & Reliable'),
      desc: t(
        'Enterprise-grade security with comprehensive permission management'
      ),
    },
    {
      id: 'global',
      title: t('Global Coverage'),
      desc: t('Multi-region deployment for stable global access'),
    },
    {
      id: 'developer',
      title: t('Developer Friendly'),
      desc: t('Compatible API routes for common AI application workflows'),
    },
  ]

  const additionalFeatures = [
    {
      title: t('High Performance'),
      desc: t('Support for high concurrency with automatic load balancing'),
    },
    {
      title: t('Transparent Billing'),
      desc: t('Pay-as-you-go with real-time usage monitoring'),
    },
    {
      title: t('Team Collaboration'),
      desc: t('Multi-user management with flexible permission allocation'),
    },
    {
      title: t('Open Source'),
      desc: t('Community driven, self-hosted, and extensible'),
    },
  ]

  return (
    <section className='relative z-10 px-6 py-20 md:py-28'>
      <div className='mx-auto max-w-6xl'>
        <AnimateInView className='mb-12 max-w-xl md:mb-16'>
          <h2 className='text-2xl leading-tight font-semibold tracking-tight md:text-3xl'>
            {t('Built for developers,')}
            <br />
            {t('designed for scale')}
          </h2>
        </AnimateInView>

        <div className='border-border grid gap-px overflow-hidden rounded-lg border md:grid-cols-2'>
          {features.map((f, i) => (
            <AnimateInView
              key={f.id}
              delay={i * 80}
              animation='fade-in'
              className='bg-background p-6 md:p-8'
            >
              <h3 className='text-base font-semibold'>{f.title}</h3>
              <p className='text-muted-foreground mt-2 text-sm leading-relaxed'>
                {f.desc}
              </p>
            </AnimateInView>
          ))}
        </div>

        <div className='border-border mt-10 overflow-hidden rounded-lg border'>
          <div className='border-border bg-muted/20 border-b px-6 py-3'>
            <p className='text-muted-foreground font-mono text-xs'>
              {t('Compatible API routes for common AI application workflows')}
            </p>
          </div>
          <ul className='divide-border divide-y'>
            {PROTOCOL_ROUTES.map((route) => (
              <li
                key={route.id}
                className='flex items-baseline gap-4 px-6 py-3 font-mono text-sm'
              >
                <span className='text-primary w-10 shrink-0'>{route.method}</span>
                <span className='text-foreground min-w-16 shrink-0'>
                  {route.label}
                </span>
                <span className='text-muted-foreground truncate'>
                  {route.path}
                </span>
              </li>
            ))}
          </ul>
        </div>

        <div className='mt-12 grid grid-cols-1 gap-8 sm:grid-cols-2 md:grid-cols-4 md:gap-10'>
          {additionalFeatures.map((f, i) => (
            <AnimateInView
              key={f.title}
              delay={i * 80}
              animation='fade-in'
              className='flex flex-col'
            >
              <h3 className='text-sm font-semibold'>{f.title}</h3>
              <p className='text-muted-foreground mt-1.5 text-sm leading-relaxed'>
                {f.desc}
              </p>
            </AnimateInView>
          ))}
        </div>
      </div>
    </section>
  )
}
