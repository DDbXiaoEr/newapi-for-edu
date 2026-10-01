<div align="center">

![new-api](/web/public/logo.png)

# New API - 大学カスタム版

🍥 **NEWAPIベースの大学環境カスタム版 - AIゲートウェイと資産管理システム**

###
注意：本プロジェクトの現在の変更はすべてAIによって行われています。作者は昨年、脳出血による半身麻痺で片手のみ使用可能な状態です。そのため、実装されていない機能が多くあります。
毎日リハビリに通っており時間と体力に限りがあり、収入がないため、トークンは各プラットフォームの新規ユーザー枠や招待特典に頼っています。現在、開発効率は限界に達しています。トークンに余裕のある方はぜひご支援ください）
###GLMとアリババクラウド百煉、DeepSeek（梁文谷）そして[親友](https://github.com/FireSpoonYZ)が支援してくれたgrokに感謝します

> ⚠️ **現在の進捗**：キャンパス LDAP / CAS ログイン、グループ×チャネル固定、トークンの所属グループ固定、JWT セッション、Syslog、Kubernetes / Docker Compose デプロイは実装済みです。コース管理やクラス管理はまだ計画中です。貢献を歓迎します！

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
> - リポジトリ名に `edu` があり、文書は大学向けに書かれていますが、中核は汎用の AI API ゲートウェイです。企業、研究機関、その他の業種でもそのまま導入でき、教育業務に固定されていません。
> - 本プロジェクトは **NEWAPI** ベースの大学環境カスタム版であり、大学の既存環境向けに特別に最適化されています。キャンパス LDAP / CAS、グループ×チャネル固定、トークンの所属グループ固定は、企業の LDAP/AD、SSO、権限分離にもそのまま使えます。
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

