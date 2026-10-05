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
import { ArrowRight } from 'lucide-react'
import { useTranslation } from 'react-i18next'

import { AnimateInView } from '@/components/animate-in-view'
import { Button } from '@/components/ui/button'

interface CTAProps {
  isAuthenticated?: boolean
}

export function CTA(props: CTAProps) {
  const { t } = useTranslation()

  if (props.isAuthenticated) {
    return null
  }

  return (
    <section className='border-border relative z-10 border-t px-6 py-20 md:py-24'>
      <AnimateInView
        className='mx-auto flex max-w-6xl flex-col gap-6 md:flex-row md:items-end md:justify-between'
        animation='fade-in'
      >
        <div className='max-w-xl'>
          <h2 className='text-2xl leading-tight font-semibold tracking-tight md:text-3xl'>
            {t('Ready to simplify')} {t('your AI integration?')}
          </h2>
          <p className='text-muted-foreground mt-3 max-w-md text-sm leading-relaxed md:text-base'>
            {t(
              'Deploy your own gateway and start routing requests through your configured upstream services.'
            )}
          </p>
        </div>
        <div className='flex flex-wrap items-center gap-3'>
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
        </div>
      </AnimateInView>
    </section>
  )
}
