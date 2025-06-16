package entity

type OrderStatus interface {
	Model() OrderStatusModel
}

type OrderStatusModel struct {
	ID   int
	Name string
}
