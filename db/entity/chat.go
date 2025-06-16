package entity

type Chat interface {
	Model() ChatModel
}

type ChatModel struct {
	ID   int
	Name string
}
