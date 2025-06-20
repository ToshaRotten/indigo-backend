package entity

type User interface {
	Model() UserModel
}

type UserModel struct {
	ID             int
	Username       string
	FullName       string
	PasswordHash   string
	Email          string
	UserCategoryID int
}
