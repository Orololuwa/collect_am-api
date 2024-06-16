package driver

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	SQL *sql.DB
	Gorm *gorm.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

func ConnectSQL (dsn string)(*DB, error){
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)


	gormDB, err := NewGormDBConnection(dsn)
	if err != nil {
		panic(err)
	}

	dbConn.SQL = d
	dbConn.Gorm = gormDB

	err = testDb(d)
	if err != nil {
		return nil, err
	}

	return dbConn, err
}

func testDb(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}

	return nil
}

func NewDatabase (dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}


	return db, nil
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


func CreateTestDBInstance() *sql.DB {
    db, err := sql.Open("pgx", ":memory:")
    if err != nil {
        log.Fatalf("Error opening database connection: %v", err)
    }
	
    return db
}