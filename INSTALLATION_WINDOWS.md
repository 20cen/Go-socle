# 🪟 Installation de Go-Scaffold sur Windows

Guide complet pour installer Go et go-scaffold sur Windows.

## 📥 Étape 1 : Installer Go

### Téléchargement

1. Allez sur : https://go.dev/dl/
2. Téléchargez la version Windows : `go1.21.x.windows-amd64.msi` (ou plus récente)
3. Exécutez l'installeur MSI
4. Suivez l'assistant d'installation (garder les options par défaut)

### Vérification de l'installation

Ouvrez PowerShell ou CMD et tapez :

```powershell
go version
```

Vous devriez voir quelque chose comme :
```
go version go1.21.x windows/amd64
```

Si la commande n'est pas reconnue, redémarrez votre terminal ou votre PC.

## 📦 Étape 2 : Installer go-scaffold

### Option A : Compilation manuelle (Recommandé)

1. **Extraire l'archive**
   - Faites clic droit sur `go-scaffold.tar.gz`
   - Extraire avec 7-Zip, WinRAR ou Windows (extraction native)
   - Ou utilisez PowerShell :

```powershell
# Si tar est disponible (Windows 10+)
tar -xzf go-scaffold.tar.gz
cd go-scaffold
```

2. **Télécharger les dépendances**

```powershell
go mod download
```

3. **Compiler le projet**

```powershell
go build -o go-scaffold.exe main.go
```

4. **Tester l'exécutable**

```powershell
.\go-scaffold.exe --help
```

Vous devriez voir le menu d'aide !

### Option B : Installation globale (Optionnel)

Pour utiliser `go-scaffold` depuis n'importe où :

1. **Créer un dossier pour les binaires Go** (si pas déjà fait)

```powershell
# Créer le dossier
mkdir C:\Go\bin

# Copier l'exécutable
copy go-scaffold.exe C:\Go\bin\
```

2. **Ajouter au PATH**

   **Méthode 1 - Interface graphique :**
   - Ouvrir "Paramètres système avancés"
   - Cliquer sur "Variables d'environnement"
   - Dans "Variables système", trouver "Path"
   - Cliquer "Modifier"
   - Ajouter `C:\Go\bin`
   - Cliquer OK partout

   **Méthode 2 - PowerShell (Admin) :**
   ```powershell
   [Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\Go\bin", "Machine")
   ```

3. **Redémarrer le terminal** et tester :

```powershell
go-scaffold --help
```

## 🚀 Étape 3 : Créer votre premier projet

### 1. Initialiser un projet

```powershell
# Avec installation globale
go-scaffold init mon-api

# Sans installation globale (depuis le dossier go-scaffold)
.\go-scaffold.exe init mon-api
```

### 2. Accéder au projet

```powershell
cd mon-api
```

### 3. Créer un schéma

```powershell
# Avec installation globale
go-scaffold make:schema produit

# Sans installation globale
..\go-scaffold.exe make:schema produit
```

### 4. Éditer le schéma

Ouvrez `database\schemas\produit.yaml` avec votre éditeur préféré (VSCode, Notepad++, etc.)

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

  - name: prix
    type: float
    nullable: false

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
```

### 5. Générer le code

```powershell
# Avec installation globale
go-scaffold generate database\schemas\produit.yaml

# Sans installation globale
..\go-scaffold.exe generate database\schemas\produit.yaml
```

### 6. Installer PostgreSQL

**Télécharger PostgreSQL :**
- Site : https://www.postgresql.org/download/windows/
- Télécharger l'installeur
- Installer avec les options par défaut
- Noter le mot de passe que vous choisissez pour l'utilisateur `postgres`

**Créer la base de données :**

Ouvrir pgAdmin ou psql et créer votre base :

```sql
CREATE DATABASE mon_api;
```

### 7. Configurer l'environnement

Copier `.env.example` vers `.env` :

```powershell
copy .env.example .env
```

Éditer `.env` avec Notepad :

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=votre_mot_de_passe
DB_NAME=mon_api
SERVER_PORT=8080
```

### 8. Installer les dépendances du projet

```powershell
go mod download
```

### 9. Lancer l'application

```powershell
go run main.go
```

Votre API sera disponible sur : **http://localhost:8080** 🎉

