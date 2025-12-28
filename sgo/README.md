# sgo

Projet Go généré automatiquement avec go-scaffold.

## Installation

```bash
go mod download
```

## Configuration

Copiez le fichier .env.example vers .env et modifiez les valeurs selon votre configuration.

```bash
cp .env.example .env
```

## Utilisation

Pour générer du code à partir d'un schéma:

```bash
go-scaffold generate database/schemas/votre_schema.yaml
```

Pour créer un nouveau schéma:

```bash
go-scaffold make:schema nom_table
```

## Démarrage

```bash
go run main.go
```
