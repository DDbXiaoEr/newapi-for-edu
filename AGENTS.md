# AGENTS.md — Project Conventions for new-api

## Rule 0: DO NOT Explore Project Structure — Read from PROJECT_STRUCTURE.md

- `PROJECT_STRUCTURE.md` contains the complete project file tree.
- **MUST read `PROJECT_STRUCTURE.md` first** to understand the project layout. Never run `find`/`ls -R`/`tree` or glob patterns to explore the entire project directory tree.
- When adding, deleting, or restructuring files/directories, **MUST update `PROJECT_STRUCTURE.md`** to keep it in sync.

## Overview

AI API gateway/proxy built with Go. Aggregates 40+ upstream AI providers (OpenAI, Claude, Gemini, Azure, AWS Bedrock, etc.) behind a unified OpenAI-compatible API, with user management, billing, rate limiting, and an admin dashboard.

## Tech Stack

- **Backend**: Go 1.22+, Gin web framework, GORM v2 ORM
- **Frontend**: React 19, TypeScript, Rsbuild + TanStack Router (`web/default/`)
- **Databases**: MySQL, PostgreSQL (both must be supported); ClickHouse is optional for logs only
- **Cache**: Redis (go-redis) + in-memory cache
- **Auth**: JWT, WebAuthn/Passkeys, OAuth (GitHub, Discord, OIDC, etc.)
- **Frontend package manager**: Bun

## Architecture

Layered architecture: Router -> Controller -> Service -> Model

| Layer | Directory | Description |
|-------|-----------|-------------|
| Entry | `main.go`, `main-backend.go` | Application startup |
| Router | `router/` | HTTP routing (API, relay, dashboard, web, video) |
| Controller | `controller/` | Request handlers |
| Service | `service/` | Business logic |
| Model | `model/` | GORM data models |
| Relay | `relay/` | AI provider adapters (40+ providers in `relay/channel/`) |
| Middleware | `middleware/` | Auth, rate limiting, CORS, logging, distribution |
| Setting | `setting/` | Configuration management |
| Common | `common/` | Shared utilities (JSON, crypto, Redis, rate-limit) |
| DTO | `dto/` | Data transfer objects |
| Constant | `constant/` | Constants and enums |
| Types | `types/` | Core type definitions |
| Frontend | `web/default/` (Rsbuild+TS+TanStack) | React frontend embedded by `main.go` |

## Key Conventions

### JSON Package — Use `common/json.go`

All JSON marshal/unmarshal MUST use wrappers in `common/json.go`:
- `common.Marshal(v any) ([]byte, error)`
- `common.Unmarshal(data []byte, v any) error`
- `common.UnmarshalJsonStr(data string, v any) error`
- `common.DecodeJson(reader io.Reader, v any) error`
- `common.GetJsonType(data json.RawMessage) string`

Do NOT call `encoding/json` marshal/unmarshal directly. `json.RawMessage` etc. as types are fine.

### Database Compatibility — MySQL >= 5.7.8, PostgreSQL >= 9.6

- Prefer GORM methods over raw SQL.
- Use `commonGroupCol`, `commonKeyCol` for reserved-word columns.
- Use `commonTrueVal`/`commonFalseVal` instead of hardcoded booleans.
- Use `common.UsingMainDatabase(common.DatabaseTypePostgreSQL)` / `common.UsingMainDatabase(common.DatabaseTypeMySQL)` for DB-specific branching.
- Avoid DB-specific functions/operators without fallback.

### Frontend — Prefer Bun

Use `bun` for the frontend (`web/default/`):
- `bun install`, `bun run dev`, `bun run build`, `bun run i18n:*`

### New Channel StreamOptions

When implementing a new channel, confirm StreamOptions support and add to `streamSupportedChannels` if supported.

### Protected Project Information — DO NOT Modify

The following are strictly protected:
- References to **nеw-аρi** (project name/identity)
- References to **QuаntumΝоuѕ** (organization/author identity)
- README files, license headers, copyright, package metadata, Docker image names, etc.

### Upstream Relay Request DTOs — Preserve Explicit Zero Values

For request structs parsed from client JSON and re-marshaled upstream:
- Optional scalar fields MUST use pointer types with `omitempty` (e.g. `*int`, `*bool`).
- Field absent => `nil` => omitted. Field explicitly zero/false => non-`nil` pointer => sent upstream.

## Internationalization (i18n)

### Backend (`i18n/`)
- Library: `nicksnyder/go-i18n/v2`
- Languages: en, zh

### Frontend (`web/default/src/i18n/`)
- Library: `i18next` + `react-i18next`
- Languages: zh (fallback), en, fr, ru, ja, vi
- Translation files: `web/default/src/i18n/locales/{lang}.json` (flat JSON, keys are Chinese source strings)
- Usage: `useTranslation()` hook, call `t('中文key')`
- CLI tools: `bun run i18n:extract`, `bun run i18n:sync`, `bun run i18n:lint`
