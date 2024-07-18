package dbrepo

import (
	"errors"

	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
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
	if user.FirstName == "fail" { //failed insert operation
		return id, errors.New("failed to insert")
	}
	return id, nil
}

func (o *testUserDBRepo) UpdateUser(user models.User, tx ...*gorm.DB) (err error) {
	return nil
}

// Business
func (m *testBusinessDBRepo) GetOneById(id uint) (business models.Business, err error) {
	if id == 0 {
		return business, errors.New("error getting business")
	}
	if id == 1 {
		return business, errors.New("record not found")
	}
	return business, nil
}

func (m *testBusinessDBRepo) GetOneByUserId(userId uint) (businesses models.Business, err error) {
	return businesses, nil
}

func (o *testBusinessDBRepo) InsertBusiness(business models.Business, tx ...*gorm.DB) (id uint, err error) {
	if business.Email == "invalid" {
		return id, errors.New("failed to insert business") //fail case
	}

	return id, nil //success case
}

func (o *testBusinessDBRepo) UpdateBusiness(updateData map[string]interface{}, where models.Business, tx ...*gorm.DB) (err error) {
	if updateData["name"] == "invalid" {
		return errors.New("failed to insert business") //fail case
	}

	return nil //success case
}

// kyc
func (o *testKycDBRepo) GetOneByID(id uint) (kyc models.Kyc, err error) {
	return kyc, nil
}

func (o *testKycDBRepo) InsertKyc(kyc models.Kyc, tx ...*gorm.DB) (id uint, err error) {
	if kyc.BVN == "invalid" {
		return id, errors.New("failed to insert kyc")
	}
	return id, nil
}

func (o *testKycDBRepo) UpdateKyc(updateData map[string]interface{}, where models.Kyc, tx ...*gorm.DB) (err error) {
	if updateData["bvn"] == "invalid" {
		return errors.New("failed to insert kyc") //fail case
	}

	return nil
}

// products
func (o *testProductDBRepo) CreateProduct(createData map[string]interface{}, where models.Product, tx ...*gorm.DB) (id uint, err error) {
	return id, err
}
func (o *testProductDBRepo) InsertProduct(product models.Product, tx ...*gorm.DB) (id uint, err error) {
	if product.Code == "invalid" { //case for failed operation
		return id, errors.New("failed to create product")
	}
	return id, err
}
func (o *testProductDBRepo) UpdateProduct(where repository.FindOneBy, product models.Product, tx ...*gorm.DB) (err error) {
	if product.Category == "invalid" { //case for failed operation
		return errors.New("failed to create product")
	}
	return err
}
func (p *testProductDBRepo) FindAllWithPagination(query map[string]interface{}) (products []models.Product, pagination repository.Pagination, err error) {
	if page, exists := query["page"]; exists && page == 1 { //case for failed operation
		return products, pagination, errors.New("failed to get all product")
	}
	return products, pagination, err
}
func (o *testProductDBRepo) FindOneById(findOneBy repository.FindOneBy) (product models.Product, err error) {
	product.Price = 5000.00

	if findOneBy.ID == 1 {
		return product, errors.New("failed to get product")
	}
	return product, err
}
func (p *testProductDBRepo) FindOneBy(findOneBy models.Product) (product models.Product, err error) {
	product.Price = 5000.00

	if findOneBy.Code == "exists" { //email exists
		return models.Product{ID: 1}, nil
	}
	return product, err
}

// Customers
func (o *testCustomerDBRepo) InsertCustomer(customer models.Customer, tx ...*gorm.DB) (id uint, err error) {
	if customer.Email == "invalid" {
		return id, errors.New("failed to insert customer")
	}
	return id, err
}
func (o *testCustomerDBRepo) UpdateCustomer(where repository.FindOneBy, customer models.Customer, tx ...*gorm.DB) (err error) {
	if where.ID == 1 {
		return errors.New("failed to update customer")
	}
	return err
}
func (p *testCustomerDBRepo) FindAllWithPagination(query map[string]interface{}) (customers []models.Customer, pagination repository.Pagination, err error) {
	if page, exists := query["page"]; exists && page == 1 { //case for failed operation
		return customers, pagination, errors.New("failed to get all customer")
	}
	return customers, pagination, err
}
func (o *testCustomerDBRepo) FindOneById(findOneBy repository.FindOneBy) (customer models.Customer, err error) {
	if findOneBy.ID == 1 {
		return customer, errors.New("failed to get customer")
	}
	return customer, err
}
func (p *testCustomerDBRepo) FindOneBy(findOneBy models.Customer) (customer models.Customer, err error) {
	return customer, err
}

