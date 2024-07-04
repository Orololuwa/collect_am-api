package repository

import (
	"database/sql"
	"net/url"
	"strconv"

	"github.com/Orololuwa/collect_am-api/src/models"
	"gorm.io/gorm"
)

type FilterQueryPagination struct {
	PageSize   int  `json:"pageSize"`
	Page       int  `json:"page"`
	BusinessId uint `json:"businessId,omitempty"`
}

func CleanQueryParams(queryParams url.Values) (filter FilterQueryPagination) {
	filter.Page = 1
	filter.PageSize = 10

	if pageStr := queryParams.Get("page"); pageStr != "" {
		filter.Page, _ = strconv.Atoi(pageStr)
	}

	if pageSizeStr := queryParams.Get("pageSize"); pageSizeStr != "" {
		filter.PageSize, _ = strconv.Atoi(pageSizeStr)
	}

	if businessIdStr := queryParams.Get("businessId"); businessIdStr != "" {
		businessId, _ := strconv.ParseUint(businessIdStr, 10, 32)
		filter.BusinessId = uint(businessId)
	}

	return filter
}

type Pagination struct {
	HasPrev  bool `json:"hasPrev"`
	PrevPage int  `json:"prevPage"`
	HasNext  bool `json:"hasNext"`
	NextPage int  `json:"nextPage"`
	CurrPage int  `json:"currPage"`
	PageSize int  `json:"pageSize"`
	LastPage int  `json:"lastPage"`
	Total    int  `json:"total"`
}

type FindOneBy struct {
	ID         uint `json:"id"`
	BusinessId uint `json:"businessId,omitempty"`
}

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
	UpdateProduct(where models.Product, product models.Product, tx ...*gorm.DB) (err error)
	FindAllWithPagination(query FilterQueryPagination) (products []models.Product, pagination Pagination, err error)
	FindOneById(findOneBy FindOneBy) (product models.Product, err error)
}

type CustomerDBRepo interface {
	InsertCustomer(customer models.Customer, tx ...*gorm.DB) (id uint, err error)
	UpdateCustomer(where models.Customer, customer models.Customer, tx ...*gorm.DB) (err error)
	FindAllWithPagination(query FilterQueryPagination) (customers []models.Customer, pagination Pagination, err error)
	FindOneById(findOneBy FindOneBy) (customer models.Customer, err error)
}

type AddressDBRepo interface {
	InsertAddress(address models.Address, tx ...*gorm.DB) (id uint, err error)
	UpdateAddress(where models.Address, address models.Address, tx ...*gorm.DB) (err error)
	FindAllWithPagination(query FilterQueryPagination) (addresses []models.Address, pagination Pagination, err error)
	FindOneById(findOneBy FindOneBy) (address models.Address, err error)
}
