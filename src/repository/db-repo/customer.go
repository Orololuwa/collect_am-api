package dbrepo

import (
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

func (o *customerOrm) UpdateCustomer(where repository.FindOneBy, customer models.Customer, tx ...*gorm.DB) (err error) {
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

func (p *customerOrm) FindAllWithPagination(query map[string]interface{}) (customers []models.Customer, pagination repository.Pagination, err error) {
	page := 1
	pageSize := 10

	if pageVal, exists := query["page"].(int); exists && pageVal > 0 {
		page = pageVal
	}
	if pageSizeVal, exists := query["pageSize"].(int); exists && pageSizeVal > 0 {
		pageSize = pageSizeVal
	}

	delete(query, "page")
	delete(query, "pageSize")

	offset := (page - 1) * pageSize

	var total int64
	countResult := p.db.Model(&models.Customer{}).Where(query).Count(&total)
	if countResult.Error != nil {
		return nil, pagination, countResult.Error
	}

	result := p.db.
		Preload("Address").
		Model(&models.Customer{}).
		Where(query).
		Offset(offset).
		Limit(pageSize).
		Find(&customers)
	if result.Error != nil {
		return nil, pagination, result.Error
	}

	lastPage := int((total + int64(pageSize) - 1) / int64(pageSize)) // Calculate the last page number

	pagination = repository.Pagination{
		HasPrev:  page > 1,
		PrevPage: page - 1,
		HasNext:  page < lastPage,
		NextPage: page + 1,
		CurrPage: page,
		PageSize: pageSize,
		LastPage: lastPage,
		Total:    int(total),
	}

	return customers, pagination, nil
}

func (p *customerOrm) FindOneById(findOneBy repository.FindOneBy) (customer models.Customer, err error) {
	result := p.db.Preload("Address").Where(&findOneBy).First(&customer)
	return customer, result.Error
}

func (p *customerOrm) FindOneBy(findOneBy models.Customer) (customer models.Customer, err error) {
	result := p.db.Where(&findOneBy).First(&customer)
	return customer, result.Error
}