// Address
func (o *testAddressDBRepo) InsertAddress(address models.Address, tx ...*gorm.DB) (id uint, err error) {
	if address.UnitNumber == "invalid" {
		return id, errors.New("failed to insert address")
	}
	return id, err
}
func (o *testAddressDBRepo) UpdateAddress(where repository.FindOneBy, address models.Address, tx ...*gorm.DB) (err error) {
	return err
}
func (p *testAddressDBRepo) FindAllWithPagination(query map[string]interface{}) (addresses []models.Address, pagination repository.Pagination, err error) {
	if page, exists := query["page"]; exists && page == 1 { //case for failed operation
		return addresses, pagination, errors.New("failed to get all address")
	}
	return addresses, pagination, err
}
func (o *testAddressDBRepo) FindOneById(findOneBy repository.FindOneBy) (address models.Address, err error) {
	if findOneBy.ID == 1 {
		return address, errors.New("failed to get address")
	}
	return address, err
}

// invoices
func (o *testInvoiceDBRepo) Insert(invoice models.Invoice, tx ...*gorm.DB) (id uint, err error) {
	if invoice.Code == "invalid" { //case for failed operation
		return id, errors.New("failed to create invoice")
	}
	return id, err
}
func (o *testInvoiceDBRepo) Update(where repository.FindOneBy, invoice models.Invoice, tx ...*gorm.DB) (err error) {
	return err
}
func (p *testInvoiceDBRepo) FindAllWithPagination(query map[string]interface{}) (invoices []models.Invoice, pagination repository.Pagination, err error) {
	if page, exists := query["page"]; exists && page == 1 { //case for failed operation
		return invoices, pagination, errors.New("failed to get all invoice")
	}
	return invoices, pagination, err
}
func (o *testInvoiceDBRepo) FindOneById(findOneBy repository.FindOneBy) (invoice models.Invoice, err error) {
	if findOneBy.ID == 1 {
		return invoice, errors.New("failed to get invoice")
	}
	return invoice, err
}
func (p *testInvoiceDBRepo) FindOneBy(findOneBy models.Invoice) (invoice models.Invoice, err error) {
	if findOneBy.Code == "exists" { //email exists
		return models.Invoice{ID: 1}, nil
	}
	return invoice, err
}

// listedProducts
func (o *testListedProductDBRepo) Create(createData map[string]interface{}, where models.ListedProduct, tx ...*gorm.DB) (id uint, err error) {
	return id, err
}
func (o *testListedProductDBRepo) Insert(listedProduct models.ListedProduct, tx ...*gorm.DB) (id uint, err error) {
	return id, err
}
func (o *testListedProductDBRepo) Update(where repository.FindOneBy, listedProduct models.ListedProduct, tx ...*gorm.DB) (err error) {
	return err
}
func (p *testListedProductDBRepo) FindAllWithPagination(query map[string]interface{}) (listedProducts []models.ListedProduct, pagination repository.Pagination, err error) {
	if page, exists := query["page"]; exists && page == 1 { //case for failed operation
		return listedProducts, pagination, errors.New("failed to get all listedProduct")
	}
	return listedProducts, pagination, err
}
func (o *testListedProductDBRepo) FindOneById(findOneBy repository.FindOneBy) (listedProduct models.ListedProduct, err error) {
	if findOneBy.ID == 1 {
		return listedProduct, errors.New("failed to get listedProduct")
	}
	return listedProduct, err
}
func (o *testListedProductDBRepo) BatchInsert(listedProducts []models.ListedProduct, tx ...*gorm.DB) (ids []uint, err error) {
	if len(listedProducts) == 0 {
		return ids, errors.New("failed to batch insert listedProducts")
	}
	return ids, nil
}
