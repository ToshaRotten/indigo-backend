package entity

import "time"

type Order interface {
	Model() OrderModel
}

type OrderModel struct {
	ID                  int
	OrderDate           time.Time
	DesiredDeliveryDate time.Time
	TotalAmount         float32
	StatusID            int
	UserID              int
	AddressID           int
}
