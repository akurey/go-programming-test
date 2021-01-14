package in_memory

import (
	"fmt"
	"sync"
	"time"

	"github.com/akurey/go-programming-test/inventory"
)

type inventoryService struct {
	sync.RWMutex
	items map[int]*inventory.Item
}

func NewInventoryService() inventory.Service {
	return &inventoryService{
		items: testData(),
	}
}

func (s *inventoryService) Sell(id int, price float32) error {
	s.Lock()
	defer s.Unlock()
	if item, ok := s.items[id]; ok {
		if item.Status != inventory.ItemStatusSold {
			item.Status = inventory.ItemStatusSold
			item.SoldPrice = price
			return nil
		}
		return fmt.Errorf("Could not sell item")
	}
	return fmt.Errorf("Item does not exist")
}

func (s *inventoryService) Find(id int) (inventory.Item, error) {
	s.RLock()
	defer s.RUnlock()
	item, ok := s.items[id]
	if ok {
		return *item, nil
	}
	return inventory.Item{}, fmt.Errorf("Item not found")
}

func (s *inventoryService) ListItems() []inventory.Item {
	s.RLock()
	defer s.RUnlock()
	values := make([]inventory.Item, 0, len(s.items))
	for _, v := range s.items {
		values = append(values, *v)
	}
	return values
}

func (s *inventoryService) Close() error {
	return nil
}

func testData() map[int]*inventory.Item {
	return map[int]*inventory.Item{
		1:  {ID: 1, Status: inventory.ItemStatusSold, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		2:  {ID: 2, Status: inventory.ItemStatusSold, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		3:  {ID: 3, Status: inventory.ItemStatusSold, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		4:  {ID: 4, Status: inventory.ItemStatusSold, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		5:  {ID: 5, Status: inventory.ItemStatusAvailable, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		6:  {ID: 6, Status: inventory.ItemStatusAvailable, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		7:  {ID: 7, Status: inventory.ItemStatusAvailable, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		8:  {ID: 8, Status: inventory.ItemStatusAvailable, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		9:  {ID: 9, Status: inventory.ItemStatusAvailable, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		10: {ID: 10, Status: inventory.ItemStatusReserved, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		11: {ID: 11, Status: inventory.ItemStatusSold, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		12: {ID: 12, Status: inventory.ItemStatusReserved, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		13: {ID: 13, Status: inventory.ItemStatusReserved, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		14: {ID: 14, Status: inventory.ItemStatusReserved, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		15: {ID: 15, Status: inventory.ItemStatusReserved, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		16: {ID: 16, Status: inventory.ItemStatusReserved, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		17: {ID: 17, Status: inventory.ItemStatusReserved, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		18: {ID: 18, Status: inventory.ItemStatusAvailable, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		19: {ID: 19, Status: inventory.ItemStatusAvailable, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		20: {ID: 20, Status: inventory.ItemStatusAvailable, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		21: {ID: 21, Status: inventory.ItemStatusAvailable, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		22: {ID: 22, Status: inventory.ItemStatusAvailable, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		23: {ID: 23, Status: inventory.ItemStatusAvailable, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		24: {ID: 24, Status: inventory.ItemStatusAvailable, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		25: {ID: 25, Status: inventory.ItemStatusAvailable, SKU: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		26: {ID: 26, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		27: {ID: 27, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		28: {ID: 28, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		29: {ID: 29, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		30: {ID: 30, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		31: {ID: 31, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		32: {ID: 32, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		33: {ID: 33, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		34: {ID: 34, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		35: {ID: 35, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		36: {ID: 36, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		37: {ID: 37, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		38: {ID: 38, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		39: {ID: 39, Status: inventory.ItemStatusAvailable, SKU: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		40: {ID: 40, Status: inventory.ItemStatusAvailable, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		41: {ID: 41, Status: inventory.ItemStatusAvailable, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		42: {ID: 42, Status: inventory.ItemStatusAvailable, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		43: {ID: 43, Status: inventory.ItemStatusAvailable, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		44: {ID: 44, Status: inventory.ItemStatusAvailable, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		45: {ID: 45, Status: inventory.ItemStatusReserved, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		46: {ID: 46, Status: inventory.ItemStatusReserved, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		47: {ID: 47, Status: inventory.ItemStatusReserved, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		48: {ID: 48, Status: inventory.ItemStatusReserved, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		49: {ID: 49, Status: inventory.ItemStatusReserved, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		50: {ID: 50, Status: inventory.ItemStatusReserved, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		51: {ID: 51, Status: inventory.ItemStatusReserved, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		52: {ID: 52, Status: inventory.ItemStatusReserved, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		53: {ID: 53, Status: inventory.ItemStatusReserved, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		54: {ID: 54, Status: inventory.ItemStatusReserved, SKU: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
}
