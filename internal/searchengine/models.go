package searchengine

type Product struct {
	ProductID   string  `json:"product_id"`
	Brand       string  `json:"brand"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int32   `json:"stock"`
}
