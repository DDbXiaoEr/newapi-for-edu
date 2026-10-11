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
import { describe, expect, test } from 'vitest'

import { parseQuotaFromDollars } from '@/lib/format'

import {
  transformFormDataToPayload,
  USER_FORM_DEFAULT_VALUES,
} from '../user-form'

describe('transformFormDataToPayload', () => {
  test('includes quota when creating a user', () => {
    const payload = transformFormDataToPayload({
      ...USER_FORM_DEFAULT_VALUES,
      username: 'quota-user',
      password: 'password1',
      quota_dollars: 1.5,
    })
    expect(payload.quota).toBe(parseQuotaFromDollars(1.5))
    expect(payload.role).toBe(1)
  })

  test('does not send quota on the user update payload', () => {
    const payload = transformFormDataToPayload(
      {
        ...USER_FORM_DEFAULT_VALUES,
        username: 'quota-user',
        quota_dollars: 1.5,
        group: 'default',
      },
      7
    )
    expect(payload.quota).toBeUndefined()
    expect(payload.id).toBe(7)
  })
})
