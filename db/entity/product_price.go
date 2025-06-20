package entity

import "time"

type ProductPrice interface {
	Model() ProductPriceModel
}

// Модель для пересмотров цены
type ProductPriceModel struct {
	ID               int
	ModificationDate time.Time
	NewPrice         int
	ProductID        int
}
