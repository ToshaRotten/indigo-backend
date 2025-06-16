package entity

type ShoppingCart interface {
	Model() ShoppingCartModel
}

type ShoppingCartModel struct {
	ID        int
	ProductID int
	Quantity  int
	UserID    int
}
