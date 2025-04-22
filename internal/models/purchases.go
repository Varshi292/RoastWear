package models

// PurchaseSwagger is a Swagger-friendly version of the Purchase model
type Purchase struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Username  string `json:"username"`
	Products  string `json:"products"` // Format: productID:quantity#productID:quantity
	Total     string `json:"total"`
}
