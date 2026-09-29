<div align="center">

![new-api](/web/public/logo.png)

# New API - 高校定制版

🍥 **基于 NEWAPI 的高校环境定制版 - AI 网关与资产管理系统**

###
Note: 本项目目前所有的更改全部由AI完成，由于本人去年脑出血偏瘫。目前只有一只手可以用，所以这个项目有很多还没实现
由于我每天要去医院康复，时间精力有限，而且由于没工作没收入，token全靠各大平台新用户额度和邀请赠送，目前开发效率已经到极限了。token富裕的大哥可以自己上）
###在此感谢GLM 阿里云百炼

> ⚠️ **当前进度**：校园 LDAP / CAS 登录、分组渠道绑定、令牌锁定账号分组、JWT 会话、Syslog 与 Kubernetes / Docker Compose 部署已落地。课程管理、教师班级管理等教学业务仍在规划中。欢迎贡献！

<p align="center">
  <strong>简体中文</strong> |
  <a href="./README.zh_TW.md">繁體中文</a> |
  <a href="./README.md">English</a> |
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
  <a href="#-快速开始">快速开始</a> •
  <a href="#-项目特点">项目特点</a> •
  <a href="#-部署">部署</a> •
  <a href="#-高校定制功能">高校定制功能</a> •
  <a href="#-帮助支持">帮助</a>
</p>

</div>

## 📝 项目说明

