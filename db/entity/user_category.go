package entity

type UserCategory interface {
	Model() UserCategoryModel
}

type UserCategoryModel struct {
	ID   int
	Name string
}
