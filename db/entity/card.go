package entity

type Card interface {
	Model() CardModel
}

type CardModel struct {
	ID   int
	Name string
}
