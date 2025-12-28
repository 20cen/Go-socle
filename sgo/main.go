package main

import (
	"log"
	"sgo/config"
	"sgo/routes"
	
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
