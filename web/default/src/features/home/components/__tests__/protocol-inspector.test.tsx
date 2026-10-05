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
import { cleanup, render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { afterEach, describe, expect, it } from 'vitest'

import { PROTOCOL_ROUTES } from '../../lib/protocol-routes'
import { HeroTerminalDemo } from '../hero-terminal-demo'

afterEach(() => {
  cleanup()
})

describe('protocol inspector', () => {
  it('shows the first compatible route and updates the endpoint when another tab is selected', async () => {
    const user = userEvent.setup()
    render(<HeroTerminalDemo />)

    const tabs = screen.getAllByRole('tab')
    expect(tabs.map((tab) => tab.textContent)).toEqual(
      PROTOCOL_ROUTES.map((route) => route.label)
    )
    expect(tabs[0]).toHaveAttribute('aria-selected', 'true')
    expect(screen.getByTestId('protocol-endpoint')).toHaveTextContent(
      PROTOCOL_ROUTES[0].path
    )

    await user.click(screen.getByRole('tab', { name: PROTOCOL_ROUTES[2].label }))
    await waitFor(() => {
      expect(screen.getByRole('tab', { name: PROTOCOL_ROUTES[2].label })).toHaveAttribute(
        'aria-selected',
        'true'
      )
      expect(screen.getByTestId('protocol-endpoint')).toHaveTextContent(
        PROTOCOL_ROUTES[2].path
      )
    })
  })
})
