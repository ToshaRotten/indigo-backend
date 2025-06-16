package entity

type User interface {
	Model() UserModel
}

type UserModel struct {
	ID             int
	FullName       string
	PasswordHash   string
	Email          string
	UserCategoryID int
}
