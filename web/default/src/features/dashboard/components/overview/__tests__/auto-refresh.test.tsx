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
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import {
  createMemoryHistory,
  createRootRoute,
  createRouter,
  RouterProvider,
} from '@tanstack/react-router'
import { act, cleanup, render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

import { DASHBOARD_AUTO_REFRESH_STORAGE_KEY } from '@/features/dashboard/constants'
import { api } from '@/lib/api'
import { useAuthStore } from '@/stores/auth-store'
import { useSystemConfigStore } from '@/stores/system-config-store'

import { DashboardAutoRefreshProvider } from '../../../hooks/use-dashboard-auto-refresh'
import { OverviewDashboard } from '../overview-dashboard'

let client: QueryClient
const usageRequests: Array<{ start?: number; end?: number }> = []

function usageCallCount() {
  return usageRequests.length
}

beforeEach(() => {
  window.localStorage.clear()
  usageRequests.length = 0
  useSystemConfigStore.setState(useSystemConfigStore.getInitialState(), true)
  useAuthStore.getState().auth.setUser({
    id: 1,
    username: 'dashboard-user',
    role: 1,
    quota: 1000000,
    used_quota: 1000,
    request_count: 1,
  })
  client = new QueryClient({
    defaultOptions: { queries: { retry: false } },
  })
  vi.spyOn(api, 'get').mockImplementation(async (url, config) => {
    switch (url) {
      case '/api/token/?p=1&size=10':
        return {
          data: {
            success: true,
            data: {
              items: [{ id: 1, name: 'App key', key: 'masked', status: 1 }],
            },
          },
        }
      case '/api/status':
        return {
          data: {
            data: {
              api_info_enabled: false,
              announcements_enabled: false,
              faq_enabled: false,
              uptime_kuma_enabled: false,
            },
          },
        }
      case '/api/user/models':
        return { data: { success: true, data: ['gpt-4o-mini'] } }
      case '/api/user/self':
        return {
          data: {
            success: true,
            data: useAuthStore.getState().auth.user,
          },
        }
      case '/api/data/self': {
        const params = (
          config as { params?: { start_timestamp?: number; end_timestamp?: number } }
        )?.params
        usageRequests.push({
          start: params?.start_timestamp,
          end: params?.end_timestamp,
        })
        return { data: { success: true, data: [] } }
      }
      default:
        throw new Error(`Unexpected dashboard request: ${url}`)
    }
  })
})

afterEach(() => {
  cleanup()
  client.clear()
  useAuthStore.setState(useAuthStore.getInitialState(), true)
  useSystemConfigStore.setState(useSystemConfigStore.getInitialState(), true)
  window.localStorage.clear()
  vi.useRealTimers()
  vi.restoreAllMocks()
})

async function renderOverview() {
  const router = createRouter({
    routeTree: createRootRoute({ component: OverviewDashboard }),
    history: createMemoryHistory({ initialEntries: ['/'] }),
  })
  await router.load()
  return render(
    <QueryClientProvider client={client}>
      <DashboardAutoRefreshProvider>
        <RouterProvider router={router} />
      </DashboardAutoRefreshProvider>
    </QueryClientProvider>
  )
}

describe('overview auto refresh', () => {
  it('keeps auto refresh on by default and remembers turning it off', async () => {
    const user = userEvent.setup()
    const first = await renderOverview()
    const toggle = await screen.findByRole('switch', {
      name: 'Auto refresh (30s)',
    })
    expect(toggle).toBeChecked()

    await user.click(toggle)
    expect(toggle).not.toBeChecked()
    expect(window.localStorage.getItem(DASHBOARD_AUTO_REFRESH_STORAGE_KEY)).toBe(
      'off'
    )
    first.unmount()

    await renderOverview()
    expect(
      await screen.findByRole('switch', { name: 'Auto refresh (30s)' })
    ).not.toBeChecked()
  })

  it('polls usage while enabled and stops after the switch is turned off', async () => {
    vi.useFakeTimers({ toFake: ['setInterval', 'clearInterval'] })
    const user = userEvent.setup({ advanceTimers: vi.advanceTimersByTime })
    const view = await renderOverview()
    await screen.findByRole('switch', { name: 'Auto refresh (30s)' })
    await act(async () => {
      await vi.waitFor(() => expect(usageCallCount()).toBe(1))
    })
    await act(async () => {
      await vi.advanceTimersByTimeAsync(30_000)
    })
    await act(async () => {
      await vi.waitFor(() => expect(usageCallCount()).toBe(2))
    })

    await user.click(
      screen.getByRole('switch', { name: 'Auto refresh (30s)' })
    )
    const pausedCount = usageCallCount()
    await act(async () => {
      await vi.advanceTimersByTimeAsync(60_000)
    })
    expect(usageCallCount()).toBe(pausedCount)
    view.unmount()
  })

  it('slides the rolling usage window forward on each refresh', async () => {
    vi.useFakeTimers({ toFake: ['setInterval', 'clearInterval', 'Date'] })
    vi.setSystemTime(new Date('2026-10-02T12:00:00.000Z'))
    await renderOverview()
    await act(async () => {
      await vi.waitFor(() => expect(usageCallCount()).toBe(1))
    })
    const firstRange = usageRequests[0]

    vi.setSystemTime(new Date('2026-10-02T12:10:00.000Z'))
    await act(async () => {
      await vi.advanceTimersByTimeAsync(30_000)
    })
    await act(async () => {
      await vi.waitFor(() => expect(usageCallCount()).toBe(2))
    })
    expect(usageRequests[1]?.end).toBeGreaterThan(firstRange?.end ?? 0)
    expect((usageRequests[1]?.end ?? 0) - (usageRequests[1]?.start ?? 0)).toBe(
      (firstRange?.end ?? 0) - (firstRange?.start ?? 0)
    )
  })
})
