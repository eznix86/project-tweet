package repository

import (
	"errors"

	"github.com/theArtechnology/project-tweet/src/entity"
)

// UserRepository is a Repository (meaning Storage) for all my users
type UserRepository struct {
	UserList []entity.User // in memory userList
}

// Find is
func (r *UserRepository) Find(id int) (*entity.User, error) {
	for _, user := range r.UserList {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

// Add is
func (r *UserRepository) Add(entity entity.User) error {
	r.UserList = append(r.UserList, entity)
	return nil
}

// Remove is
func (r *UserRepository) Remove(id int) error {
	for index, user := range r.UserList {
		if user.ID == id {
			r.UserList = append(r.UserList[:index], r.UserList[index:]...)
			return nil
		}
	}

	return errors.New("User not found")
}

func (r *UserRepository) getAll(entity entity.User) ([]entity.User, error) {
	return r.UserList, nil
}
