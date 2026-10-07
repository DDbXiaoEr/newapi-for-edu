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
import {
  Bot,
  Code2,
  Globe,
  MessageSquare,
  SquareTerminal,
  type LucideIcon,
} from 'lucide-react'

import type { ClientOption } from '../lib/clients'

const BUILTIN_ICONS: Record<string, LucideIcon> = {
  opencode: SquareTerminal,
  'claude-code': Bot,
  codex: Code2,
}

export function getClientIcon(client: ClientOption): LucideIcon {
  if (client.kind === 'preset') {
    return client.preset.type === 'web' ? Globe : MessageSquare
  }
  return BUILTIN_ICONS[client.builtin.id] ?? SquareTerminal
}
