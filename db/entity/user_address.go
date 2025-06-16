package entity

type UserAddress interface {
	Model() UserAddressModel
}

type UserAddressModel struct {
	ID   int
	Name string
}
