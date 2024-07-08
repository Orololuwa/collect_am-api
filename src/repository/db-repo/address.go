package dbrepo

import (
	"log"

	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	"gorm.io/gorm"
)

type addressOrm struct {
	db *gorm.DB
}

func NewAddressDBRepo(db *driver.DB) repository.AddressDBRepo {
	return &addressOrm{
		db: db.Gorm,
	}
}

type testAddressDBRepo struct{}

func NewAddressTestingDBRepo() repository.AddressDBRepo {
	return &testAddressDBRepo{}
}

func (o *addressOrm) InsertAddress(address models.Address, tx ...*gorm.DB) (id uint, err error) {
	db := o.db
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}

	result := db.Model(&models.Address{}).Create(&address)
	return address.ID, result.Error
}

func (o *addressOrm) UpdateAddress(where repository.FindOneBy, address models.Address, tx ...*gorm.DB) (err error) {
	db := o.db
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}

	result := db.
		Model(&models.Address{}).
		Where(&where).
		Model(&address).
		Updates(&address)

	return result.Error
}

func (p *addressOrm) FindAllWithPagination(query map[string]interface{}) (addresses []models.Address, pagination repository.Pagination, err error) {
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
	countResult := p.db.Model(&models.Address{}).Where(query).Count(&total)
	if countResult.Error != nil {
		return nil, pagination, countResult.Error
	}

	result := p.db.
		Model(&models.Address{}).
		Where(query).
		Offset(offset).
		Limit(pageSize).
		Find(&addresses)
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

	return addresses, pagination, nil
}

func (p *addressOrm) FindOneById(findOneBy repository.FindOneBy) (address models.Address, err error) {
	log.Println(findOneBy)
	result := p.db.First(&address, findOneBy.ID)
	return address, result.Error
}
