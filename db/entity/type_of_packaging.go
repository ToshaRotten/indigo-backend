package entity

type TypeOfPackaging interface {
	Model() TypeOfPackagingModel
}

type TypeOfPackagingModel struct {
	ID   int
	Name string
}
