package in_memory

import (
	"github.com/akurey/go-programming-test/inventory"
	"github.com/akurey/go-programming-test/order"
	"github.com/akurey/go-programming-test/product"
)

func NewOrderService(ps product.Service, is inventory.Service) order.Service {
	return &orderService{
		productService:   ps,
		inventoryService: is,
		repository:       newOrderRepository(),
	}
}

type orderService struct {
	productService   product.Service
	inventoryService inventory.Service
	repository       order.Repository
}

func (s *orderService) CreateOrder(items []int) (int, []order.RejectedItem, error) {
	var rejectedItems []order.RejectedItem
	ord := &order.Order{}

	for _, id := range items {
		item, err := s.inventoryService.Find(id)
		if err != nil {
			rejectedItems = append(rejectedItems, order.RejectedItem{ID: id, ReasonMessage: err.Error()})
			continue
		}

		product, err := s.productService.Find(item.SKU)
		if err != nil {
			rejectedItems = append(rejectedItems, order.RejectedItem{ID: id, ReasonMessage: err.Error()})
			continue
		}

		finalPrice := product.RetailPrice
		err = s.inventoryService.Sell(id, finalPrice)
		if err != nil {
			rejectedItems = append(rejectedItems, order.RejectedItem{ID: id, ReasonMessage: err.Error()})
			continue
		}
		ord.Items = append(ord.Items, &order.Item{
			ID:      item.ID,
			SoldFor: finalPrice,
		})
	}
	id, err := s.repository.New(ord)
	return id, rejectedItems, err
}

func (s *orderService) ListOrders() []order.Order {
	return s.repository.FindAll()
}

func (s *orderService) Find(id int) (order.Order, error) {
	return s.repository.Find(id)
}
