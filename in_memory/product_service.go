package in_memory

import (
	"fmt"
	"sync"

	"github.com/akurey/go-programming-test/product"
)

type productService struct {
	skus map[int]product.Product
	sync.RWMutex
}

func (s *productService) Find(skuID int) (product.Product, error) {
	s.RLock()
	defer s.RUnlock()
	prod, ok := s.skus[skuID]
	if ok {
		return prod, nil
	}
	return product.Product{}, fmt.Errorf("Product not found")
}

func (s *productService) ListProducts() []product.Product {
	s.RLock()
	defer s.RUnlock()
	values := make([]product.Product, 0, len(s.skus))
	for _, v := range s.skus {
		values = append(values, v)
	}
	return values
}

func NewProductService() product.Service {
	return &productService{
		skus: map[int]product.Product{
			1: {ID: 1, Name: "Green grapes", Category: product.ProductCategoryFood, RetailPrice: 60},
			2: {ID: 2, Name: "Nike v Puma Hat", Category: product.ProductCategoryClothing, RetailPrice: 60},
			3: {ID: 3, Name: "PlayStation 5 Digital Edition", Category: product.ProductCategoryElectronic, RetailPrice: 45},
		},
	}
}
