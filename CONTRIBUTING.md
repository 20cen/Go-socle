# Guide de Contribution

Merci de votre intérêt pour contribuer à go-scaffold ! Ce document vous guidera à travers le processus de contribution.

## Table des matières

1. [Code de conduite](#code-de-conduite)
2. [Comment contribuer](#comment-contribuer)
3. [Structure du projet](#structure-du-projet)
4. [Standards de code](#standards-de-code)
5. [Processus de Pull Request](#processus-de-pull-request)
6. [Rapport de bugs](#rapport-de-bugs)
7. [Suggestions de fonctionnalités](#suggestions-de-fonctionnalités)

## Code de conduite

En participant à ce projet, vous vous engagez à maintenir un environnement respectueux et inclusif pour tous.

## Comment contribuer

### Prérequis

- Go 1.21 ou supérieur
- Git
- Connaissance de base de Go, GORM et Gin
- (Optionnel) golangci-lint pour le linting

### Configuration de l'environnement de développement

1. Fork le repository
2. Clone votre fork :
   ```bash
   git clone https://github.com/votre-username/go-scaffold.git
   cd go-scaffold
   ```

3. Installer les dépendances :
   ```bash
   go mod download
   ```

4. Créer une branche pour vos modifications :
   ```bash
   git checkout -b feature/ma-fonctionnalite
   ```

## Structure du projet

```
go-scaffold/
├── cmd/                    # Commandes CLI
│   ├── root.go            # Commande racine
│   ├── init.go            # Commande init
│   ├── make.go            # Commandes make:*
│   └── generate.go        # Commande generate
├── internal/
│   ├── generator/         # Générateurs de code
│   │   ├── generator.go   # Générateur principal
│   │   ├── repository.go  # Générateur de repositories
│   │   ├── controller.go  # Générateur de contrôleurs
│   │   ├── request.go     # Générateur de requests
│   │   └── routes.go      # Générateur de routes
│   └── parser/            # Parser de schémas YAML
│       └── parser.go
├── examples/              # Exemples de schémas
├── main.go               # Point d'entrée
└── README.md
```

## Standards de code

### Style de code

- Suivre les conventions Go standards (gofmt, golint)
- Utiliser des noms descriptifs pour les variables et fonctions
- Commenter les fonctions publiques
- Garder les fonctions courtes et focalisées

### Formatage

Avant de commit, assurez-vous que votre code est formaté :

```bash
make fmt
# ou
go fmt ./...
```

### Linting

```bash
make lint
# ou
golangci-lint run
```

### Tests

Ajoutez des tests pour toute nouvelle fonctionnalité :

```bash
go test ./... -v
```

### Commits

Utilisez des messages de commit clairs et descriptifs :

```
type(scope): description courte

Description détaillée si nécessaire

Fixes #123
```

Types de commit :
- `feat`: Nouvelle fonctionnalité
- `fix`: Correction de bug
- `docs`: Documentation
- `style`: Formatage, point-virgules manquants, etc.
- `refactor`: Refactoring du code
- `test`: Ajout de tests
- `chore`: Maintenance

Exemples :
```
feat(generator): ajouter le support de MySQL
fix(parser): corriger le parsing des relations many-to-many
docs(readme): ajouter des exemples de validation
```

## Processus de Pull Request

1. **Mettez à jour votre fork**
   ```bash
   git remote add upstream https://github.com/original/go-scaffold.git
   git fetch upstream
   git rebase upstream/main
   ```

2. **Testez vos modifications**
   ```bash
   make test
   make fmt
   make lint
   ```

3. **Poussez vos modifications**
   ```bash
   git push origin feature/ma-fonctionnalite
   ```

4. **Créez une Pull Request**
   - Allez sur GitHub et créez une PR depuis votre branche
   - Décrivez clairement vos modifications
   - Liez les issues concernées
   - Attendez la review

5. **Checklist de la PR**
   - [ ] Le code compile sans erreurs
   - [ ] Les tests passent
   - [ ] Le code est formaté (gofmt)
   - [ ] Le code est linté (golangci-lint)
   - [ ] La documentation est à jour
   - [ ] Les exemples fonctionnent
   - [ ] CHANGELOG.md est mis à jour

## Rapport de bugs

### Avant de rapporter un bug

- Vérifiez que le bug n'a pas déjà été rapporté
- Vérifiez que vous utilisez la dernière version
- Collectez des informations sur votre environnement

### Comment rapporter un bug

Créez une issue avec :

**Titre** : Description courte et claire

**Description** :
- Description détaillée du problème
- Étapes pour reproduire
- Comportement attendu
- Comportement actuel
- Captures d'écran (si applicable)

**Environnement** :
- Version de go-scaffold
- Version de Go
- Système d'exploitation
- Base de données utilisée

**Exemple** :

```markdown
### Description
Le générateur échoue lors de la création d'une relation many-to-many

### Étapes pour reproduire
1. Créer un schéma avec une relation many-to-many
2. Exécuter `go-scaffold generate schema.yaml`
3. Observer l'erreur

### Comportement attendu
Le code devrait être généré sans erreur

### Comportement actuel
Erreur : "pivot_table is required for many_to_many"

### Environnement
- go-scaffold: v1.0.0
- Go: 1.21.0
- OS: Ubuntu 22.04
- DB: PostgreSQL 14
```

## Suggestions de fonctionnalités

### Avant de suggérer

- Vérifiez que la fonctionnalité n'existe pas déjà
- Vérifiez qu'elle n'est pas déjà proposée
- Réfléchissez à son utilité générale

### Comment suggérer

Créez une issue avec :

**Titre** : Description claire de la fonctionnalité

**Description** :
- Problème que cette fonctionnalité résout
- Solution proposée
- Alternatives considérées
- Exemples d'utilisation

**Exemple** :

```markdown
### Problème
Actuellement, le générateur ne supporte que PostgreSQL

### Solution proposée
Ajouter le support de MySQL et SQLite

### Utilisation proposée
bash
go-scaffold init mon-projet --database=mysql

### Bénéfices
- Plus de flexibilité
- Support de plus de projets
- Adoption plus large
```

## Types de contributions recherchées

### Priorité haute
- Corrections de bugs
- Amélioration de la documentation
- Ajout de tests
- Support de nouvelles bases de données

### Priorité moyenne
- Nouvelles fonctionnalités
- Optimisations de performance
- Amélioration de l'UX du CLI

### Priorité basse
- Refactoring
- Nouvelles options de configuration

## Questions ?

N'hésitez pas à :
- Ouvrir une issue pour poser une question
- Rejoindre nos discussions
- Consulter la documentation

## Licence

En contribuant à go-scaffold, vous acceptez que vos contributions soient sous licence MIT.

---

Merci pour vos contributions ! 🎉
