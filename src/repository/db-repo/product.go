package dbrepo

import (
	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	"gorm.io/gorm"
)

type productOrm struct {
	db *gorm.DB
}

func NewProductDBRepo(db *driver.DB) repository.ProductDBRepo {
	return &productOrm{
		db: db.Gorm,
	}
}

type testProductDBRepo struct {
}

func NewProductTestingDBRepo() repository.ProductDBRepo {
	return &testProductDBRepo{}
}

func (o *productOrm) InsertProduct(product models.Product, tx ...*gorm.DB) (id uint, err error) {
	db := o.db
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}

	result := db.Model(&models.Product{}).Create(&product)
	return product.ID, result.Error
}

func (o *productOrm) CreateProduct(createData map[string]interface{}, where models.Product, tx ...*gorm.DB) (id uint, err error) {
	db := o.db
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}

	product := &models.Product{}
	result := db.Model(&models.Product{}).Create(createData).Scan(product)

	if result.Error != nil {
		return 0, result.Error
	}

	return product.ID, result.Error
}

func (o *productOrm) UpdateProduct(where repository.FindOneBy, product models.Product, tx ...*gorm.DB) (err error) {
	db := o.db
	if len(tx) > 0 && tx[0] != nil {
		db = tx[0]
	}

	result := db.
		Model(&models.Product{}).
		Where(&where).
		Model(&product).
		Updates(&product)

	return result.Error
}

func (p *productOrm) FindAllWithPagination(query map[string]interface{}) (products []models.Product, pagination repository.Pagination, err error) {
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
	countResult := p.db.Model(&models.Product{}).Where(query).Count(&total)
	if countResult.Error != nil {
		return nil, pagination, countResult.Error
	}

	result := p.db.
		Model(&models.Product{}).
		Where(query).
		Offset(offset).
		Limit(pageSize).
		Find(&products)
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

	return products, pagination, nil
}

func (p *productOrm) FindOneById(findOneBy repository.FindOneBy) (product models.Product, err error) {
	result := p.db.Where(&findOneBy).First(&product)
	return product, result.Error
}
