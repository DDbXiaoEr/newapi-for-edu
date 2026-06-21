<div align="center">

![new-api](/web/public/logo.png)

# New API - University Custom Edition

🍥 **University environment custom edition based on NEWAPI - AI Gateway & Asset Management System**

###
Note: All current modifications to this project have been done by AI. The author suffered a cerebral hemorrhage last year resulting in hemiplegia, and currently only has the use of one hand, so many features have not yet been implemented.
I go to the hospital for rehabilitation every day, with limited time and energy. Without a job or income, tokens rely entirely on free new-user quotas and invitation rewards from various platforms. Development efficiency has reached its limit. Those with abundant tokens are welcome to contribute)
###Thanks to GLM and Alibaba Cloud Bailian

> ⚠️ **Current Progress**: Only campus account integration (student/faculty ID login) has been implemented. All other university custom features listed below are planned but NOT YET developed. Contributions are welcome!

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
> - This project is a university environment custom version based on **NEWAPI**, specially optimized for existing campus environments
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

### 🏫 Educational Scenario Optimization

| Feature Module | Status | Description |
|---------|--------|------|
| 🎓 Educational Account Management | ✅ Done | Supports unified authentication via student/faculty ID, integrated with campus systems |
| 🔐 Campus Security Integration | 🔜 Planned | Supports campus unified identity authentication systems |

### 🛠️ Technical Customization

| Custom Item | Status | Description |
|---------|--------|------|
| 🏗️ Network Adaptation | 🔜 Planned | Optimized for campus intranet environments, supporting proxy and firewall configuration |
| 💾 Database Compatibility | 🔜 Planned | Deep optimization for commonly used databases (MySQL, PostgreSQL, SQLite) |
| 🔄 Interface Adaptation | 🔜 Planned | Provides standard interfaces with other campus systems |
| 📱 Mobile Adaptation | 🔜 Planned | Optimized mobile access experience, supporting campus app integration |

---

## 🚀 Quick Start

### Using Docker Compose (Recommended)

```bash
# Clone the project
git clone [project address]
cd [project directory]

# Edit docker-compose.yml configuration
nano docker-compose.yml

# Start the service
docker-compose up -d
```

<details>
<summary><strong>Using Docker Command</strong></summary>

```bash
# Pull the latest image
docker pull [custom image name]

# Using SQLite (default)
docker run --name new-api-edu -d --restart always \
  -p 3000:3000 \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  [custom image name]:latest

# Using MySQL
docker run --name new-api-edu -d --restart always \
  -p 3000:3000 \
  -e SQL_DSN="root:123456@tcp(localhost:3306)/oneapi" \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  [custom image name]:latest
```

> **💡 Tip:** `-v ./data:/data` will store data in the `data` folder of the current directory. You can also use an absolute path like `-v /your/custom/path:/data`

</details>

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
**Custom Feature Documentation:**
- University account management configuration
- Campus system integration guide
- Educational permission settings
- Data statistics and analysis features

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

### 🎓 University Custom New Features

- 🏫 **Campus Identity Authentication** (✅ Done): Supports student/faculty ID login
- 📚 **Course Management** (🔜 Planned): Assign AI usage permissions by course
- 👨‍🏫 **Teacher Management** (🔜 Planned): Teachers can manage class students
- 📊 **Educational Statistics** (🔜 Planned): AI usage data analysis
- 🔐 **Security Audit** (🔜 Planned): Complete operation logs and audit features

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

### 📋 Deployment Requirements

| Component | Requirement |
|------|------|
| **Local Database** | SQLite (Docker needs to mount `/data` directory) |
| **Remote Database** | MySQL ≥ 5.7.8 or PostgreSQL ≥ 9.6 |
| **Container Engine** | Docker / Docker Compose |
| **Network Environment** | Supports campus intranet environment configuration |

### ⚙️ University Custom Environment Variables

| Variable | Description | Default |
|--------|------|--------|
| `EDU_MODE` | Enable university mode | `true` |
| `CAMPUS_AUTH_URL` | Campus authentication address | - |
| `CAMPUS_API_KEY` | Campus API key | - |
| `EDU_DOMAIN` | Education domain restriction | - |
| `ALLOWED_DOMAINS` | List of allowed domains | - |
| `JWT_SECRET` | Secret key for JWT token signing (defaults to `SESSION_SECRET` if not set) | `uuid` |
| `JWT_EXPIRATION_SECONDS` | JWT token expiration time in seconds (default 7 days) | `604800` |

📖 **Full Configuration**: Please refer to NEWAPI environment variable documentation + university custom configuration guide

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