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
import { useState } from 'react'
import { useTranslation } from 'react-i18next'

import { SectionPageLayout } from '@/components/layout'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { SettingsPageProvider } from '@/features/system-settings/components/settings-page-context'
import { ROLE } from '@/lib/roles'
import { useAuthStore } from '@/stores/auth-store'

import { GroupSettingsPanel } from './components/group-settings-panel'
import { GroupUsersPanel } from './components/group-users-panel'

export function Groups() {
  const { t } = useTranslation()
  const role = useAuthStore((state) => state.auth.user?.role ?? 0)
  const canEditSettings = role >= ROLE.SUPER_ADMIN
  const [actionsContainer, setActionsContainer] =
    useState<HTMLDivElement | null>(null)
  const [tab, setTab] = useState(canEditSettings ? 'settings' : 'members')

  return (
    <SettingsPageProvider actionsContainer={actionsContainer}>
      <SectionPageLayout fixedContent>
        <SectionPageLayout.Title>{t('Groups')}</SectionPageLayout.Title>
        <SectionPageLayout.Actions>
          <div
            ref={setActionsContainer}
            className='flex flex-wrap items-center justify-end gap-2'
          />
        </SectionPageLayout.Actions>
        <SectionPageLayout.Content>
          {canEditSettings ? (
            <Tabs
              value={tab}
              onValueChange={setTab}
              className='flex h-full min-h-0 flex-col gap-4'
            >
              <TabsList className='w-fit'>
                <TabsTrigger value='settings'>
                  {t('Group settings')}
                </TabsTrigger>
                <TabsTrigger value='members'>{t('Group members')}</TabsTrigger>
              </TabsList>
              {tab === 'settings' ? (
                <TabsContent value='settings' className='min-h-0 overflow-auto'>
                  <GroupSettingsPanel />
                </TabsContent>
              ) : (
                <TabsContent
                  value='members'
                  className='flex min-h-0 flex-1 flex-col'
                >
                  <GroupUsersPanel />
                </TabsContent>
              )}
            </Tabs>
          ) : (
            <GroupUsersPanel />
          )}
        </SectionPageLayout.Content>
      </SectionPageLayout>
    </SettingsPageProvider>
  )
}
