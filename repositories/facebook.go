package repositories

import "github.com/suttapak/cacf/models"

type FacebookRespository interface {
	Get() (*models.Facebook, error)
	Create(facebook models.Facebook) error
	Update(facebook models.Facebook) error
	Delete(facebook models.Facebook) error
}
