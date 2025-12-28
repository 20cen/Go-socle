package routes

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
		RegisterArticleRoutes(api)
}
