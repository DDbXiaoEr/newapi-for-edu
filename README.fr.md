<div align="center">

![new-api](/web/public/logo.png)

# New API - Édition universitaire personnalisée

🍥 **Version personnalisée pour l'environnement universitaire basée sur NEWAPI - Passerelle IA et système de gestion d'actifs**

###
Remarque : Toutes les modifications actuelles de ce projet ont été effectuées par une IA. L'auteur a subi une hémorragie cérébrale l'année dernière, entraînant une hémiplégie, et ne peut actuellement utiliser qu'une seule main. De nombreuses fonctionnalités ne sont donc pas encore implémentées.
Je vais à l'hôpital chaque jour pour la rééducation, mon temps et mon énergie sont limités, et sans travail ni revenu, les tokens dépendent entièrement des quotas pour nouveaux utilisateurs et des invitations offertes par les plateformes. L'efficacité du développement a atteint sa limite. Les personnes disposant de tokens en abondance sont les bienvenues pour contribuer)
###Remerciements à GLM et Alibaba Cloud Bailian

> ⚠️ **Avancement actuel** : Seule l'intégration des comptes de campus (connexion par numéro d'étudiant/employé) est actuellement implémentée. Toutes les autres fonctionnalités universitaires personnalisées sont prévues mais PAS ENCORE développées. Les contributions sont les bienvenues !

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
> - Ce projet est une version personnalisée pour l'environnement universitaire basée sur **NEWAPI**, spécialement optimisée pour les environnements existants des établissements d'enseignement supérieur
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

### 🏫 Optimisation pour le contexte éducatif

| Module fonctionnel | Statut | Description |
|---------|------|------|
| 🎓 Gestion des comptes éducatifs | ✅ Fait | Authentification unifiée par numéro d'étudiant/employé, intégration avec le système du campus |
| 🔐 Intégration de la sécurité du campus | 🔜 Prévu | Prise en charge du système d'authentification unifiée du campus |

### 🛠️ Personnalisation technique

| Élément personnalisé | Statut | Description |
|---------|------|------|
| 🏗️ Adaptation réseau | 🔜 Prévu | Optimisation de l'environnement intranet universitaire, prise en charge du proxy et de la configuration du pare-feu |
| 💾 Compatibilité des bases de données | 🔜 Prévu | Optimisation approfondie pour les bases de données courantes (MySQL, PostgreSQL, SQLite) |
| 🔄 Adaptation des interfaces | 🔜 Prévu | Fourniture d'interfaces standard avec les autres systèmes de l'établissement |
| 📱 Adaptation mobile | 🔜 Prévu | Optimisation de l'expérience d'accès mobile, prise en charge de l'intégration avec l'application du campus |

---

## 🚀 Démarrage rapide

### Utilisation de Docker Compose (recommandé)

```bash
# Cloner le projet
git clone [adresse du projet]
cd [répertoire du projet]

# Modifier la configuration docker-compose.yml
nano docker-compose.yml

# Démarrer le service
docker-compose up -d
```

<details>
<summary><strong>Utilisation des commandes Docker</strong></summary>

```bash
# Tirer la dernière image
docker pull [nom de l'image personnalisée]

# Utilisation de SQLite (par défaut)
docker run --name new-api-edu -d --restart always \
  -p 3000:3000 \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  [nom de l'image personnalisée]:latest

# Utilisation de MySQL
docker run --name new-api-edu -d --restart always \
  -p 3000:3000 \
  -e SQL_DSN="root:123456@tcp(localhost:3306)/oneapi" \
  -e TZ=Asia/Shanghai \
  -v ./data:/data \
  [nom de l'image personnalisée]:latest
```

> **💡 Astuce:** `-v ./data:/data` sauvegardera les données dans le dossier `data` du répertoire actuel, vous pouvez également le changer en chemin absolu comme `-v /your/custom/path:/data`

</details>

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
**Documentation des fonctionnalités personnalisées :**
- Configuration de la gestion des comptes universitaires
- Guide d'intégration des systèmes du campus
- Instructions de configuration des autorisations éducatives
- Fonctionnalités d'analyse statistique des données

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

### 🎓 Nouvelles fonctionnalités personnalisées pour l'université

- 🏫 **Authentification du campus** (✅ Fait) : prise en charge de la connexion par numéro d'étudiant/employé
- 📚 **Gestion des cours** (🔜 Prévu) : attribution des droits d'utilisation de l'IA par cours
- 👨‍🏫 **Gestion des enseignants** (🔜 Prévu) : les enseignants peuvent gérer les étudiants de leur classe
- 📊 **Statistiques éducatives** (🔜 Prévu) : analyse des données d'utilisation de l'IA
- 🔐 **Audit de sécurité** (🔜 Prévu) : journalisation complète des opérations et fonctionnalités d'audit

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
| **Base de données locale** | SQLite (Docker doit monter le répertoire `/data`) |
| **Base de données distante** | MySQL ≥ 5.7.8 ou PostgreSQL ≥ 9.6 |
| **Moteur de conteneur** | Docker / Docker Compose |
| **Environnement réseau** | Prise en charge de la configuration de l'environnement intranet universitaire |

### ⚙️ Variables d'environnement personnalisées pour l'université

| Nom de variable | Description | Valeur par défaut |
|--------|------|--------|
| `EDU_MODE` | Activer le mode universitaire | `true` |
| `CAMPUS_AUTH_URL` | Adresse d'authentification du campus | - |
| `CAMPUS_API_KEY` | Clé API du campus | - |
| `EDU_DOMAIN` | Restriction de domaine éducatif | - |
| `ALLOWED_DOMAINS` | Liste des domaines autorisés | - |

📖 **Configuration complète** : veuillez consulter la documentation des variables d'environnement NEWAPI + la description de la configuration personnalisée pour l'université

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

Si ce projet vous est utile, n'hésitez pas à nous donner une ⭐️ Étoile !

**[Documentation officielle NEWAPI](https://docs.newapi.pro/zh/docs)** • **[Retours sur les problèmes](https://github.com/Calcium-Ion/new-api/issues)** • **[Dernière version](https://github.com/Calcium-Ion/new-api/releases)**

<sub>Construit sur la base de NEWAPI ❤️ QuantumNous</sub>

</div>