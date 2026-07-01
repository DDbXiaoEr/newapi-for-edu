<div align="center">

![new-api](/web/public/logo.png)

# New API - 大学カスタム版

🍥 **NEWAPIベースの大学環境カスタム版 - AIゲートウェイと資産管理システム**

###
注意：本プロジェクトの現在の変更はすべてAIによって行われています。作者は昨年、脳出血による半身麻痺で片手のみ使用可能な状態です。そのため、実装されていない機能が多くあります。
毎日リハビリに通っており時間と体力に限りがあり、収入がないため、トークンは各プラットフォームの新規ユーザー枠や招待特典に頼っています。現在、開発効率は限界に達しています。トークンに余裕のある方はぜひご支援ください）
###GLMとアリババクラウド百煉に感謝します

> ⚠️ **現在の進捗**：現在はキャンパスアカウント統合（学籍番号/職員番号ログイン）のみ実装されています。その他の大学カスタム機能はすべて計画中で、まだ開発されていません。貢献を歓迎します！

<p align="center">
  <a href="./README.zh_CN.md">简体中文</a> |
  <a href="./README.zh_TW.md">繁體中文</a> |
  <a href="./README.md">English</a> |
  <a href="./README.fr.md">Français</a> |
  <strong>日本語</strong>
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
  <a href="#-クイックスタート">クイックスタート</a> •
  <a href="#-プロジェクトの特徴">プロジェクトの特徴</a> •
  <a href="#-デプロイ">デプロイ</a> •
  <a href="#-大学カスタム機能">大学カスタム機能</a> •
  <a href="#-ヘルプサポート">ヘルプ</a>
</p>

</div>

## 📝 プロジェクト説明

