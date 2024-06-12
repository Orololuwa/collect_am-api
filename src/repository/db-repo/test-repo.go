package dbrepo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Orololuwa/collect_am-api/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User
func (m *testUserDBRepo) CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error){
	var newId int

	if user.FirstName == "fail" {
		return newId, errors.New("")
	}

	return newId, nil
}

func (m *testUserDBRepo) GetAUser(ctx context.Context, tx *sql.Tx, u models.User) (*models.User, error) {
	// user := &models.User{ID: 2}
	var user *models.User

	if u.Email == "johndoe@fail.com" {
		return user, errors.New("")
	}
	if u.Email == "johndoe@exists.com" {
		return &models.User{Model: gorm.Model{
			ID: 1,
		},}, nil
	}
	if u.Email == "johndoe@null.com" {
		return user, nil
	}	
	if u.Email == "test_fail@test.com" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Testpass123###"), bcrypt.DefaultCost)
		return &models.User{Password: string(hashedPassword)}, nil
	}	
	if u.Email == "test_correct@test.com" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Testpass123###"), bcrypt.DefaultCost)
		return &models.User{Password: string(hashedPassword)}, nil
	}

	if u.Phone == "+2340000000000" {
		return user, errors.New("")
	}
	if u.Phone == "+2340000000001" {
		return &models.User{Model: gorm.Model{
			ID: 1,
		},}, nil
	}
	if u.Phone == "+2340000000002" {
		return user, nil
	}

	user = &models.User{Model: gorm.Model{
        ID: 1,
    },}
	return user, nil
}

func (o *testUserDBRepo) GetOneByID(id uint) (user models.User, err error) {
	return user, nil
}

func (o *testUserDBRepo) GetOneByEmail(email string) (user models.User, err error) {
	return user, nil
}

func (o *testUserDBRepo) GetOneByPhone(phone string) (user models.User, err error) {
	return user, nil
}

func (o *testUserDBRepo) InsertUser(user models.User, tx ...*gorm.DB) (id uint, err error) {
	return id, nil
}

func (o *testUserDBRepo) UpdateUser(user models.User, tx ...*gorm.DB) (err error) {
	return nil
}

// Business
func (m *testBusinessDBRepo) CreateBusiness(ctx context.Context, tx *sql.Tx, business models.Business) (int, error){
	var id int
	return id, nil
}

func (m *testBusinessDBRepo) GetUserBusiness(ctx context.Context, tx *sql.Tx, userId int, b models.Business) (*models.Business, error) {
	var business *models.Business

	return business, nil
}

func (m *testBusinessDBRepo) UpdateBusinessOld(ctx context.Context, tx *sql.Tx, business models.Business) error{
	return nil
} 

func (m *testBusinessDBRepo) GetOneByUserId(userId uint) (businesses models.Business, err error){
	return businesses, nil
}

func (o *testBusinessDBRepo) InsertBusiness(business models.Business, tx ...*gorm.DB) (id uint, err error) {
	return id, nil
}

func (o *testBusinessDBRepo) UpdateBusiness(updateData map[string]interface{},  where models.Business, tx ...*gorm.DB) (err error) {
	return nil
}

// kyc
func (m *testKycDBRepo) CreateKyc(ctx context.Context, tx *sql.Tx, kyc models.Kyc) (int, error){
	var id int
	return id, nil
}

func (m *testKycDBRepo) GetBusinessKyc(ctx context.Context, tx *sql.Tx, business_id int, b models.Kyc) (*models.Kyc, error){
	var kyc models.Kyc

	return &kyc, nil
}

func (o *testKycDBRepo) GetOneByID(id uint) (kyc models.Kyc, err error) {
	return kyc, nil
}

func (o *testKycDBRepo) InsertKyc(kyc models.Kyc, tx ...*gorm.DB) (id uint, err error) {
	return id, nil
}

func (o *testKycDBRepo) UpdateKyc(updateData map[string]interface{}, where models.Kyc, tx ...*gorm.DB) (err error) {
	return nil
}