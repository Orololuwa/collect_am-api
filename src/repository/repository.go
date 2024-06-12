package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Orololuwa/collect_am-api/src/models"
	"gorm.io/gorm"
)

type DatabaseRepo interface {
	Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error 
	InsertReservation(ctx context.Context, tx *sql.Tx, res models.Reservation) (int, error)
	InsertRoomRestriction(ctx context.Context, tx *sql.Tx, r models.RoomRestriction) error
	SearchAvailabilityForDatesByRoomId(ctx context.Context, tx *sql.Tx, start, end time.Time, roomId int) (bool, error)
	SearchAvailabilityForAllRooms(ctx context.Context, tx *sql.Tx, start, end time.Time) ([]models.Room, error)
	GetRoomById(ctx context.Context, tx *sql.Tx, id int) (models.Room, error)
	GetAllRooms(ctx context.Context, tx *sql.Tx, id int, room_name string, created_at string, updated_at string)([]models.Room, error)
}

type UserDBRepo interface {
	CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error)
	GetAUser(ctx context.Context, tx *sql.Tx, u models.User) (*models.User, error)

	GetOneByID(id uint) (user models.User, err error)
	GetOneByEmail(email string) (user models.User, err error)
	GetOneByPhone(phone string) (user models.User, err error)
	InsertUser(user models.User, tx ...*gorm.DB) (id uint, err error)
	UpdateUser(user models.User, tx ...*gorm.DB) (err error)
}

type BusinessDBRepo interface {
	CreateBusiness(ctx context.Context, tx *sql.Tx, business models.Business) (int, error)
	GetUserBusiness(ctx context.Context, tx *sql.Tx, userId int, b models.Business) (*models.Business, error)	
	UpdateBusinessOld(ctx context.Context, tx *sql.Tx, business models.Business) (error)


	GetOneByUserId(userId uint) (businesses models.Business, err error)
	InsertBusiness(business models.Business, tx ...*gorm.DB) (id uint, err error)
	UpdateBusiness(updateData map[string]interface{}, where models.Business, tx ...*gorm.DB) (err error)
}

type KycDBRepo interface {
	CreateKyc(ctx context.Context, tx *sql.Tx, kyc models.Kyc) (int, error)
	GetBusinessKyc(ctx context.Context, tx *sql.Tx, business_id int, b models.Kyc) (*models.Kyc, error)

	InsertKyc(kyc models.Kyc, tx ...*gorm.DB) (id uint, err error)
	UpdateKyc(updateData map[string]interface{}, where models.Kyc, tx ...*gorm.DB) (err error)
}