package repositories

type Message interface {
	Get(permission string) (*Message, error)
	Create(message Message) error
	Update(message Message) error
	Delete(message Message) error
}
