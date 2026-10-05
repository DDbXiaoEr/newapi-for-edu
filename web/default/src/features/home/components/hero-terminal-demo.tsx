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
import { useState, useEffect, useRef, type ReactNode } from 'react'
import { useTranslation } from 'react-i18next'

import { cn } from '@/lib/utils'

import { PROTOCOL_ROUTES } from '../lib/protocol-routes'

interface ApiDemoConfig {
  id: string
  label: string
  method: 'POST' | 'GET'
  endpoint: string
  headers: string[]
  request: string[]
  response: string[]
  tokens: number
  latency: number
}

const API_DEMOS: ApiDemoConfig[] = [
  {
    id: PROTOCOL_ROUTES[0].id,
    label: PROTOCOL_ROUTES[0].label,
    method: PROTOCOL_ROUTES[0].method,
    endpoint: PROTOCOL_ROUTES[0].path,
    headers: ['"Authorization: Bearer sk-••••"'],
    request: [
      '"model": "your-model",',
      '"messages": [',
      '  { "role": "user", "content": "..." }',
      ']',
    ],
    response: [
      '{',
      '  "choices": [{ "message": { "content": <text> } }],',
      '  "usage": { "total_tokens": <tokens> }',
      '}',
    ],
    tokens: 27,
    latency: 142,
  },
  {
    id: PROTOCOL_ROUTES[1].id,
    label: PROTOCOL_ROUTES[1].label,
    method: PROTOCOL_ROUTES[1].method,
    endpoint: PROTOCOL_ROUTES[1].path,
    headers: ['"Authorization: Bearer sk-••••"'],
    request: ['"model": "your-model",', '"input": "..."'],
    response: [
      '{',
      '  "output": [{ "type": "output_text", "text": <text> }],',
      '  "usage": { "total_tokens": <tokens> }',
      '}',
    ],
    tokens: 31,
    latency: 168,
  },
  {
    id: PROTOCOL_ROUTES[2].id,
    label: PROTOCOL_ROUTES[2].label,
    method: PROTOCOL_ROUTES[2].method,
    endpoint: PROTOCOL_ROUTES[2].path,
    headers: ['"x-api-key: sk-••••"', '"anthropic-version: 2023-06-01"'],
    request: [
      '"model": "your-model",',
      '"max_tokens": 1024,',
      '"messages": [',
      '  { "role": "user", "content": "..." }',
      ']',
    ],
    response: [
      '{',
      '  "content": [{ "type": "text", "text": <text> }],',
      '  "usage": { "input_tokens": <in>, "output_tokens": <out> }',
      '}',
    ],
    tokens: 29,
    latency: 156,
  },
  {
    id: PROTOCOL_ROUTES[3].id,
    label: PROTOCOL_ROUTES[3].label,
    method: PROTOCOL_ROUTES[3].method,
    endpoint: PROTOCOL_ROUTES[3].path,
    headers: ['"x-goog-api-key: sk-••••"'],
    request: [
      '"contents": [',
      '  { "role": "user",',
      '    "parts": [{ "text": "..." }] }',
      ']',
    ],
    response: [
      '{',
      '  "candidates": [{ "content": { "parts": [{ "text": <text> }] } }],',
      '  "usageMetadata": { "totalTokenCount": <tokens> }',
      '}',
    ],
    tokens: 25,
    latency: 93,
  },
]

const CYCLE_INTERVAL = 4500
const TRANSITION_MS = 220

const RESPONSE_TEXT: Record<string, string> = {
  'gpt-chat': 'Chat request routed.',
  responses: 'Response workflow ready.',
  claude: 'Claude message routed.',
  gemini: 'Gemini request served.',
}

interface HeroTerminalDemoProps {
  className?: string
}

