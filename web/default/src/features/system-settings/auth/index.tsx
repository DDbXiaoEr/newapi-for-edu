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
import { SettingsPage } from '../components/settings-page'
import type { AuthSettings } from '../types'
import {
  AUTH_DEFAULT_SECTION,
  getAuthSectionContent,
  getAuthSectionMeta,
} from './section-registry.tsx'

const defaultAuthSettings: AuthSettings = {
  PasswordLoginEnabled: true,
  PasswordRegisterEnabled: true,
  EmailVerificationEnabled: false,
  RegisterEnabled: true,
  EmailDomainRestrictionEnabled: false,
  EmailAliasRestrictionEnabled: false,
  EmailDomainWhitelist: '',
  ServerAddress: '',
  GitHubOAuthEnabled: false,
  GitHubClientId: '',
  GitHubClientSecret: '',
  'discord.enabled': false,
  'discord.client_id': '',
  'discord.client_secret': '',
  'oidc.enabled': false,
  'oidc.display_name': '',
  'oidc.client_id': '',
  'oidc.client_secret': '',
  'oidc.well_known': '',
  'oidc.authorization_endpoint': '',
  'oidc.token_endpoint': '',
  'oidc.user_info_endpoint': '',
  TelegramOAuthEnabled: false,
  'telegram.client_id': '',
  'telegram.client_secret': '',
  TelegramBotToken: '',
  TelegramBotName: '',
  LinuxDOOAuthEnabled: false,
  LinuxDOClientId: '',
  LinuxDOClientSecret: '',
  LinuxDOMinimumTrustLevel: '0',
  WeChatAuthEnabled: false,
  WeChatServerAddress: '',
  WeChatServerToken: '',
  WeChatAccountQRCodeImageURL: '',
  'cas.enabled': false,
  'cas.server_url': '',
  'cas.service_id': '',
  'cas.username_attribute': '',
  'cas.display_name_attribute': '',
  'cas.email_attribute': '',
  'cas.access_attribute': '',
  'cas.access_attribute_value': '',
  'cas.group_assignment_rules': '[]',
  TurnstileCheckEnabled: false,
  TurnstileSiteKey: '',
  TurnstileSecretKey: '',
  'passkey.enabled': false,
  'passkey.rp_display_name': '',
  'passkey.rp_id': '',
  'passkey.legacy_rp_ids': '',
  'passkey.origins': '',
  'passkey.allow_insecure_origin': false,
  'passkey.user_verification': 'preferred',
  'passkey.attachment_preference': '',
  'ldap.enabled': false,
  'ldap.server_url': '',
  'ldap.bind_dn': '',
  'ldap.bind_password': '',
  'ldap.base_dn': '',
  'ldap.user_filter': '',
  'ldap.username_attribute': 'uid',
  'ldap.display_name_attribute': 'cn',
  'ldap.mail_attribute': 'mail',
  'ldap.start_tls': false,
  'ldap.skip_tls_verify': false,
  'ldap.timeout_seconds': '5',
  'ldap.group_assignment_rules': '[]',
}

export function AuthSettings() {
  return (
    <SettingsPage
      routePath='/_authenticated/system-settings/auth/$section'
      defaultSettings={defaultAuthSettings}
      defaultSection={AUTH_DEFAULT_SECTION}
      getSectionContent={getAuthSectionContent}
      getSectionMeta={getAuthSectionMeta}
    />
  )
}
