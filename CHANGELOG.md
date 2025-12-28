# Changelog

Toutes les modifications notables de ce projet seront documentées dans ce fichier.

Le format est basé sur [Keep a Changelog](https://keepachangelog.com/fr/1.0.0/),
et ce projet adhère au [Semantic Versioning](https://semver.org/lang/fr/).

## [1.0.0] - 2024-01-XX

### Ajouté
- ✨ Commande `init` pour initialiser un nouveau projet Go
- ✨ Commande `make:schema` pour créer des fichiers de schéma YAML
- ✨ Commande `make:migration` pour créer des migrations
- ✨ Commande `generate` pour générer le code à partir des schémas
- ✨ Génération automatique des models GORM
- ✨ Génération automatique des repositories avec pattern Repository
- ✨ Génération automatique des contrôleurs RESTful
- ✨ Génération automatique des structs de validation (requests)
- ✨ Génération automatique des routes Gin
- ✨ Support des relations (belongsTo, hasMany, hasOne, manyToMany)
- ✨ Support des validations personnalisées
- ✨ Support des index de base de données
- ✨ Pagination automatique dans les contrôleurs
- ✨ Documentation Swagger intégrée dans les contrôleurs
- ✨ Support de PostgreSQL par défaut
- ✨ Structure de projet organisée et professionnelle
- ✨ Fichiers d'exemple (user_schema.yaml, post_schema.yaml)
- ✨ Makefile pour faciliter l'utilisation
- ✨ Script d'installation automatique
- ✨ Documentation complète (README, QUICKSTART)

### Types de colonnes supportés
- `string`, `text` - Types texte
- `int`, `integer`, `bigint`, `smallint` - Types numériques entiers
- `float`, `double`, `decimal` - Types numériques décimaux
- `boolean`, `bool` - Type booléen
- `date`, `datetime`, `timestamp`, `time` - Types temporels
- `uuid` - Identifiants uniques universels
- `json`, `jsonb` - Types JSON

### Règles de validation supportées
- `required` - Champ requis
- `min` - Longueur/valeur minimale
- `max` - Longueur/valeur maximale
- `email` - Format email
- `url` - Format URL
- `in` - Valeur dans une liste
- `regex` - Expression régulière personnalisée

### Relations supportées
- `belongs_to` - Many-to-One
- `has_many` - One-to-Many
- `has_one` - One-to-One
- `many_to_many` - Many-to-Many avec table pivot

## [À venir]

### Prévu pour v1.1.0
- 🔄 Support de MySQL et SQLite
- 🔄 Génération de tests unitaires
- 🔄 Génération de seeders
- 🔄 Migration automatique des tables
- 🔄 Support de l'authentification JWT
- 🔄 Génération de documentation API (OpenAPI/Swagger)
- 🔄 CLI interactive pour la création de schémas
- 🔄 Templates personnalisables
- 🔄 Support des événements (observers)
- 🔄 Support des jobs/queues

### Prévu pour v1.2.0
- 🔄 Support de GraphQL
- 🔄 Génération de clients API
- 🔄 Support des websockets
- 🔄 Génération de webhooks
- 🔄 Support du caching (Redis)
- 🔄 Monitoring et logging avancés

### Idées futures
- 📝 Dashboard web pour gérer les schémas
- 📝 Support de Docker/Kubernetes
- 📝 Intégration CI/CD
- 📝 Générateur de frontend (React, Vue)
- 📝 Support de microservices
- 📝 Support de gRPC

## Notes de version

### v1.0.0 - Version initiale
Première version stable de go-scaffold avec toutes les fonctionnalités de base pour générer rapidement des APIs RESTful en Go. Cette version fournit une base solide pour le développement rapide d'applications backend en Go avec une architecture propre et maintenable.

Le générateur crée automatiquement tout le code boilerplate nécessaire, permettant aux développeurs de se concentrer sur la logique métier plutôt que sur la configuration et la structure du projet.

### Comparaison avec Laravel Artisan
- ✅ Génération de models (équivalent à `php artisan make:model`)
- ✅ Génération de contrôleurs (équivalent à `php artisan make:controller`)
- ✅ Génération de migrations (équivalent à `php artisan make:migration`)
- ✅ Génération de requests (équivalent à `php artisan make:request`)
- ✅ Support des relations Eloquent
- ✅ Support des validations
- ✅ Structure de projet organisée

## Contributeurs

Merci à tous les contributeurs qui ont participé à ce projet !

---

Pour toute question ou suggestion, n'hésitez pas à ouvrir une issue sur GitHub.
