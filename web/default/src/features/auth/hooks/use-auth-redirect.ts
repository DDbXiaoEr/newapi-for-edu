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
import { useNavigate } from '@tanstack/react-router'
import i18n from 'i18next'
import { useAuthStore, type AuthUser } from '@/stores/auth-store'
import { getSelf } from '@/lib/api'
import { markSessionVerified } from '@/lib/session-flag'
import type { User } from '@/features/users/types'
import { saveUserId, saveToken } from '../lib/storage'
import type { LoginSuccessData } from '../types'

function getSavedLanguage(user: User): string | undefined {
  const userData = user as Record<string, unknown>
  if (typeof userData.language === 'string') {
    return userData.language
  }

  if (typeof userData.setting !== 'string') {
    return undefined
  }

  try {
    const setting = JSON.parse(userData.setting) as { language?: unknown }
    return typeof setting.language === 'string' ? setting.language : undefined
  } catch {
    return undefined
  }
}

function buildMinimalUser(data: LoginSuccessData): AuthUser {
  return {
    id: data.id ?? 0,
    username: data.username ?? '',
    display_name: data.display_name,
    role: data.role ?? 1,
    status: data.status ?? 1,
    group: data.group ?? 'default',
    token: data.token,
  }
}

/**
 * Hook for handling authentication redirects and user data management
 */
export function useAuthRedirect() {
  const navigate = useNavigate()
  const { auth } = useAuthStore()

  /**
   * Handle successful login
   * @param userData - User data from login response (setupLogin fields)
   * @param redirectTo - Redirect path after login
   */
  const handleLoginSuccess = async (
    userData?: LoginSuccessData | null,
    redirectTo?: string
  ) => {
    if (userData?.id) {
      saveUserId(userData.id)
    }

    if (userData?.token) {
      saveToken(userData.token)
    }

    if (userData?.id && userData?.username && userData?.token) {
      auth.setUser(buildMinimalUser(userData))
      markSessionVerified()
    } else {
      try {
        const self = await getSelf()
        if (self?.success && self.data) {
          const user = self.data as User
          auth.setUser(user)

          if (user.id) {
            saveUserId(user.id)
          }

          const savedLang = getSavedLanguage(user)
          if (savedLang && savedLang !== i18n.language) {
            i18n.changeLanguage(savedLang)
          }

          markSessionVerified()
        }
      } catch (error) {
        // eslint-disable-next-line no-console
        console.error('Failed to fetch user data:', error)
      }
    }

    const targetPath = redirectTo || '/dashboard'
    navigate({ to: targetPath, replace: true })
  }

  /**
   * Redirect to 2FA page
   */
  const redirectTo2FA = () => {
    navigate({ to: '/otp', replace: true })
  }

  /**
   * Redirect to login page
   */
  const redirectToLogin = () => {
    navigate({ to: '/sign-in', replace: true })
  }

  /**
   * Redirect to register page
   */
  const redirectToRegister = () => {
    navigate({ to: '/sign-up', replace: true })
  }

  return {
    handleLoginSuccess,
    redirectTo2FA,
    redirectToLogin,
    redirectToRegister,
  }
}
