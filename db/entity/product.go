package entity

type Product interface {
	Model() ProductModel
}

type ProductModel struct {
	ID                 int
	Name               string
	Weight             float64
	ProductComposition string
	MinStorageTemp     int
	MaxStorageTemp     int
	ShelfLife          string
	Picture            string
	ProductCategoryID  int
	TypeOfPackagingID  int
}
