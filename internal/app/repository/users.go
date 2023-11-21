package repository

import (
	"rip2023/internal/app/ds"
)

func (r *Repository) UsersList() (*[]ds.Users, error) {
	var users []ds.Users
	result := r.db.Find(&users)
	return &users, result.Error
}

func (r *Repository) Register(user *ds.Users) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByLogin(login string) (*ds.Users, error) {
	user := &ds.Users{}

	if err := r.db.Where("login = ?", login).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
