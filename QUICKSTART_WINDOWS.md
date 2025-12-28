# 🚀 Démarrage Rapide - Windows

Guide ultra-rapide pour démarrer avec go-scaffold sur Windows en 10 minutes !

## ⚡ Installation Express (5 minutes)

### Étape 1 : Installer Go

1. **Télécharger Go**
   - Aller sur : https://go.dev/dl/
   - Cliquer sur `go1.21.x.windows-amd64.msi`
   - Exécuter l'installeur
   - Cliquer "Next" → "Next" → "Install"

2. **Vérifier l'installation**
   ```powershell
   # Ouvrir PowerShell et taper :
   go version
   ```
   
   Si ça ne fonctionne pas, redémarrez PowerShell.

### Étape 2 : Compiler go-scaffold

1. **Extraire l'archive**
   - Clic droit sur `go-scaffold.tar.gz`
   - Extraire tout
   - Ouvrir le dossier `go-scaffold` dans PowerShell

2. **Option A - Script automatique (Recommandé)**
   ```powershell
   # Dans PowerShell
   .\install.ps1
   ```
   
   Suivez les instructions à l'écran !

3. **Option B - Compilation manuelle**
   ```powershell
   # Télécharger les dépendances
   go mod download
   
   # Compiler
   go build -o go-scaffold.exe main.go
   
   # Tester
   .\go-scaffold.exe --help
   ```

### Étape 3 : Installer PostgreSQL

1. **Télécharger**
   - Aller sur : https://www.postgresql.org/download/windows/
   - Télécharger l'installeur
   - Exécuter et suivre l'assistant
   - **IMPORTANT** : Noter le mot de passe choisi !

2. **Créer la base de données**
   - Ouvrir pgAdmin (installé avec PostgreSQL)
   - Clic droit sur "Databases" → "Create" → "Database"
   - Nom : `blog_api`
   - Cliquer "Save"

## 🎯 Premier Projet (5 minutes)

### Créer le projet

```powershell
# Créer un nouveau projet
.\go-scaffold.exe init blog-api

# Accéder au projet
cd blog-api
```

### Créer un schéma

```powershell
# Créer le fichier de schéma
..\go-scaffold.exe make:schema article
```

### Éditer le schéma

Ouvrir `database\schemas\article.yaml` avec Notepad ou VSCode :

```yaml
table: articles
model: Article

columns:
  - name: id
    type: uuid
    primary: true
    nullable: false

  - name: titre
    type: string
    size: 255
    nullable: false

  - name: contenu
    type: text
    nullable: false

  - name: auteur
    type: string
    size: 100
    nullable: false

  - name: publie
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
  - field: titre
    rules:
      required: true
      min: 5
      max: 255

  - field: contenu
    rules:
      required: true
      min: 50

  - field: auteur
    rules:
      required: true
      min: 3
```

### Générer le code

```powershell
# Générer tout le code automatiquement
..\go-scaffold.exe generate database\schemas\article.yaml
```

**Cela crée automatiquement :**
- ✅ `app\models\article.go`
- ✅ `app\repositories\article_repository.go`
- ✅ `app\controllers\article_controller.go`
- ✅ `app\requests\article_request.go`
- ✅ `routes\article_routes.go`

### Configurer la base de données

```powershell
# Copier le fichier d'exemple
copy .env.example .env

# Éditer avec Notepad
notepad .env
```

