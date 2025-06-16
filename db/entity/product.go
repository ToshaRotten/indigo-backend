package entity

type Product interface {
	Model() ProductModel
}

type ProductModel struct {
	ID                 int
	Name               string
	Weight             float32
	ProductComposition float32
	MinStorageTemp     float32
	MaxStorageTemp     float32
	ShelfLife          float32
	Picture            string
	ProductCategoryID  int
	TypeOfPackagingID  int
}
