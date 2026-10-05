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
import { cleanup, render, screen, waitFor } from '@testing-library/react'
import i18next from 'i18next'
import { afterEach, describe, expect, it } from 'vitest'

import { toIntlLocale } from '@/i18n/languages'
import { formatNumber } from '@/lib/format'

import { Stats } from '../stats'

const INTERFACE_LANGUAGES = [
  'zhCN',
  'zhTW',
  'en',
  'fr',
  'ru',
  'ja',
  'vi',
] as const

afterEach(async () => {
  cleanup()
  await i18next.changeLanguage('en')
})

async function expectFormattedStat(language: string, value: number) {
  await i18next.changeLanguage(language)
  render(<Stats />)
  const locale = toIntlLocale(i18next.resolvedLanguage || i18next.language)
  const expected = `${formatNumber(value, locale)}+`
  await waitFor(() => {
    expect(screen.getAllByText(expected).length).toBeGreaterThan(0)
  })
}

describe('home stats locale formatting', () => {
  it.each(INTERFACE_LANGUAGES)(
    'formats counters for interface language %s',
    async (language) => {
      await expectFormattedStat(language, 50)
    }
  )

  it('falls back when the interface language is not a valid Intl locale', async () => {
    await expectFormattedStat('not-a-locale', 50)
  })

  it('updates formatted counters after the interface language changes', async () => {
    await i18next.changeLanguage('en')
    const { rerender } = render(<Stats />)
    await waitFor(() => {
      expect(
        screen.getAllByText(`${formatNumber(100, toIntlLocale('en'))}+`).length
      ).toBeGreaterThan(0)
    })

    await i18next.changeLanguage('zhCN')
    rerender(<Stats />)
    await waitFor(() => {
      expect(
        screen.getAllByText(
          `${formatNumber(100, toIntlLocale('zhCN'))}+`
        ).length
      ).toBeGreaterThan(0)
    })
  })
})
