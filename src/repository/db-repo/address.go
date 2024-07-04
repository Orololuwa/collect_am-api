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

func (o *addressOrm) UpdateAddress(where models.Address, address models.Address, tx ...*gorm.DB) (err error) {
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

func (p *addressOrm) FindAllWithPagination(query repository.FilterQueryPagination) (addresss []models.Address, pagination repository.Pagination, err error) {
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 10
	}

	offset := (query.Page - 1) * query.PageSize

	var total int64
	countResult := p.db.Model(&models.Address{}).Count(&total)
	if countResult.Error != nil {
		return nil, pagination, countResult.Error
	}

	result := p.db.
		Model(&models.Address{}).
		Offset(offset).
		Limit(query.PageSize).
		Find(&addresss)
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

	return addresss, pagination, nil
}

func (p *addressOrm) FindOneById(findOneBy repository.FindOneBy) (address models.Address, err error) {
	log.Println(findOneBy)
	result := p.db.First(&address, findOneBy.ID)
	return address, result.Error
}
