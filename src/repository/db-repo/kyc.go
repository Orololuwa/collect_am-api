package dbrepo

import (
	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	"gorm.io/gorm"
)

type kycOrm struct {
	db *gorm.DB
}
func NewKycDBRepo(db *driver.DB) repository.KycDBRepo {
	return &kycOrm{
		db: db.Gorm,
	}
}

type testKycDBRepo struct {
}
func NewKycTestingDBRepo() repository.KycDBRepo {
	return &testKycDBRepo{
	}
}

func (o *kycOrm) InsertKyc(kyc models.Kyc, tx ...*gorm.DB) (id uint, err error) {
	db := o.db
    if len(tx) > 0 && tx[0] != nil {
        db = tx[0]
    }

	result := db.Model(&models.Kyc{}).Create(&kyc)
	return kyc.ID, result.Error
}

func (o *kycOrm) UpdateKyc(updateData map[string]interface{}, where models.Kyc, tx ...*gorm.DB) (err error) {
	db := o.db
    if len(tx) > 0 && tx[0] != nil {
        db = tx[0]
    }

	result := db.
			Model(&models.Kyc{}).
			Where(&where).
			Updates(updateData)
	return result.Error
}