<div align="center">

![new-api](/web/public/logo.png)

# New API - University Custom Edition

🍥 **University environment custom edition based on NEWAPI - AI Gateway & Asset Management System**

###
Note: All current modifications to this project have been done by AI. The author suffered a cerebral hemorrhage last year resulting in hemiplegia, and currently only has the use of one hand, so many features have not yet been implemented.
I go to the hospital for rehabilitation every day, with limited time and energy. Without a job or income, tokens rely entirely on free new-user quotas and invitation rewards from various platforms. Development efficiency has reached its limit. Those with abundant tokens are welcome to contribute)
###Thanks to GLM and Alibaba Cloud Bailian

> ⚠️ **Current Progress**: Campus LDAP / CAS login, group-channel bindings, tokens locked to the account group, JWT sessions, syslog, and Kubernetes / Docker Compose deployment are implemented. Course and class-management features are still planned. Contributions are welcome!

<p align="center">
  <a href="./README.zh_CN.md">简体中文</a> |
  <a href="./README.zh_TW.md">繁體中文</a> |
  <strong>English</strong> |
  <a href="./README.fr.md">Français</a> |
  <a href="./README.ja.md">日本語</a>
</p>

<p align="center">
  <a href="https://raw.githubusercontent.com/Calcium-Ion/new-api/main/LICENSE">
    <img src="https://img.shields.io/github/license/Calcium-Ion/new-api?color=brightgreen" alt="license">
  </a><!--
  --><a href="https://github.com/Calcium-Ion/new-api/releases/latest">
    <img src="https://img.shields.io/github/v/release/Calcium-Ion/new-api?color=brightgreen&include_prereleases" alt="release">
  </a><!--
  --><a href="https://hub.docker.com/r/CalciumIon/new-api">
    <img src="https://img.shields.io/badge/docker-dockerHub-blue" alt="docker">
  </a>
</p>

<p align="center">
  <a href="#-quick-start">Quick Start</a> •
  <a href="#-project-features">Features</a> •
  <a href="#-deployment">Deployment</a> •
  <a href="#-university-custom-features">University Features</a> •
  <a href="#-help-support">Help</a>
</p>

</div>

## 📝 Project Description

