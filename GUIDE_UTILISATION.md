# Go-Scaffold - Générateur de Code Automatique pour Go

## 📦 Contenu du Package

Vous avez téléchargé **go-scaffold**, un générateur de code automatique pour Go similaire à Laravel Artisan, mais adapté pour l'écosystème Go avec Gin, GORM et PostgreSQL.

### Fichiers inclus

```
go-scaffold/
├── main.go                          # Point d'entrée du générateur
├── go.mod                           # Dépendances Go
├── cmd/                             # Commandes CLI
│   ├── root.go                      # Commande racine
│   ├── init.go                      # Commande d'initialisation
│   ├── make.go                      # Commandes make:*
│   └── generate.go                  # Commande de génération
├── internal/
│   ├── generator/                   # Générateurs de code
│   │   ├── generator.go             # Générateur principal
│   │   ├── model.go                 # Génère les models
│   │   ├── repository.go            # Génère les repositories
│   │   ├── controller.go            # Génère les contrôleurs
│   │   ├── request.go               # Génère les validations
│   │   └── routes.go                # Génère les routes
│   └── parser/
│       └── parser.go                # Parse les schémas YAML
├── examples/                        # Exemples de schémas
│   ├── user_schema.yaml             # Exemple utilisateur
│   └── post_schema.yaml             # Exemple post
├── install.sh                       # Script d'installation
├── Makefile                         # Commandes make
├── README.md                        # Documentation complète
├── QUICKSTART.md                    # Guide de démarrage rapide
├── ARCHITECTURE.md                  # Documentation d'architecture
├── CONTRIBUTING.md                  # Guide de contribution
├── CHANGELOG.md                     # Historique des versions
└── LICENSE                          # Licence MIT
```

## 🚀 Installation Rapide

### Méthode 1 : Script d'installation (Recommandé)

```bash
# Extraire l'archive
tar -xzf go-scaffold.tar.gz
cd go-scaffold

# Rendre le script exécutable et l'exécuter
chmod +x install.sh
./install.sh
```

### Méthode 2 : Installation manuelle

```bash
# Extraire l'archive
tar -xzf go-scaffold.tar.gz
cd go-scaffold

# Compiler
go build -o go-scaffold main.go

# Installer globalement (optionnel)
sudo mv go-scaffold /usr/local/bin/
```

### Méthode 3 : Avec Make

```bash
# Extraire l'archive
tar -xzf go-scaffold.tar.gz
cd go-scaffold

# Compiler et installer
make install
```

## 📖 Utilisation

### 1. Créer un nouveau projet

```bash
go-scaffold init mon-api
cd mon-api
```

Cela crée une structure complète :
- Configuration de la base de données
- Structure MVC
- Fichiers de configuration
- Point d'entrée main.go

### 2. Créer un schéma de table

```bash
go-scaffold make:schema produit
```

Éditez `database/schemas/produit.yaml` :

```yaml
table: produits
model: Produit

columns:
  - name: id
    type: uuid
    primary: true
    nullable: false

  - name: nom
    type: string
    size: 255
    nullable: false

  - name: description
    type: text
    nullable: true

  - name: prix
    type: float
    nullable: false

  - name: quantite
    type: integer
    nullable: false
    default: 0

  - name: actif
    type: boolean
    nullable: false
    default: true

  - name: created_at
    type: timestamp
    nullable: false

  - name: updated_at
    type: timestamp
    nullable: false

validations:
  - field: nom
    rules:
      required: true
      min: 3
      max: 255

  - field: prix
    rules:
      required: true
      min: 0

  - field: quantite
    rules:
      required: true
      min: 0
```

### 3. Générer le code

```bash
go-scaffold generate database/schemas/produit.yaml
```

Cela génère automatiquement :
- ✅ Model GORM complet
- ✅ Repository avec CRUD
- ✅ Contrôleur RESTful
- ✅ Validations de requêtes
- ✅ Routes configurées

### 4. Configurer et lancer

```bash
# Copier la configuration
cp .env.example .env

# Éditer .env avec vos infos
nano .env

# Installer les dépendances
go mod download

# Lancer l'application
go run main.go
```

Votre API sera disponible sur `http://localhost:8080` 🎉

## 🔥 Fonctionnalités

### Génération automatique

- **Models** : Structures Go avec GORM, tags JSON et validation
- **Repositories** : Pattern Repository avec méthodes CRUD
- **Contrôleurs** : Endpoints RESTful avec Gin
- **Requests** : Validation automatique avec go-playground/validator
- **Routes** : Configuration automatique des routes

### Relations supportées

- `belongs_to` - Many-to-One
- `has_many` - One-to-Many
- `has_one` - One-to-One
- `many_to_many` - Many-to-Many avec table pivot

