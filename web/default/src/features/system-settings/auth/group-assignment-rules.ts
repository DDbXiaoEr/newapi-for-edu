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

export const GROUP_FILTER_ATTRIBUTES = [
  { value: 'uid', labelKey: 'Username' },
  { value: 'mail', labelKey: 'Email' },
  { value: 'cn', labelKey: 'Display Name' },
  { value: 'dn', labelKey: 'DN' },
] as const

export type GroupAssignmentRule = {
  attribute: string
  pattern: string
  group: string
}

export type GroupAssignmentEditorValue = {
  attribute: string
  rules: Array<{ id: string; pattern: string; group: string }>
}

function parseJsonRules(value: string): GroupAssignmentRule[] | null {
  try {
    const parsed = JSON.parse(value) as unknown
    if (!Array.isArray(parsed)) {
      return null
    }
    return parsed
      .map((rule) => ({
        attribute: String(
          (rule as GroupAssignmentRule)?.attribute ?? ''
        ).trim(),
        pattern: String((rule as GroupAssignmentRule)?.pattern ?? '').trim(),
        group: String((rule as GroupAssignmentRule)?.group ?? '').trim(),
      }))
      .filter((rule) => rule.attribute && rule.pattern && rule.group)
  } catch {
    return null
  }
}

export function parseGroupAssignmentRules(
  value: string | GroupAssignmentRule[] | null | undefined
): GroupAssignmentRule[] {
  if (Array.isArray(value)) {
    return value
      .map((rule) => ({
        attribute: String(rule?.attribute ?? '').trim(),
        pattern: String(rule?.pattern ?? '').trim(),
        group: String(rule?.group ?? '').trim(),
      }))
      .filter((rule) => rule.attribute && rule.pattern && rule.group)
  }
  const text = String(value ?? '').trim()
  if (!text) {
    return []
  }
  if (text.startsWith('[')) {
    return parseJsonRules(text) ?? []
  }
  return text
    .split('\n')
    .map((line) => line.trim())
    .filter((line) => line && !line.startsWith('#'))
    .flatMap((line) => {
      const parts = line.split('|').map((part) => part.trim())
      if (parts.length < 3) {
        return []
      }
      const [attribute, pattern, group] = parts
      if (!attribute || !pattern || !group) {
        return []
      }
      return [{ attribute, pattern, group }]
    })
}

export function editorValueFromRules(
  value: string | GroupAssignmentRule[] | null | undefined
): GroupAssignmentEditorValue {
  const rules = parseGroupAssignmentRules(value)
  const attribute = rules[0]?.attribute || 'uid'
  return {
    attribute,
    rules:
      rules.length > 0
        ? rules.map((rule, index) => ({
            id: `${rule.attribute}-${rule.pattern}-${rule.group}-${index}`,
            pattern: rule.pattern,
            group: rule.group,
          }))
        : [{ id: 'empty-0', pattern: '', group: '' }],
  }
}

export function stringifyGroupAssignmentEditorValue(
  value: GroupAssignmentEditorValue
): string {
  const attribute = value.attribute.trim() || 'uid'
  return JSON.stringify(
    value.rules
      .map((rule) => ({
        attribute,
        pattern: rule.pattern.trim(),
        group: rule.group.trim(),
      }))
      .filter((rule) => rule.pattern && rule.group)
  )
}