export function HeroTerminalDemo(props: HeroTerminalDemoProps) {
  const { t } = useTranslation()
  const [activeIndex, setActiveIndex] = useState(0)
  const [transitioning, setTransitioning] = useState(false)
  const intervalRef = useRef<ReturnType<typeof setInterval>>(undefined)
  const timeoutRef = useRef<ReturnType<typeof setTimeout>>(undefined)

  useEffect(() => {
    const mq = window.matchMedia('(prefers-reduced-motion: reduce)')
    if (mq.matches) return

    intervalRef.current = setInterval(() => {
      setTransitioning(true)
      timeoutRef.current = setTimeout(() => {
        setActiveIndex((prev) => (prev + 1) % API_DEMOS.length)
        setTransitioning(false)
      }, TRANSITION_MS)
    }, CYCLE_INTERVAL)

    return () => {
      if (intervalRef.current) clearInterval(intervalRef.current)
      if (timeoutRef.current) clearTimeout(timeoutRef.current)
    }
  }, [])

  const handleSelect = (index: number) => {
    if (index === activeIndex) return
    if (intervalRef.current) clearInterval(intervalRef.current)
    if (timeoutRef.current) clearTimeout(timeoutRef.current)
    setTransitioning(true)
    timeoutRef.current = setTimeout(() => {
      setActiveIndex(index)
      setTransitioning(false)
    }, TRANSITION_MS)
  }

  const demo = API_DEMOS[activeIndex]

  return (
    <div className={cn('mx-auto w-full max-w-2xl', props.className)}>
      <div className='border-border bg-card overflow-hidden rounded-lg border'>
        <div
          role='tablist'
          aria-label={t(
            'Compatible API routes for common AI application workflows'
          )}
          className='border-border flex items-center gap-1 border-b px-2 sm:gap-1.5 sm:px-3'
        >
          {API_DEMOS.map((item, index) => {
            const isActive = index === activeIndex
            return (
              <button
                key={item.id}
                type='button'
                role='tab'
                aria-selected={isActive}
                onClick={() => handleSelect(index)}
                className={cn(
                  'relative -mb-px border-b-2 px-2.5 py-2.5 font-mono text-[11px] font-medium transition-colors sm:px-3 sm:text-xs',
                  isActive
                    ? 'border-primary text-primary'
                    : 'text-muted-foreground hover:text-foreground border-transparent'
                )}
              >
                {item.label}
              </button>
            )
          })}
          <div className='ml-auto flex items-center gap-2 pr-2 sm:pr-3'>
            <span className='bg-success inline-block size-1.5 rounded-full' />
            <span className='text-muted-foreground font-mono text-[10px] tracking-wide'>
              200 ok
            </span>
          </div>
        </div>

        <div className='border-border flex items-center gap-2.5 border-b px-5 py-3'>
          <span className='bg-primary/10 text-primary rounded-md px-1.5 py-0.5 font-mono text-[10px] font-semibold tracking-wider'>
            {demo.method}
          </span>
          <code
            data-testid='protocol-endpoint'
            className={cn(
              'text-foreground/80 truncate font-mono text-[12.5px] transition-opacity duration-200',
              transitioning ? 'opacity-0' : 'opacity-100'
            )}
          >
            {demo.endpoint}
          </code>
        </div>

        <div className='grid h-[400px] grid-rows-[235px_minmax(0,1fr)] font-mono text-[12.5px] leading-[1.55]'>
          <RequestBlock demo={demo} transitioning={transitioning} />
          <ResponseBlock demo={demo} transitioning={transitioning} />
        </div>

        <div className='border-border bg-muted/30 flex items-center justify-between border-t px-5 py-2.5'>
          <div className='text-muted-foreground flex items-center gap-3 text-[10px] tabular-nums'>
            <span className='flex items-center gap-1'>
              <span className='font-mono'>{demo.latency}</span>
              <span>ms</span>
            </span>
            <span className='bg-border size-1 rounded-full' />
            <span className='flex items-center gap-1'>
              <span className='font-mono'>{demo.tokens}</span>
              <span>tokens</span>
            </span>
            <span className='bg-border size-1 rounded-full' />
            <span className='flex items-center gap-1'>
              <span>cost</span>
              <span className='font-mono'>
                ${(demo.tokens * 0.00003).toFixed(5)}
              </span>
            </span>
          </div>
          <span className='text-muted-foreground/70 font-mono text-[10px]'>
            stream · sse
          </span>
        </div>
      </div>
    </div>
  )
}

function RequestBlock(props: { demo: ApiDemoConfig; transitioning: boolean }) {
  const { demo, transitioning } = props

  return (
    <div className='relative px-5 py-4'>
      <SectionLabel>Request</SectionLabel>
      <div
        className={cn(
          'mt-2 transition-opacity duration-200',
          transitioning ? 'opacity-0' : 'opacity-100'
        )}
      >
        <CodeLine>
          <Command>curl</Command> <Flag>-X</Flag> <Flag>POST</Flag>{' '}
          <StringText>&quot;{demo.endpoint}&quot;</StringText>{' '}
          <Muted>{'\\'}</Muted>
        </CodeLine>
        {demo.headers.map((header) => (
          <CodeLine key={header} indent={2}>
            <Flag>-H</Flag> <StringText>{header}</StringText>{' '}
            <Muted>{'\\'}</Muted>
          </CodeLine>
        ))}
        <CodeLine indent={2}>
          <Flag>-d</Flag> <StringText>&apos;{'{'}</StringText>
        </CodeLine>
        {demo.request.map((line) => (
          <CodeLine key={line} indent={4}>
            {renderJsonLine(line)}
          </CodeLine>
        ))}
        <CodeLine indent={2}>
          <StringText>{'}'}&apos;</StringText>
        </CodeLine>
      </div>
    </div>
  )
}

