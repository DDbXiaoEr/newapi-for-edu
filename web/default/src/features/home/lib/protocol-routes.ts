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
export const PROTOCOL_ROUTES = [
  {
    id: 'gpt-chat',
    label: 'Chat',
    method: 'POST',
    path: '/v1/chat/completions',
  },
  {
    id: 'responses',
    label: 'Responses',
    method: 'POST',
    path: '/v1/responses',
  },
  {
    id: 'claude',
    label: 'Claude',
    method: 'POST',
    path: '/v1/messages',
  },
  {
    id: 'gemini',
    label: 'Gemini',
    method: 'POST',
    path: '/v1beta/models/{model}:generateContent',
  },
] as const

export type ProtocolRoute = (typeof PROTOCOL_ROUTES)[number]
