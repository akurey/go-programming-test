package inventory

import (
	"io"
	"time"
)

type ItemStatus string

const (
	ItemStatusAvailable ItemStatus = "available"
	ItemStatusSold                 = "sold"
	ItemStatusReserved             = "reserved"
)

type Item struct {
	ID        int
	SKU       int
	Status    ItemStatus
	Price     float32
	SoldPrice float32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Find(ID int) (Item, error)
	ListItems() []Item
	Sell(ID int, soldPrice float32) error
	io.Closer
}
