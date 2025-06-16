package storage

import (
	"context"

	"main/db/ent"
	"main/db/entity"
)

func NewUserStorage(client *ent.Client) user {
	return user{
		Client: client,
	}
}

type user struct {
	Client *ent.Client
}

func (u user) Create(ctx context.Context, user entity.UserModel) error {
	create := u.Client.User.Create().
		SetFullName(user.FullName).
		SetPassword(user.PasswordHash).
		SetEmail(user.Email).
		SetUserCategoryID(user.UserCategoryID)

	err := create.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u user) Get(ctx context.Context, userID int) (entity.UserModel, error) {
	user, err := u.Client.User.Get(ctx, userID)
	if err != nil {
		return entity.UserModel{}, err
	}

	var userData entity.UserModel
	userData.FullName = user.FullName
	userData.Email = user.Email
	userData.UserCategoryID = user.UserCategoryID

	return userData, nil
}