> [!IMPORTANT]
> - The repository name includes `edu` and the docs are written for campuses, but the core is still a general-purpose AI API gateway. Enterprises, research labs, and other industries can deploy it as-is; it is not locked to education workflows.
> - This project is a university environment custom version based on **NEWAPI**, specially optimized for existing campus environments. Campus LDAP / CAS, group-channel bindings, and tokens locked to the account group also map cleanly to enterprise LDAP/AD, SSO, and access isolation.
> - This project is for personal learning purposes only, with no guarantee of stability and no technical support
> - Users must comply with OpenAI's [Terms of Use](https://openai.com/policies/terms-of-use) and applicable **laws and regulations**, and must not use it for illegal purposes
> - In accordance with the [Interim Measures for the Management of Generative Artificial Intelligence Services](http://www.cac.gov.cn/2023-07/13/c_1690898327029107.htm), please do not provide any unregistered generative AI services to the public in China

---

## 🏗️ Project Architecture

This project is developed on top of **NEWAPI** ([GitHub - Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api)), retaining all core features of the original project while being deeply customized for university environments.

### 📚 Original Project Documentation
- **NEWAPI Official Documentation**: [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs)
- **NEWAPI GitHub**: [https://github.com/Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api)
- **NEWAPI Deployment Guide**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

---

## 🎯 University Custom Features

Deltas versus upstream [NEWAPI](https://github.com/Calcium-Ion/new-api). Everything else (model access, billing, console) is inherited.

### 🏫 Campus identity and access

| Feature | Status | Description |
|---------|--------|-------------|
| LDAP login | ✅ Done | Campus LDAP / AD; server URL, Bind DN, user filter, attribute mapping, StartTLS |
| CAS login | ✅ Done | Campus CAS SSO; attribute mapping and access-attribute restrictions |
| Directory group assignment | ✅ Done | Map LDAP / CAS attributes to existing groups with regex rules |
| Group-channel bindings | ✅ Done | Pin channels per group+model; unbound models keep upstream routing |
| Tokens locked to account group | ✅ Done | Create/update forces the account group; users cannot pick a group or cross-group retry |
| Batch assign users to group | ✅ Done | Admins can move selected users into a target group |
| Hide user group in UI | ✅ Done | Profile, keys, and model square hide group/user ID; model square shows only the current group's models |
| Course / class management | 🔜 Planned | Per-course grants and teacher class management |

### 🛠️ Deployment and operations

| Feature | Status | Description |
|---------|--------|-------------|
| No SQLite | ✅ Done | MySQL ≥ 5.7.8 or PostgreSQL ≥ 9.6 only; `SQL_DSN` is required at startup |
| JWT sessions | ✅ Done | Cookie sessions removed; console auth is JWT (`JWT_SECRET` / `JWT_EXPIRATION_SECONDS`) |
| Syslog | ✅ Done | Optional remote/local syslog; dial timeout so an unreachable server cannot block startup |
| Docker Compose stack | ✅ Done | `docker-compose/` ships PostgreSQL, Redis, ClickHouse, OpenLDAP, and this service |
| Kubernetes | ✅ Done | `kubernetes/` includes Deployment, Service, HPA, ConfigMap/Secret, and external dependency manifests |
| amd64 / arm64 builds | ✅ Done | `make` cross-compiles Linux amd64/arm64, plus pure-backend and all-in-one images |

---

## 🚀 Quick Start

### Using Docker Compose (Recommended)

`docker-compose/` starts **New API Edu + PostgreSQL + Redis + ClickHouse + OpenLDAP**:

```bash
git clone https://gitee.com/ddbxiaoer/newapi_2_-edu.git
cd newapi_2_-edu

# Build the all-in-one image (frontend embedded)
make docker-allinone

# Change passwords and JWT_SECRET in docker-compose/docker-compose.yml first
docker compose -f docker-compose/docker-compose.yml up -d
```

Pure-backend image (no embedded frontend): `make docker-backend` → `newapi-edu-pure`.

> **⚠️ Change** PostgreSQL / Redis / LDAP passwords and `JWT_SECRET` before production. Startup requires `SQL_DSN` and `LOG_SQL_DSN` (set `REQUIRED_ENV_VARS=""` to skip the check).

---

🎉 After deployment, visit `http://localhost:3000` to start using!

---

## 📚 Documentation

### 📖 Basic Documentation (Based on NEWAPI)
Since this project is developed on top of NEWAPI, most basic functionality documentation can be referenced directly:
- **NEWAPI Official Documentation**: [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs)
- **NEWAPI API Documentation**: [https://docs.newapi.pro/zh/docs/api](https://docs.newapi.pro/zh/docs/api)
- **NEWAPI Deployment Guide**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

### 🎓 University Custom Edition Documentation
- LDAP / CAS: System Settings → Authentication
- Group-channel bindings: billing / group settings
- Deployment manifests: `docker-compose/`, `kubernetes/`

---

## ✨ Key Features

> For detailed features, please refer to **NEWAPI Official Documentation**: [https://docs.newapi.pro/zh/docs/guide/wiki/basic-concepts/features-introduction](https://docs.newapi.pro/zh/docs/guide/wiki/basic-concepts/features-introduction)

### 🎨 Core Features (Inherited from NEWAPI)

| Feature | Description |
|------|------|
| 🎨 Brand New UI | Modern user interface design |
| 🌍 Multilingual | Supports Chinese, English, French, Japanese |
| 🔄 Data Compatibility | Fully compatible with the original One API database |
| 📈 Data Dashboard | Visual console and statistical analysis |
| 🔒 Permission Management | Token grouping, model restrictions, user management |

### 💰 Payment & Billing (Inherited from NEWAPI)

- ✅ Online top-up (EPay, Stripe)
- ✅ Per-use model pricing
- ✅ Cache billing support (OpenAI, Azure, DeepSeek, Claude, Qwen and all supported models)
- ✅ Flexible billing policy configuration

### 🎓 Deltas versus upstream

- 🏫 **LDAP / CAS campus login**: student/faculty ID against campus SSO
- 🧭 **Directory group assignment**: map LDAP / CAS attributes onto existing groups
- 🔗 **Group-channel bindings**: pin channels per group+model; unbound models keep upstream routing
- 🔒 **Tokens locked to account group**: users cannot pick a token group; model square shows only the current group's models
- 👥 **Batch assign users to group**: admins can move selected users into a target group
- 🧾 **JWT + syslog**: cookie sessions removed; optional syslog with a dial timeout
- 🚢 **Campus deployment**: no SQLite; Docker Compose stack and Kubernetes + HPA
- 📚 **Course / class management** (🔜 Planned)

---

## 🤖 Model Support

> For details, please refer to **NEWAPI API Documentation**: [https://docs.newapi.pro/zh/docs/api](https://docs.newapi.pro/zh/docs/api)

| Model Type | Description | Documentation |
|---------|------|------|
| 🤖 OpenAI-Compatible | OpenAI compatible models | [NEWAPI Docs](https://docs.newapi.pro/zh/docs/api/ai-model/chat/openai/createchatcompletion) |
| 🤖 OpenAI Responses | OpenAI Responses format | [NEWAPI Docs](https://docs.newapi.pro/zh/docs/api/ai-model/chat/openai/createresponse) |
| 🎨 Midjourney-Proxy | Midjourney proxy | [NEWAPI Docs](https://doc.newapi.pro/api/midjourney-proxy-image) |
| 🎵 Suno-API | Suno music generation | [NEWAPI Docs](https://doc.newapi.pro/api/suno-music) |
| 💬 Claude | Claude Messages format | [NEWAPI Docs](https://docs.newapi.pro/zh/docs/api/ai-model/chat/createmessage) |
| 🌐 Gemini | Google Gemini | [NEWAPI Docs](https://docs.newapi.pro/zh/docs/api/ai-model/chat/gemini/geminirelayv1beta) |

---

## 🚢 Deployment

> [!TIP]
> **Basic Deployment Method**: Please refer to **NEWAPI Official Deployment Guide**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

### 🛠️ Build from Source

```bash
# Clone the project
git clone [project address]
cd [project directory]

# Install dependencies
make prepare

# Build for current platform
make build

# Build for Linux (cross-compile)
make build-backend-linux          # Linux amd64
make build-backend-linux-arm64    # Linux arm64

# Build backend only (no embedded frontend)
make build-backend-pure           # Current platform
make build-backend-pure-linux     # Linux amd64
make build-backend-pure-linux-arm64  # Linux arm64
```

### 📋 Deployment Requirements

| Component | Requirement |
|------|------|
| **Main database** | MySQL ≥ 5.7.8 or PostgreSQL ≥ 9.6 (`SQL_DSN` **required**; SQLite removed) |
| **Log database** | `LOG_SQL_DSN`; supports a separate DB / ClickHouse |
| **Cache** | Redis (required when nodes share rate limits) |
| **Containers / orchestration** | Docker Compose (`docker-compose/`) or Kubernetes (`kubernetes/`, includes HPA) |

Kubernetes example:

```bash
kubectl apply -k kubernetes/
```

### ⚙️ University Custom Environment Variables

| Variable | Description | Default |
|--------|------|--------|
| `SQL_DSN` | Main database DSN (required) | - |
| `LOG_SQL_DSN` | Log database DSN (required by default) | - |
| `REQUIRED_ENV_VARS` | Startup required-var list; empty string disables the check | `SQL_DSN,LOG_SQL_DSN` |
| `JWT_SECRET` | JWT signing secret (falls back to `SESSION_SECRET`) | `uuid` |
| `JWT_EXPIRATION_SECONDS` | JWT TTL in seconds | `604800` |
| `SYSLOG_ENABLED` | Enable syslog | `false` |
| `SYSLOG_NETWORK` | syslog protocol (`udp`/`tcp`; empty = local socket) | - |
| `SYSLOG_ADDR` | Remote syslog address | - |
| `SYSLOG_TAG` | syslog tag | `newapi` |

LDAP / CAS are configured in the console under System Settings → Authentication, not via the env vars above.

📖 **Full configuration**: NEWAPI environment-variable docs plus `docker-compose/` and `kubernetes/` in this repo

---

## 🔗 Related Projects

### Upstream Projects

| Project | Description |
|------|------|
| [NEWAPI](https://github.com/Calcium-Ion/new-api) | **Base Project** - This project is developed on top of this |
| [One API](https://github.com/songquanpeng/one-api) | Foundation project of NEWAPI |

### Companion Tools

| Project | Description |
|------|------|
| [neko-api-key-tool](https://github.com/Calcium-Ion/neko-api-key-tool) | Key quota query tool |
| [new-api-horizon](https://github.com/Calcium-Ion/new-api-horizon) | NEWAPI high-performance optimized edition |

---

## 💬 Help & Support

### 📖 Documentation Resources

| Resource | Link |
|------|------|
| 📘 **NEWAPI Official Documentation** | [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs) |
| 🎓 **University Custom Documentation** | Project documentation directory |
| 💬 **Community** | [NEWAPI Community Channels](https://docs.newapi.pro/zh/docs/support/community-interaction) |
| 🐛 **Issue Feedback** | [NEWAPI Issue Feedback](https://github.com/Calcium-Ion/new-api/issues) |

### 🤝 Contribution Guide

All forms of contributions are welcome!

- 🐛 Report bugs
- 💡 Propose new features
- 📝 Improve documentation
- 🔧 Submit code

---

## 📜 License

This project is licensed under the [GNU Affero General Public License v3.0 (AGPLv3)](./LICENSE).

This is an open-source project developed on top of **NEWAPI** ([https://github.com/Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api)) (based on MIT license).

If your organization's policies do not permit the use of AGPLv3-licensed software, or if you wish to avoid the open-source obligations of AGPLv3, please contact us at: [support@quantumnous.com](mailto:support@quantumnous.com)

---

## 🌟 Star History

<div align="center">

[![Star History Chart](https://api.star-history.com/svg?repos=Calcium-Ion/new-api&type=Date)](https://star-history.com/#Calcium-Ion/new-api&Date)

</div>

---

<div align="center">

### 💖 Thank You for Using New API University Custom Edition

If this project helps you, please give us a ⭐️ Star!

**[NEWAPI Official Documentation](https://docs.newapi.pro/zh/docs)** • **[Issue Feedback](https://github.com/Calcium-Ion/new-api/issues)** • **[Latest Release](https://github.com/Calcium-Ion/new-api/releases)**

<sub>Built on NEWAPI ❤️ QuantumNous</sub>

</div>