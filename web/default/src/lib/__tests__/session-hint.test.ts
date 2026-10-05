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
import { afterEach, describe, expect, test } from 'vitest'

import {
  hasSessionHint,
  readCookie,
  SESSION_HINT_COOKIE_NAME,
} from '../session-hint'

function clearCookies() {
  for (const part of document.cookie.split(';')) {
    const name = part.split('=')[0]?.trim()
    if (name) {
      document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/`
    }
  }
}

afterEach(() => {
  clearCookies()
})

describe('session hint cookie parsing', () => {
  test('reads the named cookie and ignores a similar suffix name', () => {
    expect(readCookie('', SESSION_HINT_COOKIE_NAME)).toBeNull()
    expect(
      readCookie('prefix_new_api_has_session=1', SESSION_HINT_COOKIE_NAME)
    ).toBeNull()
    expect(
      readCookie(
        'theme=dark; new_api_has_session=1; sid=abc',
        SESSION_HINT_COOKIE_NAME
      )
    ).toBe('1')
    expect(
      readCookie(` ${SESSION_HINT_COOKIE_NAME} = 1 `, SESSION_HINT_COOKIE_NAME)
    ).toBe('1')
  })

  test('treats any present hint cookie as a session, including an empty value', () => {
    expect(hasSessionHint()).toBe(false)

    document.cookie = `${SESSION_HINT_COOKIE_NAME}=1`
    expect(hasSessionHint()).toBe(true)

    clearCookies()
    document.cookie = `${SESSION_HINT_COOKIE_NAME}=`
    expect(hasSessionHint()).toBe(true)
  })
})
