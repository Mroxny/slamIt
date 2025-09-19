package repository

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
)

type UserRepository struct {
	users  []api.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  []api.User{},
		nextID: 1,
	}
}

func (r *UserRepository) GetAll() []api.User {
	return r.users
}

func (r *UserRepository) GetByID(id string) (*api.User, error) {
	for _, u := range r.users {
		if *u.Id == id {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) Create(user *api.User) (*api.User, error) {
	r.users = append(r.users, *user)
	return user, nil
}

func (r *UserRepository) Update(id string, updated api.User) (*api.User, error) {
	for i, u := range r.users {
		if *u.Id == id {
			if *updated.Name == "" || *updated.Email == "" {
				return nil, errors.New("invalid input")
			}
			r.users[i].Name = updated.Name
			r.users[i].Email = updated.Email
			return &r.users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) Delete(id string) error {
	for i, u := range r.users {
		if *u.Id == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *UserRepository) GetByEmail(email string) (*api.User, error) {
	for _, user := range r.users {
		if *user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
