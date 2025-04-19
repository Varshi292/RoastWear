// Package models ...
package models

type CartItem struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
