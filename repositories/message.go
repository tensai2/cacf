package repositories

import "github.com/suttapak/cacf/models"

type MessageRepository interface {
	Get(permission string) (*models.Message, error)
	Create(message models.Message) error
	Update(message models.Message) error
	Delete(message models.Message) error
}
