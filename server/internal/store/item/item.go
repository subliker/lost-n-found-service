package item

import "github.com/subliker/server/internal/models"

type Store interface {
	Find(items *[]models.Item) error
	Create(item *models.Item) error
	FindWithPagination(items *[]models.Item, page int, pageSize int) error
	Close()
}
