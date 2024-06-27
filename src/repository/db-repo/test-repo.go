package dbrepo

import (
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
	if email == "johndoe@null.com" { //email doesn not exists
		return user, errors.New("record not found")
	}
	if email == "hash_fail@test.com" {
		return user, errors.New("hashing error")
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

func (o *testUserDBRepo) UpdateUser(user models.User, tx ...*gorm.DB) (err error) {
	return nil
}

// Business
func (m *testBusinessDBRepo) GetOneById(id uint) (business models.Business, err error){
	if id == 0 {
		return business, errors.New("error getting business")
	}
	if id == 1 {
		return business, errors.New("record not found")
	}
	return business, nil
}

func (m *testBusinessDBRepo) GetOneByUserId(userId uint) (businesses models.Business, err error){
	return businesses, nil
}

func (o *testBusinessDBRepo) InsertBusiness(business models.Business, tx ...*gorm.DB) (id uint, err error) {
	if (business.Email == "invalid"){
		return id, errors.New("failed to insert business") //fail case
	}

	return id, nil //success case
}

func (o *testBusinessDBRepo) UpdateBusiness(updateData map[string]interface{},  where models.Business, tx ...*gorm.DB) (err error) {
	if (updateData["name"] == "invalid"){
		return errors.New("failed to insert business") //fail case
	}

	return nil //success case
}

// kyc
func (o *testKycDBRepo) GetOneByID(id uint) (kyc models.Kyc, err error) {
	return kyc, nil
}

func (o *testKycDBRepo) InsertKyc(kyc models.Kyc, tx ...*gorm.DB) (id uint, err error) {
	if (kyc.BVN == "invalid"){
		return id, errors.New("failed to insert kyc")
	}
	return id, nil
}

func (o *testKycDBRepo) UpdateKyc(updateData map[string]interface{}, where models.Kyc, tx ...*gorm.DB) (err error) {
	if (updateData["bvn"] == "invalid"){
		return errors.New("failed to insert kyc") //fail case
	}

	return nil
}

// products
func (o *testProductDBRepo) GetOneById(id uint) (product models.Product, err error){
	return product, err
}
func (o *testProductDBRepo) CreateProduct(createData map[string]interface{},  where models.Product, tx ...*gorm.DB) (id uint, err error){
	return id, err
}
func (o *testProductDBRepo) InsertProduct(product models.Product, tx ...*gorm.DB) (id uint, err error){
	if product.Code == "invalid" { //case for failed operation
		return id, errors.New("failed to create product")
	}
	return id, err
}
func (o *testProductDBRepo) UpdateProduct(updateData map[string]interface{}, where models.Product, tx ...*gorm.DB) (err error){
	return err
}