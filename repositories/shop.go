package repositories

import "github.com/suttapak/cacf/models"

type ShopRespository interface {
	Get() (*models.Shop, error)
	Create(shop models.Shop) error
	Update(shop models.Shop) error
	Delete(shop models.Shop) error
}
