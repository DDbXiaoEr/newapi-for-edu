<div align="center">

![new-api](/web/public/logo.png)

# New API - 高校定制版

🍥 **基於 NEWAPI 的高校環境定制版 - AI 網關與資產管理系統**

###
注意：本項目目前所有的更改全部由AI完成，由於本人去年腦出血偏癱。目前只有一隻手可以用，所以這個項目有很多還沒實現
由於我每天要去醫院康復，時間精力有限，而且由於沒工作沒收入，token全靠各大平台新用戶額度和邀請贈送，目前開發效率已經到極限了。token富裕的大哥可以自己上）
###在此感謝GLM 阿里云百煉

> ⚠️ **目前進度**：校園 LDAP / CAS 登入、分組渠道綁定、令牌鎖定帳號分組、JWT 會話、Syslog 與 Kubernetes / Docker Compose 部署已落地。課程管理、教師班級管理等教學業務仍在規劃中。歡迎貢獻！

<p align="center">
  <a href="./README.zh_CN.md">简体中文</a> |
  <strong>繁體中文</strong> |
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
  <a href="#-快速開始">快速開始</a> •
  <a href="#-項目特點">項目特點</a> •
  <a href="#-部署">部署</a> •
  <a href="#-高校定制功能">高校定制功能</a> •
  <a href="#-幫助支援">幫助</a>
</p>

</div>

## 📝 項目說明

