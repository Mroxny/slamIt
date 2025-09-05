package repository

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/model"
)

type UserRepository struct {
	users  []model.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  []model.User{},
		nextID: 1,
	}
}

func (r *UserRepository) GetAll() []model.User {
	return r.users
}

func (r *UserRepository) GetByID(id int) (*model.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) Create(user model.User) model.User {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)
	return user
}
