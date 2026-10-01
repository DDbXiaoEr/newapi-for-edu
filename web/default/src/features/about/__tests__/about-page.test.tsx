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
import { cleanup, render, screen } from '@testing-library/react'
import type { ReactNode } from 'react'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'

import { getAboutContent } from '../api'
import { About } from '../index'

vi.mock('@/components/layout', () => ({
  PublicLayout: (props: { children: ReactNode }) => <div>{props.children}</div>,
}))

vi.mock('../api', () => ({
  getAboutContent: vi.fn(),
}))

vi.mock('@/components/rich-content', () => ({
  RichContent: (props: { mode?: string; content: string }) => (
    <div data-testid='rich-content' data-mode={props.mode}>
      {props.content}
    </div>
  ),
}))

let client: QueryClient

beforeEach(() => {
  client = new QueryClient({
    defaultOptions: { queries: { retry: false } },
  })
})

afterEach(() => {
  cleanup()
  client.clear()
})

function showAbout() {
  return render(
    <QueryClientProvider client={client}>
      <About />
    </QueryClientProvider>
  )
}

describe('about page', () => {
  it('keeps an admin-configured URL iframe same-origin so the embedded SPA can load its own CSS and scripts', async () => {
    vi.mocked(getAboutContent).mockResolvedValue({
      success: true,
      message: '',
      data: 'https://www.9thnet.cn',
      content_type: 'html',
    })

    showAbout()

    const frame = await screen.findByTitle('About')
    expect(frame).toHaveAttribute('src', 'https://www.9thnet.cn')
    expect(frame).not.toHaveAttribute('srcDoc')
    expect(frame.getAttribute('sandbox')?.split(/\s+/)).toEqual(
      expect.arrayContaining(['allow-scripts', 'allow-same-origin'])
    )
  })

  it('embeds HTML markup in an iframe srcDoc when the format is HTML', async () => {
    vi.mocked(getAboutContent).mockResolvedValue({
      success: true,
      message: '',
      data: '<p>About us</p>',
      content_type: 'html',
    })

    showAbout()

    const frame = await screen.findByTitle('About')
    expect(frame).toHaveAttribute('srcDoc', '<p>About us</p>')
    expect(frame).not.toHaveAttribute('src')
    expect(frame.getAttribute('sandbox')?.split(/\s+/)).not.toEqual(
      expect.arrayContaining(['allow-same-origin'])
    )
  })

  it('renders Markdown content instead of an iframe when the format is Markdown', async () => {
    vi.mocked(getAboutContent).mockResolvedValue({
      success: true,
      message: '',
      data: '# About us',
      content_type: 'markdown',
    })

    showAbout()

    const content = await screen.findByTestId('rich-content')
    expect(content).toHaveAttribute('data-mode', 'markdown')
    expect(content).toHaveTextContent('# About us')
    expect(screen.queryByTitle('About')).not.toBeInTheDocument()
  })

  it('shows the empty state when the administrator has not set about content', async () => {
    vi.mocked(getAboutContent).mockResolvedValue({
      success: true,
      message: '',
      data: '',
    })

    showAbout()

    expect(
      await screen.findByRole('heading', { name: 'No About Content Set' })
    ).toBeVisible()
    expect(screen.queryByTitle('About')).not.toBeInTheDocument()
  })
})