### 10. Tester l'API

**Dans un nouveau terminal ou avec un outil comme Postman :**

```powershell
# Tester l'endpoint de santé
curl http://localhost:8080/api/health

# Créer un produit (avec PowerShell)
Invoke-WebRequest -Uri "http://localhost:8080/api/produits" -Method POST -ContentType "application/json" -Body '{"nom":"Ordinateur","prix":999.99}'

# Lister les produits
curl http://localhost:8080/api/produits
```

## 🛠️ Outils recommandés pour Windows

### Éditeurs de code
- **Visual Studio Code** (https://code.visualstudio.com/)
  - Extension Go officielle
  - Extension YAML
  - Extension REST Client (pour tester l'API)

### Gestionnaire de base de données
- **pgAdmin** (inclus avec PostgreSQL)
- **DBeaver** (https://dbeaver.io/)
- **TablePlus** (https://tableplus.com/)

### Terminal
- **Windows Terminal** (Microsoft Store) - Recommandé
- **PowerShell 7** (https://github.com/PowerShell/PowerShell)
- **Git Bash** (inclus avec Git for Windows)

### Test d'API
- **Postman** (https://www.postman.com/)
- **Insomnia** (https://insomnia.rest/)
- Extension VSCode : REST Client

## 📝 Commandes PowerShell utiles

### Compilation

```powershell
# Compiler
go build -o go-scaffold.exe main.go

# Compiler pour Linux (si vous déployez sur Linux)
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o go-scaffold main.go
```

### Gestion des dépendances

```powershell
# Télécharger les dépendances
go mod download

# Nettoyer les dépendances inutilisées
go mod tidy

# Vérifier les dépendances
go mod verify
```

### Compilation du projet généré

```powershell
# Dans le dossier de votre projet (mon-api)
go build -o mon-api.exe main.go

# Exécuter le binaire
.\mon-api.exe
```

## 🔧 Résolution de problèmes

### Problème : "go n'est pas reconnu"

**Solution :**
1. Vérifiez que Go est installé : Panneau de configuration → Programmes
2. Redémarrez votre terminal
3. Vérifiez le PATH : `echo $env:Path`
4. Réinstallez Go si nécessaire

### Problème : "Erreur de connexion à la base de données"

**Solution :**
1. Vérifiez que PostgreSQL est démarré (Services Windows)
2. Vérifiez vos identifiants dans `.env`
3. Testez la connexion avec pgAdmin
4. Vérifiez le firewall

### Problème : "Module not found"

**Solution :**
```powershell
go mod download
go mod tidy
```

### Problème : Port 8080 déjà utilisé

**Solution :**
```powershell
# Trouver le processus utilisant le port
netstat -ano | findstr :8080

# Tuer le processus (remplacer PID)
taskkill /PID <numéro_pid> /F

# Ou changer le port dans .env
# SERVER_PORT=8081
```

## 📚 Ressources supplémentaires

### Apprentissage de Go
- Tour de Go (Français) : https://go-tour-fr.appspot.com/
- Go by Example : https://gobyexample.com/
- Documentation officielle : https://go.dev/doc/

### Communauté
- Forum Go : https://forum.golangbridge.org/
- Reddit : r/golang
- Discord Go : https://discord.gg/golang

## 🎯 Prochaines étapes

1. ✅ Installer Go
2. ✅ Compiler go-scaffold
3. ✅ Créer votre premier projet
4. ✅ Installer PostgreSQL
5. ✅ Générer votre première API
6. 📖 Lire le guide complet dans README.md
7. 💡 Explorer les exemples
8. 🚀 Créer vos propres schémas

## 💡 Conseils

- **Utilisez Windows Terminal** pour une meilleure expérience
- **Installez VSCode** avec l'extension Go
- **Utilisez Git for Windows** pour Git Bash
- **Sauvegardez vos schémas** dans Git
- **Testez avec Postman** pour faciliter les tests d'API

## 🆘 Besoin d'aide ?

Si vous rencontrez des problèmes :
1. Consultez la section "Résolution de problèmes" ci-dessus
2. Lisez la documentation complète (README.md)
3. Vérifiez les exemples fournis
4. Ouvrez une issue sur GitHub

---

**Bon développement avec go-scaffold sur Windows !** 🪟🚀
