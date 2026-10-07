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
import type { ChatPreset } from '@/features/chat/lib/chat-links'
import type { CCSwitchApp } from '@/features/keys/lib/cc-switch'

export type QuickAccessContext = {
  serverAddress: string
  /** Active API key, or the `sk-...` placeholder when none is available. */
  apiKey: string
  /** First available model, used to seed ready-to-use configuration. */
  model?: string
}

export type ConfigSnippet = {
  /** i18n key describing what the snippet configures. */
  labelKey: string
  /** Optional human-readable file path shown next to the label. */
  filePath?: string
  language: 'bash' | 'json' | 'toml'
  value: string
}

export type BuiltinClient = {
  id: string
  name: string
  descriptionKey: string
  docsUrl: string
  installCommand?: string
  /** Clients that expose a CC Switch one-click import deeplink. */
  quickImport?: { app: CCSwitchApp; name: string }
  snippets: (ctx: QuickAccessContext) => ConfigSnippet[]
}

export type ClientOption =
  | { id: string; name: string; kind: 'builtin'; builtin: BuiltinClient }
  | { id: string; name: string; kind: 'preset'; preset: ChatPreset }

const FALLBACK_API_KEY = 'sk-...'
const FALLBACK_MODEL = 'your-model'

function apiKeyOf(ctx: QuickAccessContext): string {
  return ctx.apiKey.trim() || FALLBACK_API_KEY
}

function modelOf(ctx: QuickAccessContext): string {
  return ctx.model?.trim() || FALLBACK_MODEL
}

function openCodeConfig(ctx: QuickAccessContext): string {
  return JSON.stringify(
    {
      $schema: 'https://opencode.ai/config.json',
      provider: {
        'new-api': {
          npm: '@ai-sdk/openai-compatible',
          name: 'New API',
          options: {
            baseURL: `${ctx.serverAddress}/v1`,
            apiKey: apiKeyOf(ctx),
          },
          models: {
            [modelOf(ctx)]: {},
          },
        },
      },
    },
    null,
    2
  )
}

function claudeSettings(ctx: QuickAccessContext): string {
  return JSON.stringify(
    {
      env: {
        ANTHROPIC_BASE_URL: ctx.serverAddress,
        ANTHROPIC_AUTH_TOKEN: apiKeyOf(ctx),
      },
    },
    null,
    2
  )
}

function codexConfig(ctx: QuickAccessContext): string {
  return [
    `model = "${modelOf(ctx)}"`,
    'model_provider = "new-api"',
    '',
    '[model_providers.new-api]',
    'name = "New API"',
    `base_url = "${ctx.serverAddress}/v1"`,
    'env_key = "NEW_API_KEY"',
  ].join('\n')
}

export const BUILTIN_CLIENTS: BuiltinClient[] = [
  {
    id: 'opencode',
    name: 'OpenCode',
    descriptionKey:
      'OpenCode is an open-source AI coding agent for the terminal. Register this service as an OpenAI-compatible provider in its config file.',
    docsUrl: 'https://opencode.ai/docs/',
    installCommand: 'curl -fsSL https://opencode.ai/install | bash',
    snippets: (ctx) => [
      {
        labelKey: 'Configuration File',
        filePath: '~/.config/opencode/opencode.json',
        language: 'json',
        value: openCodeConfig(ctx),
      },
    ],
  },
  {
    id: 'claude-code',
    name: 'Claude Code',
    descriptionKey:
      "Claude Code is Anthropic's coding agent for the terminal. Import it into CC Switch in one click, or point it at this service with environment variables.",
    docsUrl: 'https://docs.claude.com/en/docs/claude-code',
    installCommand: 'npm install -g @anthropic-ai/claude-code',
    quickImport: { app: 'claude', name: 'New API (Claude Code)' },
    snippets: (ctx) => [
      {
        labelKey: 'Environment variables',
        language: 'bash',
        value: [
          `export ANTHROPIC_BASE_URL="${ctx.serverAddress}"`,
          `export ANTHROPIC_AUTH_TOKEN="${apiKeyOf(ctx)}"`,
        ].join('\n'),
      },
      {
        labelKey: 'Configuration File',
        filePath: '~/.claude/settings.json',
        language: 'json',
        value: claudeSettings(ctx),
      },
    ],
  },
  {
    id: 'codex',
    name: 'Codex',
    descriptionKey:
      "Codex is OpenAI's coding CLI. Import it into CC Switch in one click, or configure an OpenAI-compatible provider that points at this service.",
    docsUrl: 'https://github.com/openai/codex',
    installCommand: 'npm install -g @openai/codex',
    quickImport: { app: 'codex', name: 'New API (Codex)' },
    snippets: (ctx) => [
      {
        labelKey: 'Configuration File',
        filePath: '~/.codex/config.toml',
        language: 'toml',
        value: codexConfig(ctx),
      },
      {
        labelKey: 'Environment variables',
        language: 'bash',
        value: `export NEW_API_KEY="${apiKeyOf(ctx)}"`,
      },
    ],
  },
]
