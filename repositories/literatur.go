package repositories

import (
	"literature/models"

	"gorm.io/gorm"
)

type LiteraturRepository interface {
	FindLiteraturs() ([]models.Literatur, error)
	GetLiteratur(ID int) (models.Literatur, error)
	CreateLiteratur(Literatur models.Literatur) (models.Literatur, error)
}

type repositoryLiteratur struct {
	db *gorm.DB
}

func RepositoryLiteratur(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateLiteratur(literatur models.Literatur) (models.Literatur, error) {
	err := r.db.Preload("User").Create(&literatur).Error

	return literatur, err
}

func (r *repository) FindLiteraturs() ([]models.Literatur, error) {
	var literaturs []models.Literatur
	err := r.db.Find(&literaturs).Error

	return literaturs, err
}

func (r *repository) GetLiteratur(ID int) (models.Literatur, error) {
	var literatur models.Literatur
	err := r.db.Preload("User").First(&literatur, ID).Error

	return literatur, err
}

func (r *repository) UpdateLiteratur(literatur models.Literatur, ID int) (models.Literatur, error) {
	err := r.db.Save(&literatur).Error

	return literatur, err
}

func (r *repository) DeleteLiteratur(literatur models.Literatur, ID int) (models.Literatur, error) {
	err := r.db.Delete(&literatur).Error

	return literatur, err
}
