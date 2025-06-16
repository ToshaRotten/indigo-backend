package entity

type OrderDetail interface {
	Model() OrderDetailModel
}

type OrderDetailModel struct {
	ID        int
	ProductID int
	OrderID   int
}
