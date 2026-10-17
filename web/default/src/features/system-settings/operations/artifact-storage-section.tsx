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
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { useTranslation } from 'react-i18next'
import * as z from 'zod'

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

import {
  SettingsForm,
  SettingsSwitchContent,
  SettingsSwitchItem,
} from '../components/settings-form-layout'
import { SettingsPageFormActions } from '../components/settings-page-context'
import { SettingsSection } from '../components/settings-section'
import { useResetForm } from '../hooks/use-reset-form'
import { useUpdateOption } from '../hooks/use-update-option'

const createArtifactStorageSchema = (t: (key: string) => string) =>
  z.object({
    enabled: z.boolean(),
    endpoint: z.string().refine((value) => {
      const trimmed = value.trim()
      if (!trimmed) return true
      return /^https?:\/\/[^\s]+$/.test(trimmed)
    }, t('Provide a valid URL starting with http:// or https://')),
    bucket: z.string().refine((value) => {
      const trimmed = value.trim()
      if (!trimmed) return true
      return /^[a-z0-9][a-z0-9.-]{1,61}[a-z0-9]$/.test(trimmed)
    }, t('Use lowercase letters, numbers, dots, and hyphens')),
    region: z.string(),
    accessKey: z.string(),
    secretKey: z.string(),
    prefix: z.string().refine((value) => {
      const trimmed = value.trim()
      return (
        trimmed === '' || (!trimmed.startsWith('/') && !trimmed.includes('\\'))
      )
    }, t('The prefix must not start with a slash or contain backslashes')),
    pathStyle: z.boolean(),
  })

type ArtifactStorageFormValues = z.infer<
  ReturnType<typeof createArtifactStorageSchema>
>

type ArtifactStorageSectionProps = {
  defaultValues: {
    'task_artifact_store.mode': string
    'task_artifact_store.endpoint': string
    'task_artifact_store.bucket': string
    'task_artifact_store.region': string
    'task_artifact_store.access_key': string
    'task_artifact_store.prefix': string
    'task_artifact_store.path_style': boolean
  }
}

const buildFormDefaults = (
  defaults: ArtifactStorageSectionProps['defaultValues']
): ArtifactStorageFormValues => ({
  enabled: defaults['task_artifact_store.mode'] === 's3',
  endpoint: defaults['task_artifact_store.endpoint'],
  bucket: defaults['task_artifact_store.bucket'],
  region: defaults['task_artifact_store.region'],
  accessKey: defaults['task_artifact_store.access_key'],
  secretKey: '',
  prefix: defaults['task_artifact_store.prefix'],
  pathStyle: defaults['task_artifact_store.path_style'] ?? true,
})

