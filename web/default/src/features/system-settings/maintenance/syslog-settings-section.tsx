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
import * as z from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { useTranslation } from 'react-i18next'
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Switch } from '@/components/ui/switch'
import {
  SettingsForm,
  SettingsSwitchContent,
  SettingsSwitchItem,
} from '../components/settings-form-layout'
import { SettingsPageFormActions } from '../components/settings-page-context'
import { SettingsSection } from '../components/settings-section'
import { useResetForm } from '../hooks/use-reset-form'
import { useUpdateOption } from '../hooks/use-update-option'

const syslogSchema = z.object({
  SyslogEnabled: z.boolean(),
  SyslogNetwork: z.string(),
  SyslogAddr: z.string(),
  SyslogTag: z.string(),
})

type SyslogFormValues = z.infer<typeof syslogSchema>

type SyslogSettingsSectionProps = {
  defaultValues: SyslogFormValues
}

export function SyslogSettingsSection({
  defaultValues,
}: SyslogSettingsSectionProps) {
  const { t } = useTranslation()
  const updateOption = useUpdateOption()

  const form = useForm({
    resolver: zodResolver(syslogSchema),
    defaultValues,
  })

  useResetForm(form, defaultValues)

  const onSubmit = async (data: SyslogFormValues) => {
    const updates = Object.entries(data).filter(
      ([key, value]) => value !== defaultValues[key as keyof SyslogFormValues]
    )

    for (const [key, value] of updates) {
      await updateOption.mutateAsync({ key, value })
    }
  }

  return (
    <SettingsSection title={t('Syslog Settings')}>
      <Form {...form}>
        <SettingsForm onSubmit={form.handleSubmit(onSubmit)}>
          <SettingsPageFormActions
            onSave={form.handleSubmit(onSubmit)}
            isSaving={updateOption.isPending}
            saveLabel='Save Syslog Settings'
          />
          <FormField
            control={form.control}
            name='SyslogEnabled'
            render={({ field }) => (
              <SettingsSwitchItem>
                <SettingsSwitchContent>
                  <FormLabel>{t('Enable Syslog')}</FormLabel>
                </SettingsSwitchContent>
                <FormControl>
                  <Switch
                    checked={field.value}
                    onCheckedChange={field.onChange}
                  />
                </FormControl>
                <FormMessage />
              </SettingsSwitchItem>
            )}
          />
          <FormField
            control={form.control}
            name='SyslogTag'
            render={({ field }) => (
              <SettingsSwitchItem>
                <SettingsSwitchContent>
                  <FormLabel>{t('Syslog Tag')}</FormLabel>
                </SettingsSwitchContent>
                <FormControl>
                  <Input placeholder='newapi' {...field} />
                </FormControl>
                <FormMessage />
              </SettingsSwitchItem>
            )}
          />
          <FormField
            control={form.control}
            name='SyslogNetwork'
            render={({ field }) => (
              <SettingsSwitchItem>
                <SettingsSwitchContent>
                  <FormLabel>{t('Syslog Network Protocol')}</FormLabel>
                  <FormDescription>
                    {t(
                      'Supports udp / tcp, leave empty for local syslog'
                    )}
                  </FormDescription>
                </SettingsSwitchContent>
                <FormControl>
                  <Input
                    placeholder={t('Leave empty to use local socket')}
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </SettingsSwitchItem>
            )}
          />
          <FormField
            control={form.control}
            name='SyslogAddr'
            render={({ field }) => (
              <SettingsSwitchItem>
                <SettingsSwitchContent>
                  <FormLabel>{t('Syslog Server Address')}</FormLabel>
                  <FormDescription>
                    {t(
                      'Remote syslog server address, e.g. 192.168.1.1:514'
                    )}
                  </FormDescription>
                </SettingsSwitchContent>
                <FormControl>
                  <Input placeholder='localhost:514' {...field} />
                </FormControl>
                <FormMessage />
              </SettingsSwitchItem>
            )}
          />
        </SettingsForm>
      </Form>
    </SettingsSection>
  )
}
