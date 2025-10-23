# 🕒 Time Manager

![Go Version](https://img.shields.io/badge/Go-1.25.1-blue?logo=go)
![Angular](https://img.shields.io/badge/Angular-18-red?logo=angular)
![Node.js](https://img.shields.io/badge/Node.js-22.13-green?logo=node.js)
![Built with](https://img.shields.io/badge/Built%20With-Docker-blue?logo=docker)
[![CI Build](https://github.com/AristideWafo/Time-Manager/actions/workflows/ci.yml/badge.svg)](https://github.com/AristideWafo/Time-Manager/actions/workflows/ci.yml)
[![CodeQL](https://github.com/AristideWafo/Time-Manager/actions/workflows/codeql.yml/badge.svg)](https://github.com/AristideWafo/Time-Manager/actions/workflows/codeql.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> 🚀 **Application de gestion du temps moderne** construite avec Go (Gin), Angular, et PostgreSQL. Solution complète pour le suivi du temps de travail avec authentification sécurisée et interface utilisateur intuitive.

## 📋 Table des matières

- [✨ Fonctionnalités](#-fonctionnalités)
- [🏗️ Architecture](#️-architecture)
- [🚀 Démarrage rapide](#-démarrage-rapide)
- [📦 Installation](#-installation)
- [🛠️ Développement](#️-développement)
- [🧪 Tests](#-tests)
- [🚢 Déploiement](#-déploiement)
- [📚 API Documentation](#-api-documentation)
- [🤝 Contribution](#-contribution)
- [📄 License](#-license)

## ✨ Fonctionnalités

### 🔐 Authentification & Autorisation
- Système d'authentification JWT sécurisé
- Gestion des rôles utilisateurs (Admin, Employé, Manager)
- Middleware de protection des routes

### ⏱️ Gestion du temps
- **Horloges de travail** : Pointage d'entrée/sortie
- **Suivi des présences** : Historique complet des temps de travail
- **Tableaux de bord** : Visualisation des données de temps
- **Rapports** : Génération de rapports personnalisés

### 👥 Gestion des utilisateurs
- Profils utilisateurs complets
- Gestion des équipes et départements
- Interface d'administration

### 🎨 Interface moderne
- Design responsive et intuitif
- Progressive Web App (PWA)
- Mode sombre/clair
- Notifications en temps réel

## 🏗️ Architecture

```
Time-Manager/
├── 🎯 front/           # Application Angular
│   ├── src/
│   │   ├── app/        # Composants et services
│   │   ├── assets/     # Ressources statiques
│   │   └── environments/
│   ├── Dockerfile
│   └── package.json
├── 🚀 back/            # API Go (Gin)
│   ├── controllers/    # Contrôleurs API
│   ├── middleware/     # Middlewares
│   ├── models/         # Modèles de données
│   ├── routes/         # Définition des routes
│   ├── repository/     # Couche d'accès aux données
│   ├── Dockerfile
│   └── go.mod
├── 🌐 nginx/           # Reverse proxy
│   ├── default.conf
│   └── Dockerfile
├── 🐳 docker-compose-dev.yml
├── 🐳 docker-compose.yml
└── 📋 .github/workflows/  # CI/CD Pipelines
```

### 🛠️ Stack technique

**Backend**
- ![Go](https://img.shields.io/badge/Go-1.25.1-blue?logo=go) **Go** avec framework **Gin**
- ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-blue?logo=postgresql) **PostgreSQL** pour la base de données
- ![JWT](https://img.shields.io/badge/JWT-Auth-orange) **JWT** pour l'authentification
- ![Swagger](https://img.shields.io/badge/Swagger-API-green) **Swagger** pour la documentation API

**Frontend**
- ![Angular](https://img.shields.io/badge/Angular-18-red?logo=angular) **Angular 18** avec TypeScript
- ![RxJS](https://img.shields.io/badge/RxJS-Reactive-purple) **RxJS** pour la programmation réactive
- ![Angular Material](https://img.shields.io/badge/Material-UI-blue) **Angular Material** pour les composants UI

**DevOps**
- ![Docker](https://img.shields.io/badge/Docker-Containerization-blue?logo=docker) **Docker** & Docker Compose
- ![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-CI/CD-blue?logo=github) **GitHub Actions** pour CI/CD
- ![Nginx](https://img.shields.io/badge/Nginx-Proxy-green?logo=nginx) **Nginx** comme reverse proxy

## 🚀 Démarrage rapide

### Prérequis
- ![Docker](https://img.shields.io/badge/Docker-20+-blue?logo=docker) Docker & Docker Compose
- ![Git](https://img.shields.io/badge/Git-Latest-orange?logo=git) Git

### Lancement avec Docker

```bash
# Cloner le repository
git clone https://github.com/AristideWafo/Time-Manager.git
cd Time-Manager

# Lancer l'application complète
docker compose up --build

# Ou en mode développement avec live reload
docker compose -f docker-compose-dev.yml up --build
```

🌐 **Accès aux services :**
- **Application** : http://localhost
- **API Backend** : http://localhost:8080
- **Frontend Dev** : http://localhost:4200
- **API Documentation** : http://localhost:8080/swagger/index.html

## 📦 Installation

### Installation locale (Développement)

<details>
<summary>🔧 Setup Backend (Go)</summary>

```bash
cd back

# Installer les dépendances
go mod download

# Copier la configuration
cp .env.example .env

# Configurer la base de données dans .env
DATABASE_URL=postgresql://user:password@localhost:5432/timemanager

# Lancer l'API
go run main.go
```
</details>

<details>
<summary>🎨 Setup Frontend (Angular)</summary>

```bash
cd front

# Installer les dépendances
npm install

# Lancer en mode développement
npm start

# Ou avec configuration spécifique
ng serve --configuration development
```
</details>

<details>
<summary>🗄️ Setup Base de données</summary>

```bash
# Avec Docker
docker run --name postgres-timemanager \
  -e POSTGRES_DB=timemanager \
  -e POSTGRES_USER=admin \
  -e POSTGRES_PASSWORD=password \
  -p 5432:5432 -d postgres:16

# Ou installer PostgreSQL localement
# Créer la base de données
createdb timemanager
```
</details>

## 🛠️ Développement

### Scripts disponibles

**Frontend (Angular)**
```bash
npm run start          # Serveur de développement
npm run build          # Build de production
npm run test           # Tests unitaires
npm run lint           # Linting du code
npm run lint:fix       # Correction automatique du linting
npm run e2e            # Tests end-to-end
```

**Backend (Go)**
```bash
go run main.go         # Lancer l'API
go test ./...          # Tests unitaires
go mod tidy            # Nettoyer les dépendances
air                    # Live reload (si Air est installé)
```

### 🔧 Configuration de l'environnement

<details>
<summary>Variables d'environnement Backend (.env)</summary>

```env
# Base de données
DATABASE_URL=postgresql://user:password@localhost:5432/timemanager
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=password
DB_NAME=timemanager

# JWT
JWT_SECRET=your-super-secret-key
JWT_EXPIRES_IN=24h

# Serveur
PORT=8080
ENV=development

# Cors
CORS_ORIGINS=http://localhost:4200,http://localhost:3000
```
</details>

<details>
<summary>Configuration Frontend (environment.ts)</summary>

```typescript
export const environment = {
  production: false,
  apiUrl: 'http://localhost:8080',
  apiVersion: 'v1',
  enableDevTools: true
};
```
</details>

## 🧪 Tests

### Tests automatisés
```bash
# Lancer tous les tests
npm run test:all

# Tests backend
cd back && go test -v ./...

# Tests frontend avec couverture
cd front && npm run test:coverage

# Tests E2E
cd front && npm run e2e
```

### Qualité du code
```bash
# Linting
npm run lint:fix

# Analyse de sécurité
npm audit
go list -json -m all | nancy sleuth
```

## 🚢 Déploiement

### 🐳 Déploiement Docker

```bash
# Build des images
docker compose build

# Déploiement en production
docker compose -f docker-compose.yml up -d

# Monitoring
docker compose logs -f
```

### ☁️ Déploiement Cloud

<details>
<summary>Déploiement sur des plateformes cloud</summary>

**Heroku**
```bash
# Backend
heroku create timemanager-api
git subtree push --prefix=back heroku main

# Frontend
heroku create timemanager-web
git subtree push --prefix=front heroku main
```

**AWS/GCP/Azure**
- Utilisation des images Docker construites
- Configuration des variables d'environnement
- Setup de la base de données cloud
</details>

## 📚 API Documentation

La documentation interactive de l'API est disponible via Swagger :

🔗 **[Documentation API](http://localhost:8080/swagger/index.html)**

### Endpoints principaux

<details>
<summary>📋 Liste des endpoints</summary>

```
Authentication
POST   /api/auth/login
POST   /api/auth/register
POST   /api/auth/logout
GET    /api/auth/me

Users
GET    /api/users
GET    /api/users/:id
POST   /api/users
PUT    /api/users/:id
DELETE /api/users/:id

Work Times
GET    /api/worktimes
GET    /api/worktimes/:id
POST   /api/worktimes
PUT    /api/worktimes/:id
DELETE /api/worktimes/:id

Clocks
POST   /api/clocks/:userID
GET    /api/clocks/:userID
```
</details>

## 🔄 CI/CD

Le projet utilise **GitHub Actions** pour l'intégration et le déploiement continus :

- ✅ **Tests automatiques** sur chaque PR
- ✅ **Linting** et vérification de la qualité du code
- ✅ **Analyse de sécurité** avec CodeQL
- ✅ **Build Docker** automatique
- ✅ **Déploiement** automatique en staging
- ✅ **Release** automatique avec semantic-versioning

### Badges de statut
[![Tests](https://github.com/AristideWafo/Time-Manager/actions/workflows/test.yml/badge.svg)](https://github.com/AristideWafo/Time-Manager/actions/workflows/test.yml)
[![Build](https://github.com/AristideWafo/Time-Manager/actions/workflows/build.yml/badge.svg)](https://github.com/AristideWafo/Time-Manager/actions/workflows/build.yml)
[![Security](https://github.com/AristideWafo/Time-Manager/actions/workflows/codeql.yml/badge.svg)](https://github.com/AristideWafo/Time-Manager/actions/workflows/codeql.yml)

## 🤝 Contribution

Les contributions sont les bienvenues ! Veuillez suivre ces étapes :

1. **Forkez** le repository
2. **Créez** une branche feature (`git checkout -b feat/amazing-feature`)
3. **Committez** vos changements (`git commit -m 'feat: add amazing feature'`)
4. **Pushez** vers la branche (`git push origin feat/amazing-feature`)
5. **Ouvrez** une Pull Request

### Conventions de commits
Nous utilisons [Conventional Commits](https://www.conventionalcommits.org/) :

```
feat: nouvelle fonctionnalité
fix: correction de bug
docs: documentation
style: formatage, points-virgules manquants, etc.
refactor: refactoring du code
test: ajout de tests
chore: maintenance
```

## 📊 Roadmap

- [ ] 📱 Application mobile (React Native/Flutter)
- [ ] 📈 Analytics avancés et rapports
- [ ] 🔔 Notifications push
- [ ] 📅 Intégration calendrier
- [ ] 🌍 Internationalisation (i18n)
- [ ] 🤖 API webhooks
- [ ] 📊 Dashboard temps réel

## 🐛 Signaler un bug

Si vous trouvez un bug, veuillez [créer une issue](https://github.com/AristideWafo/Time-Manager/issues/new) avec :
- Description détaillée du problème
- Étapes pour reproduire
- Screenshots si applicable
- Informations sur votre environnement

## 📄 License

Ce projet est sous licence **MIT**. Voir le fichier [LICENSE](LICENSE) pour plus de détails.

---

<div align="center">


⭐ **N'oubliez pas de donner une étoile si ce projet vous a aidé !** ⭐

</div>