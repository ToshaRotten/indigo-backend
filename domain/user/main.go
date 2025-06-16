package user

type Service struct {
	storage Storage
}

func CreateService(storage Storage) Service {
	return Service{
		storage: storage,
	}
}
