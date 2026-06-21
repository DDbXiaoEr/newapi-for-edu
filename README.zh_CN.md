<div align="center">

![new-api](/web/public/logo.png)

# New API - 高校定制版

🍥 **基于 NEWAPI 的高校环境定制版 - AI 网关与资产管理系统**

###
Note: 本项目目前所有的更改全部由AI完成，由于本人去年脑出血偏瘫。目前只有一只手可以用，所以这个项目有很多还没实现
由于我每天要去医院康复，时间精力有限，而且由于没工作没收入，token全靠各大平台新用户额度和邀请赠送，目前开发效率已经到极限了。token富裕的大哥可以自己上）
###在此感谢GLM 阿里云百炼

> ⚠️ **当前进度**：目前仅实现了校园账号集成（学号/工号登录），其余高校定制功能均为计划中，尚未开发。欢迎贡献！

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

### 🏫 教育场景优化

| 功能模块 | 状态 | 说明 |
|---------|------|------|
| 🎓 教育账号管理 | ✅ 已完成 | 支持学号/工号统一认证，与校园系统集成 |
| 🔐 校园安全集成 | 🔜 计划中 | 支持校园统一身份认证系统 |

### 🛠️ 技术定制

| 定制项目 | 状态 | 说明 |
|---------|------|------|
| 🏗️ 网络适配 | 🔜 计划中 | 优化高校内网环境，支持代理和防火墙配置 |
| 💾 数据库兼容 | 🔜 计划中 | 针对高校常用数据库（MySQL、PostgreSQL、SQLite）深度优化 |
| 🔄 接口适配 | 🔜 计划中 | 提供与高校其他系统的标准接口 |
| 📱 移动端适配 | 🔜 计划中 | 优化移动端访问体验，支持校园APP集成 |

---

## 🚀 快速开始

### 使用 Docker Compose（推荐）

```bash
# 克隆项目
git clone [本项目地址]
cd [项目目录]

# 编辑 docker-compose.yml 配置
nano docker-compose.yml

# 启动服务
docker-compose up -d
```

<details>
<summary><strong>使用 Docker 命令</strong></summary>

```bash
# 拉取最新镜像
docker pull [定制版镜像名称]

# 使用 SQLite（默认）
docker run --name new-api-edu -d --restart always \
  -p 3000:3000 \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  [定制版镜像名称]:latest

# 使用 MySQL
docker run --name new-api-edu -d --restart always \
  -p 3000:3000 \
  -e SQL_DSN="root:123456@tcp(localhost:3306)/oneapi" \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  [定制版镜像名称]:latest
```

> **💡 提示：** `-v ./data:/data` 会将数据保存在当前目录的 `data` 文件夹中，你也可以改为绝对路径如 `-v /your/custom/path:/data`

</details>

---

🎉 部署完成后，访问 `http://localhost:3000` 即可使用！

---

## 📚 文档

### 📖 基础文档（基于 NEWAPI）
由于本项目基于 NEWAPI 开发，大部分基础功能文档可直接参考：
- **NEWAPI 官方文档**: [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs)
- **NEWAPI API 文档**: [https://docs.newapi.pro/zh/docs/api](https://docs.newapi.pro/zh/docs/api)
- **NEWAPI 部署指南**: [https://docs.newapi.pro/zh/docs/installation](httpshttps://docs.newapi.pro/zh/docs/installation)

### 🎓 高校定制版文档
**定制功能文档：**
- 高校账号管理配置
- 校园系统集成指南
- 教育权限设置说明
- 数据统计分析功能

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

### 🎓 高校定制新增功能

- 🏫 **校园身份认证**（✅ 已完成）：支持学号/工号登录
- 📚 **课程管理**（🔜 计划中）：按课程分配 AI 使用权限
- 👨‍🏫 **教师管理**（🔜 计划中）：教师可管理班级学生
- 📊 **教育统计**（🔜 计划中）：AI 使用数据分析
- 🔐 **安全审计**（🔜 计划中）：完整的操作日志和审计功能

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

### 📋 部署要求

| 组件 | 要求 |
|------|------|
| **本地数据库** | SQLite（Docker 需挂载 `/data` 目录）|
| **远程数据库** | MySQL ≥ 5.7.8 或 PostgreSQL ≥ 9.6 |
| **容器引擎** | Docker / Docker Compose |
| **网络环境** | 支持高校内网环境配置 |

### ⚙️ 高校定制环境变量

| 变量名 | 说明 | 默认值 |
|--------|------|--------|
| `EDU_MODE` | 启用高校模式 | `true` |
| `CAMPUS_AUTH_URL` | 校园认证地址 | - |
| `CAMPUS_API_KEY` | 校园 API 密钥 | - |
| `EDU_DOMAIN` | 教育域名限制 | - |
| `ALLOWED_DOMAINS` | 允许访问的域名列表 | - |

📖 **完整配置**：请参考 NEWAPI 环境变量文档 + 高校定制配置说明

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