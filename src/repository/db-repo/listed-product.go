package dbrepo

import (
	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	"gorm.io/gorm"
)

type listedProductOrm struct {
	db *gorm.DB
}

func NewListedProductDBRepo(db *driver.DB) repository.ListedProductDBRepo {
	return &listedProductOrm{
		db: db.Gorm,
	}
}

type testListedProductDBRepo struct {
}

func NewListedProductTestingDBRepo() repository.ListedProductDBRepo {
	return &testListedProductDBRepo{}
}

func (o *listedProductOrm) Insert(listedProduct models.ListedProduct, tx ...*gorm.DB) (id uint, err error) {
	db := o.db
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}

	result := db.Model(&models.ListedProduct{}).Create(&listedProduct)
	return listedProduct.ID, result.Error
}

func (o *listedProductOrm) Update(where repository.FindOneBy, listedProduct models.ListedProduct, tx ...*gorm.DB) (err error) {
	db := o.db
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}

	result := db.
		Model(&models.ListedProduct{}).
		Where(&where).
		Model(&listedProduct).
		Updates(&listedProduct)

	return result.Error
}

func (p *listedProductOrm) FindAllWithPagination(query map[string]interface{}) (listedProducts []models.ListedProduct, pagination repository.Pagination, err error) {
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
	countResult := p.db.Model(&models.ListedProduct{}).Where(query).Count(&total)
	if countResult.Error != nil {
		return nil, pagination, countResult.Error
	}

	result := p.db.
		Model(&models.ListedProduct{}).
		Where(query).
		Offset(offset).
		Limit(pageSize).
		Find(&listedProducts)
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

	return listedProducts, pagination, nil
}

func (p *listedProductOrm) FindOneById(findOneBy repository.FindOneBy) (listedProduct models.ListedProduct, err error) {
	result := p.db.Where(&findOneBy).First(&listedProduct)
	return listedProduct, result.Error
}

func (p *listedProductOrm) BatchInsert(listedProducts []models.ListedProduct, tx ...*gorm.DB) (ids []uint, err error) {
	db := p.db
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}

	result := db.Model(&models.ListedProduct{}).Create(&listedProducts)
	if result.Error != nil {
		return nil, result.Error
	}

	// Extract the IDs of the inserted records
	for _, lp := range listedProducts {
		ids = append(ids, lp.ID)
	}

	return ids, nil
}
