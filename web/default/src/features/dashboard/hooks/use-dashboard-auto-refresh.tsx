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
/* eslint-disable react/only-export-components */
import {
  createContext,
  useCallback,
  useContext,
  useMemo,
  useState,
  type ReactNode,
} from 'react'

import {
  DASHBOARD_AUTO_REFRESH_INTERVAL_MS,
  DASHBOARD_AUTO_REFRESH_STORAGE_KEY,
} from '@/features/dashboard/constants'

interface DashboardAutoRefreshValue {
  enabled: boolean
  setEnabled: (enabled: boolean) => void
  refetchInterval: number | false
}

const DISABLED_AUTO_REFRESH: DashboardAutoRefreshValue = {
  enabled: false,
  setEnabled: () => undefined,
  refetchInterval: false,
}

const DashboardAutoRefreshContext =
  createContext<DashboardAutoRefreshValue | null>(null)

function getSavedAutoRefresh(): boolean {
  if (typeof window === 'undefined') return true
  const saved = window.localStorage.getItem(DASHBOARD_AUTO_REFRESH_STORAGE_KEY)
  if (saved === 'off') return false
  if (saved === 'on') return true
  return true
}

function saveAutoRefresh(enabled: boolean): void {
  if (typeof window === 'undefined') return
  window.localStorage.setItem(
    DASHBOARD_AUTO_REFRESH_STORAGE_KEY,
    enabled ? 'on' : 'off'
  )
}

export function DashboardAutoRefreshProvider(props: { children: ReactNode }) {
  const [enabled, setEnabledState] = useState(getSavedAutoRefresh)
  const setEnabled = useCallback((next: boolean) => {
    setEnabledState(next)
    saveAutoRefresh(next)
  }, [])
  const value = useMemo<DashboardAutoRefreshValue>(
    () => ({
      enabled,
      setEnabled,
      refetchInterval: enabled ? DASHBOARD_AUTO_REFRESH_INTERVAL_MS : false,
    }),
    [enabled, setEnabled]
  )

  return (
    <DashboardAutoRefreshContext.Provider value={value}>
      {props.children}
    </DashboardAutoRefreshContext.Provider>
  )
}

export function useDashboardAutoRefresh(): DashboardAutoRefreshValue {
  return useContext(DashboardAutoRefreshContext) ?? DISABLED_AUTO_REFRESH
}

export function useDashboardQueryRefresh() {
  const { refetchInterval } = useDashboardAutoRefresh()
  return {
    refetchInterval,
    refetchIntervalInBackground: false as const,
  }
}
