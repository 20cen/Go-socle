package routes

import (
	"sgo/app/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterArticleRoutes enregistre les routes pour Article
func RegisterArticleRoutes(router *gin.RouterGroup) {
	ctrl := controllers.NewArticleController()

	// Routes RESTful pour article
	articleGroup := router.Group("/articles")
	{
		articleGroup.GET("", ctrl.Index)        // GET /articles
		articleGroup.POST("", ctrl.Store)       // POST /articles
		articleGroup.GET("/:id", ctrl.Show)     // GET /articles/:id
		articleGroup.PUT("/:id", ctrl.Update)   // PUT /articles/:id
		articleGroup.DELETE("/:id", ctrl.Delete) // DELETE /articles/:id
	}
}
