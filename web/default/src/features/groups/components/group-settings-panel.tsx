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
import { ErrorState } from '@/components/error-state'
import { LoadingState } from '@/components/loading-state'
import {
  getOptionValue,
  useSystemOptions,
} from '@/features/system-settings/hooks/use-system-options'
import {
  RatioSettingsCard,
  type GroupFormValues,
} from '@/features/system-settings/models/ratio-settings-card'

const emptyModelDefaults = {
  ModelPrice: '',
  ModelRatio: '',
  CacheRatio: '',
  CreateCacheRatio: '',
  CompletionRatio: '',
  ImageRatio: '',
  AudioRatio: '',
  AudioCompletionRatio: '',
  ExposeRatioEnabled: false,
  BillingMode: '{}',
  BillingExpr: '{}',
  PluginBillingExpr: '{}',
}

const defaultGroupSettings: GroupFormValues = {
  GroupRatio: '',
  TopupGroupRatio: '',
  UserUsableGroups: '',
  GroupGroupRatio: '',
  AutoGroups: '',
  MaxTokenAutoGroups: 5,
  DefaultUseAutoGroup: false,
  GroupSpecialUsableGroup: '{}',
}

export function GroupSettingsPanel() {
  const query = useSystemOptions()
  const settings = getOptionValue(query.data?.data, {
    GroupRatio: '',
    TopupGroupRatio: '',
    UserUsableGroups: '',
    GroupGroupRatio: '',
    AutoGroups: '',
    MaxTokenAutoGroups: 5,
    DefaultUseAutoGroup: false,
    'group_ratio_setting.group_special_usable_group': '{}',
  })

  if (query.isLoading) return <LoadingState />
  if (query.isError && !query.data) {
    return <ErrorState onRetry={() => void query.refetch()} />
  }

  return (
    <RatioSettingsCard
      titleKey='Group Pricing'
      modelDefaults={emptyModelDefaults}
      groupDefaults={{
        ...defaultGroupSettings,
        GroupRatio: settings.GroupRatio,
        TopupGroupRatio: settings.TopupGroupRatio,
        UserUsableGroups: settings.UserUsableGroups,
        GroupGroupRatio: settings.GroupGroupRatio,
        AutoGroups: settings.AutoGroups,
        MaxTokenAutoGroups: settings.MaxTokenAutoGroups,
        DefaultUseAutoGroup: settings.DefaultUseAutoGroup,
        GroupSpecialUsableGroup:
          settings['group_ratio_setting.group_special_usable_group'],
      }}
      toolPricesDefault='{}'
      visibleTabs={['groups']}
    />
  )
}
