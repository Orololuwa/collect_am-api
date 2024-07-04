package dbrepo

import (
	"log"

	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	"gorm.io/gorm"
)

type customerOrm struct {
	db *gorm.DB
}

func NewCustomerDBRepo(db *driver.DB) repository.CustomerDBRepo {
	return &customerOrm{
		db: db.Gorm,
	}
}

type testCustomerDBRepo struct {
}

func NewCustomerTestingDBRepo() repository.CustomerDBRepo {
	return &testCustomerDBRepo{}
}

func (o *customerOrm) InsertCustomer(customer models.Customer, tx ...*gorm.DB) (id uint, err error) {
	db := o.db
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}

	result := db.Model(&models.Customer{}).Create(&customer)
	return customer.ID, result.Error
}

func (o *customerOrm) UpdateCustomer(where models.Customer, customer models.Customer, tx ...*gorm.DB) (err error) {
	db := o.db
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}

	result := db.
		Model(&models.Customer{}).
		Where(&where).
		Model(&customer).
		Updates(&customer)

	return result.Error
}

func (p *customerOrm) FindAllWithPagination(query repository.FilterQueryPagination) (customers []models.Customer, pagination repository.Pagination, err error) {
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 10
	}

	offset := (query.Page - 1) * query.PageSize

	var total int64
	countResult := p.db.Model(&models.Customer{BusinessID: query.BusinessId}).Count(&total)
	if countResult.Error != nil {
		return nil, pagination, countResult.Error
	}

	result := p.db.
		Model(&models.Customer{BusinessID: query.BusinessId}).
		Offset(offset).
		Limit(query.PageSize).
		Find(&customers)
	if result.Error != nil {
		return nil, pagination, result.Error
	}

	lastPage := int((total + int64(query.PageSize) - 1) / int64(query.PageSize)) // Calculate the last page number

	pagination = repository.Pagination{
		HasPrev:  query.Page > 1,
		PrevPage: query.Page - 1,
		HasNext:  query.Page < lastPage,
		NextPage: query.Page + 1,
		CurrPage: query.Page,
		PageSize: query.PageSize,
		LastPage: lastPage,
		Total:    int(total),
	}

	return customers, pagination, nil
}

func (p *customerOrm) FindOneById(findOneBy repository.FindOneBy) (customer models.Customer, err error) {
	log.Println(findOneBy)
	result := p.db.First(&customer, findOneBy.ID)
	return customer, result.Error
}
