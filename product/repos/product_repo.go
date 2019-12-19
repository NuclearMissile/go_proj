package repos

import "github.com/go_proj/product/models"

type IProduct interface {
	Conn() error
	Insert(*models.Product) (int64, error)
	Delete(int64) bool
	Update(*models.Product) error
	SelectByKey(int64) (*models.Product, error)
	SelectAll() ([]*models.Product, error)
}

type ProductManager struct {

}
