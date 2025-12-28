package repositories

import (
	"errors"
	"sgo/app/models"
	"sgo/config"

	"gorm.io/gorm"
)

// ArticleInterface définit les méthodes du repository
type ArticleInterface interface {
	Create(article *models.Article) error
	FindByID(id string) (*models.Article, error)
	FindAll(page, pageSize int) ([]models.Article, int64, error)
	Update(article *models.Article) error
	Delete(id string) error
}

// ArticleRepository implémente ArticleInterface
type ArticleRepository struct {
	db *gorm.DB
}

// NewArticleRepository crée une nouvelle instance du repository
func NewArticleRepository() ArticleInterface {
	return &ArticleRepository{
		db: config.GetDB(),
	}
}

// Create crée un nouveau article
func (r *ArticleRepository) Create(article *models.Article) error {
	return r.db.Create(article).Error
}

// FindByID trouve un article par son ID
func (r *ArticleRepository) FindByID(id string) (*models.Article, error) {
	var article models.Article
	err := r.db.First(&article, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("enregistrement non trouvé")
		}
		return nil, err
	}
	return &article, nil
}

// FindAll récupère tous les articles avec pagination
func (r *ArticleRepository) FindAll(page, pageSize int) ([]models.Article, int64, error) {
	var articles []models.Article
	var total int64

	// Compter le total
	if err := r.db.Model(&models.Article{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculer l'offset
	offset := (page - 1) * pageSize

	// Récupérer les données avec pagination
	err := r.db.
		Offset(offset).
		Limit(pageSize).
		Find(&articles).Error

	if err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

// Update met à jour un article
func (r *ArticleRepository) Update(article *models.Article) error {
	return r.db.Save(article).Error
}

// Delete supprime un article
func (r *ArticleRepository) Delete(id string) error {
	return r.db.Delete(&models.Article{}, "id = ?", id).Error
}

