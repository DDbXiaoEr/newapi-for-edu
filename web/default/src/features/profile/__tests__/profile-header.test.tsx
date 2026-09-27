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
import { render, screen } from '@testing-library/react'
import { createInstance } from 'i18next'
import { I18nextProvider, initReactI18next } from 'react-i18next'
import { expect, it } from 'vitest'

import { ProfileHeader } from '../components/profile-header'
import type { UserProfile } from '../types'

const i18n = createInstance()
await i18n.use(initReactI18next).init({
  lng: 'en',
  resources: { en: { translation: {} } },
})

const profile: UserProfile = {
  id: 1,
  username: 'alice',
  display_name: 'Alice',
  role: 1,
  email: 'alice@example.com',
  group: 'vip-hidden-group',
  quota: 1000000,
  used_quota: 0,
  request_count: 0,
  status: 1,
  aff_count: 0,
  aff_quota: 0,
  aff_history_quota: 0,
  created_time: 0,
}

it('does not show the user group or user id on the profile header', () => {
  render(
    <I18nextProvider i18n={i18n}>
      <ProfileHeader profile={profile} loading={false} />
    </I18nextProvider>
  )

  expect(screen.getByText('@alice')).toBeInTheDocument()
  expect(screen.getByText('alice@example.com')).toBeInTheDocument()
  expect(screen.queryByText('vip-hidden-group')).not.toBeInTheDocument()
  expect(screen.queryByText(/User ID/)).not.toBeInTheDocument()
  expect(screen.queryByText('1')).not.toBeInTheDocument()
})