> [!IMPORTANT]
> - 本项目是基于 **NEWAPI** 的高校环境定制版本，针对高校现有环境做了特别优化
> - 本项目仅供个人学习使用，不保证稳定性，且不提供任何技术支持
> - 使用者必须在遵循 OpenAI 的 [使用条款](https://openai.com/policies/terms-of-use) 以及**法律法规**的情况下使用，不得用于非法用途
> - 根据 [《生成式人工智能服务管理暂行办法》](http://www.cac.gov.cn/2023-07/13/c_1690898327029107.htm) 的要求，请勿对中国地区公众提供一切未经备案的生成式人工智能服务

---

## 🏗️ 项目架构

本项目基于 **NEWAPI** ([GitHub - Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api)) 进行二次开发，保留了原项目的所有核心功能，并针对高校环境进行了深度定制。

### 📚 原项目文档
- **NEWAPI 官方文档**: [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs)
- **NEWAPI GitHub**: [https://github.com/Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api)
- **NEWAPI 部署指南**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

---

## 🎯 高校定制功能

相对上游 [NEWAPI](https://github.com/Calcium-Ion/new-api) 的增量能力如下。未列出的能力（模型接入、计费、控制台等）仍继承上游。

### 🏫 校园身份与权限

| 功能 | 状态 | 说明 |
|------|------|------|
| LDAP 登录 | ✅ 已完成 | 对接校园 LDAP / AD；可配置服务器、Bind DN、用户过滤器、属性映射、StartTLS |
| CAS 登录 | ✅ 已完成 | 对接校园 CAS 统一身份认证；支持属性映射与访问属性限制 |
| 目录分组自动分配 | ✅ 已完成 | 按 LDAP / CAS 属性正则规则，把用户映射到已有分组 |
| 分组渠道绑定 | ✅ 已完成 | 按「分组 + 模型」钉死可用渠道；未绑定模型回退上游原有调度 |
| 令牌锁定账号分组 | ✅ 已完成 | 创建/更新令牌时强制继承账号分组，用户无法自选分组或跨组重试 |
| 批量分配用户分组 | ✅ 已完成 | 管理端用户页可批量把用户划入指定分组 |
| 隐藏用户分组信息 | ✅ 已完成 | 个人资料、令牌页和模型广场不展示分组/用户 ID，模型广场只显示当前分组模型 |
| 课程 / 班级管理 | 🔜 计划中 | 按课程授权、教师管理班级学生 |

### 🛠️ 部署与运维

| 功能 | 状态 | 说明 |
|------|------|------|
| 去掉 SQLite | ✅ 已完成 | 仅支持 MySQL ≥ 5.7.8 / PostgreSQL ≥ 9.6；启动必须提供 `SQL_DSN` |
| JWT 会话 | ✅ 已完成 | 去掉 Cookie Session，控制台鉴权改为 JWT（`JWT_SECRET` / `JWT_EXPIRATION_SECONDS`） |
| Syslog | ✅ 已完成 | 可选输出到远程/本地 syslog，连接超时避免阻塞启动 |
| Docker Compose 全栈 | ✅ 已完成 | `docker-compose/` 提供 PostgreSQL、Redis、ClickHouse、OpenLDAP 与本服务 |
| Kubernetes | ✅ 已完成 | `kubernetes/` 含 Deployment、Service、HPA、ConfigMap/Secret 及外部依赖清单 |
| amd64 / arm64 构建 | ✅ 已完成 | `make` 可交叉编译 Linux amd64/arm64，以及纯后端 / 全栈镜像 |

---

## 🚀 快速开始

### 使用 Docker Compose（推荐）

本仓库 `docker-compose/` 会拉起 **New API Edu + PostgreSQL + Redis + ClickHouse + OpenLDAP**：

```bash
git clone https://gitee.com/ddbxiaoer/newapi_2_-edu.git
cd newapi_2_-edu

# 先构建镜像（全栈镜像内嵌前端）
make docker-allinone

# 按需修改 docker-compose/docker-compose.yml 中的密码与 JWT_SECRET
docker compose -f docker-compose/docker-compose.yml up -d
```

仅后端镜像（不含内嵌前端）用 `make docker-backend`，镜像名为 `newapi-edu-pure`。

> **⚠️ 生产环境务必修改** PostgreSQL / Redis / LDAP 默认密码，以及 `JWT_SECRET`。启动必须提供 `SQL_DSN` 与 `LOG_SQL_DSN`（可用 `REQUIRED_ENV_VARS=""` 关闭检查）。

---

🎉 部署完成后，访问 `http://localhost:3000` 即可使用！

---

## 📚 文档

### 📖 基础文档（基于 NEWAPI）
由于本项目基于 NEWAPI 开发，大部分基础功能文档可直接参考：
- **NEWAPI 官方文档**: [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs)
- **NEWAPI API 文档**: [https://docs.newapi.pro/zh/docs/api](https://docs.newapi.pro/zh/docs/api)
- **NEWAPI 部署指南**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

### 🎓 高校定制版文档
- LDAP / CAS：系统设置 → 认证
- 分组渠道绑定：计费 / 分组设置
- 部署清单：`docker-compose/`、`kubernetes/`

---

## ✨ 主要特性

> 详细特性请参考 **NEWAPI 官方文档**: [https://docs.newapi.pro/zh/docs/guide/wiki/basic-concepts/features-introduction](https://docs.newapi.pro/zh/docs/guide/wiki/basic-concepts/features-introduction)

### 🎨 核心功能（继承自 NEWAPI）

| 特性 | 说明 |
|------|------|
| 🎨 全新 UI | 现代化的用户界面设计 |
| 🌍 多语言 | 支持中文、英文、法语、日语 |
| 🔄 数据兼容 | 完全兼容原版 One API 数据库 |
| 📈 数据看板 | 可视化控制台与统计分析 |
| 🔒 权限管理 | 令牌分组、模型限制、用户管理 |

### 💰 支付与计费（继承自 NEWAPI）

- ✅ 在线充值（易支付、Stripe）
- ✅ 模型按次数收费
- ✅ 缓存计费支持（OpenAI、Azure、DeepSeek、Claude、Qwen等所有支持的模型）
- ✅ 灵活的计费策略配置

### 🎓 相对上游的新增能力

- 🏫 **LDAP / CAS 校园登录**：学号/工号对接校园统一身份认证
- 🧭 **目录分组自动分配**：按 LDAP / CAS 属性正则映射到已有分组
- 🔗 **分组渠道绑定**：按分组+模型钉死渠道，未绑定则走上游调度
- 🔒 **令牌锁定账号分组**：用户不能自选令牌分组；模型广场只显示当前分组模型
- 👥 **批量分配用户分组**：管理端可批量划入指定分组
- 🧾 **JWT + Syslog**：去掉 Cookie Session；可选 syslog，超时不阻塞启动
- 🚢 **高校部署清单**：去掉 SQLite；提供 Docker Compose 全栈与 Kubernetes + HPA
- 📚 **课程 / 班级管理**（🔜 计划中）

---

## 🤖 模型支持

> 详情请参考 **NEWAPI 接口文档**: [https://docs.newapi.pro/zh/docs/api](https://docs.newapi.pro/zh/docs/api)

| 模型类型 | 说明 | 文档 |
|---------|------|------|
| 🤖 OpenAI-Compatible | OpenAI 兼容模型 | [NEWAPI 文档](https://docs.newapi.pro/zh/docs/api/ai-model/chat/openai/createchatcompletion) |
| 🤖 OpenAI Responses | OpenAI Responses 格式 | [NEWAPI 文档](https://docs.newapi.pro/zh/docs/api/ai-model/chat/openai/createresponse) |
| 🎨 Midjourney-Proxy | Midjourney 代理 | [NEWAPI 文档](https://doc.newapi.pro/api/midjourney-proxy-image) |
| 🎵 Suno-API | Suno 音乐生成 | [NEWAPI 文档](https://doc.newapi.pro/api/suno-music) |
| 💬 Claude | Claude Messages 格式 | [NEWAPI 文档](https://docs.newapi.pro/zh/docs/api/ai-model/chat/createmessage) |
| 🌐 Gemini | Google Gemini | [NEWAPI 文档](https://docs.newapi.pro/zh/docs/api/ai-model/chat/gemini/geminirelayv1beta) |

---

## 🚢 部署

> [!TIP]
> **基础部署方式**：请参考 **NEWAPI 官方部署指南**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

### 🛠️ 从源码构建

```bash
# 克隆项目
git clone [项目地址]
cd [项目目录]

# 安装依赖
make prepare

# 构建当前平台
make build

# 交叉编译 Linux 平台
make build-backend-linux          # Linux amd64
make build-backend-linux-arm64    # Linux arm64

# 仅构建后端（不含内嵌前端）
make build-backend-pure           # 当前平台
make build-backend-pure-linux     # Linux amd64
make build-backend-pure-linux-arm64  # Linux arm64
```

### 📋 部署要求

| 组件 | 要求 |
|------|------|
| **主数据库** | MySQL ≥ 5.7.8 或 PostgreSQL ≥ 9.6（**必须**设置 `SQL_DSN`，已去掉 SQLite） |
| **日志库** | 通过 `LOG_SQL_DSN` 配置，支持独立库 / ClickHouse |
| **缓存** | Redis（多节点共享限流时必须） |
| **容器 / 编排** | Docker Compose（`docker-compose/`）或 Kubernetes（`kubernetes/`，含 HPA） |

Kubernetes 示例：

```bash
kubectl apply -k kubernetes/
```

### ⚙️ 高校定制环境变量

| 变量名 | 说明 | 默认值 |
|--------|------|--------|
| `SQL_DSN` | 主库连接串（必填） | - |
| `LOG_SQL_DSN` | 日志库连接串（默认必填） | - |
| `REQUIRED_ENV_VARS` | 启动必填变量列表；设为空字符串可关闭检查 | `SQL_DSN,LOG_SQL_DSN` |
| `JWT_SECRET` | JWT 签名密钥（未设置时回退 `SESSION_SECRET`） | `uuid` |
| `JWT_EXPIRATION_SECONDS` | JWT 过期时间（秒） | `604800` |
| `SYSLOG_ENABLED` | 启用 syslog | `false` |
| `SYSLOG_NETWORK` | syslog 协议（`udp`/`tcp`，空则本地 socket） | - |
| `SYSLOG_ADDR` | 远程 syslog 地址 | - |
| `SYSLOG_TAG` | syslog 标识 | `newapi` |

LDAP / CAS 在控制台「系统设置 → 认证」中配置，不走上述环境变量。

📖 **完整配置**：请参考 NEWAPI 环境变量文档 + 本仓库 `docker-compose/`、`kubernetes/`

---

## 🔗 相关项目

### 上游项目

| 项目 | 说明 |
|------|------|
| [NEWAPI](https://github.com/Calcium-Ion/new-api) | **基础项目** - 本项目基于此开发 |
| [One API](https://github.com/songquanpeng/one-api) | NEWAPI 的基础项目 |

### 配套工具

| 项目 | 说明 |
|------|------|
| [neko-api-key-tool](https://github.com/Calcium-Ion/neko-api-key-tool) | Key 额度查询工具 |
| [new-api-horizon](https://github.com/Calcium-Ion/new-api-horizon) | NEWAPI 高性能优化版 |

---

## 💬 帮助支持

### 📖 文档资源

| 资源 | 链接 |
|------|------|
| 📘 **NEWAPI 官方文档** | [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs) |
| 🎓 **高校定制文档** | 本项目文档目录 |
| 💬 **社区交流** | [NEWAPI 交流渠道](https://docs.newapi.pro/zh/docs/support/community-interaction) |
| 🐛 **问题反馈** | [NEWAPI 问题反馈](https://github.com/Calcium-Ion/new-api/issues) |

### 🤝 贡献指南

欢迎各种形式的贡献！

- 🐛 报告 Bug
- 💡 提出新功能
- 📝 改进文档
- 🔧 提交代码

---

## 📜 许可证

本项目采用 [GNU Affero 通用公共许可证 v3.0 (AGPLv3)](./LICENSE) 授权。

本项目为开源项目，在 **NEWAPI** ([https://github.com/Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api))（基于 MIT 许可证）的基础上进行二次开发。

如果您所在的组织政策不允许使用 AGPLv3 许可的软件，或您希望规避 AGPLv3 的开源义务，请发送邮件至：[support@quantumnous.com](mailto:support@quantumnous.com)

---

## 🌟 Star History

<div align="center">

[![Star History Chart](https://api.star-history.com/svg?repos=Calcium-Ion/new-api&type=Date)](https://star-history.com/#Calcium-Ion/new-api&Date)

</div>

---

<div align="center">

### 💖 感谢使用 New API 高校定制版

如果这个项目对你有帮助，欢迎给我们一个 ⭐️ Star！

**[NEWAPI 官方文档](https://docs.newapi.pro/zh/docs)** • **[问题反馈](https://github.com/Calcium-Ion/new-api/issues)** • **[最新发布](https://github.com/Calcium-Ion/new-api/releases)**

<sub>基于 NEWAPI 构建 ❤️ QuantumNous</sub>

</div>