package gorm

import (
	"final_project/server/models"
	"final_project/server/repositories"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repositories.UserRepo {
	return &userRepo{db: db}
}

func (u *userRepo) CreateUser(orders *models.User) error {
	return u.db.Create(orders).Error
}

func (u *userRepo) GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := u.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) GetUserByUser(username string) (*models.User, error) {
	user := models.User{}
	err := u.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) UpdateUser(id int, userReq *models.ReqUser) error {
	user := models.User{}
	err := u.db.Model(&user).Where("id = ?", id).Updates(models.User{Email: userReq.Email, Username: userReq.Username}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) DeleteUser(id int) error {
	user := models.User{}
	err := u.db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