### Validations

- `required`, `min`, `max`
- `email`, `url`
- `in` (enum)
- `regex` (custom)

### Fonctionnalités avancées

- Pagination automatique
- Gestion des erreurs
- Documentation Swagger intégrée
- Support des index de base de données
- Hooks GORM (BeforeCreate, BeforeUpdate)

## 📚 Documentation

### Guides disponibles

1. **README.md** - Documentation complète avec toutes les fonctionnalités
2. **QUICKSTART.md** - Guide de démarrage rapide (5 minutes)
3. **ARCHITECTURE.md** - Architecture technique détaillée
4. **CONTRIBUTING.md** - Guide pour contribuer au projet
5. **CHANGELOG.md** - Historique des versions

### Exemples

Le dossier `examples/` contient :
- `user_schema.yaml` - Schéma complet d'utilisateur avec toutes les fonctionnalités
- `post_schema.yaml` - Schéma de post avec relations

## 🎯 Exemple complet

### Créer une API de blog

```bash
# 1. Initialiser le projet
go-scaffold init blog-api
cd blog-api

# 2. Créer les schémas
go-scaffold make:schema user
go-scaffold make:schema post
go-scaffold make:schema comment

# 3. Éditer les schémas (voir exemples/)

# 4. Générer tout le code
go-scaffold generate --all

# 5. Configurer la base de données
cp .env.example .env
# Éditer .env

# 6. Lancer
go run main.go
```

### Endpoints générés

Pour chaque model, vous obtenez automatiquement :

```
GET    /api/users              # Liste avec pagination
POST   /api/users              # Créer
GET    /api/users/:id          # Afficher un
PUT    /api/users/:id          # Mettre à jour
DELETE /api/users/:id          # Supprimer
```

### Exemple de requête

```bash
# Créer un utilisateur
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "secret123"
  }'

# Liste avec pagination
curl "http://localhost:8080/api/users?page=1&page_size=10"
```

## 🛠️ Commandes disponibles

```bash
# Initialiser un projet
go-scaffold init [nom]

# Créer un schéma
go-scaffold make:schema [nom]

# Créer une migration
go-scaffold make:migration [nom]

# Générer le code
go-scaffold generate [chemin-schema]
go-scaffold generate --all

# Aide
go-scaffold --help
go-scaffold [commande] --help
```

## 💡 Astuces

### Utiliser le Makefile

Le projet généré inclut un Makefile pratique :

```bash
make build          # Compiler
make run            # Compiler et exécuter
make test           # Exécuter les tests
make fmt            # Formater le code
make clean          # Nettoyer
```

### Commandes rapides avec Make

Dans le répertoire go-scaffold :

```bash
make install                    # Installer globalement
make generate-all               # Générer tous les schémas
make make-schema NAME=category  # Créer un schéma
```

### Personnalisation

Après génération, vous pouvez :
- Ajouter des méthodes dans les repositories
- Personnaliser les contrôleurs
- Ajouter des middlewares
- Modifier les validations
- Ajouter de la logique métier

## 🔄 Workflow recommandé

1. **Design** : Concevez votre base de données
2. **Schémas** : Créez les fichiers YAML
3. **Génération** : Générez le code
4. **Personnalisation** : Ajoutez votre logique métier
5. **Tests** : Testez vos endpoints
6. **Déploiement** : Déployez votre API

## 🌟 Comparaison avec Laravel

| Laravel | Go-Scaffold |
|---------|-------------|
| `php artisan make:model` | `go-scaffold make:schema` |
| `php artisan make:controller` | Généré automatiquement |
| `php artisan make:migration` | `go-scaffold make:migration` |
| `php artisan make:request` | Généré automatiquement |
| Eloquent ORM | GORM |
| Laravel Routes | Gin Routes |
| Blade Templates | (API REST seulement) |

## 🤝 Support

### Besoin d'aide ?

1. Consultez la documentation (README.md, QUICKSTART.md)
2. Vérifiez les exemples dans `examples/`
3. Lisez ARCHITECTURE.md pour comprendre le fonctionnement
4. Ouvrez une issue sur GitHub

### Contribuer

Consultez CONTRIBUTING.md pour savoir comment contribuer au projet.

## 📝 Licence

MIT - Voir le fichier LICENSE

## 🎉 Prochaines étapes

1. Lisez QUICKSTART.md pour démarrer en 5 minutes
2. Explorez les exemples dans `examples/`
3. Créez votre premier projet
4. Expérimentez avec les relations et validations
5. Partagez vos retours et suggestions

---

**Bon codage avec go-scaffold !** 🚀

Si vous avez des questions ou des suggestions, n'hésitez pas à contribuer au projet.
