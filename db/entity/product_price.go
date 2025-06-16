package entity

import "time"

type ProductPrice interface {
	Model() ProductPriceModel
}

type ProductPriceModel struct {
	ID               int
	ModificationDate time.Time
	NewPrice         int
	ProductID        int
}
