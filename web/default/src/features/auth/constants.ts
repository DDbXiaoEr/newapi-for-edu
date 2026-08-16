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
import { z } from 'zod'

// ============================================================================
// Validation Constants
// ============================================================================

export const PASSWORD_MIN_LENGTH = 8
export const PASSWORD_MAX_LENGTH = 20
export const OTP_LENGTH = 6
export const BACKUP_CODE_LENGTH = 9 // XXXX-XXXX format
export const BACKUP_CODE_REGEX = /^[A-Z0-9]{4}-[A-Z0-9]{4}$/i
export const OTP_REGEX = /^\d{6}$/

// ============================================================================
// Form Schemas
// ============================================================================

// 登录不限制密码长度：密码只是凭证，LDAP/本地账号都可能是任意长度
export function getLoginFormSchema() {
  return z.object({
    username: z.string().min(1, 'Please enter your username or email'),
    password: z.string().min(1, 'Please enter your password'),
  })
}

export const loginFormSchema = getLoginFormSchema()

export function getRegisterFormSchema(ldapEnabled = false) {
  return z
    .object({
      username: z.string().min(1, 'Please enter your username'),
      email: z.string().optional(),
      password: z
        .string()
        .min(1, 'Please enter your password')
        .refine(
          (value) => ldapEnabled || value.length >= PASSWORD_MIN_LENGTH,
          { message: 'Password must be at least 8 characters long' }
        )
        .max(20, 'Password must be at most 20 characters long'),
      confirmPassword: z.string().min(1, 'Please confirm your password'),
    })
    .refine((data) => data.password === data.confirmPassword, {
      message: "Passwords don't match.",
      path: ['confirmPassword'],
    })
}

export const registerFormSchema = getRegisterFormSchema(false)

export const forgotPasswordFormSchema = z.object({
  email: z.string().email({
    message: 'Please enter a valid email address',
  }),
})

export const otpFormSchema = z.object({
  otp: z.string().min(1, 'Please enter a code.'),
})

// ============================================================================
// Countdown Constants
// ============================================================================

export const EMAIL_VERIFICATION_COUNTDOWN = 30 // seconds
export const PASSWORD_RESET_COUNTDOWN = 30 // seconds

// ============================================================================
// OAuth Constants
// ============================================================================

export const OAUTH_BIND_STORAGE_KEY = 'oauth:binding:result'
