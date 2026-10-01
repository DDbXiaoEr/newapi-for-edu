<div align="center">

![new-api](/web/public/logo.png)

# New API - Édition universitaire personnalisée

🍥 **Version personnalisée pour l'environnement universitaire basée sur NEWAPI - Passerelle IA et système de gestion d'actifs**

###
Remarque : Toutes les modifications actuelles de ce projet ont été effectuées par une IA. L'auteur a subi une hémorragie cérébrale l'année dernière, entraînant une hémiplégie, et ne peut actuellement utiliser qu'une seule main. De nombreuses fonctionnalités ne sont donc pas encore implémentées.
Je vais à l'hôpital chaque jour pour la rééducation, mon temps et mon énergie sont limités, et sans travail ni revenu, les tokens dépendent entièrement des quotas pour nouveaux utilisateurs et des invitations offertes par les plateformes. L'efficacité du développement a atteint sa limite. Les personnes disposant de tokens en abondance sont les bienvenues pour contribuer)
###Remerciements à GLM, Alibaba Cloud Bailian, DeepSeek (Liang Wengu) et mon [bon ami](https://github.com/FireSpoonYZ) qui a fourni un accès à grok

> ⚠️ **Avancement actuel** : Connexion LDAP / CAS campus, liaison groupe-canal, jetons verrouillés sur le groupe du compte, sessions JWT, syslog et déploiement Kubernetes / Docker Compose sont implémentés. La gestion des cours et des classes reste prévue. Les contributions sont les bienvenues !

<p align="center">
  <a href="./README.zh_CN.md">简体中文</a> |
  <a href="./README.zh_TW.md">繁體中文</a> |
  <a href="./README.md">English</a> |
  <strong>Français</strong> |
  <a href="./README.ja.md">日本語</a>
</p>

<p align="center">
  <a href="https://raw.githubusercontent.com/Calcium-Ion/new-api/main/LICENSE">
    <img src="https://img.shields.io/github/license/Calcium-Ion/new-api?color=brightgreen" alt="licence">
  </a><!--
  --><a href="https://github.com/Calcium-Ion/new-api/releases/latest">
    <img src="https://img.shields.io/github/v/release/Calcium-Ion/new-api?color=brightgreen&include_prereleases" alt="version">
  </a><!--
  --><a href="https://hub.docker.com/r/CalciumIon/new-api">
    <img src="https://img.shields.io/badge/docker-dockerHub-blue" alt="docker">
  </a>
</p>

<p align="center">
  <a href="#-démarrage-rapide">Démarrage rapide</a> •
  <a href="#-caractéristiques-du-projet">Caractéristiques</a> •
  <a href="#-déploiement">Déploiement</a> •
  <a href="#-fonctionnalités-personnalisées-pour-luniversité">Fonctionnalités universitaires</a> •
  <a href="#-aide-support">Aide</a>
</p>

</div>

## 📝 Description du projet

