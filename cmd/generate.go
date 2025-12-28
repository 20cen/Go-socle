package cmd

import (
	"fmt"
	"os"

	"go-scaffold/internal/generator"
	"go-scaffold/internal/parser"

	"github.com/spf13/cobra"
)

var (
	generateAll bool
)

var generateCmd = &cobra.Command{
	Use:   "generate [chemin-schema]",
	Short: "Générer le code à partir d'un fichier de schéma",
	Long: `Génère automatiquement les models, contrôleurs, routes, requests et repositories
à partir d'un fichier de schéma YAML.`,
	Run: func(cmd *cobra.Command, args []string) {
		var schemaFiles []string

		if generateAll {
			// Générer pour tous les schémas dans database/schemas
			files, err := os.ReadDir("database/schemas")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erreur de lecture des schémas: %v\n", err)
				os.Exit(1)
			}
			for _, file := range files {
				if !file.IsDir() && (file.Name()[len(file.Name())-5:] == ".yaml" || file.Name()[len(file.Name())-4:] == ".yml") {
					schemaFiles = append(schemaFiles, "database/schemas/"+file.Name())
				}
			}
		} else if len(args) > 0 {
			schemaFiles = []string{args[0]}
		} else {
			fmt.Println("Usage: go-scaffold generate [chemin-schema] ou go-scaffold generate --all")
			os.Exit(1)
		}

		if len(schemaFiles) == 0 {
			fmt.Println("Aucun fichier de schéma trouvé.")
			os.Exit(1)
		}

		for _, schemaFile := range schemaFiles {
			if err := generateFromSchema(schemaFile); err != nil {
				fmt.Fprintf(os.Stderr, "Erreur lors de la génération de %s: %v\n", schemaFile, err)
				continue
			}
			fmt.Printf("✓ Code généré avec succès pour %s\n", schemaFile)
		}
	},
}

func init() {
	generateCmd.Flags().BoolVarP(&generateAll, "all", "a", false, "Générer pour tous les schémas")
}

func generateFromSchema(schemaFile string) error {
	// Parser le schéma
	schema, err := parser.ParseSchema(schemaFile)
	if err != nil {
		return fmt.Errorf("erreur de parsing du schéma: %w", err)
	}

	// Créer le générateur
	gen := generator.NewGenerator(schema)

	// Générer le model
	if err := gen.GenerateModel(); err != nil {
		return fmt.Errorf("erreur de génération du model: %w", err)
	}

	// Générer le repository
	if err := gen.GenerateRepository(); err != nil {
		return fmt.Errorf("erreur de génération du repository: %w", err)
	}

	// Générer le contrôleur
	if err := gen.GenerateController(); err != nil {
		return fmt.Errorf("erreur de génération du contrôleur: %w", err)
	}

	// Générer les requests
	if err := gen.GenerateRequests(); err != nil {
		return fmt.Errorf("erreur de génération des requests: %w", err)
	}

	// Générer les routes
	if err := gen.GenerateRoutes(); err != nil {
		return fmt.Errorf("erreur de génération des routes: %w", err)
	}

	return nil
}