上流 [NEWAPI](https://github.com/Calcium-Ion/new-api) からの差分です。記載のない能力（モデル接続、課金、コンソール等）は上流を継承します。

### 🏫 キャンパス認証と権限

| 機能 | 状態 | 説明 |
|------|------|------|
| LDAP ログイン | ✅ 完了 | キャンパス LDAP / AD。サーバー、Bind DN、ユーザーフィルター、属性マッピング、StartTLS |
| CAS ログイン | ✅ 完了 | キャンパス CAS SSO。属性マッピングとアクセス属性制限 |
| ディレクトリグループ自動割当 | ✅ 完了 | LDAP / CAS 属性の正規表現ルールで既存グループへ割当 |
| グループ×チャネル固定 | ✅ 完了 | グループ+モデルごとにチャネルを固定。未設定モデルは上流のルーティング |
| トークンの所属グループ固定 | ✅ 完了 | 作成/更新時にアカウントグループを強制。ユーザーはグループ選択や跨グループ再試行不可 |
| ユーザーの一括グループ割当 | ✅ 完了 | 管理画面のユーザーページから対象グループへ一括移動 |
| UI からグループ情報を隠す | ✅ 完了 | プロフィール、キー、モデル広場でグループ/ユーザー ID を非表示。モデル広場は所属グループのモデルのみ |
| コース / クラス管理 | 🔜 計画中 | コース単位の権限と教員によるクラス管理 |

### 🛠️ デプロイと運用

| 機能 | 状態 | 説明 |
|------|------|------|
| SQLite 廃止 | ✅ 完了 | MySQL ≥ 5.7.8 / PostgreSQL ≥ 9.6 のみ。起動時に `SQL_DSN` 必須 |
| JWT セッション | ✅ 完了 | Cookie Session を廃止し、コンソール認証を JWT に変更（`JWT_SECRET` / `JWT_EXPIRATION_SECONDS`） |
| Syslog | ✅ 完了 | リモート/ローカル syslog。接続タイムアウトで起動ブロックを回避 |
| Docker Compose スタック | ✅ 完了 | `docker-compose/` に PostgreSQL、Redis、ClickHouse、OpenLDAP と本サービス |
| Kubernetes | ✅ 完了 | `kubernetes/` に Deployment、Service、HPA、ConfigMap/Secret と外部依存マニフェスト |
| amd64 / arm64 ビルド | ✅ 完了 | `make` で Linux amd64/arm64 クロスコンパイル、純バックエンド / オールインワンイメージ |

---

## 🚀 クイックスタート

### Docker Composeを使用（推奨）

`docker-compose/` は **New API Edu + PostgreSQL + Redis + ClickHouse + OpenLDAP** を起動します：

```bash
git clone https://gitee.com/ddbxiaoer/newapi_2_-edu.git
cd newapi_2_-edu

# オールインワンイメージをビルド（フロントエンド埋め込み）
make docker-allinone

# docker-compose/docker-compose.yml のパスワードと JWT_SECRET を変更してから起動
docker compose -f docker-compose/docker-compose.yml up -d
```

純バックエンドイメージ（フロントエンドなし）は `make docker-backend` → `newapi-edu-pure`。

> **⚠️ 本番では** PostgreSQL / Redis / LDAP のデフォルトパスワードと `JWT_SECRET` を必ず変更してください。起動には `SQL_DSN` と `LOG_SQL_DSN` が必要です（`REQUIRED_ENV_VARS=""` でチェックを無効化できます）。

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
- LDAP / CAS：システム設定 → 認証
- グループ×チャネル固定：課金 / グループ設定
- デプロイマニフェスト：`docker-compose/`、`kubernetes/`

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

### 🎓 上流からの差分

- 🏫 **LDAP / CAS キャンパスログイン**：学籍番号/職員番号でキャンパス SSO に接続
- 🧭 **ディレクトリグループ自動割当**：LDAP / CAS 属性の正規表現で既存グループへ割当
- 🔗 **グループ×チャネル固定**：グループ+モデルでチャネルを固定。未設定は上流ルーティング
- 🔒 **トークンの所属グループ固定**：ユーザーはトークングループを選べない。モデル広場は所属グループのモデルのみ
- 👥 **ユーザーの一括グループ割当**：管理画面から対象グループへ一括移動
- 🧾 **JWT + syslog**：Cookie Session 廃止。任意の syslog、接続タイムアウトで起動ブロック回避
- 🚢 **大学向けデプロイ**：SQLite 廃止。Docker Compose スタックと Kubernetes + HPA
- 📚 **コース / クラス管理**（🔜 計画中）

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
| **メインデータベース** | MySQL ≥ 5.7.8 または PostgreSQL ≥ 9.6（`SQL_DSN` **必須**、SQLite 廃止） |
| **ログデータベース** | `LOG_SQL_DSN`。独立 DB / ClickHouse 対応 |
| **キャッシュ** | Redis（ノード間でレート制限を共有する場合必須） |
| **コンテナ / オーケストレーション** | Docker Compose（`docker-compose/`）または Kubernetes（`kubernetes/`、HPA 含む） |

Kubernetes 例：

```bash
kubectl apply -k kubernetes/
```

### ⚙️ 大学カスタム環境変数

| 変数名 | 説明 | デフォルト |
|--------|------|--------|
| `SQL_DSN` | メイン DB 接続文字列（必須） | - |
| `LOG_SQL_DSN` | ログ DB 接続文字列（デフォルト必須） | - |
| `REQUIRED_ENV_VARS` | 起動時必須変数リスト。空文字でチェック無効 | `SQL_DSN,LOG_SQL_DSN` |
| `JWT_SECRET` | JWT 署名シークレット（未設定時は `SESSION_SECRET`） | `uuid` |
| `JWT_EXPIRATION_SECONDS` | JWT 有効期限（秒） | `604800` |
| `SYSLOG_ENABLED` | syslog を有効化 | `false` |
| `SYSLOG_NETWORK` | syslog プロトコル（`udp`/`tcp`、空ならローカル socket） | - |
| `SYSLOG_ADDR` | リモート syslog アドレス | - |
| `SYSLOG_TAG` | syslog タグ | `newapi` |

LDAP / CAS はコンソールの「システム設定 → 認証」で設定し、上記の環境変数は使いません。

📖 **完全な設定**：NEWAPI 環境変数ドキュメント + 本リポジトリの `docker-compose/`、`kubernetes/`

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