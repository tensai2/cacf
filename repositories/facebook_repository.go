package repositories

import (
	"github.com/suttapak/cacf/models"
	"gorm.io/gorm"
)

type facebookRepository struct {
	db *gorm.DB
}

func NewFacebookRepository(db *gorm.DB) FacebookRespository {
	return &facebookRepository{db: db}
}
func (r facebookRepository) Get() (*models.Facebook, error) {
	facebook := models.Facebook{}
	if err := r.db.First(&facebook).Error; err != nil {
		return nil, err
	}
	return &facebook, nil
}
func (r facebookRepository) Create(facebook models.Facebook) error {
	if err := r.db.Model(&facebook).FirstOrCreate(&facebook).Error; err != nil {
		return err
	}
	return nil
}
func (r facebookRepository) Update(facebook models.Facebook) error {
	if err := r.db.Updates(&facebook).Error; err != nil {
		return err
	}
	return nil
}
func (r facebookRepository) Delete(facebook models.Facebook) error {
	if err := r.db.Where("id = ?", facebook.ID).Delete(&facebook).Error; err != nil {
		return err
	}
	return nil
}
