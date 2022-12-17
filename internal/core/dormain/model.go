package dormain

type Product struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	Category   string  `json:"category"`
	CategoryID int     `json:"category_id"`
}
