package dbrepo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Orololuwa/collect_am-api/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Transactions
func (m *testDBRepo) Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error {
	if err := operation(ctx, nil); err != nil {
        return err
    }

    return nil
}

// Reservations
func (m *testDBRepo) InsertReservation(ctx context.Context, tx *sql.Tx, res models.Reservation) (int, error) {
	// fail if roomId is 2
	if res.RoomID == 2 {
		return 0, errors.New("failed to insert reservation")
	}

	return 1, nil
}

// Room restrictions
func (m *testDBRepo) InsertRoomRestriction(ctx context.Context, tx *sql.Tx, r models.RoomRestriction) error {
	// fail if i try to insert a room restriction for room id of 1000
	if r.RoomID == 1000 {
		return errors.New("failed to insert room restriction")
	}

 	return nil
}

// Rooms
// SearchAvailabilityForAllRooms returns a slice of rooms for a given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(ctx context.Context, tx *sql.Tx, start, end time.Time) ([]models.Room, error){
	var rooms = make([]models.Room, 0)

	// return an error when the year in the startDate is 1960
	if start.Year() < 1960 {
		return rooms, errors.New("error searching rooms")
	}

	return rooms, nil
}

// SearchAvailabilityForDatesByRoomId returns true if availability exists for a room_id and false if no availability exists
func (m *testDBRepo) SearchAvailabilityForDatesByRoomId(ctx context.Context, tx *sql.Tx, start, end time.Time, roomId int) (bool, error){
	// simulate failure for roomId 2
	if roomId == 2 {
		return false, errors.New("reservation for room not found")
	}

	return true, nil
}

func (m *testDBRepo) GetAllRooms(ctx context.Context, tx *sql.Tx, id int, room_name string, created_at string, updated_at string) ([]models.Room, error){
	var rooms = make([]models.Room, 0)

	// simulate failure for roomId 2
	if id == 2 {
		return rooms, errors.New("error getting rooms")
	}

	return rooms, nil	
}

func (m *testDBRepo) GetRoomById(ctx context.Context, tx *sql.Tx, id int) (models.Room, error) {
	var room models.Room

	if id == 1000 {
		return room, errors.New("error getting room")
	}

	return room, nil
}

// User
func (m *testUserDBRepo) CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error){
	var newId int

	if user.FirstName == "fail" {
		return newId, errors.New("")
	}

	return newId, nil
}

func (m *testUserDBRepo) GetAUser(ctx context.Context, tx *sql.Tx, u models.User) (*models.User, error) {
	// user := &models.User{ID: 2}
	var user *models.User

	if u.Email == "johndoe@fail.com" {
		return user, errors.New("")
	}
	if u.Email == "johndoe@exists.com" {
		return &models.User{Model: gorm.Model{
			ID: 1,
		},}, nil
	}
	if u.Email == "johndoe@null.com" {
		return user, nil
	}	
	if u.Email == "test_fail@test.com" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Testpass123###"), bcrypt.DefaultCost)
		return &models.User{Password: string(hashedPassword)}, nil
	}	
	if u.Email == "test_correct@test.com" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Testpass123###"), bcrypt.DefaultCost)
		return &models.User{Password: string(hashedPassword)}, nil
	}

	if u.Phone == "+2340000000000" {
		return user, errors.New("")
	}
	if u.Phone == "+2340000000001" {
		return &models.User{Model: gorm.Model{
			ID: 1,
		},}, nil
	}
	if u.Phone == "+2340000000002" {
		return user, nil
	}

	user = &models.User{Model: gorm.Model{
        ID: 1,
    },}
	return user, nil
}

// Business
func (m *testBusinessDBRepo) CreateBusiness(ctx context.Context, tx *sql.Tx, business models.Business) (int, error){
	var id int
	return id, nil
}

func (m *testBusinessDBRepo) GetUserBusiness(ctx context.Context, tx *sql.Tx, userId int, b models.Business) (*models.Business, error) {
	var business *models.Business

	return business, nil
}

func (m *testBusinessDBRepo) UpdateBusiness(ctx context.Context, tx *sql.Tx, business models.Business) error{
	return nil
}

// kyc
func (m *testKycDBRepo) CreateKyc(ctx context.Context, tx *sql.Tx, kyc models.Kyc) (int, error){
	var id int
	return id, nil
}

func (m *testKycDBRepo) GetBusinessKyc(ctx context.Context, tx *sql.Tx, business_id int, b models.Kyc) (*models.Kyc, error){
	var kyc models.Kyc

	return &kyc, nil
}