package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-scaffold",
	Short: "Générateur de code automatique pour Go",
	Long: `Un outil CLI pour générer automatiquement des models, contrôleurs, routes et validations
à partir de fichiers de schéma de base de données.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(makeCmd)
}
