package mocks

import (
	"database/sql"

	"gorm.io/gorm"
)

// MockDB is a struct that embeds gorm.DB and overrides the Transaction method.
type MockDB struct {
	*gorm.DB
	TransactionFunc func(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
}

func NewMockDB () *MockDB{
	mockDB := &MockDB{}
	mockDB.TransactionFunc = func(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
		tx := &gorm.DB{} // Mock the transaction DB
		return fc(tx)
	}

	return mockDB
}

// Transaction overrides the gorm.DB Transaction method.
func (db *MockDB) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	if db.TransactionFunc != nil {
		return db.TransactionFunc(fc, opts...)
	}
	return db.DB.Transaction(fc, opts...)
}
