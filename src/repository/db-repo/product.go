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

func (p *productOrm) GetOneById(id uint) (product models.Product, err error) {
	result := p.db.Model(&models.Product{}).Where("id = ?", id).First(&product)
	return product, result.Error
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

func (o *productOrm) UpdateProduct(where models.Product, product models.Product, tx ...*gorm.DB) (err error) {
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

func (p *productOrm) FindAllWithPagination(query repository.FilterQueryPagination) (products []models.Product, pagination repository.Pagination, err error) {
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 10
	}

	offset := (query.Page - 1) * query.PageSize

	var total int64
	countResult := p.db.Model(&models.Product{BusinessID: query.BusinessId}).Count(&total)
	if countResult.Error != nil {
		return nil, pagination, countResult.Error
	}

	result := p.db.
		Model(&models.Product{BusinessID: query.BusinessId}).
		Offset(offset).
		Limit(query.PageSize).
		Find(&products)
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

	return products, pagination, nil
}
