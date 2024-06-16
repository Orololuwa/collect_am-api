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
func (o *testUserDBRepo) GetOneByID(id uint) (user models.User, err error) {
	return user, nil
}

func (o *testUserDBRepo) GetOneByEmail(email string) (user models.User, err error) {
	if email == "johndoe@exists.com" { //email exists
		return models.User{ID: 1}, nil
	}

	if email == "test_fail@test.com" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Testpass123###"), bcrypt.DefaultCost)
		return models.User{Password: string(hashedPassword)}, nil
	}	
	if email == "test_correct@test.com" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Testpass123###"), bcrypt.DefaultCost)
		return models.User{Password: string(hashedPassword)}, nil
	}
	return user, nil
}

func (o *testUserDBRepo) GetOneByPhone(phone string) (user models.User, err error) {
	if phone == "+2340000000001" { //phone exists
		return models.User{ID: 1}, nil
	}
	return user, nil
}

func (o *testUserDBRepo) InsertUser(user models.User, tx ...*gorm.DB) (id uint, err error) {
	if user.FirstName == "fail" {//failed insert operation
		return id, errors.New("failed to insert")
	}
	return id, nil
}

func (m *testUserDBRepo) CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error){
	var newId int

	if user.FirstName == "fail" {
		return newId, errors.New("")
	}

	return newId, nil
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