package user

import (
	"ecommerce/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserPort {
	return &UserRepository{db: db}
}


// createUser implements [UserPort].
func (u *UserRepository) createUser(user domain.User) error {
	return u.db.Create(&user).Error
}

// deleteUser implements [UserPort].
func (u *UserRepository) deleteUser(id string) error {
	return u.db.Delete(&domain.User{}, id).Error
}

// getAllUsers implements [UserPort].
func (u *UserRepository) getAllUsers() ([]domain.User, error) {
	var users []domain.User
	return users, u.db.Find(&users).Error
}

// getUserByID implements [UserPort].
func (u *UserRepository) getUserByID(id string) (domain.User, error) {
	var user domain.User
	return user, u.db.First(&user, id).Error
}

// rowCount implements [UserPort].
func (u *UserRepository) rowCount() (int64, error) {
	var count int64
	err := u.db.Model(&domain.User{}).Count(&count).Error
	return count, err
}

// updateUser implements [UserPort].
func (u *UserRepository) updateUser(id string, user domain.User) error {
	return u.db.Model(&domain.User{}).Where("id = ?", id).Updates(user).Error
}

