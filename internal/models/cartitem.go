package models

// CartItemSwagger represents a Swagger-friendly cart item
type CartItem struct {
	ID         uint   `json:"id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Username   string `json:"username"`
	ProductID  int    `json:"productid"`
	Quantity   int    `json:"quantity"`
	TotalPrice string `json:"total_price"`
}
