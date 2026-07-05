package user

import "ecommerce/domain"

type UserServe interface {
	store(user domain.User) error
	singleUser(id string) (domain.User, error)
	userlist() ([]domain.User, error)
	updateUser(id string, user domain.User) error
	deleteUser(id string) error
	rowCount() (int64, error)
	findByEmail(email string) (*domain.User, error)
	existsByEmail(email string) (bool, error)
}

type userService struct {
	repository UserPort
}

// findByEmail implements [UserServe].
func (u *userService) findByEmail(email string) (*domain.User, error) {
	return u.repository.FindByEmail(email)
}

// existsByEmail implements [UserServe].
func (u *userService) existsByEmail(email string) (bool, error) {
	return u.repository.ExistsByEmail(email)
}

func NewUserService(repository UserPort) UserServe {
	return &userService{repository: repository}
}

// deleteUser implements [UserServe].
func (u *userService) deleteUser(id string) error {
	return u.repository.deleteUser(id)
}

// rowCount implements [UserServe].
func (u *userService) rowCount() (int64, error) {
	return u.repository.rowCount()
}

// singleUser implements [UserServe].
func (u *userService) singleUser(id string) (domain.User, error) {
	return u.repository.getUserByID(id)
}

// store implements [UserServe].
func (u *userService) store(user domain.User) error {
	return u.repository.createUser(user)
}

// updateUser implements [UserServe].
func (u *userService) updateUser(id string, user domain.User) error {
	return u.repository.updateUser(id, user)
}

// userlist implements [UserServe].
func (u *userService) userlist() ([]domain.User, error) {
	return u.repository.getAllUsers()
}
