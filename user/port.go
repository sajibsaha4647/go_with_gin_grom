package user

import "ecommerce/domain"

type UserPort interface {
	createUser(user domain.User) error
	getUserByID(id string) (domain.User, error)
	getAllUsers() ([]domain.User, error)
	updateUser(id string, user domain.User) error
	deleteUser(id string) error
	rowCount() (int64, error)
}
