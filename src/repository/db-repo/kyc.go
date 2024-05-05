package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
)

type kyc struct {
	DB *sql.DB
}
func NewKycDBRepo(conn *sql.DB) repository.KycDBRepo {
	return &kyc{
		DB: conn,
	}
}

type testKycDBRepo struct {
	DB *sql.DB
}
func NewKycTestingDBRepo() repository.KycDBRepo {
	return &testKycDBRepo{
	}
}

func (m *kyc) CreateKyc(ctx context.Context, tx *sql.Tx, kyc models.KYC) (int, error){
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var newId int
	var err error

	queryFields := ""
	queryPlaceholders := ""
	var args []interface{}

	kycType := reflect.TypeOf(kyc)
	kycValue := reflect.ValueOf(kyc)

	for i := 0; i < kycType.NumField(); i++ {
		field := kycType.Field(i)
		value := kycValue.Field(i)
		tagValue := field.Tag.Get("db")

        if value.IsZero() || tagValue == "" {
            continue
        }

		if queryFields == "" {
			queryFields += tagValue
		}else{
			queryFields += ", " + tagValue
		}

		if queryPlaceholders == "" {
			queryPlaceholders += "$" + strconv.Itoa(len(args) + 1)
		}else{
			queryPlaceholders += ", $" + strconv.Itoa(len(args) + 1)
		}

		args = append(args, value.Interface())
	}


	query := fmt.Sprintf(`
		INSERT INTO kyc
			(%s)
		VALUES
			(%s)
		RETURNING id;
	`, queryFields, queryPlaceholders)

	if tx != nil {
		err = m.DB.QueryRowContext(ctx, query, args...).Scan(&newId)
	}else{
		err = tx.QueryRowContext(ctx, query, args...).Scan(&newId)
	}

	if err != nil {
		return 0, err
	}

	return newId, nil
}