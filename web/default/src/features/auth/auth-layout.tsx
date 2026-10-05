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
import { useTranslation } from 'react-i18next'

import { Skeleton } from '@/components/ui/skeleton'
import { PROTOCOL_ROUTES } from '@/features/home/lib/protocol-routes'
import { useSystemConfig } from '@/hooks/use-system-config'

type AuthLayoutProps = {
  children: React.ReactNode
}

export function AuthLayout({ children }: AuthLayoutProps) {
  const { t } = useTranslation()
  const { systemName, logo, loading } = useSystemConfig()

  const brand = (
    <Link
      to='/'
      className='flex items-center gap-2 transition-opacity hover:opacity-80'
    >
      <div className='relative h-8 w-8'>
        {loading ? (
          <Skeleton className='absolute inset-0 rounded-full' />
        ) : (
          <img
            src={logo}
            alt={t('Logo')}
            className='h-8 w-8 rounded-full object-cover'
          />
        )}
      </div>
      {loading ? (
        <Skeleton className='h-6 w-24' />
      ) : (
        <span className='text-xl font-medium'>{systemName}</span>
      )}
    </Link>
  )

  return (
    <div className='relative grid min-h-svh lg:grid-cols-[minmax(0,1fr)_32rem]'>
      <aside className='border-border bg-muted/20 hidden flex-col justify-between border-r p-10 lg:flex'>
        {brand}
        <div>
          <p className='text-muted-foreground max-w-sm text-sm leading-relaxed'>
            {t('Compatible API routes for common AI application workflows')}
          </p>
          <ul className='mt-8 space-y-2.5 font-mono text-sm'>
            {PROTOCOL_ROUTES.map((route) => (
              <li key={route.id} className='flex items-baseline gap-3'>
                <span className='text-primary w-10 shrink-0'>{route.method}</span>
                <span className='text-foreground/80 truncate'>{route.path}</span>
              </li>
            ))}
          </ul>
        </div>
        <p className='text-muted-foreground text-xs'>{systemName}</p>
      </aside>
      <div className='flex flex-col'>
        <div className='px-4 pt-4 sm:px-8 sm:pt-8 lg:hidden'>{brand}</div>
        <div className='flex flex-1 items-center'>
          <div className='mx-auto flex w-full flex-col justify-center space-y-2 px-4 py-8 sm:w-[480px] sm:p-8'>
            {children}
          </div>
        </div>
      </div>
    </div>
  )
}
