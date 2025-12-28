package models

import (
	"time"
	"gorm.io/gorm"
)

// Article représente la table articles
type Article struct {
	Id string `json:"id" gorm:"primaryKey;not null;column:id"`
	Titre string `json:"titre" gorm:"not null;size:255;column:titre" validate:"required,max=255"`
	Contenu string `json:"contenu" gorm:"not null;column:contenu" validate:"required"`
	Auteur string `json:"auteur" gorm:"not null;size:100;column:auteur" validate:"required,max=100"`
	Publie bool `json:"publie" gorm:"not null;default:false;column:publie" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;column:updated_at"`
}

// TableName retourne le nom de la table
func (Article) TableName() string {
	return "articles"
}

// BeforeCreate hook GORM
func (m *Article) BeforeCreate(tx *gorm.DB) error {
	// Logique avant création
	return nil
}

// BeforeUpdate hook GORM
func (m *Article) BeforeUpdate(tx *gorm.DB) error {
	// Logique avant mise à jour
	return nil
}
