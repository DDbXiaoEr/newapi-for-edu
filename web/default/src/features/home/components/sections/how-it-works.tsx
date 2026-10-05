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

export function HowItWorks() {
  const { t } = useTranslation()

  const steps = [
    {
      num: '1',
      title: t('Configure'),
      desc: t(
        'Add your API keys, set up channels and configure access permissions'
      ),
    },
    {
      num: '2',
      title: t('Connect'),
      desc: t(
        'Connect through OpenAI, Claude, Gemini, and other compatible API routes'
      ),
    },
    {
      num: '3',
      title: t('Monitor'),
      desc: t('Track usage, costs and performance with real-time analytics'),
    },
  ]

  return (
    <section className='border-border relative z-10 border-t px-6 py-20 md:py-28'>
      <div className='mx-auto max-w-6xl'>
        <AnimateInView className='mb-12 md:mb-16'>
          <h2 className='text-2xl font-semibold tracking-tight md:text-3xl'>
            {t('Three steps to get started')}
          </h2>
        </AnimateInView>

        <ol className='grid gap-0 md:grid-cols-3'>
          {steps.map((step, i) => (
            <AnimateInView
              key={step.num}
              as='li'
              delay={i * 80}
              animation='fade-in'
              className='border-border relative border-l py-1 pl-6 md:border-t md:border-l-0 md:pt-6 md:pb-0 md:pl-0 md:pr-10'
            >
              <span className='text-muted-foreground font-mono text-sm tabular-nums'>
                {step.num}
              </span>
              <h3 className='mt-3 text-base font-semibold'>{step.title}</h3>
              <p className='text-muted-foreground mt-2 max-w-sm text-sm leading-relaxed'>
                {step.desc}
              </p>
            </AnimateInView>
          ))}
        </ol>
      </div>
    </section>
  )
}
