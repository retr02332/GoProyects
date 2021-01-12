package store

import "../product"

// Store ...
type Store struct {
	Name        string
	Description string
	Product     product.Product
}
