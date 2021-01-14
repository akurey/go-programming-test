package order

type Order struct {
	// ID of the order
	ID int
	// The items that were sold
	Items []*Item
}

type Item struct {
	ID      int
	SoldFor float32
}

type RejectedItem struct {
	ID            int
	ReasonMessage string
}

type Service interface {
	CreateOrder(items []int) (orderID int, rejectedItems []RejectedItem, err error)
	ListOrders() []Order
	Find(orderID int) (Order, error)
}
