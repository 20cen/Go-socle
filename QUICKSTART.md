# Guide de Démarrage Rapide 🚀

Ce guide vous aidera à démarrer rapidement avec go-scaffold.

## Installation en 5 minutes

### 1. Installer le générateur

```bash
# Cloner le repository
git clone <votre-repo>
cd go-scaffold

# Compiler et installer
make install
```

### 2. Créer votre premier projet

```bash
# Initialiser un nouveau projet
go-scaffold init blog-api
cd blog-api
```

### 3. Créer votre premier schéma

```bash
# Créer un schéma pour les articles
go-scaffold make:schema article
```

Éditez `database/schemas/article.yaml` :

```yaml
table: articles
model: Article

columns:
  - name: id
    type: uuid
    primary: true
    nullable: false

  - name: title
    type: string
    size: 255
    nullable: false

  - name: content
    type: text
    nullable: false

  - name: author
    type: string
    size: 100
    nullable: false

  - name: published
    type: boolean
    nullable: false
    default: false

  - name: created_at
    type: timestamp
    nullable: false

  - name: updated_at
    type: timestamp
    nullable: false

validations:
  - field: title
    rules:
      required: true
      min: 5
      max: 255

  - field: content
    rules:
      required: true
      min: 50

  - field: author
    rules:
      required: true
      min: 3
      max: 100
```

### 4. Générer le code

```bash
go-scaffold generate database/schemas/article.yaml
```

Cette commande va générer automatiquement :
- ✅ `app/models/article.go`
- ✅ `app/repositories/article_repository.go`
- ✅ `app/controllers/article_controller.go`
- ✅ `app/requests/article_request.go`
- ✅ `routes/article_routes.go`

### 5. Configurer la base de données

```bash
# Copier le fichier d'environnement
cp .env.example .env

# Éditer .env avec vos informations
nano .env
```

Exemple de configuration :

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=monmotdepasse
DB_NAME=blog_api
SERVER_PORT=8080
```

### 6. Installer les dépendances

```bash
go mod download
```

### 7. Créer la base de données

```bash
# Créer la base de données PostgreSQL
createdb blog_api

# Ou avec psql
psql -U postgres -c "CREATE DATABASE blog_api;"
```

### 8. Lancer l'application

```bash
go run main.go
```

Votre API sera disponible sur `http://localhost:8080` 🎉

## Tester votre API

### Créer un article

```bash
curl -X POST http://localhost:8080/api/articles \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Mon premier article",
    "content": "Ceci est le contenu de mon premier article. Il doit contenir au moins 50 caractères pour passer la validation.",
    "author": "John Doe",
    "published": true
  }'
```

### Lister les articles

```bash
curl http://localhost:8080/api/articles?page=1&page_size=10
```

### Afficher un article

```bash
curl http://localhost:8080/api/articles/{id}
```

### Mettre à jour un article

```bash
curl -X PUT http://localhost:8080/api/articles/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Mon article mis à jour",
    "published": true
  }'
```

### Supprimer un article

```bash
curl -X DELETE http://localhost:8080/api/articles/{id}
```

## Prochaines étapes

### Ajouter d'autres modèles

```bash
# Créer un schéma pour les commentaires
go-scaffold make:schema comment

# Générer le code
go-scaffold generate database/schemas/comment.yaml
```

### Ajouter des relations

Éditez votre schéma pour ajouter des relations :

```yaml
relations:
  - type: belongs_to
    model: Article
    foreign_key: article_id
```

### Personnaliser le code

Vous pouvez modifier les fichiers générés pour ajouter :
- Des méthodes personnalisées dans les repositories
- Des middlewares d'authentification
- Des transformations de données
- Des hooks GORM
- De la logique métier dans les services

### Ajouter l'authentification

```bash
# Créer un schéma utilisateur
go-scaffold make:schema user

# Générer le code
go-scaffold generate database/schemas/user.yaml

# Ajouter JWT et bcrypt dans go.mod
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
```

### Ajouter des middlewares

Créez `middleware/auth.go` :

```go
package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Token requis",
            })
            c.Abort()
            return
        }
        
        // Vérifier le token ici
        
        c.Next()
    }
}
```

Puis dans `routes/routes.go` :

```go
import "votre-projet/middleware"

// Protéger les routes
api.Use(middleware.Auth())
```

## Astuces

### Utiliser le Makefile

```bash
# Compiler
make build

# Installer globalement
make install

# Générer pour tous les schémas
make generate-all

# Créer un nouveau schéma
make make-schema NAME=category

# Formater le code
make fmt
```

### Générer pour plusieurs schémas

```bash
# Générer pour tous les schémas en une commande
go-scaffold generate --all
```

### Structure recommandée des schémas

```
database/schemas/
├── user.yaml          # Utilisateurs
├── post.yaml          # Articles/Posts
├── comment.yaml       # Commentaires
├── category.yaml      # Catégories
└── tag.yaml           # Tags
```

## Besoin d'aide ?

- 📖 Consultez le [README complet](README.md)
- 💬 Ouvrez une issue sur GitHub
- 📧 Contactez l'équipe de support

Bon codage ! 🚀
