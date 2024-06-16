package driver

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Gorm *gorm.DB
}

var dbConn = &DB{}

func ConnectSQL (dsn string)(*DB, error){
	gormDB, err := NewGormDBConnection(dsn)
	if err != nil {
		panic(err)
	}

	dbConn.Gorm = gormDB

	return dbConn, err
}

func NewGormDBConnection (dsn string)(*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        // Logger: logger.Default.LogMode(logger.Info),
		// DryRun: true,
	})


	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
