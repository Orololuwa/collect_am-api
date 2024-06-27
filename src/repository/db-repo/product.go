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
