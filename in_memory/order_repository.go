package in_memory

import (
	"fmt"
	"sync"

	"github.com/akurey/go-programming-test/order"
)

type orderRepository struct {
	sync.RWMutex
	orders map[int]*order.Order
}

func newOrderRepository() order.Repository {
	return &orderRepository{
		orders: make(map[int]*order.Order),
	}
}

func (r *orderRepository) New(order *order.Order) (int, error) {
	r.Lock()
	defer r.Unlock()
	id := len(r.orders) + 1
	order.ID = id
	r.orders[id] = order
	return id, nil
}

func (r *orderRepository) Find(id int) (order.Order, error) {
	r.RLock()
	defer r.RUnlock()
	c, ok := r.orders[id]
	if !ok {
		return order.Order{}, fmt.Errorf("No order found")
	}
	return *c, nil
}

func (r *orderRepository) FindAll() []order.Order {
	r.RLock()
	defer r.RUnlock()
	values := make([]order.Order, 0, len(r.orders))
	for _, v := range r.orders {
		values = append(values, *v)
	}
	return values
}