> [!IMPORTANT]
> - 本プロジェクトは **NEWAPI** ベースの大学環境カスタム版であり、大学の既存環境向けに特別に最適化されています
> - 本プロジェクトは個人学習用のみであり、安定性の保証や技術サポートは提供しません。
> - ユーザーは、OpenAIの[利用規約](https://openai.com/policies/terms-of-use)および**法律法規**を遵守する必要があり、違法な目的で使用してはいけません。
> - [《生成式人工智能服务管理暂行办法》](http://www.cac.gov.cn/2023-07/13/c_1690898327029107.htm)の要求に従い、中国地域の公衆に未登録の生成式AIサービスを提供しないでください。

---

## 🏗️ プロジェクトアーキテクチャ

本プロジェクトは **NEWAPI** ([GitHub - Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api)) をベースに二次開発され、元プロジェクトのすべてのコア機能を保持しつつ、大学環境向けに深くカスタマイズされています。

### 📚 元プロジェクトのドキュメント
- **NEWAPI 公式ドキュメント**: [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs)
- **NEWAPI GitHub**: [https://github.com/Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api)
- **NEWAPI デプロイガイド**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

---

## 🎯 大学カスタム機能

### 🏫 教育シーンの最適化

| 機能モジュール | 状態 | 説明 |
|---------|------|------|
| 🎓 教育アカウント管理 | ✅ 完了 | 学籍番号/職員番号の統一認証、キャンパスシステムとの統合 |
| 🔐 キャンパスセキュリティ統合 | 🔜 計画中 | キャンパス統一ID認証システムのサポート |

### 🛠️ 技術カスタマイズ

| カスタム項目 | 状態 | 説明 |
|---------|------|------|
| 🏗️ ネットワーク適応 | 🔜 計画中 | 大学内ネットワーク環境の最適化、プロキシとファイアウォール設定のサポート |
| 💾 データベース互換性 | 🔜 計画中 | 大学でよく使われるデータベース（MySQL、PostgreSQL、SQLite）への深い最適化 |
| 🔄 インターフェース適応 | 🔜 計画中 | 大学内の他システムとの標準インターフェース提供 |
| 📱 モバイル適応 | 🔜 計画中 | モバイルアクセス体験の最適化、キャンパスアプリ統合のサポート |

---

## 🚀 クイックスタート

### Docker Composeを使用（推奨）

```bash
# プロジェクトをクローン
git clone [本プロジェクトアドレス]
cd [プロジェクトディレクトリ]

# docker-compose.yml 設定を編集
nano docker-compose.yml

# サービスを起動
docker-compose up -d
```

<details>
<summary><strong>Dockerコマンドを使用</strong></summary>

```bash
# 最新のイメージをプル
docker pull [カスタム版イメージ名]

# SQLiteを使用（デフォルト）
docker run --name new-api-edu -d --restart always \
  -p 3000:3000 \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  [カスタム版イメージ名]:latest

# MySQLを使用
docker run --name new-api-edu -d --restart always \
  -p 3000:3000 \
  -e SQL_DSN="root:123456@tcp(localhost:3306)/oneapi" \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  [カスタム版イメージ名]:latest
```

> **💡 ヒント:** `-v ./data:/data` は現在のディレクトリの `data` フォルダにデータを保存します。絶対パスに変更することもできます：`-v /your/custom/path:/data`

</details>

---

🎉 デプロイが完了したら、`http://localhost:3000` にアクセスして使用を開始してください！

---

## 📚 ドキュメント

### 📖 基本ドキュメント（NEWAPIベース）
本プロジェクトは NEWAPI をベースに開発されているため、ほとんどの基本機能ドキュメントは以下を直接参照できます：
- **NEWAPI 公式ドキュメント**: [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs)
- **NEWAPI API ドキュメント**: [https://docs.newapi.pro/zh/docs/api](https://docs.newapi.pro/zh/docs/api)
- **NEWAPI デプロイガイド**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

### 🎓 大学カスタム版ドキュメント
**カスタム機能ドキュメント：**
- 大学アカウント管理設定
- キャンパスシステム統合ガイド
- 教育権限設定説明
- データ統計分析機能

---

## ✨ 主な機能

> 詳細な機能については **NEWAPI 公式ドキュメント** を参照: [https://docs.newapi.pro/zh/docs/guide/wiki/basic-concepts/features-introduction](https://docs.newapi.pro/zh/docs/guide/wiki/basic-concepts/features-introduction)

### 🎨 コア機能（NEWAPIから継承）

| 機能 | 説明 |
|------|------|
| 🎨 新しいUI | モダンなユーザーインターフェースデザイン |
| 🌍 多言語 | 中国語、英語、フランス語、日本語をサポート |
| 🔄 データ互換性 | オリジナルのOne APIデータベースと完全に互換性あり |
| 📈 データダッシュボード | ビジュアルコンソールと統計分析 |
| 🔒 権限管理 | トークングループ化、モデル制限、ユーザー管理 |

### 💰 支払いと課金（NEWAPIから継承）

- ✅ オンライン充電（EPay、Stripe）
- ✅ モデルの従量課金
- ✅ キャッシュ課金サポート（OpenAI、Azure、DeepSeek、Claude、Qwenなどすべてのサポートされているモデル）
- ✅ 柔軟な課金ポリシー設定

### 🎓 大学カスタム新機能

- 🏫 **キャンパスID認証**（✅ 完了）：学籍番号/職員番号ログインをサポート
- 📚 **コース管理**（🔜 計画中）：コースごとにAI使用権限を割り当て
- 👨‍🏫 **教員管理**（🔜 計画中）：教員はクラスの学生を管理可能
- 📊 **教育統計**（🔜 計画中）：AI使用データ分析
- 🔐 **セキュリティ監査**（🔜 計画中）：完全な操作ログと監査機能

---

## 🤖 モデルサポート

> 詳細については **NEWAPI API ドキュメント** を参照: [https://docs.newapi.pro/zh/docs/api](https://docs.newapi.pro/zh/docs/api)

| モデルタイプ | 説明 | ドキュメント |
|---------|------|------|
| 🤖 OpenAI-Compatible | OpenAI互換モデル | [NEWAPI ドキュメント](https://docs.newapi.pro/zh/docs/api/ai-model/chat/openai/createchatcompletion) |
| 🤖 OpenAI Responses | OpenAI Responsesフォーマット | [NEWAPI ドキュメント](https://docs.newapi.pro/zh/docs/api/ai-model/chat/openai/createresponse) |
| 🎨 Midjourney-Proxy | Midjourneyプロキシ | [NEWAPI ドキュメント](https://doc.newapi.pro/api/midjourney-proxy-image) |
| 🎵 Suno-API | Suno音楽生成 | [NEWAPI ドキュメント](https://doc.newapi.pro/api/suno-music) |
| 💬 Claude | Claude Messagesフォーマット | [NEWAPI ドキュメント](https://docs.newapi.pro/zh/docs/api/ai-model/chat/createmessage) |
| 🌐 Gemini | Google Gemini | [NEWAPI ドキュメント](https://docs.newapi.pro/zh/docs/api/ai-model/chat/gemini/geminirelayv1beta) |

---

## 🚢 デプロイ

> [!TIP]
> **基本デプロイ方法**：**NEWAPI 公式デプロイガイド** を参照してください: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

### 🛠️ ソースからビルド

```bash
# プロジェクトをクローン
git clone [プロジェクトアドレス]
cd [プロジェクトディレクトリ]

# 依存関係をインストール
make prepare

# 現在のプラットフォームでビルド
make build

# Linux向けクロスコンパイル
make build-backend-linux          # Linux amd64
make build-backend-linux-arm64    # Linux arm64

# バックエンドのみビルド（フロントエンド埋め込みなし）
make build-backend-pure           # 現在のプラットフォーム
make build-backend-pure-linux     # Linux amd64
make build-backend-pure-linux-arm64  # Linux arm64
```

### 📋 デプロイ要件

| コンポーネント | 要件 |
|------|------|
| **ローカルデータベース** | SQLite（Dockerは `/data` ディレクトリをマウントする必要があります）|
| **リモートデータベース** | MySQL ≥ 5.7.8 または PostgreSQL ≥ 9.6 |
| **コンテナエンジン** | Docker / Docker Compose |
| **ネットワーク環境** | 大学内ネットワーク環境設定をサポート |

### ⚙️ 大学カスタム環境変数

| 変数名 | 説明 | デフォルト値 |
|--------|------|--------|
| `EDU_MODE` | 大学モードを有効にする | `true` |
| `CAMPUS_AUTH_URL` | キャンパス認証アドレス | - |
| `CAMPUS_API_KEY` | キャンパスAPIキー | - |
| `EDU_DOMAIN` | 教育ドメイン制限 | - |
| `ALLOWED_DOMAINS` | アクセス許可ドメインリスト | - |

📖 **完全な設定**：NEWAPI 環境変数ドキュメント + 大学カスタム設定説明を参照してください

---

## 🔗 関連プロジェクト

### 上流プロジェクト

| プロジェクト | 説明 |
|------|------|
| [NEWAPI](https://github.com/Calcium-Ion/new-api) | **基本プロジェクト** - 本プロジェクトはこれに基づいて開発されています |
| [One API](https://github.com/songquanpeng/one-api) | NEWAPIのベースプロジェクト |

### 補助ツール

| プロジェクト | 説明 |
|------|------|
| [neko-api-key-tool](https://github.com/Calcium-Ion/neko-api-key-tool) | キー使用量クォータ照会ツール |
| [new-api-horizon](https://github.com/Calcium-Ion/new-api-horizon) | NEWAPI高性能最適化版 |

---

## 💬 ヘルプサポート

### 📖 ドキュメントリソース

| リソース | リンク |
|------|------|
| 📘 **NEWAPI 公式ドキュメント** | [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs) |
| 🎓 **大学カスタムドキュメント** | 本プロジェクトドキュメントディレクトリ |
| 💬 **コミュニティ交流** | [NEWAPI 交流チャネル](https://docs.newapi.pro/zh/docs/support/community-interaction) |
| 🐛 **問題フィードバック** | [NEWAPI 問題フィードバック](https://github.com/Calcium-Ion/new-api/issues) |

### 🤝 貢献ガイド

あらゆる形の貢献を歓迎します！

- 🐛 バグを報告する
- 💡 新しい機能を提案する
- 📝 ドキュメントを改善する
- 🔧 コードを提出する

---

## 📜 ライセンス

本プロジェクトは [GNU Affero General Public License v3.0 (AGPLv3)](./LICENSE) の下でライセンスされています。

本プロジェクトは、**NEWAPI** ([https://github.com/Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api))（MITライセンスに基づく）をベースに二次開発されたオープンソースプロジェクトです。

お客様の組織のポリシーがAGPLv3ライセンスのソフトウェアの使用を許可していない場合、またはAGPLv3のオープンソース義務を回避したい場合は、こちらまでお問い合わせください：[support@quantumnous.com](mailto:support@quantumnous.com)

---

## 🌟 スター履歴

<div align="center">

[![スター履歴チャート](https://api.star-history.com/svg?repos=Calcium-Ion/new-api&type=Date)](https://star-history.com/#Calcium-Ion/new-api&Date)

</div>

---

<div align="center">

### 💖 New API 大学カスタム版をご利用いただきありがとうございます

このプロジェクトがあなたのお役に立てたなら、ぜひ ⭐️ スターをください！

**[NEWAPI 公式ドキュメント](https://docs.newapi.pro/zh/docs)** • **[問題フィードバック](https://github.com/Calcium-Ion/new-api/issues)** • **[最新リリース](https://github.com/Calcium-Ion/new-api/releases)**

<sub>NEWAPIベースで構築 ❤️ QuantumNous</sub>

</div>