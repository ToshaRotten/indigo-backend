package entity

type UserToken interface {
	Model() UserCategoryModel
}

type UserTokenModel struct {
	ID   int
	Name string
}
