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
import { describe, expect, it } from 'vitest'

import { normalizeAboutContentType } from '../content-format'

describe('normalizeAboutContentType', () => {
  it('keeps an explicit HTML or Markdown selection', () => {
    expect(normalizeAboutContentType('html', '# About')).toBe('html')
    expect(normalizeAboutContentType('markdown', '<p>About</p>')).toBe(
      'markdown'
    )
  })

  it('infers HTML from a URL or HTML markup when the format is unset', () => {
    expect(normalizeAboutContentType('', 'https://example.com')).toBe('html')
    expect(normalizeAboutContentType(undefined, '<p>About us</p>')).toBe(
      'html'
    )
  })

  it('infers Markdown when the format is unset and the content is not HTML', () => {
    expect(normalizeAboutContentType('', '# About us')).toBe('markdown')
    expect(normalizeAboutContentType(undefined, '')).toBe('markdown')
  })
})
