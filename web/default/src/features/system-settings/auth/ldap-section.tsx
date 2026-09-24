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
import { useEffect, useMemo, useRef } from 'react'
import * as z from 'zod'
import { useTranslation } from 'react-i18next'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { toast } from 'sonner'
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Switch } from '@/components/ui/switch'
import { FormNavigationGuard } from '../components/form-navigation-guard'
import { GroupAssignmentRulesEditor } from './group-assignment-rules-editor'
import {
  editorValueFromRules,
  stringifyGroupAssignmentEditorValue,
} from './group-assignment-rules'
import {
  SettingsForm,
  SettingsSwitchContent,
  SettingsSwitchItem,
} from '../components/settings-form-layout'
import { SettingsPageFormActions } from '../components/settings-page-context'
import { SettingsSection } from '../components/settings-section'
import { useUpdateOption } from '../hooks/use-update-option'

const ldapSchema = z.object({
  ldap: z.object({
    enabled: z.boolean(),
    server_url: z.string(),
    bind_dn: z.string(),
    bind_password: z.string(),
    base_dn: z.string(),
    user_filter: z.string(),
    username_attribute: z.string(),
    display_name_attribute: z.string(),
    mail_attribute: z.string(),
    start_tls: z.boolean(),
    skip_tls_verify: z.boolean(),
    timeout_seconds: z.string(),
    group_assignment_rules: z.object({
      attribute: z.string(),
      rules: z.array(
        z.object({
          id: z.string(),
          pattern: z.string(),
          group: z.string(),
        })
      ),
    }),
  }),
})

type LdapFormValues = z.infer<typeof ldapSchema>

type FlatLdapDefaults = {
  'ldap.enabled': boolean
  'ldap.server_url': string
  'ldap.bind_dn': string
  'ldap.bind_password': string
  'ldap.base_dn': string
  'ldap.user_filter': string
  'ldap.username_attribute': string
  'ldap.display_name_attribute': string
  'ldap.mail_attribute': string
  'ldap.start_tls': boolean
  'ldap.skip_tls_verify': boolean
  'ldap.timeout_seconds': string
  'ldap.group_assignment_rules': string
}

const buildFormDefaults = (defaults: FlatLdapDefaults): LdapFormValues => ({
  ldap: {
    enabled: defaults['ldap.enabled'],
    server_url: defaults['ldap.server_url'] ?? '',
    bind_dn: defaults['ldap.bind_dn'] ?? '',
    bind_password: defaults['ldap.bind_password'] ?? '',
    base_dn: defaults['ldap.base_dn'] ?? '',
    user_filter: defaults['ldap.user_filter'] ?? '',
    username_attribute: defaults['ldap.username_attribute'] ?? '',
    display_name_attribute: defaults['ldap.display_name_attribute'] ?? '',
    mail_attribute: defaults['ldap.mail_attribute'] ?? '',
    start_tls: defaults['ldap.start_tls'],
    skip_tls_verify: defaults['ldap.skip_tls_verify'],
    timeout_seconds: defaults['ldap.timeout_seconds'] || '5',
    group_assignment_rules: editorValueFromRules(
      defaults['ldap.group_assignment_rules']
    ),
  },
})

const normalizeFormValues = (values: LdapFormValues): FlatLdapDefaults => ({
  'ldap.enabled': values.ldap.enabled,
  'ldap.server_url': values.ldap.server_url,
  'ldap.bind_dn': values.ldap.bind_dn,
  'ldap.bind_password': values.ldap.bind_password,
  'ldap.base_dn': values.ldap.base_dn,
  'ldap.user_filter': values.ldap.user_filter,
  'ldap.username_attribute': values.ldap.username_attribute,
  'ldap.display_name_attribute': values.ldap.display_name_attribute,
  'ldap.mail_attribute': values.ldap.mail_attribute,
  'ldap.start_tls': values.ldap.start_tls,
  'ldap.skip_tls_verify': values.ldap.skip_tls_verify,
  'ldap.timeout_seconds': values.ldap.timeout_seconds || '5',
  'ldap.group_assignment_rules': stringifyGroupAssignmentEditorValue(
    values.ldap.group_assignment_rules
  ),
})

type LdapSectionProps = {
  defaultValues: FlatLdapDefaults
}

