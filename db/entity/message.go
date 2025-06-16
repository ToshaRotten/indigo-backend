package entity

type Message interface {
	Model() MessageModel
}

type MessageModel struct {
	ID   int
	Name string
}
