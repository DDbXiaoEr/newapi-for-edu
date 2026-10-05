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

import { PROTOCOL_ROUTES } from '../../../lib/protocol-routes'
import { CTA } from '../cta'
import { Features } from '../features'
import { Hero } from '../hero'
import { HowItWorks } from '../how-it-works'

vi.mock('@tanstack/react-router', () => ({
  Link: (props: { to: string; children?: ReactNode }) => (
    <a href={props.to}>{props.children}</a>
  ),
}))

vi.mock('@/hooks/use-status', () => ({
  useStatus: () => ({ status: { docs_link: '/docs' }, loading: false }),
}))

vi.mock('@lobehub/icons', () => ({
  CherryStudio: { Color: () => <span /> },
}))

afterEach(() => {
  cleanup()
})

describe('home landing layout', () => {
  it('lists compatible protocol routes in the hero instead of a gradient badge', () => {
    render(<Hero />)

    expect(
      screen.getByRole('heading', { name: /Unified API Gateway for/ })
    ).toBeInTheDocument()
    for (const route of PROTOCOL_ROUTES) {
      expect(screen.getAllByText(route.path).length).toBeGreaterThan(0)
    }
    expect(screen.queryByText('01')).not.toBeInTheDocument()
    expect(
      screen.getByRole('link', { name: 'Get Started' })
    ).toHaveAttribute('href', '/sign-up')
    expect(
      screen.getByRole('link', { name: 'View Pricing' })
    ).toHaveAttribute('href', '/pricing')
  })

  it('sends authenticated visitors to the dashboard', () => {
    render(<Hero isAuthenticated />)

    expect(
      screen.getByRole('link', { name: 'Go to Dashboard' })
    ).toHaveAttribute('href', '/dashboard')
    expect(
      screen.queryByRole('link', { name: 'Get Started' })
    ).not.toBeInTheDocument()
  })

  it('renders features as a protocol table without numbered sequence markers', () => {
    render(<Features />)

    expect(screen.getByText('Lightning Fast')).toBeInTheDocument()
    expect(screen.getByText('Developer Friendly')).toBeInTheDocument()
    for (const route of PROTOCOL_ROUTES) {
      expect(screen.getByText(route.path)).toBeInTheDocument()
    }
    expect(screen.queryByText('01')).not.toBeInTheDocument()
    expect(screen.queryByText('02')).not.toBeInTheDocument()
  })

  it('keeps how-it-works as an ordered sequence of three steps', () => {
    render(<HowItWorks />)

    const items = screen.getAllByRole('listitem')
    expect(items).toHaveLength(3)
    expect(items[0]).toHaveTextContent('Configure')
    expect(items[1]).toHaveTextContent('Connect')
    expect(items[2]).toHaveTextContent('Monitor')
  })

  it('hides the signup call to action after sign-in', () => {
    const { rerender } = render(<CTA />)
    expect(screen.getByRole('link', { name: 'Get Started' })).toBeInTheDocument()
    rerender(<CTA isAuthenticated />)
    expect(
      screen.queryByRole('link', { name: 'Get Started' })
    ).not.toBeInTheDocument()
  })
})
