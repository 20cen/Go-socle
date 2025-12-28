package requests

import (
	"sgo/app/models"
)

// CreateArticleRequest représente les données pour créer un article
type CreateArticleRequest struct {
	Titre string `json:"titre" validate:"required,min=5,max=255"`
	Contenu string `json:"contenu" validate:"required,min=50"`
	Auteur string `json:"auteur" validate:"min=3,max=100,required"`
	Publie bool `json:"publie" validate:"required"`
}

// ToModel convertit la requête en model
func (r *CreateArticleRequest) ToModel() models.Article {
	return models.Article{
		Titre: r.Titre,
		Contenu: r.Contenu,
		Auteur: r.Auteur,
		Publie: r.Publie,
	}
}

// UpdateArticleRequest représente les données pour mettre à jour un article
type UpdateArticleRequest struct {
	Titre *string `json:"titre,omitempty" validate:"min=5,max=255"`
	Contenu *string `json:"contenu,omitempty" validate:"min=50"`
	Auteur *string `json:"auteur,omitempty" validate:"min=3,max=100"`
	Publie *bool `json:"publie,omitempty" `
}

// UpdateModel met à jour le model avec les données de la requête
func (r *UpdateArticleRequest) UpdateModel(m *models.Article) {
	if r.Titre != nil {
		m.Titre = *r.Titre
	}
	if r.Contenu != nil {
		m.Contenu = *r.Contenu
	}
	if r.Auteur != nil {
		m.Auteur = *r.Auteur
	}
	if r.Publie != nil {
		m.Publie = *r.Publie
	}
}
