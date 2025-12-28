package controllers

import (
	"net/http"
	"strconv"
	
	"sgo/app/repositories"
	"sgo/app/requests"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ArticleController gère les requêtes HTTP pour Article
type ArticleController struct {
	repo repositories.ArticleInterface
	validate *validator.Validate
}

// NewArticleController crée une nouvelle instance du contrôleur
func NewArticleController() *ArticleController {
	return &ArticleController{
		repo: repositories.NewArticleRepository(),
		validate: validator.New(),
	}
}

// Index récupère la liste des articles
// @Summary Liste des articles
// @Description Récupère tous les articles avec pagination
// @Tags Article
// @Accept json
// @Produce json
// @Param page query int false "Numéro de page" default(1)
// @Param page_size query int false "Taille de page" default(10)
// @Success 200 {object} map[string]interface{}
// @Router /articles [get]
func (ctrl *ArticleController) Index(c *gin.Context) {
	// Paramètres de pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	articles, total, err := ctrl.repo.FindAll(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erreur lors de la récupération des données",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": articles,
		"pagination": gin.H{
			"page": page,
			"page_size": pageSize,
			"total": total,
			"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}

// Show récupère un article par ID
// @Summary Afficher un article
// @Description Récupère un article spécifique par son ID
// @Tags Article
// @Accept json
// @Produce json
// @Param id path string true "ID du article"
// @Success 200 {object} models.Article
// @Router /articles/{id} [get]
func (ctrl *ArticleController) Show(c *gin.Context) {
	id := c.Param("id")

	article, err := ctrl.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Enregistrement non trouvé",
		})
		return
	}

	c.JSON(http.StatusOK, article)
}

// Store crée un nouveau article
// @Summary Créer un article
// @Description Crée un nouveau article
// @Tags Article
// @Accept json
// @Produce json
// @Param article body requests.CreateArticleRequest true "Données du article"
// @Success 201 {object} models.Article
// @Router /articles [post]
func (ctrl *ArticleController) Store(c *gin.Context) {
	var req requests.CreateArticleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Données invalides",
			"details": err.Error(),
		})
		return
	}

	if err := ctrl.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation échouée",
			"details": err.Error(),
		})
		return
	}

	article := req.ToModel()

	if err := ctrl.repo.Create(&article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erreur lors de la création",
		})
		return
	}

	c.JSON(http.StatusCreated, article)
}

// Update met à jour un article
// @Summary Mettre à jour un article
// @Description Met à jour un article existant
// @Tags Article
// @Accept json
// @Produce json
// @Param id path string true "ID du article"
// @Param article body requests.UpdateArticleRequest true "Nouvelles données"
// @Success 200 {object} models.Article
// @Router /articles/{id} [put]
func (ctrl *ArticleController) Update(c *gin.Context) {
	id := c.Param("id")

	article, err := ctrl.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Enregistrement non trouvé",
		})
		return
	}

	var req requests.UpdateArticleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Données invalides",
			"details": err.Error(),
		})
		return
	}

	if err := ctrl.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation échouée",
			"details": err.Error(),
		})
		return
	}

	req.UpdateModel(article)

	if err := ctrl.repo.Update(article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erreur lors de la mise à jour",
		})
		return
	}

	c.JSON(http.StatusOK, article)
}

// Delete supprime un article
// @Summary Supprimer un article
// @Description Supprime un article par son ID
// @Tags Article
// @Accept json
// @Produce json
// @Param id path string true "ID du article"
// @Success 204
// @Router /articles/{id} [delete]
func (ctrl *ArticleController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erreur lors de la suppression",
		})
		return
	}

	c.Status(http.StatusNoContent)
}
