<div align="center">

![new-api](/web/public/logo.png)

# New API - 高校定制版

🍥 **基於 NEWAPI 的高校環境定制版 - AI 網關與資產管理系統**

###
注意：本項目目前所有的更改全部由AI完成，由於本人去年腦出血偏癱。目前只有一隻手可以用，所以這個項目有很多還沒實現
由於我每天要去醫院康復，時間精力有限，而且由於沒工作沒收入，token全靠各大平台新用戶額度和邀請贈送，目前開發效率已經到極限了。token富裕的大哥可以自己上）
###在此感謝GLM 阿里云百煉

> ⚠️ **目前進度**：目前僅實現了校園帳號整合（學號/工號登入），其餘高校定制功能均為規劃中，尚未開發。歡迎貢獻！

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

### 🏫 教育場景優化

| 功能模組 | 狀態 | 說明 |
|---------|------|------|
| 🎓 教育帳號管理 | ✅ 已完成 | 支援學號/工號統一認證，與校園系統整合 |
| 🔐 校園安全整合 | 🔜 規劃中 | 支援校園統一身份認證系統 |

### 🛠️ 技術定制

| 定制項目 | 狀態 | 說明 |
|---------|------|------|
| 🏗️ 網路適配 | 🔜 規劃中 | 優化高校內網環境，支援代理和防火牆配置 |
| 💾 資料庫兼容 | 🔜 規劃中 | 針對高校常用資料庫（MySQL、PostgreSQL、SQLite）深度優化 |
| 🔄 介面適配 | 🔜 規劃中 | 提供與高校其他系統的標準介面 |
| 📱 移動端適配 | 🔜 規劃中 | 優化移動端訪問體驗，支援校園APP整合 |

---

## 🚀 快速開始

### 使用 Docker Compose（推薦）

```bash
# 克隆項目
git clone [本項目地址]
cd [項目目錄]

# 編輯 docker-compose.yml 配置
nano docker-compose.yml

# 啟動服務
docker-compose up -d
```

<details>
<summary><strong>使用 Docker 命令</strong></summary>

```bash
# 拉取最新鏡像
docker pull [定制版鏡像名稱]

# 使用 SQLite（預設）
docker run --name new-api-edu -d --restart always \
  -p 3000:3000 \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  [定制版鏡像名稱]:latest

# 使用 MySQL
docker run --name new-api-edu -d --restart always \
  -p 3000:3000 \
  -e SQL_DSN="root:123456@tcp(localhost:3306)/oneapi" \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  [定制版鏡像名稱]:latest
```

> **💡 提示：** `-v ./data:/data` 會將數據保存在當前目錄的 `data` 資料夾中，你也可以改為絕對路徑如 `-v /your/custom/path:/data`

</details>

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
**定制功能文檔：**
- 高校帳號管理配置
- 校園系統整合指南
- 教育權限設置說明
- 數據統計分析功能

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

### 🎓 高校定制新增功能

- 🏫 **校園身份認證**（✅ 已完成）：支援學號/工號登錄
- 📚 **課程管理**（🔜 規劃中）：按課程分配 AI 使用權限
- 👨‍🏫 **教師管理**（🔜 規劃中）：教師可管理班級學生
- 📊 **教育統計**（🔜 規劃中）：AI 使用數據分析
- 🔐 **安全審計**（🔜 規劃中）：完整的操作日誌和審計功能

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
| **本地資料庫** | SQLite（Docker 需掛載 `/data` 目錄）|
| **遠端資料庫** | MySQL ≥ 5.7.8 或 PostgreSQL ≥ 9.6 |
| **容器引擎** | Docker / Docker Compose |
| **網路環境** | 支援高校內網環境配置 |

### ⚙️ 高校定制環境變數

| 變數名 | 說明 | 預設值 |
|--------|------|--------|
| `EDU_MODE` | 啟用高校模式 | `true` |
| `CAMPUS_AUTH_URL` | 校園認證地址 | - |
| `CAMPUS_API_KEY` | 校園 API 密鑰 | - |
| `EDU_DOMAIN` | 教育域名限制 | - |
| `ALLOWED_DOMAINS` | 允許訪問的域名列表 | - |

📖 **完整配置**：請參考 NEWAPI 環境變數文檔 + 高校定制配置說明

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