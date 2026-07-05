package user

import "ecommerce/domain"

type UserServe interface {
	Store(user domain.User) error
	SingleUser(id string) (domain.User, error)
	Userlist() ([]domain.User, error)
	UpdateUser(id string, user domain.User) error
	DeleteUser(id string) error
	RowCount() (int64, error)
	FindByEmail(email string) (*domain.User, error)
	ExistsByEmail(email string) (bool, error)
}

type userService struct {
	repository UserPort
}



func NewUserService(repository UserPort) UserServe {
	return &userService{repository: repository}
}

// deleteUser implements [UserServe].
func (u *userService) DeleteUser(id string) error {
	return u.repository.deleteUser(id)
}

// rowCount implements [UserServe].
func (u *userService) RowCount() (int64, error) {
	return u.repository.rowCount()
}

// singleUser implements [UserServe].
func (u *userService) SingleUser(id string) (domain.User, error) {
	return u.repository.getUserByID(id)
}

// store implements [UserServe].
func (u *userService) Store(user domain.User) error {
	return u.repository.createUser(user)
}

// updateUser implements [UserServe].
func (u *userService) UpdateUser(id string, user domain.User) error {
	return u.repository.updateUser(id, user)
}

// userlist implements [UserServe].
func (u *userService) Userlist() ([]domain.User, error) {
	return u.repository.getAllUsers()
}
// findByEmail implements [UserServe].
func (u *userService) FindByEmail(email string) (*domain.User, error) {
	return u.repository.FindByEmail(email)
}

// existsByEmail implements [UserServe].
func (u *userService) ExistsByEmail(email string) (bool, error) {
	return u.repository.ExistsByEmail(email)
}