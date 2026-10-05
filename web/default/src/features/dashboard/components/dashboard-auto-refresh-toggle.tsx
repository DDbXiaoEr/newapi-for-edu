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
import { useId } from 'react'
import { useTranslation } from 'react-i18next'

import { Label } from '@/components/ui/label'
import { Switch } from '@/components/ui/switch'
import { useDashboardAutoRefresh } from '@/features/dashboard/hooks/use-dashboard-auto-refresh'

export function DashboardAutoRefreshToggle() {
  const { t } = useTranslation()
  const autoRefreshId = useId()
  const { enabled, setEnabled } = useDashboardAutoRefresh()

  return (
    <div className='flex items-center gap-2'>
      <Switch
        id={autoRefreshId}
        size='sm'
        checked={enabled}
        onCheckedChange={(checked) => setEnabled(checked === true)}
      />
      <Label
        htmlFor={autoRefreshId}
        className='text-muted-foreground text-xs font-medium'
      >
        {t('Auto refresh (30s)')}
      </Label>
    </div>
  )
}