export function LdapSection(props: LdapSectionProps) {
  const { t } = useTranslation()
  const updateOption = useUpdateOption()

  const formDefaults = useMemo(
    () => buildFormDefaults(props.defaultValues),
    [props.defaultValues]
  )

  const form = useForm<LdapFormValues>({
    resolver: zodResolver(ldapSchema),
    defaultValues: formDefaults,
  })

  const baselineRef = useRef<FlatLdapDefaults>(props.defaultValues)
  const baselineSerializedRef = useRef<string>(
    JSON.stringify(props.defaultValues)
  )

  useEffect(() => {
    const serialized = JSON.stringify(props.defaultValues)
    if (serialized === baselineSerializedRef.current) return
    baselineRef.current = props.defaultValues
    baselineSerializedRef.current = serialized
    form.reset(buildFormDefaults(props.defaultValues))
  }, [props.defaultValues, form])

  const onSubmit = async (values: LdapFormValues) => {
    const normalized = normalizeFormValues(values)
    const changedKeys = (
      Object.keys(normalized) as Array<keyof FlatLdapDefaults>
    ).filter((key) => {
      if (key === 'ldap.bind_password') {
        return (
          normalized[key] !== '' &&
          normalized[key] !== baselineRef.current[key]
        )
      }
      return normalized[key] !== baselineRef.current[key]
    })

    if (changedKeys.length === 0) {
      toast.info(t('No changes to save'))
      return
    }

    for (const key of changedKeys) {
      await updateOption.mutateAsync({
        key,
        value: normalized[key],
      })
    }

    baselineRef.current = normalized
    baselineSerializedRef.current = JSON.stringify(normalized)
    form.reset(buildFormDefaults(normalized))
  }

  const handleReset = () => {
    form.reset(buildFormDefaults(baselineRef.current))
    toast.success(t('Form reset to saved values'))
  }

  return (
    <>
      <FormNavigationGuard when={form.formState.isDirty} />

      <SettingsSection title={t('LDAP Configuration')}>
        <div className='text-muted-foreground text-sm'>
          {t(
            'Support LDAP login, fallback to LDAP verification when local password does not match'
          )}
        </div>
        <Form {...form}>
          <SettingsForm onSubmit={form.handleSubmit(onSubmit)}>
            <SettingsPageFormActions
              onSave={form.handleSubmit(onSubmit)}
              onReset={handleReset}
              isSaving={updateOption.isPending}
              isResetDisabled={!form.formState.isDirty}
            />

            <FormField
              control={form.control}
              name='ldap.enabled'
              render={({ field }) => (
                <SettingsSwitchItem>
                  <SettingsSwitchContent>
                    <FormLabel>{t('Enable LDAP Login')}</FormLabel>
                    <FormDescription>
                      {t('Allow users to sign in with LDAP')}
                    </FormDescription>
                  </SettingsSwitchContent>
                  <FormControl>
                    <Switch
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                  </FormControl>
                </SettingsSwitchItem>
              )}
            />

            <FormField
              control={form.control}
              name='ldap.server_url'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('LDAP Server URL')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder={t('e.g. ldaps://ldap.example.com:636')}
                      autoComplete='off'
                      value={field.value ?? ''}
                      onChange={(event) =>
                        field.onChange(event.target.value)
                      }
                      name={field.name}
                      onBlur={field.onBlur}
                      ref={field.ref}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='ldap.base_dn'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Base DN')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder={t('e.g. dc=example,dc=com')}
                      autoComplete='off'
                      value={field.value ?? ''}
                      onChange={(event) =>
                        field.onChange(event.target.value)
                      }
                      name={field.name}
                      onBlur={field.onBlur}
                      ref={field.ref}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='ldap.bind_dn'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Bind DN')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder={t('Optional, used for directory search')}
                      autoComplete='off'
                      value={field.value ?? ''}
                      onChange={(event) =>
                        field.onChange(event.target.value)
                      }
                      name={field.name}
                      onBlur={field.onBlur}
                      ref={field.ref}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='ldap.bind_password'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Bind Password')}</FormLabel>
                  <FormControl>
                    <Input
                      type='password'
                      placeholder={t(
                        'Sensitive information will not be sent to the frontend'
                      )}
                      autoComplete='new-password'
                      value={field.value ?? ''}
                      onChange={(event) =>
                        field.onChange(event.target.value)
                      }
                      name={field.name}
                      onBlur={field.onBlur}
                      ref={field.ref}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='ldap.user_filter'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('User Filter')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder={t(
                        'Defaults to (|(uid={username})(mail={username}))'
                      )}
                      autoComplete='off'
                      value={field.value ?? ''}
                      onChange={(event) =>
                        field.onChange(event.target.value)
                      }
                      name={field.name}
                      onBlur={field.onBlur}
                      ref={field.ref}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='ldap.timeout_seconds'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Connection Timeout (seconds)')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder='5'
                      autoComplete='off'
                      value={field.value ?? ''}
                      onChange={(event) =>
                        field.onChange(event.target.value)
                      }
                      name={field.name}
                      onBlur={field.onBlur}
                      ref={field.ref}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='ldap.username_attribute'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Username Attribute')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder='uid'
                      autoComplete='off'
                      value={field.value ?? ''}
                      onChange={(event) =>
                        field.onChange(event.target.value)
                      }
                      name={field.name}
                      onBlur={field.onBlur}
                      ref={field.ref}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='ldap.display_name_attribute'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Display Name Attribute')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder='cn'
                      autoComplete='off'
                      value={field.value ?? ''}
                      onChange={(event) =>
                        field.onChange(event.target.value)
                      }
                      name={field.name}
                      onBlur={field.onBlur}
                      ref={field.ref}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='ldap.mail_attribute'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Email Attribute')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder='mail'
                      autoComplete='off'
                      value={field.value ?? ''}
                      onChange={(event) =>
                        field.onChange(event.target.value)
                      }
                      name={field.name}
                      onBlur={field.onBlur}
                      ref={field.ref}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='ldap.group_assignment_rules'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('New user group assignment rules')}</FormLabel>
                  <FormControl>
                    <GroupAssignmentRulesEditor
                      value={field.value}
                      onChange={field.onChange}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='ldap.start_tls'
              render={({ field }) => (
                <SettingsSwitchItem>
                  <SettingsSwitchContent>
                    <FormLabel>{t('Enable StartTLS')}</FormLabel>
                  </SettingsSwitchContent>
                  <FormControl>
                    <Switch
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                  </FormControl>
                </SettingsSwitchItem>
              )}
            />

            <FormField
              control={form.control}
              name='ldap.skip_tls_verify'
              render={({ field }) => (
                <SettingsSwitchItem>
                  <SettingsSwitchContent>
                    <FormLabel>{t('Skip TLS Certificate Verification')}</FormLabel>
                  </SettingsSwitchContent>
                  <FormControl>
                    <Switch
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                  </FormControl>
                </SettingsSwitchItem>
              )}
            />
          </SettingsForm>
        </Form>
      </SettingsSection>
    </>
  )
}
