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
import { describe, expect, it, vi } from 'vitest'

import { PublicLayout } from '../components/public-layout'

vi.mock('../components/public-header', () => ({
  PublicHeader: () => <header>Header</header>,
}))

describe('public layout', () => {
  it('renders the page content', () => {
    render(
      <PublicLayout>
        <p>Public content</p>
      </PublicLayout>
    )

    expect(screen.getByText('Public content')).toBeInTheDocument()
  })

  it('offers the AGPLv3 corresponding source link on every public page', () => {
    render(
      <PublicLayout showMainContainer={false}>
        <p>Custom home</p>
      </PublicLayout>
    )

    expect(screen.getByRole('link', { name: '源代码' })).toHaveAttribute(
      'href',
      'https://github.com/DDbXiaoEr/newapi-for-edu'
    )
  })

  it('retains the required upstream attribution and link', () => {
    render(
      <PublicLayout showMainContainer={false}>
        <p>Custom home</p>
      </PublicLayout>
    )

    const upstream = screen.getByRole('link', {
      name: 'New API contributors',
    })
    expect(upstream).toHaveAttribute(
      'href',
      'https://github.com/QuantumNous/new-api'
    )
    expect(upstream.parentElement).toHaveTextContent(
      'Frontend design and development by New API contributors.'
    )
  })
})
