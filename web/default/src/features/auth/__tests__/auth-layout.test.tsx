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
import { cleanup, render, screen } from '@testing-library/react'
import type { ReactNode } from 'react'
import { afterEach, describe, expect, it, vi } from 'vitest'

import { PROTOCOL_ROUTES } from '@/features/home/lib/protocol-routes'

import { AuthLayout } from '../auth-layout'

vi.mock('@tanstack/react-router', () => ({
  Link: (props: { to: string; children?: ReactNode }) => (
    <a href={props.to}>{props.children}</a>
  ),
}))

vi.mock('@/hooks/use-system-config', () => ({
  useSystemConfig: () => ({
    systemName: 'New API',
    logo: '/logo.png',
    loading: false,
  }),
}))

afterEach(() => {
  cleanup()
})

describe('auth layout', () => {
  it('keeps the brand link home and lists compatible protocol routes beside the form', () => {
    render(
      <AuthLayout>
        <h2>Sign in</h2>
      </AuthLayout>
    )

    const brandLinks = screen.getAllByRole('link', { name: 'Logo New API' })
    expect(brandLinks.length).toBeGreaterThan(0)
    for (const link of brandLinks) {
      expect(link).toHaveAttribute('href', '/')
    }
    expect(screen.getByRole('heading', { name: 'Sign in' })).toBeInTheDocument()
    for (const route of PROTOCOL_ROUTES) {
      expect(screen.getByText(route.path)).toBeInTheDocument()
    }
  })
})
