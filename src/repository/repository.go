package repository

import (
	"database/sql"

	"github.com/Orololuwa/collect_am-api/src/models"
	"gorm.io/gorm"
)

type DBInterface interface {
	Transaction(func(tx *gorm.DB) error, ...*sql.TxOptions) error
}

type UserDBRepo interface {
	GetOneByID(id uint) (user models.User, err error)
	GetOneByEmail(email string) (user models.User, err error)
	GetOneByPhone(phone string) (user models.User, err error)
	InsertUser(user models.User, tx ...*gorm.DB) (id uint, err error)
	UpdateUser(user models.User, tx ...*gorm.DB) (err error)
}

type BusinessDBRepo interface {
	GetOneById(id uint) (business models.Business, err error)
	GetOneByUserId(userId uint) (businesses models.Business, err error)
	InsertBusiness(business models.Business, tx ...*gorm.DB) (id uint, err error)
	UpdateBusiness(updateData map[string]interface{}, where models.Business, tx ...*gorm.DB) (err error)
}

type KycDBRepo interface {
	InsertKyc(kyc models.Kyc, tx ...*gorm.DB) (id uint, err error)
	UpdateKyc(updateData map[string]interface{}, where models.Kyc, tx ...*gorm.DB) (err error)
}

type ProductDBRepo interface {
	CreateProduct(createData map[string]interface{}, where models.Product, tx ...*gorm.DB) (id uint, err error)
	InsertProduct(product models.Product, tx ...*gorm.DB) (id uint, err error)
	UpdateProduct(where FindOneBy, product models.Product, tx ...*gorm.DB) (err error)
	FindAllWithPagination(query map[string]interface{}) (products []models.Product, pagination Pagination, err error)
	FindOneById(findOneBy FindOneBy) (product models.Product, err error)
	FindOneBy(findOneBy models.Product) (product models.Product, err error)
}

type CustomerDBRepo interface {
	InsertCustomer(customer models.Customer, tx ...*gorm.DB) (id uint, err error)
	UpdateCustomer(where FindOneBy, customer models.Customer, tx ...*gorm.DB) (err error)
	FindAllWithPagination(query map[string]interface{}) (customers []models.Customer, pagination Pagination, err error)
	FindOneById(findOneBy FindOneBy) (customer models.Customer, err error)
	FindOneBy(findOneBy models.Customer) (customer models.Customer, err error)
}

type AddressDBRepo interface {
	InsertAddress(address models.Address, tx ...*gorm.DB) (id uint, err error)
	UpdateAddress(where FindOneBy, address models.Address, tx ...*gorm.DB) (err error)
	FindAllWithPagination(query map[string]interface{}) (addresses []models.Address, pagination Pagination, err error)
	FindOneById(findOneBy FindOneBy) (address models.Address, err error)
}

type InvoiceDBRepo interface {
	Insert(invoice models.Invoice, tx ...*gorm.DB) (id uint, err error)
	Update(where FindOneBy, invoice models.Invoice, tx ...*gorm.DB) (err error)
	FindAllWithPagination(query map[string]interface{}) (invoices []models.Invoice, pagination Pagination, err error)
	FindOneById(findOneBy FindOneBy) (invoice models.Invoice, err error)
	FindOneBy(findOneBy models.Invoice) (invoice models.Invoice, err error)
}

type ListedProductDBRepo interface {
	Insert(listedProduct models.ListedProduct, tx ...*gorm.DB) (id uint, err error)
	Update(where FindOneBy, listedProduct models.ListedProduct, tx ...*gorm.DB) (err error)
	FindAllWithPagination(query map[string]interface{}) (listedProducts []models.ListedProduct, pagination Pagination, err error)
	FindOneById(findOneBy FindOneBy) (listedProduct models.ListedProduct, err error)
	BatchInsert(listedProducts []models.ListedProduct, tx ...*gorm.DB) (ids []uint, err error)
}
