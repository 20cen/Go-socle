package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Créer des fichiers de base (schéma, migration, etc.)",
}

var makeSchemaCmd = &cobra.Command{
	Use:   "schema [nom]",
	Short: "Créer un nouveau fichier de schéma",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := createSchema(name); err != nil {
			fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✓ Schéma '%s' créé avec succès!\n", name)
	},
}

var makeMigrationCmd = &cobra.Command{
	Use:   "migration [nom]",
	Short: "Créer un nouveau fichier de migration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := createMigration(name); err != nil {
			fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✓ Migration '%s' créée avec succès!\n", name)
	},
}

func init() {
	makeCmd.AddCommand(makeSchemaCmd)
	makeCmd.AddCommand(makeMigrationCmd)
}

func createSchema(name string) error {
	// Normaliser le nom (singulier, snake_case)
	schemaName := toSnakeCase(name)
	filename := filepath.Join("database", "schemas", schemaName+".yaml")

	// Vérifier si le fichier existe déjà
	if _, err := os.Stat(filename); err == nil {
		return fmt.Errorf("le schéma %s existe déjà", filename)
	}

	// Créer le répertoire s'il n'existe pas
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return err
	}

	// Template de schéma
	template := `# Schéma pour la table ` + schemaName + `
table: ` + schemaName + `
model: ` + toPascalCase(name) + `

# Définir les colonnes de la table
columns:
  - name: id
    type: uuid
    primary: true
    auto_increment: false
    nullable: false

  - name: name
    type: string
    size: 255
    nullable: false
    unique: false

  - name: description
    type: text
    nullable: true

  - name: status
    type: string
    size: 50
    nullable: false
    default: "active"

  - name: created_at
    type: timestamp
    nullable: false

  - name: updated_at
    type: timestamp
    nullable: false

# Relations (optionnel)
relations:
  # Exemple de relation belongsTo (many-to-one)
  # - type: belongs_to
  #   model: User
  #   foreign_key: user_id
  #   references: id

  # Exemple de relation hasMany (one-to-many)
  # - type: has_many
  #   model: Comment
  #   foreign_key: ` + schemaName + `_id

  # Exemple de relation many-to-many
  # - type: many_to_many
  #   model: Tag
  #   pivot_table: ` + schemaName + `_tags
  #   foreign_key: ` + schemaName + `_id
  #   related_key: tag_id

# Indexes (optionnel)
indexes:
  - name: idx_` + schemaName + `_status
    columns: [status]
    unique: false

# Validations pour les requests
validations:
  - field: name
    rules:
      - required: true
      - min: 3
      - max: 255

  - field: description
    rules:
      - max: 1000

  - field: status
    rules:
      - required: true
      - in: [active, inactive, pending]
`

	return os.WriteFile(filename, []byte(template), 0644)
}

func createMigration(name string) error {
	// Créer le répertoire s'il n'existe pas
	if err := os.MkdirAll(filepath.Join("database", "migrations"), 0755); err != nil {
		return err
	}

	// Générer un nom de fichier avec timestamp
	timestamp := fmt.Sprintf("%d", os.Getpid()) // Simplification pour l'exemple
	filename := filepath.Join("database", "migrations", fmt.Sprintf("%s_%s.go", timestamp, toSnakeCase(name)))

	template := `package migrations

import (
	"gorm.io/gorm"
)

func init() {
	RegisterMigration(&Migration{
		Name: "` + name + `",
		Up: func(db *gorm.DB) error {
			// Votre code de migration ici
			return db.Exec("` + "``" + `
				-- Exemple de migration SQL
				-- CREATE TABLE ...
			` + "``" + `").Error
		},
		Down: func(db *gorm.DB) error {
			// Votre code de rollback ici
			return db.Exec("` + "``" + `
				-- Exemple de rollback SQL
				-- DROP TABLE ...
			` + "``" + `").Error
		},
	})
}
`

	return os.WriteFile(filename, []byte(template), 0644)
}

func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

func toPascalCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == ' '
	})
	for i, word := range words {
		words[i] = strings.Title(strings.ToLower(word))
	}
	return strings.Join(words, "")
}
