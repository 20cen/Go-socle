package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [nom-du-projet]",
	Short: "Initialiser un nouveau projet Go avec la structure de base",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		if err := initializeProject(projectName); err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors de l'initialisation: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✓ Projet '%s' initialisé avec succès!\n", projectName)
	},
}

func initializeProject(projectName string) error {
	// Créer la structure de dossiers
	dirs := []string{
		projectName,
		filepath.Join(projectName, "app", "models"),
		filepath.Join(projectName, "app", "controllers"),
		filepath.Join(projectName, "app", "requests"),
		filepath.Join(projectName, "app", "services"),
		filepath.Join(projectName, "app", "repositories"),
		filepath.Join(projectName, "routes"),
		filepath.Join(projectName, "database", "schemas"),
		filepath.Join(projectName, "database", "migrations"),
		filepath.Join(projectName, "config"),
		filepath.Join(projectName, "middleware"),
		filepath.Join(projectName, "utils"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("impossible de créer le dossier %s: %w", dir, err)
		}
	}

	// Créer les fichiers de base
	if err := createBaseFiles(projectName); err != nil {
		return err
	}

	return nil
}

func createBaseFiles(projectName string) error {
	// go.mod
	goModContent := fmt.Sprintf(`module %s

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/go-playground/validator/v10 v10.16.0
	gorm.io/gorm v1.25.5
	gorm.io/driver/postgres v1.5.4
)
`, projectName)

	if err := os.WriteFile(filepath.Join(projectName, "go.mod"), []byte(goModContent), 0644); err != nil {
		return err
	}

	// main.go
	mainContent := `package main

import (
	"log"
	"` + projectName + `/config"
	"` + projectName + `/routes"
	
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialiser la configuration
	if err := config.Init(); err != nil {
		log.Fatalf("Erreur de configuration: %v", err)
	}

	// Initialiser la base de données
	if err := config.InitDB(); err != nil {
		log.Fatalf("Erreur de connexion à la base de données: %v", err)
	}

	// Créer le routeur Gin
	router := gin.Default()

	// Enregistrer les routes
	routes.RegisterRoutes(router)

	// Démarrer le serveur
	port := config.GetConfig().ServerPort
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Serveur démarré sur le port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erreur de démarrage du serveur: %v", err)
	}
}
`

	if err := os.WriteFile(filepath.Join(projectName, "main.go"), []byte(mainContent), 0644); err != nil {
		return err
	}

	// config/config.go
	configContent := `package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

var (
	config *Config
	DB     *gorm.DB
)

func Init() error {
	config = &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "mydb"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
	return nil
}

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("échec de connexion à la base de données: %w", err)
	}

	return nil
}

func GetConfig() *Config {
	return config
}

func GetDB() *gorm.DB {
	return DB
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
`

	if err := os.WriteFile(filepath.Join(projectName, "config", "config.go"), []byte(configContent), 0644); err != nil {
		return err
	}

	// routes/routes.go
	routesContent := `package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Les routes générées seront ajoutées ici automatiquement
	api := router.Group("/api")
	{
		// Exemple de route de base
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
				"message": "Service en cours d'exécution",
			})
		})
	}
}
`

	if err := os.WriteFile(filepath.Join(projectName, "routes", "routes.go"), []byte(routesContent), 0644); err != nil {
		return err
	}

	// .env.example
	envContent := `DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=
DB_NAME=mydb
SERVER_PORT=8080
`

	if err := os.WriteFile(filepath.Join(projectName, ".env.example"), []byte(envContent), 0644); err != nil {
		return err
	}

	// README.md
	readmeContent := `# ` + projectName + `

Projet Go généré automatiquement avec go-scaffold.

## Installation

` + "```bash" + `
go mod download
` + "```" + `

## Configuration

Copiez le fichier .env.example vers .env et modifiez les valeurs selon votre configuration.

` + "```bash" + `
cp .env.example .env
` + "```" + `

## Utilisation

Pour générer du code à partir d'un schéma:

` + "```bash" + `
go-scaffold generate database/schemas/votre_schema.yaml
` + "```" + `

Pour créer un nouveau schéma:

` + "```bash" + `
go-scaffold make:schema nom_table
` + "```" + `

## Démarrage

` + "```bash" + `
go run main.go
` + "```" + `
`

	return os.WriteFile(filepath.Join(projectName, "README.md"), []byte(readmeContent), 0644)
}