> [!IMPORTANT]
> - 本項目是基於 **NEWAPI** 的高校環境定制版本，針對高校現有環境做了特別優化
> - 本項目僅供個人學習使用，不保證穩定性，且不提供任何技術支援
> - 使用者必須在遵循 OpenAI 的 [使用條款](https://openai.com/policies/terms-of-use) 以及**法律法規**的情況下使用，不得用於非法用途
> - 根據 [《生成式人工智慧服務管理暫行辦法》](http://www.cac.gov.cn/2023-07/13/c_1690898327029107.htm) 的要求，請勿對中國地區公眾提供一切未經備案的生成式人工智慧服務

---

## 🏗️ 項目架構

本項目基於 **NEWAPI** ([GitHub - Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api)) 進行二次開發，保留了原項目的所有核心功能，並針對高校環境進行了深度定制。

### 📚 原項目文檔
- **NEWAPI 官方文檔**: [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs)
- **NEWAPI GitHub**: [https://github.com/Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api)
- **NEWAPI 部署指南**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

---

## 🎯 高校定制功能

相對上游 [NEWAPI](https://github.com/Calcium-Ion/new-api) 的增量能力如下。未列出的能力（模型接入、計費、控制台等）仍繼承上游。

### 🏫 校園身份與權限

| 功能 | 狀態 | 說明 |
|------|------|------|
| LDAP 登入 | ✅ 已完成 | 對接校園 LDAP / AD；可配置伺服器、Bind DN、使用者過濾器、屬性對應、StartTLS |
| CAS 登入 | ✅ 已完成 | 對接校園 CAS 統一身份認證；支援屬性對應與存取屬性限制 |
| 目錄分組自動分配 | ✅ 已完成 | 按 LDAP / CAS 屬性正則規則，把使用者對應到既有分組 |
| 分組渠道綁定 | ✅ 已完成 | 按「分組 + 模型」釘死可用渠道；未綁定模型回退上游原有調度 |
| 令牌鎖定帳號分組 | ✅ 已完成 | 建立/更新令牌時強制繼承帳號分組，使用者無法自選分組或跨組重試 |
| 批次分配使用者分組 | ✅ 已完成 | 管理端使用者頁可批次把使用者劃入指定分組 |
| 隱藏使用者分組資訊 | ✅ 已完成 | 個人資料、令牌頁和模型廣場不展示分組/使用者 ID，模型廣場只顯示當前分組模型 |
| 課程 / 班級管理 | 🔜 規劃中 | 按課程授權、教師管理班級學生 |

### 🛠️ 部署與維運

| 功能 | 狀態 | 說明 |
|------|------|------|
| 去掉 SQLite | ✅ 已完成 | 僅支援 MySQL ≥ 5.7.8 / PostgreSQL ≥ 9.6；啟動必須提供 `SQL_DSN` |
| JWT 會話 | ✅ 已完成 | 去掉 Cookie Session，控制台鑑權改為 JWT（`JWT_SECRET` / `JWT_EXPIRATION_SECONDS`） |
| Syslog | ✅ 已完成 | 可選輸出到遠端/本機 syslog，連線逾時避免阻塞啟動 |
| Docker Compose 全棧 | ✅ 已完成 | `docker-compose/` 提供 PostgreSQL、Redis、ClickHouse、OpenLDAP 與本服務 |
| Kubernetes | ✅ 已完成 | `kubernetes/` 含 Deployment、Service、HPA、ConfigMap/Secret 及外部依賴清單 |
| amd64 / arm64 構建 | ✅ 已完成 | `make` 可交叉編譯 Linux amd64/arm64，以及純後端 / 全棧鏡像 |

---

## 🚀 快速開始

### 使用 Docker Compose（推薦）

本倉庫 `docker-compose/` 會拉起 **New API Edu + PostgreSQL + Redis + ClickHouse + OpenLDAP**：

```bash
git clone https://gitee.com/ddbxiaoer/newapi_2_-edu.git
cd newapi_2_-edu

# 先構建鏡像（全棧鏡像內嵌前端）
make docker-allinone

# 按需修改 docker-compose/docker-compose.yml 中的密碼與 JWT_SECRET
docker compose -f docker-compose/docker-compose.yml up -d
```

僅後端鏡像（不含內嵌前端）用 `make docker-backend`，鏡像名為 `newapi-edu-pure`。

> **⚠️ 生產環境務必修改** PostgreSQL / Redis / LDAP 預設密碼，以及 `JWT_SECRET`。啟動必須提供 `SQL_DSN` 與 `LOG_SQL_DSN`（可用 `REQUIRED_ENV_VARS=""` 關閉檢查）。

---

🎉 部署完成後，訪問 `http://localhost:3000` 即可使用！

---

## 📚 文檔

### 📖 基礎文檔（基於 NEWAPI）
由於本項目基於 NEWAPI 開發，大部分基礎功能文檔可直接參考：
- **NEWAPI 官方文檔**: [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs)
- **NEWAPI API 文檔**: [https://docs.newapi.pro/zh/docs/api](https://docs.newapi.pro/zh/docs/api)
- **NEWAPI 部署指南**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

### 🎓 高校定制版文檔
- LDAP / CAS：系統設置 → 認證
- 分組渠道綁定：計費 / 分組設置
- 部署清單：`docker-compose/`、`kubernetes/`

---

## ✨ 主要特性

> 詳細特性請參考 **NEWAPI 官方文檔**: [https://docs.newapi.pro/zh/docs/guide/wiki/basic-concepts/features-introduction](https://docs.newapi.pro/zh/docs/guide/wiki/basic-concepts/features-introduction)

### 🎨 核心功能（繼承自 NEWAPI）

| 特性 | 說明 |
|------|------|
| 🎨 全新 UI | 現代化的用戶界面設計 |
| 🌍 多語言 | 支援中文、英文、法語、日語 |
| 🔄 數據兼容 | 完全兼容原版 One API 資料庫 |
| 📈 數據看板 | 可視覺化控制台與統計分析 |
| 🔒 權限管理 | 令牌分組、模型限制、用戶管理 |

### 💰 支付與計費（繼承自 NEWAPI）

- ✅ 在線儲值（易支付、Stripe）
- ✅ 模型按次數收費
- ✅ 快取計費支援（OpenAI、Azure、DeepSeek、Claude、Qwen等所有支援的模型）
- ✅ 靈活的計費策略配置

### 🎓 相對上游的新增能力

- 🏫 **LDAP / CAS 校園登入**：學號/工號對接校園統一身份認證
- 🧭 **目錄分組自動分配**：按 LDAP / CAS 屬性正則對應到既有分組
- 🔗 **分組渠道綁定**：按分組+模型釘死渠道，未綁定則走上游調度
- 🔒 **令牌鎖定帳號分組**：使用者不能自選令牌分組；模型廣場只顯示當前分組模型
- 👥 **批次分配使用者分組**：管理端可批次劃入指定分組
- 🧾 **JWT + Syslog**：去掉 Cookie Session；可選 syslog，逾時不阻塞啟動
- 🚢 **高校部署清單**：去掉 SQLite；提供 Docker Compose 全棧與 Kubernetes + HPA
- 📚 **課程 / 班級管理**（🔜 規劃中）

---

## 🤖 模型支援

> 詳情請參考 **NEWAPI 接口文檔**: [https://docs.newapi.pro/zh/docs/api](https://docs.newapi.pro/zh/docs/api)

| 模型類型 | 說明 | 文檔 |
|---------|------|------|
| 🤖 OpenAI-Compatible | OpenAI 兼容模型 | [NEWAPI 文檔](https://docs.newapi.pro/zh/docs/api/ai-model/chat/openai/createchatcompletion) |
| 🤖 OpenAI Responses | OpenAI Responses 格式 | [NEWAPI 文檔](https://docs.newapi.pro/zh/docs/api/ai-model/chat/openai/createresponse) |
| 🎨 Midjourney-Proxy | Midjourney 代理 | [NEWAPI 文檔](https://doc.newapi.pro/api/midjourney-proxy-image) |
| 🎵 Suno-API | Suno 音樂生成 | [NEWAPI 文檔](https://doc.newapi.pro/api/suno-music) |
| 💬 Claude | Claude Messages 格式 | [NEWAPI 文檔](https://docs.newapi.pro/zh/docs/api/ai-model/chat/createmessage) |
| 🌐 Gemini | Google Gemini | [NEWAPI 文檔](https://docs.newapi.pro/zh/docs/api/ai-model/chat/gemini/geminirelayv1beta) |

---

## 🚢 部署

> [!TIP]
> **基礎部署方式**：請參考 **NEWAPI 官方部署指南**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

### 🛠️ 從源碼構建

```bash
# 克隆項目
git clone [項目地址]
cd [項目目錄]

# 安裝依賴
make prepare

# 構建當前平台
make build

# 交叉編譯 Linux 平台
make build-backend-linux          # Linux amd64
make build-backend-linux-arm64    # Linux arm64

# 僅構建後端（不含內嵌前端）
make build-backend-pure           # 當前平台
make build-backend-pure-linux     # Linux amd64
make build-backend-pure-linux-arm64  # Linux arm64
```

### 📋 部署要求

| 組件 | 要求 |
|------|------|
| **主資料庫** | MySQL ≥ 5.7.8 或 PostgreSQL ≥ 9.6（**必須**設置 `SQL_DSN`，已去掉 SQLite） |
| **日誌庫** | 透過 `LOG_SQL_DSN` 配置，支援獨立庫 / ClickHouse |
| **快取** | Redis（多節點共享限流時必須） |
| **容器 / 編排** | Docker Compose（`docker-compose/`）或 Kubernetes（`kubernetes/`，含 HPA） |

Kubernetes 示例：

```bash
kubectl apply -k kubernetes/
```

### ⚙️ 高校定制環境變數

| 變數名 | 說明 | 預設值 |
|--------|------|--------|
| `SQL_DSN` | 主庫連接串（必填） | - |
| `LOG_SQL_DSN` | 日誌庫連接串（預設必填） | - |
| `REQUIRED_ENV_VARS` | 啟動必填變數列表；設為空字串可關閉檢查 | `SQL_DSN,LOG_SQL_DSN` |
| `JWT_SECRET` | JWT 簽名密鑰（未設置時回退 `SESSION_SECRET`） | `uuid` |
| `JWT_EXPIRATION_SECONDS` | JWT 過期時間（秒） | `604800` |
| `SYSLOG_ENABLED` | 啟用 syslog | `false` |
| `SYSLOG_NETWORK` | syslog 協議（`udp`/`tcp`，空則本機 socket） | - |
| `SYSLOG_ADDR` | 遠端 syslog 地址 | - |
| `SYSLOG_TAG` | syslog 標識 | `newapi` |

LDAP / CAS 在控制台「系統設置 → 認證」中配置，不走上述環境變數。

📖 **完整配置**：請參考 NEWAPI 環境變數文檔 + 本倉庫 `docker-compose/`、`kubernetes/`

---

## 🔗 相關項目

### 上游項目

| 項目 | 說明 |
|------|------|
| [NEWAPI](https://github.com/Calcium-Ion/new-api) | **基礎項目** - 本項目基於此開發 |
| [One API](https://github.com/songquanpeng/one-api) | NEWAPI 的基礎項目 |

### 配套工具

| 項目 | 說明 |
|------|------|
| [neko-api-key-tool](https://github.com/Calcium-Ion/neko-api-key-tool) | Key 額度查詢工具 |
| [new-api-horizon](https://github.com/Calcium-Ion/new-api-horizon) | NEWAPI 高性能優化版 |

---

## 💬 幫助支援

### 📖 文檔資源

| 資源 | 連結 |
|------|------|
| 📘 **NEWAPI 官方文檔** | [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs) |
| 🎓 **高校定制文檔** | 本項目文檔目錄 |
| 💬 **社群交流** | [NEWAPI 交流渠道](https://docs.newapi.pro/zh/docs/support/community-interaction) |
| 🐛 **問題回饋** | [NEWAPI 問題回饋](https://github.com/Calcium-Ion/new-api/issues) |

### 🤝 貢獻指南

歡迎各種形式的貢獻！

- 🐛 報告 Bug
- 💡 提出新功能
- 📝 改進文檔
- 🔧 提交程式碼

---

## 📜 許可證

本項目採用 [GNU Affero 通用公共許可證 v3.0 (AGPLv3)](./LICENSE) 授權。

本項目為開源項目，在 **NEWAPI** ([https://github.com/Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api))（基於 MIT 許可證）的基礎上進行二次開發。

如果您所屬的組織政策不允許使用 AGPLv3 許可的軟體，或您希望規避 AGPLv3 的開源義務，請發送郵件至：[support@quantumnous.com](mailto:support@quantumnous.com)

---

## 🌟 Star History

<div align="center">

[![Star History Chart](https://api.star-history.com/svg?repos=Calcium-Ion/new-api&type=Date)](https://star-history.com/#Calcium-Ion/new-api&Date)

</div>

---

<div align="center">

### 💖 感謝使用 New API 高校定制版

如果這個項目對你有幫助，歡迎給我們一個 ⭐️ Star！

**[NEWAPI 官方文檔](https://docs.newapi.pro/zh/docs)** • **[問題回饋](https://github.com/Calcium-Ion/new-api/issues)** • **[最新發布](https://github.com/Calcium-Ion/new-api/releases)**

<sub>基於 NEWAPI 構建 ❤️ QuantumNous</sub>

</div>