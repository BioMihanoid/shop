package structures

type Product struct {
	ID          int64   `json:"id"`
	Price       float64 `json:"price"`
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
}