Modifier les valeurs :
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=votre_mot_de_passe_postgres
DB_NAME=blog_api
SERVER_PORT=8080
```

### Installer les dépendances

```powershell
go mod download
```

### Lancer l'application

```powershell
go run main.go
```

**Votre API est prête !** 🎉

## 🧪 Tester l'API

### Avec un navigateur

Ouvrir : http://localhost:8080/api/health

Vous devriez voir :
```json
{
  "status": "ok",
  "message": "Service en cours d'exécution"
}
```

### Avec PowerShell

```powershell
# Créer un article
$body = @{
    titre = "Mon premier article"
    contenu = "Ceci est le contenu de mon premier article avec go-scaffold. Il faut au moins 50 caractères pour que la validation passe."
    auteur = "Jean Dupont"
    publie = $true
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/api/articles" -Method POST -Body $body -ContentType "application/json"

# Lister les articles
Invoke-RestMethod -Uri "http://localhost:8080/api/articles"
```

### Avec Postman

1. **Ouvrir Postman**
2. **Créer une requête POST**
   - URL : `http://localhost:8080/api/articles`
   - Headers : `Content-Type: application/json`
   - Body (JSON) :
     ```json
     {
       "titre": "Mon article",
       "contenu": "Un contenu qui fait plus de 50 caractères pour passer la validation.",
       "auteur": "John Doe",
       "publie": true
     }
     ```
3. **Cliquer "Send"**

## 📚 Endpoints disponibles

Automatiquement créés pour chaque model :

| Méthode | URL | Description |
|---------|-----|-------------|
| GET | `/api/articles` | Liste avec pagination |
| POST | `/api/articles` | Créer un article |
| GET | `/api/articles/:id` | Afficher un article |
| PUT | `/api/articles/:id` | Modifier un article |
| DELETE | `/api/articles/:id` | Supprimer un article |

## 🎨 Ajouter d'autres models

```powershell
# Créer un schéma pour les commentaires
..\go-scaffold.exe make:schema commentaire

# Éditer database\schemas\commentaire.yaml
# Puis générer
..\go-scaffold.exe generate database\schemas\commentaire.yaml
```

## 🔧 Commandes utiles

### Compilation

```powershell
# Compiler votre API en binaire
go build -o blog-api.exe main.go

# Exécuter le binaire
.\blog-api.exe
```

### Gestion du projet

```powershell
# Voir tous les schémas
dir database\schemas

# Générer pour tous les schémas
..\go-scaffold.exe generate --all

# Nettoyer les dépendances
go mod tidy
```

### Base de données

```powershell
# Se connecter avec psql
psql -U postgres -d blog_api

# Voir les tables
\dt

# Quitter psql
\q
```

## 🛠️ Outils recommandés

### Éditeur de code
**Visual Studio Code** (Gratuit)
- Télécharger : https://code.visualstudio.com/
- Installer l'extension "Go"
- Installer l'extension "YAML"

```powershell
# Ouvrir le projet avec VSCode
code .
```

### Test d'API
**Postman** (Gratuit)
- Télécharger : https://www.postman.com/downloads/

### Terminal
**Windows Terminal** (Gratuit - Microsoft Store)
- Plus moderne et pratique que PowerShell classique

## ⚠️ Problèmes courants

### "go n'est pas reconnu"
**Solution :** Redémarrez PowerShell après l'installation de Go

### "Erreur de connexion à la base de données"
**Solutions :**
1. Vérifier que PostgreSQL est démarré (Services Windows)
2. Vérifier le mot de passe dans `.env`
3. Vérifier que la base existe dans pgAdmin

### "Port 8080 déjà utilisé"
**Solution :**
```powershell
# Trouver et arrêter le processus
netstat -ano | findstr :8080
taskkill /PID <numéro> /F

# Ou changer le port dans .env
# SERVER_PORT=8081
```

### "Module not found"
**Solution :**
```powershell
go mod download
go mod tidy
```

## 📖 Documentation complète

- **INSTALLATION_WINDOWS.md** - Guide complet Windows
- **README.md** - Documentation complète du projet
- **QUICKSTART.md** - Guide rapide multiplateforme
- **COMPARAISON_LARAVEL.md** - Pour les dev Laravel

## 🎯 Prochaines étapes

1. ✅ Votre première API fonctionne
2. 📚 Lisez INSTALLATION_WINDOWS.md pour plus de détails
3. 🔗 Ajoutez des relations entre vos models
4. 🎨 Personnalisez le code généré
5. 🚀 Déployez votre API

## 💡 Astuces

### Utiliser go-scaffold globalement

Si vous l'avez installé globalement :
```powershell
# Plus besoin de .\go-scaffold.exe
go-scaffold init mon-projet
go-scaffold make:schema user
```

### Générer plusieurs models rapidement

```powershell
# Créer plusieurs schémas
go-scaffold make:schema user
go-scaffold make:schema post
go-scaffold make:schema comment

# Éditer tous les YAML
# Puis générer tout en une fois
go-scaffold generate --all
```

### Créer un script de démarrage

Créer `start.bat` :
```batch
@echo off
echo Demarrage de l'API...
go run main.go
pause
```

Double-cliquez sur `start.bat` pour lancer l'API !

## 🆘 Besoin d'aide ?

1. Consultez **INSTALLATION_WINDOWS.md**
2. Lisez la documentation complète
3. Vérifiez les exemples fournis
4. Ouvrez une issue sur GitHub

---

**Félicitations ! Vous êtes prêt à développer des APIs avec go-scaffold !** 🎉

**Temps total : ~10 minutes** ⏱️
