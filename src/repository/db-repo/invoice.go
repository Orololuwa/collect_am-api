package dbrepo

import (
	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	"gorm.io/gorm"
)

type invoiceOrm struct {
	db *gorm.DB
}

func NewInvoiceDBRepo(db *driver.DB) repository.InvoiceDBRepo {
	return &invoiceOrm{
		db: db.Gorm,
	}
}

type testInvoiceDBRepo struct {
}

func NewInvoiceTestingDBRepo() repository.InvoiceDBRepo {
	return &testInvoiceDBRepo{}
}

func (o *invoiceOrm) Insert(invoice models.Invoice, tx ...*gorm.DB) (id uint, err error) {
	db := o.db
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}

	result := db.Model(&models.Invoice{}).Create(&invoice)
	return invoice.ID, result.Error
}

func (o *invoiceOrm) Update(where repository.FindOneBy, invoice models.Invoice, tx ...*gorm.DB) (err error) {
	db := o.db
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}

	result := db.
		Model(&models.Invoice{}).
		Where(&where).
		Model(&invoice).
		Updates(&invoice)

	return result.Error
}

func (p *invoiceOrm) FindAllWithPagination(query map[string]interface{}) (invoices []models.Invoice, pagination repository.Pagination, err error) {
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
	countResult := p.db.Model(&models.Invoice{}).Where(query).Count(&total)
	if countResult.Error != nil {
		return nil, pagination, countResult.Error
	}

	result := p.db.
		Model(&models.Invoice{}).
		Where(query).
		Offset(offset).
		Limit(pageSize).
		Find(&invoices)
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

	return invoices, pagination, nil
}

func (p *invoiceOrm) FindOneById(findOneBy repository.FindOneBy) (invoice models.Invoice, err error) {
	result := p.db.Where(&findOneBy).First(&invoice)
	return invoice, result.Error
}

func (p *invoiceOrm) FindOneBy(findOneBy models.Invoice) (invoice models.Invoice, err error) {
	result := p.db.Where(&findOneBy).First(&invoice)
	return invoice, result.Error
}
