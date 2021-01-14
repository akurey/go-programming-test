package order

type Repository interface {
	New(order *Order) (int, error)
	FindAll() []Order
	Find(orderID int) (Order, error)
}
