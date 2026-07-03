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
import {
  SettingsForm,
  SettingsSwitchContent,
  SettingsSwitchItem,
} from '../components/settings-form-layout'
import { SettingsPageFormActions } from '../components/settings-page-context'
import { SettingsSection } from '../components/settings-section'
import { useUpdateOption } from '../hooks/use-update-option'

const casSchema = z.object({
  cas: z.object({
    enabled: z.boolean(),
    server_url: z.string(),
    service_id: z.string(),
    username_attribute: z.string(),
    display_name_attribute: z.string(),
    email_attribute: z.string(),
    access_attribute: z.string(),
    access_attribute_value: z.string(),
  }),
})

type CasFormValues = z.infer<typeof casSchema>

type FlatCasDefaults = {
  'cas.enabled': boolean
  'cas.server_url': string
  'cas.service_id': string
  'cas.username_attribute': string
  'cas.display_name_attribute': string
  'cas.email_attribute': string
  'cas.access_attribute': string
  'cas.access_attribute_value': string
}

const buildFormDefaults = (defaults: FlatCasDefaults): CasFormValues => ({
  cas: {
    enabled: defaults['cas.enabled'],
    server_url: defaults['cas.server_url'] ?? '',
    service_id: defaults['cas.service_id'] ?? '',
    username_attribute: defaults['cas.username_attribute'] ?? '',
    display_name_attribute: defaults['cas.display_name_attribute'] ?? '',
    email_attribute: defaults['cas.email_attribute'] ?? '',
    access_attribute: defaults['cas.access_attribute'] ?? '',
    access_attribute_value: defaults['cas.access_attribute_value'] ?? '',
  },
})

const normalizeFormValues = (values: CasFormValues): FlatCasDefaults => ({
  'cas.enabled': values.cas.enabled,
  'cas.server_url': values.cas.server_url,
  'cas.service_id': values.cas.service_id,
  'cas.username_attribute': values.cas.username_attribute,
  'cas.display_name_attribute': values.cas.display_name_attribute,
  'cas.email_attribute': values.cas.email_attribute,
  'cas.access_attribute': values.cas.access_attribute,
  'cas.access_attribute_value': values.cas.access_attribute_value,
})

type CasSectionProps = {
  defaultValues: FlatCasDefaults
}

export function CasSection(props: CasSectionProps) {
  const { t } = useTranslation()
  const updateOption = useUpdateOption()

  const formDefaults = useMemo(
    () => buildFormDefaults(props.defaultValues),
    [props.defaultValues]
  )

  const form = useForm<CasFormValues>({
    resolver: zodResolver(casSchema),
    defaultValues: formDefaults,
  })

  const baselineRef = useRef<FlatCasDefaults>(props.defaultValues)
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

  const onSubmit = async (values: CasFormValues) => {
    const normalized = normalizeFormValues(values)
    const changedKeys = (
      Object.keys(normalized) as Array<keyof FlatCasDefaults>
    ).filter((key) => {
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

      <SettingsSection title={t('CAS Configuration')}>
        <div className='text-muted-foreground text-sm'>
          {t(
            'CAS (Central Authentication Service) is a single sign-on protocol for web applications'
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
              name='cas.enabled'
              render={({ field }) => (
                <SettingsSwitchItem>
                  <SettingsSwitchContent>
                    <FormLabel>{t('Enable CAS Login')}</FormLabel>
                    <FormDescription>
                      {t('Allow users to sign in with CAS')}
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
              name='cas.server_url'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('CAS Server URL')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder={t('https://cas.example.com/cas')}
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
                  <FormDescription>
                    {t('The base URL of your CAS server')}
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='cas.service_id'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Service ID (Optional)')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder={t('CAS service identifier')}
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
                  <FormDescription>
                    {t('Custom CAS service identifier, defaults to server address if empty')}
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='cas.username_attribute'
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
                  <FormDescription>
                    {t('CAS attribute name for username. Defaults to CAS principal if empty')}
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='cas.display_name_attribute'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Display Name Attribute')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder='displayname'
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
                  <FormDescription>
                    {t('CAS attribute name for display name')}
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='cas.email_attribute'
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
                  <FormDescription>
                    {t('CAS attribute name for email address')}
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='cas.access_attribute'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Access Restriction Attribute')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder={t('eduPersonAffiliation')}
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
                  <FormDescription>
                    {t('CAS attribute used for access control. Leave empty to disable')}
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='cas.access_attribute_value'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Required Attribute Value')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder={t('student')}
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
                  <FormDescription>
                    {t('Only users with this attribute value can login')}
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
          </SettingsForm>
        </Form>
      </SettingsSection>
    </>
  )
}
