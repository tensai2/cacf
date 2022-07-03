package repositories

import (
	"github.com/suttapak/cacf/models"
	"gorm.io/gorm"
)

type shopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) ShopRespository {
	return &shopRepository{db: db}
}

func (r shopRepository) Get() (shop *models.Shop, err error) {
	if err := r.db.Model(&shop).First(&shop).Error; err != nil {
		return nil, err
	}
	return shop, nil
}

func (r shopRepository) Create(shop models.Shop) error {
	if err := r.db.Model(&shop).FirstOrCreate(&shop).Error; err != nil {
		return err
	}
	return nil
}
func (r shopRepository) Update(shop models.Shop) error {
	if err := r.db.Updates(&shop).Error; err != nil {
		return err
	}
	return nil
}
func (r shopRepository) Delete(shop models.Shop) error {
	if err := r.db.Where("id = ?", shop.ID).Delete(&shop).Error; err != nil {
		return err
	}
	return nil
}
