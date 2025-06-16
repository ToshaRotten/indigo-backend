package entity

type ProductCategory interface {
	Model() ProductCategoryModel
}

type ProductCategoryModel struct {
	ID   int
	Name string
}
