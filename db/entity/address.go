package entity

type Address interface {
	Model() AddressModel
}

type AddressModel struct {
	ID   int
	Name string
}
