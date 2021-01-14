package product

type ProductCategory string

const (
	ProductCategoryFood       ProductCategory = "food"
	ProductCategoryClothing                   = "clothing"
	ProductCategoryElectronic                 = "electronic"
	ProductCategoryOther                      = "other"
)

type Product struct {
	ID          int
	Name        string
	Category    ProductCategory
	RetailPrice float32
}

type Service interface {
	Find(sku int) (Product, error)
	ListProducts() []Product
}
