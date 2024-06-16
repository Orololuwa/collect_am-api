package dbrepo

import (
	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	"gorm.io/gorm"
)

type userOrm struct {
	db *gorm.DB
}
func NewUserDBRepo(db *driver.DB) repository.UserDBRepo {
	return &userOrm{
		db: db.Gorm,
	}
}

type testUserDBRepo struct {
}
func NewUserTestingDBRepo() repository.UserDBRepo {
	return &testUserDBRepo{
	}
}


func (o *userOrm) GetOneByID(id uint) (user models.User, err error) {
	result := o.db.Model(&models.User{}).Where("id = ?", id).First(&user)
	return user, result.Error
}

func (o *userOrm) GetOneByEmail(email string) (user models.User, err error) {
	result := o.db.Model(&models.User{}).Where("email = ?", email).First(&user)
	return user, result.Error
}

func (o *userOrm) GetOneByPhone(phone string) (user models.User, err error) {
	result := o.db.Model(&models.User{}).Where("phone = ?", phone).First(&user)
	return user, result.Error
}

func (o *userOrm) InsertUser(user models.User, tx ...*gorm.DB) (id uint, err error) {
	db := o.db
    if len(tx) > 0 && tx[0] != nil {
        db = tx[0]
    }

	result := db.Model(&models.User{}).Create(&user)
	return user.ID, result.Error
}

func (o *userOrm) UpdateUser(user models.User, tx ...*gorm.DB) (err error) {
	db := o.db
    if len(tx) > 0 && tx[0] != nil {
        db = tx[0]
    }

	result := db.Model(&models.User{}).Model(&user).Updates(&user)
	return result.Error
}