package repository

import (
	"github.com/Orololuwa/collect_am-api/src/models"
	"gorm.io/gorm"
)

type UserDBRepo interface {
	GetOneByID(id uint) (user models.User, err error)
	GetOneByEmail(email string) (user models.User, err error)
	GetOneByPhone(phone string) (user models.User, err error)
	InsertUser(user models.User, tx ...*gorm.DB) (id uint, err error)
	UpdateUser(user models.User, tx ...*gorm.DB) (err error)
}

type BusinessDBRepo interface {
	GetOneByUserId(userId uint) (businesses models.Business, err error)
	InsertBusiness(business models.Business, tx ...*gorm.DB) (id uint, err error)
	UpdateBusiness(updateData map[string]interface{}, where models.Business, tx ...*gorm.DB) (err error)
}

type KycDBRepo interface {
	InsertKyc(kyc models.Kyc, tx ...*gorm.DB) (id uint, err error)
	UpdateKyc(updateData map[string]interface{}, where models.Kyc, tx ...*gorm.DB) (err error)
}