function ResponseBlock(props: { demo: ApiDemoConfig; transitioning: boolean }) {
  const { demo, transitioning } = props

  return (
    <div className='border-border bg-muted/20 relative border-t px-5 py-4'>
      <SectionLabel>Response</SectionLabel>
      <div
        className={cn(
          'mt-2 transition-opacity duration-200',
          transitioning ? 'opacity-0' : 'opacity-100'
        )}
      >
        {demo.response.map((line) => (
          <CodeLine key={line}>{renderResponseLine(line, demo)}</CodeLine>
        ))}
      </div>
    </div>
  )
}

function SectionLabel(props: { children: ReactNode }) {
  return (
    <span className='text-muted-foreground font-sans text-[10px] font-medium'>
      {props.children}
    </span>
  )
}

const STRING_RE = /"[^"]*"/g
const PLACEHOLDER_RE = /<[a-z]+>/gi

function renderJsonLine(line: string): ReactNode {
  if (!line.trim()) return <Muted> </Muted>
  return tokenize(line)
}

function renderResponseLine(line: string, demo: ApiDemoConfig): ReactNode {
  if (!line.trim()) return <Muted> </Muted>

  const segments: ReactNode[] = []
  let cursor = 0
  const matches = [...line.matchAll(PLACEHOLDER_RE)]

  if (matches.length === 0) return tokenize(line)

  matches.forEach((match) => {
    const start = match.index ?? 0
    if (start > cursor) {
      segments.push(
        <span key={`pre-${start}`}>{tokenize(line.slice(cursor, start))}</span>
      )
    }
    const placeholder = match[0]
    if (placeholder === '<text>') {
      segments.push(
        <Accent key={`ph-${start}`}>
          {`"${RESPONSE_TEXT[demo.id] ?? '...'}"`}
        </Accent>
      )
    } else if (placeholder === '<tokens>') {
      segments.push(<NumberText key={`ph-${start}`}>{demo.tokens}</NumberText>)
    } else if (placeholder === '<in>') {
      segments.push(
        <NumberText key={`ph-${start}`}>
          {Math.floor(demo.tokens * 0.4)}
        </NumberText>
      )
    } else if (placeholder === '<out>') {
      segments.push(
        <NumberText key={`ph-${start}`}>
          {Math.ceil(demo.tokens * 0.6)}
        </NumberText>
      )
    } else {
      segments.push(<Muted key={`ph-${start}`}>{placeholder}</Muted>)
    }
    cursor = start + placeholder.length
  })

  if (cursor < line.length) {
    segments.push(<span key='tail'>{tokenize(line.slice(cursor))}</span>)
  }

  return segments
}

function tokenize(input: string): ReactNode {
  const segments: ReactNode[] = []
  let cursor = 0
  const matches = [...input.matchAll(STRING_RE)]

  matches.forEach((match) => {
    const start = match.index ?? 0
    if (start > cursor) {
      segments.push(
        <Muted key={`m-${start}`}>{input.slice(cursor, start)}</Muted>
      )
    }
    const text = match[0]
    const after = input.slice(start + text.length).trimStart()
    const isKey = after.startsWith(':')
    if (isKey) {
      segments.push(<Key key={`k-${start}`}>{text}</Key>)
    } else {
      segments.push(<StringText key={`s-${start}`}>{text}</StringText>)
    }
    cursor = start + text.length
  })

  if (cursor < input.length) {
    segments.push(<Muted key='tail'>{input.slice(cursor)}</Muted>)
  }

  return segments
}

function CodeLine(props: { children: ReactNode; indent?: number }) {
  return (
    <div className='break-words whitespace-pre-wrap'>
      {props.indent ? (
        <span
          aria-hidden
          className='inline-block'
          style={{ width: `${props.indent}ch` }}
        />
      ) : null}
      {props.children}
    </div>
  )
}

function Command(props: { children: ReactNode }) {
  return <span className='text-chart-5 font-medium'>{props.children}</span>
}

function Flag(props: { children: ReactNode }) {
  return <span className='text-chart-1'>{props.children}</span>
}

function Key(props: { children: ReactNode }) {
  return <span className='text-chart-2'>{props.children}</span>
}

function StringText(props: { children: ReactNode }) {
  return <span className='text-chart-4'>{props.children}</span>
}

function NumberText(props: { children: ReactNode }) {
  return <span className='text-chart-3 font-medium'>{props.children}</span>
}

function Muted(props: { children: ReactNode }) {
  return <span className='text-muted-foreground'>{props.children}</span>
}

function Accent(props: { children: ReactNode }) {
  return <span className='text-primary font-medium'>{props.children}</span>
}
