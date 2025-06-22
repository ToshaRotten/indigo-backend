package storage

import (
	"context"

	"main/db/ent"
	"main/db/ent/user"
	"main/db/entity"
)

func NewUserStorage(client *ent.Client) userStorage {
	return userStorage{
		Client: client,
	}
}

type userStorage struct {
	Client *ent.Client
}

func (u userStorage) Create(ctx context.Context, user entity.UserModel) error {
	create := u.Client.User.Create().
		SetFullName(user.FullName).
		SetPasswordHash(user.PasswordHash).
		SetEmail(user.Email).
		SetUserCategoryID(user.UserCategoryID)

	err := create.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u userStorage) Get(ctx context.Context, userID int) (entity.UserModel, error) {
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

func (u userStorage) GetAll(ctx context.Context) ([]entity.UserModel, error) {
	users, err := u.Client.User.Query().
		WithUserCategory().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var userData []entity.UserModel
	for _, user := range users {
		var userModel entity.UserModel

		userModel.FullName = user.FullName
		userModel.Email = user.Email
		userModel.UserCategoryID = user.UserCategoryID

		userData = append(userData, userModel)
	}

	return userData, nil
}

func (u userStorage) GetByLogin(ctx context.Context, login string) (entity.UserModel, error) {
	userData, err := u.Client.User.Query().
		Where(user.UsernameEQ(login)).
		WithUserCategory().
		Only(ctx)
	if err != nil {
		return entity.UserModel{}, err
	}

	var user entity.UserModel

	user.FullName = userData.FullName
	user.Username = userData.Username
	user.Email = userData.Email
	user.PasswordHash = userData.PasswordHash
	user.UserCategoryID = userData.UserCategoryID

	return user, nil
}

func (u userStorage) Delete(ctx context.Context, userID int) error {
	remove := u.Client.User.DeleteOneID(userID)
	err := remove.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
