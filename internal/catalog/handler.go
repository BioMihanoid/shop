package catalog

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"shop/internal/structures"
)

type service interface {
	// CreateProduct(ctx context.Context, product structures.Product) (int64, error)
	// GetProduct(ctx context.Context, id int) (structures.Product, error)
	// GetProducts(ctx context.Context) ([]structures.Product, error)
}

type Handler struct {
	service  service
	products map[int64]structures.Product
}

func New(service service) *Handler {
	return &Handler{
		service: service,
		products: map[int64]structures.Product{
			1: {
				ID:          1,
				Price:       100,
				Count:       3,
				Name:        "Bubble tea",
				Description: "tea with bubble",
				Brand:       "Lipton",
			},
			2: {
				ID:          2,
				Price:       120,
				Count:       5,
				Name:        "Bubble tea blackberry",
				Description: "tea with blackberry bubble",
				Brand:       "Lipton",
			},
		},
	}
}

func (h *Handler) CreateProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /catalog/create")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("create products"))
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	log.Println("GET /catalog/product/" + idStr)
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(h.products[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /catalog/products")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /catalog/health")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("catalog health\n"))
}
