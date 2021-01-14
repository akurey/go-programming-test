package market

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/akurey/go-programming-test/inventory"
	"github.com/akurey/go-programming-test/order"
	"github.com/akurey/go-programming-test/product"
	"github.com/julienschmidt/httprouter"
)

func MakeHandler(os order.Service, is inventory.Service, ps product.Service) http.Handler {
	r := httprouter.New()
	r.POST("/orders", createOrderHandler(os))
	r.GET("/orders/:id", getOrderHandler(os, is, ps))
	r.GET("/products", listProductItemsHandler(is, ps))
	r.GET("/products/:id", getProductItemHandler(is, ps))
	return r
}

func createOrderHandler(os order.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()
		ids, rejectedIDs := readItemIDs(file)
		id, rejectedItems, err := os.CreateOrder(ids)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		type rejectedItem struct {
			ID     string `json:"id"`
			Reason string `json:"reason"`
		}
		// Create response
		res := struct {
			ID            int            `json:"id"`
			RejectedItems []rejectedItem `json:"rejected_ids"`
		}{ID: id}
		for _, item := range rejectedItems {
			res.RejectedItems = append(res.RejectedItems, rejectedItem{ID: strconv.Itoa(item.ID), Reason: item.ReasonMessage})
		}
		for _, value := range rejectedIDs {
			res.RejectedItems = append(res.RejectedItems, rejectedItem{ID: value, Reason: "Bad format"})
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(res)
	}
}

func getOrderHandler(os order.Service, is inventory.Service, ps product.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id, err := strconv.Atoi(params.ByName("id"))
		if err != nil {
			http.Error(w, "id is invalid", http.StatusBadRequest)
			return
		}
		ord, err := os.Find(id)
		if err != nil {
			http.Error(w, "order not found", http.StatusNotFound)
			return
		}
		orderItems := make([]inventory.Item, 0, len(ord.Items))
		for _, v := range ord.Items {
			item, _ := is.Find(v.ID)
			orderItems = append(orderItems, item)
		}
		response := struct {
			ID    int              `json:"id"`
			Items []productSummary `json:"items"`
		}{
			ID:    ord.ID,
			Items: itemSummaries(orderItems, is, ps),
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(response)
	}
}

func listProductItemsHandler(is inventory.Service, ps product.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Get all items
		items := is.ListItems()
		// Create summaries
		productSummaries := itemSummaries(items, is, ps)
		// Write response
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(productSummaries)
	}
}

func getProductItemHandler(is inventory.Service, ps product.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id, err := strconv.Atoi(params.ByName("id"))
		if err != nil {
			http.Error(w, "id is invalid", http.StatusBadRequest)
			return
		}
		item, err := is.Find(id)
		if err != nil {
			http.Error(w, "item not found", http.StatusNotFound)
			return
		}
		product, err := ps.Find(item.SKU)
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		prodSummary := productSummary{
			ID:       item.ID,
			Category: string(product.Category),
			Status:   string(item.Status),
			Name:     product.Name,
			Price:    product.RetailPrice,
			SoldFor:  item.SoldPrice,
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(prodSummary)
	}
}

func readItemIDs(r io.Reader) ([]int, []string) {
	var ids []int
	var rejected []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		val := scanner.Text()
		if id, err := strconv.Atoi(val); err != nil {
			rejected = append(rejected, val)
		} else {
			ids = append(ids, id)
		}
	}
	return ids, rejected
}

type productSummary struct {
	ID       int     `json:"id"`
	Category string  `json:"category"`
	Status   string  `json:"status"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	SoldFor  float32 `json:"sold_for"`
}

func itemSummaries(items []inventory.Item, is inventory.Service, ps product.Service) []productSummary {
	summaries := make([]productSummary, len(items))
	for i, item := range items {
		invItem, _ := is.Find(item.ID)
		product, _ := ps.Find(invItem.SKU)

		summaries[i] = productSummary{
			ID:       item.ID,
			Category: string(product.Category),
			Status:   string(item.Status),
			Name:     product.Name,
			Price:    product.RetailPrice,
			SoldFor:  item.SoldPrice,
		}
	}
	return summaries
}
