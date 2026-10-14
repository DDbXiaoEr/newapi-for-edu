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
import { render, screen, waitFor, within } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { createInstance, type i18n } from 'i18next'
import { useState, type ReactNode } from 'react'
import { I18nextProvider } from 'react-i18next'
import { beforeEach, describe, expect, it, vi } from 'vitest'

import { api } from '@/lib/api'

import { SettingsPageProvider } from '../../components/settings-page-context'
import { LdapSection } from '../ldap-section'

const defaults = {
  'ldap.enabled': true,
  'ldap.server_url': 'ldaps://ldap.example.com:636',
  'ldap.bind_dn': '',
  'ldap.bind_password': '',
  'ldap.base_dn': 'dc=example,dc=com',
  'ldap.user_filter': '',
  'ldap.username_attribute': 'uid',
  'ldap.display_name_attribute': 'cn',
  'ldap.mail_attribute': 'mail',
  'ldap.start_tls': false,
  'ldap.skip_tls_verify': false,
  'ldap.timeout_seconds': '5',
  'ldap.group_assignment_rules': '[]',
}

let testI18n: i18n

function Fixture(props: { children?: ReactNode }) {
  const [container, setContainer] = useState<HTMLDivElement | null>(null)
  const [client] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: { retry: false },
          mutations: { retry: false },
        },
      })
  )
  return (
    <I18nextProvider i18n={testI18n}>
      <QueryClientProvider client={client}>
        <div ref={setContainer} />
        <SettingsPageProvider actionsContainer={container}>
          {props.children ?? <LdapSection defaultValues={defaults} />}
        </SettingsPageProvider>
      </QueryClientProvider>
    </I18nextProvider>
  )
}

async function renderLdapSection() {
  const router = createRouter({
    routeTree: createRootRoute({
      component: () => <LdapSection defaultValues={defaults} />,
    }),
    history: createMemoryHistory({ initialEntries: ['/'] }),
  })
  await router.load()
  return render(
    <Fixture>
      <RouterProvider router={router} />
    </Fixture>
  )
}

beforeEach(async () => {
  testI18n = createInstance()
  await testI18n.init({
    lng: 'en',
    fallbackLng: 'en',
    resources: { en: { translation: {} } },
    interpolation: { escapeValue: false },
  })
})

describe('LDAP login test', () => {
  it('lists returned LDAP attribute names without values after a successful login test', async () => {
    const post = vi.spyOn(api, 'post').mockResolvedValue({
      data: {
        success: true,
        message: '',
        data: {
          attribute_names: ['dn', 'mail', 'memberOf', 'uid'],
        },
      },
    })
    const user = userEvent.setup()
    await renderLdapSection()

    await user.type(
      screen.getByRole('textbox', { name: 'Test username' }),
      'alice'
    )
    await user.type(screen.getByLabelText('Test password'), 'secret')
    await user.click(screen.getByRole('button', { name: 'Test login' }))

    await waitFor(() => {
      expect(post).toHaveBeenCalledWith('/api/option/ldap/test', {
        username: 'alice',
        password: 'secret',
      })
    })
    const result = await screen.findByRole('alert')
    expect(within(result).getByText('LDAP login succeeded')).toBeVisible()
    expect(within(result).getByText('Returned LDAP attributes')).toBeVisible()
    expect(within(result).getByText('dn')).toBeVisible()
    expect(within(result).getByText('mail')).toBeVisible()
    expect(within(result).getByText('memberOf')).toBeVisible()
    expect(within(result).getByText('uid')).toBeVisible()
    expect(screen.queryByText('alice@example.com')).not.toBeInTheDocument()
  })

  it('shows the server error when LDAP login test fails', async () => {
    vi.spyOn(api, 'post').mockResolvedValue({
      data: {
        success: false,
        message: 'LDAP username or password is incorrect',
      },
    })
    const user = userEvent.setup()
    await renderLdapSection()

    await user.type(
      screen.getByRole('textbox', { name: 'Test username' }),
      'alice'
    )
    await user.type(screen.getByLabelText('Test password'), 'wrong')
    await user.click(screen.getByRole('button', { name: 'Test login' }))

    expect(await screen.findByText('LDAP login test failed')).toBeVisible()
    expect(
      screen.getByText('LDAP username or password is incorrect')
    ).toBeVisible()
  })
})