export function ArtifactStorageSection({
  defaultValues,
}: ArtifactStorageSectionProps) {
  const { t } = useTranslation()
  const updateOption = useUpdateOption()
  const schema = createArtifactStorageSchema(t)
  const initialValues = buildFormDefaults(defaultValues)

  const form = useForm<ArtifactStorageFormValues>({
    resolver: zodResolver(schema),
    defaultValues: initialValues,
  })

  useResetForm(form, initialValues)

  const onSubmit = async (values: ArtifactStorageFormValues) => {
    const sanitized = {
      endpoint: values.endpoint.trim().replace(/\/+$/, ''),
      bucket: values.bucket.trim(),
      region: values.region.trim(),
      accessKey: values.accessKey.trim(),
      secretKey: values.secretKey.trim(),
      prefix: values.prefix.trim(),
      pathStyle: values.pathStyle,
    }

    if (values.enabled && (!sanitized.endpoint || !sanitized.bucket)) {
      form.setError('endpoint', {
        message: t('Endpoint and bucket are required before enabling S3'),
      })
      return
    }

    const updates: Array<{ key: string; value: string | boolean }> = []
    if (sanitized.endpoint !== initialValues.endpoint.trim()) {
      updates.push({
        key: 'task_artifact_store.endpoint',
        value: sanitized.endpoint,
      })
    }
    if (sanitized.bucket !== initialValues.bucket.trim()) {
      updates.push({
        key: 'task_artifact_store.bucket',
        value: sanitized.bucket,
      })
    }
    if (sanitized.region !== initialValues.region.trim()) {
      updates.push({
        key: 'task_artifact_store.region',
        value: sanitized.region,
      })
    }
    if (sanitized.accessKey !== initialValues.accessKey.trim()) {
      updates.push({
        key: 'task_artifact_store.access_key',
        value: sanitized.accessKey,
      })
    }
    if (sanitized.secretKey) {
      updates.push({
        key: 'task_artifact_store.secret_key',
        value: sanitized.secretKey,
      })
    }
    if (sanitized.prefix !== initialValues.prefix.trim()) {
      updates.push({
        key: 'task_artifact_store.prefix',
        value: sanitized.prefix,
      })
    }
    if (sanitized.pathStyle !== initialValues.pathStyle) {
      updates.push({
        key: 'task_artifact_store.path_style',
        value: sanitized.pathStyle,
      })
    }

    const mode = values.enabled ? 's3' : 'upstream'
    if (mode !== defaultValues['task_artifact_store.mode']) {
      updates.push({ key: 'task_artifact_store.mode', value: mode })
    }

    for (const update of updates) {
      await updateOption.mutateAsync(update)
    }
  }

  return (
    <SettingsSection title={t('Artifact Storage')}>
      <Form {...form}>
        <SettingsForm onSubmit={form.handleSubmit(onSubmit)} autoComplete='off'>
          <SettingsPageFormActions
            onSave={form.handleSubmit(onSubmit)}
            isSaving={updateOption.isPending}
            saveLabel='Save artifact storage settings'
          />
          <FormField
            control={form.control}
            name='enabled'
            render={({ field }) => (
              <SettingsSwitchItem>
                <SettingsSwitchContent>
                  <FormLabel>{t('Store generated artifacts in S3')}</FormLabel>
                  <FormDescription>
                    {t(
                      'Image and video task artifacts are copied to object storage. Downloads stay proxied by this server so S3 credentials are never exposed.'
                    )}
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
          <div className='grid gap-4 md:grid-cols-2'>
            <FormField
              control={form.control}
              name='endpoint'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('S3 Endpoint')}</FormLabel>
                  <FormControl>
                    <Input
                      type='url'
                      inputMode='url'
                      placeholder={t('https://s3.example.com')}
                      autoComplete='off'
                      {...field}
                      onChange={(event) => field.onChange(event.target.value)}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name='bucket'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Bucket')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder={t('task-artifacts')}
                      autoComplete='off'
                      {...field}
                      onChange={(event) => field.onChange(event.target.value)}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name='region'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Region')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder='us-east-1'
                      autoComplete='off'
                      {...field}
                      onChange={(event) => field.onChange(event.target.value)}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name='prefix'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Object Prefix')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder={t('optional/path')}
                      autoComplete='off'
                      {...field}
                      onChange={(event) => field.onChange(event.target.value)}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name='accessKey'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Access Key')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder={t('Access key ID')}
                      autoComplete='off'
                      {...field}
                      onChange={(event) => field.onChange(event.target.value)}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name='secretKey'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Secret Key')}</FormLabel>
                  <FormControl>
                    <Input
                      type='password'
                      placeholder={t('Enter new secret to update')}
                      autoComplete='new-password'
                      {...field}
                      onChange={(event) => field.onChange(event.target.value)}
                    />
                  </FormControl>
                  <FormDescription>
                    {t(
                      'Leave blank to keep the existing secret. Fill all fields before enabling S3.'
                    )}
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <FormField
            control={form.control}
            name='pathStyle'
            render={({ field }) => (
              <SettingsSwitchItem>
                <SettingsSwitchContent>
                  <FormLabel>{t('Use path-style addressing')}</FormLabel>
                  <FormDescription>
                    {t(
                      'Enable for MinIO and most self-hosted S3-compatible services.'
                    )}
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
        </SettingsForm>
      </Form>
    </SettingsSection>
  )
}
