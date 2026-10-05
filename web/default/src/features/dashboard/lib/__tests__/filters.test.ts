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
import { afterEach, describe, expect, test, vi } from 'vitest'

import { computeTimeRange, getRollingDateRange } from '@/lib/time'

import type { DashboardFilters } from '../../types'
import {
  getDashboardRangeDays,
  getDashboardTimeQueryKey,
  resolveDashboardTimeRange,
} from '../filters'

function granularityForDays(days: number): DashboardFilters['time_granularity'] {
  if (days <= 1) return 'hour'
  if (days >= 29) return 'week'
  return 'day'
}

function rollingFilters(days: number, now = new Date()): DashboardFilters {
  const { start, end } = getRollingDateRange(days, now)
  return {
    start_timestamp: start,
    end_timestamp: end,
    time_granularity: granularityForDays(days),
  }
}

describe('dashboard time-range helpers', () => {
  afterEach(() => {
    vi.useRealTimers()
  })

  test('returns the matching preset days for a rolling range', () => {
    expect(getDashboardRangeDays(rollingFilters(1))).toBe(1)
    expect(getDashboardRangeDays(rollingFilters(7))).toBe(7)
    expect(getDashboardRangeDays(rollingFilters(14))).toBe(14)
    expect(getDashboardRangeDays(rollingFilters(29))).toBe(29)
  })

  test('returns null for a custom range that is not a preset', () => {
    const start = new Date('2026-01-01T00:00:00.000Z')
    const end = new Date('2026-01-03T00:00:00.000Z')
    expect(
      getDashboardRangeDays({
        start_timestamp: start,
        end_timestamp: end,
        time_granularity: 'day',
      })
    ).toBeNull()
  })

  test('resolves a rolling preset against the current time', () => {
    vi.useFakeTimers()
    vi.setSystemTime(new Date('2026-10-02T12:00:00.000Z'))
    const filters = rollingFilters(1, new Date('2026-10-01T12:00:00.000Z'))
    expect(resolveDashboardTimeRange(filters)).toEqual(computeTimeRange(1))
  })

  test('keeps a custom range fixed in the query key', () => {
    const start = new Date('2026-01-01T08:00:00.000Z')
    const end = new Date('2026-01-03T08:00:00.000Z')
    const filters: DashboardFilters = {
      start_timestamp: start,
      end_timestamp: end,
      time_granularity: 'hour',
      username: 'alice',
    }
    expect(getDashboardTimeQueryKey(filters)).toEqual({
      start: start.getTime(),
      end: end.getTime(),
    })
    expect(resolveDashboardTimeRange(filters)).toEqual(
      computeTimeRange(2, start, end)
    )
  })

  test('uses rolling days in the query key for presets', () => {
    expect(getDashboardTimeQueryKey(rollingFilters(7))).toEqual({
      rollingDays: 7,
    })
  })
})
