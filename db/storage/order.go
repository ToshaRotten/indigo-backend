package storage

import (
	"main/db/ent"
)

func NewOrderStorage(client *ent.Client) order {
	return order{
		Client: client,
	}
}

type order struct {
	Client *ent.Client
}