> [!IMPORTANT]
> - Le nom du dépôt contient `edu` et la documentation vise le campus, mais le cœur reste une passerelle API d’IA généraliste. Les entreprises, laboratoires et autres secteurs peuvent le déployer tel quel ; il n’est pas verrouillé aux processus éducatifs.
> - Ce projet est une version personnalisée pour l'environnement universitaire basée sur **NEWAPI**, spécialement optimisée pour les environnements existants des établissements d'enseignement supérieur. LDAP / CAS campus, liaisons groupe-canal et jetons verrouillés sur le groupe du compte s’appliquent aussi à LDAP/AD, SSO et l’isolation d’accès en entreprise.
> - Ce projet est uniquement destiné à des fins d'apprentissage personnel, sans garantie de stabilité ni de support technique.
> - Les utilisateurs doivent se conformer aux [Conditions d'utilisation](https://openai.com/policies/terms-of-use) d'OpenAI et aux **lois et réglementations applicables**, et ne doivent pas l'utiliser à des fins illégales.
> - Conformément aux [《Mesures provisoires pour la gestion des services d'intelligence artificielle générative》](http://www.cac.gov.cn/2023-07/13/c_1690898327029107.htm), veuillez ne fournir aucun service d'IA générative non enregistré au public en Chine.

---

## 🏗️ Architecture du projet

Ce projet est développé sur la base de **NEWAPI** ([GitHub - Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api)), conservant toutes les fonctionnalités principales du projet original tout en étant profondément personnalisé pour l'environnement universitaire.

### 📚 Documentation du projet original
- **Documentation officielle NEWAPI**: [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs)
- **NEWAPI GitHub**: [https://github.com/Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api)
- **Guide de déploiement NEWAPI**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

---

## 🎯 Fonctionnalités personnalisées pour l'université

Différences par rapport à l'amont [NEWAPI](https://github.com/Calcium-Ion/new-api). Le reste (accès aux modèles, facturation, console) est hérité.

### 🏫 Identité et accès campus

| Fonctionnalité | Statut | Description |
|----------------|--------|-------------|
| Connexion LDAP | ✅ Fait | LDAP / AD campus ; URL, Bind DN, filtre utilisateur, mapping d'attributs, StartTLS |
| Connexion CAS | ✅ Fait | SSO CAS campus ; mapping d'attributs et restriction d'attribut d'accès |
| Affectation automatique de groupe | ✅ Fait | Mapper les attributs LDAP / CAS vers les groupes existants par règles regex |
| Liaison groupe-canal | ✅ Fait | Épingler les canaux par groupe+modèle ; les modèles non liés gardent le routage amont |
| Jetons verrouillés sur le groupe du compte | ✅ Fait | Création/mise à jour force le groupe du compte ; pas de choix de groupe ni de retry inter-groupes |
| Affectation groupée des utilisateurs | ✅ Fait | L'admin peut déplacer une sélection d'utilisateurs vers un groupe |
| Masquer le groupe dans l'UI | ✅ Fait | Profil, clés et place des modèles masquent groupe/ID ; la place n'affiche que les modèles du groupe courant |
| Gestion des cours / classes | 🔜 Prévu | Droits par cours et gestion des classes par les enseignants |

### 🛠️ Déploiement et exploitation

| Fonctionnalité | Statut | Description |
|----------------|--------|-------------|
| Plus de SQLite | ✅ Fait | MySQL ≥ 5.7.8 ou PostgreSQL ≥ 9.6 uniquement ; `SQL_DSN` obligatoire au démarrage |
| Sessions JWT | ✅ Fait | Cookies de session retirés ; auth console en JWT (`JWT_SECRET` / `JWT_EXPIRATION_SECONDS`) |
| Syslog | ✅ Fait | Syslog distant/local optionnel ; timeout de connexion pour ne pas bloquer le démarrage |
| Stack Docker Compose | ✅ Fait | `docker-compose/` fournit PostgreSQL, Redis, ClickHouse, OpenLDAP et ce service |
| Kubernetes | ✅ Fait | `kubernetes/` : Deployment, Service, HPA, ConfigMap/Secret et dépendances externes |
| Builds amd64 / arm64 | ✅ Fait | `make` compile Linux amd64/arm64, images backend-seul et all-in-one |

---

## 🚀 Démarrage rapide

### Utilisation de Docker Compose (recommandé)

`docker-compose/` démarre **New API Edu + PostgreSQL + Redis + ClickHouse + OpenLDAP** :

```bash
git clone https://gitee.com/ddbxiaoer/newapi_2_-edu.git
cd newapi_2_-edu

# Construire l'image all-in-one (frontend intégré)
make docker-allinone

# Modifier d'abord les mots de passe et JWT_SECRET dans docker-compose/docker-compose.yml
docker compose -f docker-compose/docker-compose.yml up -d
```

Image backend seul (sans frontend intégré) : `make docker-backend` → `newapi-edu-pure`.

> **⚠️ En production**, changez les mots de passe PostgreSQL / Redis / LDAP et `JWT_SECRET`. Le démarrage exige `SQL_DSN` et `LOG_SQL_DSN` (mettez `REQUIRED_ENV_VARS=""` pour désactiver le contrôle).

---

🎉 Après le déploiement, visitez `http://localhost:3000` pour commencer à utiliser !

---

## 📚 Documentation

### 📖 Documentation de base (basée sur NEWAPI)
Étant donné que ce projet est développé sur la base de NEWAPI, la plupart des documentations de base peuvent être consultées directement :
- **Documentation officielle NEWAPI**: [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs)
- **Documentation API NEWAPI**: [https://docs.newapi.pro/zh/docs/api](https://docs.newapi.pro/zh/docs/api)
- **Guide de déploiement NEWAPI**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

### 🎓 Documentation de l'édition universitaire personnalisée
- LDAP / CAS : Paramètres système → Authentification
- Liaison groupe-canal : facturation / paramètres de groupe
- Manifestes de déploiement : `docker-compose/`, `kubernetes/`

---

## ✨ Caractéristiques principales

> Pour les caractéristiques détaillées, veuillez consulter la **documentation officielle NEWAPI**: [https://docs.newapi.pro/zh/docs/guide/wiki/basic-concepts/features-introduction](https://docs.newapi.pro/zh/docs/guide/wiki/basic-concepts/features-introduction)

### 🎨 Fonctionnalités principales (héritées de NEWAPI)

| Caractéristique | Description |
|------|------|
| 🎨 Nouvelle interface utilisateur | Conception d'interface utilisateur moderne |
| 🌍 Multilingue | Prend en charge le chinois, l'anglais, le français et le japonais |
| 🔄 Compatibilité des données | Complètement compatible avec la base de données originale de One API |
| 📈 Tableau de bord des données | Console visuelle et analyse statistique |
| 🔒 Gestion des permissions | Regroupement de jetons, restrictions de modèles, gestion des utilisateurs |

### 💰 Paiement et facturation (hérités de NEWAPI)

- ✅ Recharge en ligne (EPay, Stripe)
- ✅ Tarification des modèles par utilisation
- ✅ Prise en charge de la facturation du cache (OpenAI, Azure, DeepSeek, Claude, Qwen et tous les modèles pris en charge)
- ✅ Configuration flexible des politiques de facturation

### 🎓 Différences par rapport à l'amont

- 🏫 **Connexion LDAP / CAS campus** : numéro d'étudiant/employé contre le SSO campus
- 🧭 **Affectation automatique de groupe** : mapper les attributs LDAP / CAS vers les groupes existants
- 🔗 **Liaison groupe-canal** : épingler les canaux par groupe+modèle ; les modèles non liés gardent le routage amont
- 🔒 **Jetons verrouillés sur le groupe du compte** : l'utilisateur ne choisit pas le groupe du jeton ; la place n'affiche que les modèles du groupe courant
- 👥 **Affectation groupée des utilisateurs** : l'admin peut déplacer une sélection vers un groupe
- 🧾 **JWT + syslog** : cookies de session retirés ; syslog optionnel avec timeout de connexion
- 🚢 **Déploiement campus** : plus de SQLite ; stack Docker Compose et Kubernetes + HPA
- 📚 **Gestion des cours / classes** (🔜 Prévu)

---

## 🤖 Modèles pris en charge

> Pour les détails, veuillez consulter la **documentation de l'API NEWAPI**: [https://docs.newapi.pro/zh/docs/api](https://docs.newapi.pro/zh/docs/api)

| Type de modèle | Description | Documentation |
|---------|------|------|
| 🤖 OpenAI-Compatible | Modèles compatibles OpenAI | [Documentation NEWAPI](https://docs.newapi.pro/zh/docs/api/ai-model/chat/openai/createchatcompletion) |
| 🤖 OpenAI Responses | Format OpenAI Responses | [Documentation NEWAPI](https://docs.newapi.pro/zh/docs/api/ai-model/chat/openai/createresponse) |
| 🎨 Midjourney-Proxy | Proxy Midjourney | [Documentation NEWAPI](https://doc.newapi.pro/api/midjourney-proxy-image) |
| 🎵 Suno-API | Génération de musique Suno | [Documentation NEWAPI](https://doc.newapi.pro/api/suno-music) |
| 💬 Claude | Format Claude Messages | [Documentation NEWAPI](https://docs.newapi.pro/zh/docs/api/ai-model/chat/createmessage) |
| 🌐 Gemini | Google Gemini | [Documentation NEWAPI](https://docs.newapi.pro/zh/docs/api/ai-model/chat/gemini/geminirelayv1beta) |

---

## 🚢 Déploiement

> [!TIP]
> **Méthode de déploiement de base** : veuillez consulter le **Guide de déploiement officiel NEWAPI**: [https://docs.newapi.pro/zh/docs/installation](https://docs.newapi.pro/zh/docs/installation)

### 🛠️ Compilation à partir des sources

```bash
# Cloner le projet
git clone [adresse du projet]
cd [répertoire du projet]

# Installer les dépendances
make prepare

# Compiler pour la plateforme actuelle
make build

# Compilation croisée pour Linux
make build-backend-linux          # Linux amd64
make build-backend-linux-arm64    # Linux arm64

# Compiler le backend uniquement (sans frontend intégré)
make build-backend-pure           # Plateforme actuelle
make build-backend-pure-linux     # Linux amd64
make build-backend-pure-linux-arm64  # Linux arm64
```

### 📋 Exigences de déploiement

| Composant | Exigence |
|------|------|
| **Base principale** | MySQL ≥ 5.7.8 ou PostgreSQL ≥ 9.6 (`SQL_DSN` **obligatoire** ; SQLite retiré) |
| **Base de logs** | `LOG_SQL_DSN` ; base séparée / ClickHouse |
| **Cache** | Redis (obligatoire si les nœuds partagent le rate limit) |
| **Conteneurs / orchestration** | Docker Compose (`docker-compose/`) ou Kubernetes (`kubernetes/`, HPA inclus) |

Exemple Kubernetes :

```bash
kubectl apply -k kubernetes/
```

### ⚙️ Variables d'environnement personnalisées pour l'université

| Variable | Description | Défaut |
|--------|------|--------|
| `SQL_DSN` | DSN de la base principale (obligatoire) | - |
| `LOG_SQL_DSN` | DSN de la base de logs (obligatoire par défaut) | - |
| `REQUIRED_ENV_VARS` | Liste des variables requises au démarrage ; chaîne vide pour désactiver | `SQL_DSN,LOG_SQL_DSN` |
| `JWT_SECRET` | Secret de signature JWT (repli sur `SESSION_SECRET`) | `uuid` |
| `JWT_EXPIRATION_SECONDS` | Durée de vie JWT en secondes | `604800` |
| `SYSLOG_ENABLED` | Activer syslog | `false` |
| `SYSLOG_NETWORK` | Protocole syslog (`udp`/`tcp` ; vide = socket local) | - |
| `SYSLOG_ADDR` | Adresse syslog distante | - |
| `SYSLOG_TAG` | Tag syslog | `newapi` |

LDAP / CAS se configurent dans la console : Paramètres système → Authentification, pas via les variables ci-dessus.

📖 **Configuration complète** : documentation des variables d'environnement NEWAPI + `docker-compose/` et `kubernetes/` de ce dépôt

---

## 🔗 Projets connexes

### Projets en amont

| Projet | Description |
|------|------|
| [NEWAPI](https://github.com/Calcium-Ion/new-api) | **Projet de base** - Ce projet est développé sur cette base |
| [One API](https://github.com/songquanpeng/one-api) | Projet de base de NEWAPI |

### Outils complémentaires

| Projet | Description |
|------|------|
| [neko-api-key-tool](https://github.com/Calcium-Ion/neko-api-key-tool) | Outil de recherche de quota d'utilisation avec une clé |
| [new-api-horizon](https://github.com/Calcium-Ion/new-api-horizon) | Version optimisée haute performance de NEWAPI |

---

## 💬 Aide et support

### 📖 Ressources documentaires

| Ressource | Lien |
|------|------|
| 📘 **Documentation officielle NEWAPI** | [https://docs.newapi.pro/zh/docs](https://docs.newapi.pro/zh/docs) |
| 🎓 **Documentation personnalisée pour l'université** | Répertoire de documentation de ce projet |
| 💬 **Interaction avec la communauté** | [Canaux de communication NEWAPI](https://docs.newapi.pro/zh/docs/support/community-interaction) |
| 🐛 **Retours sur les problèmes** | [Retours sur les problèmes NEWAPI](https://github.com/Calcium-Ion/new-api/issues) |

### 🤝 Guide de contribution

Bienvenue à toutes les formes de contribution !

- 🐛 Signaler des bogues
- 💡 Proposer de nouvelles fonctionnalités
- 📝 Améliorer la documentation
- 🔧 Soumettre du code

---

## 📜 Licence

Ce projet est sous licence [GNU Affero General Public License v3.0 (AGPLv3)](./LICENSE).

Ce projet est un projet open-source développé sur la base de **NEWAPI** ([https://github.com/Calcium-Ion/new-api](https://github.com/Calcium-Ion/new-api)) (sous licence MIT).

Si les politiques de votre organisation ne permettent pas l'utilisation de logiciels sous licence AGPLv3, ou si vous souhaitez éviter les obligations open-source de l'AGPLv3, veuillez nous contacter à : [support@quantumnous.com](mailto:support@quantumnous.com)

---

## 🌟 Historique des étoiles

<div align="center">

[![Graphique de l'historique des étoiles](https://api.star-history.com/svg?repos=Calcium-Ion/new-api&type=Date)](https://star-history.com/#Calcium-Ion/new-api&Date)

</div>

---

<div align="center">

### 💖 Merci d'utiliser New API Édition universitaire personnalisée

Si ce projet vous est utile, n'hésitez pas à donner une ⭐️ Étoile à moi ou à [new-api](https://github.com/Calcium-Ion/new-api) !

**[Documentation officielle NEWAPI](https://docs.newapi.pro/zh/docs)** • **[Retours sur les problèmes](https://github.com/Calcium-Ion/new-api/issues)** • **[Dernière version](https://github.com/Calcium-Ion/new-api/releases)**

<sub>Construit sur la base de NEWAPI ❤️ QuantumNous</sub>

</div>