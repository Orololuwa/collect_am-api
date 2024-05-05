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

type business struct {
	DB *sql.DB
}
func NewBusinessDBRepo(conn *sql.DB) repository.BusinessDBRepo {
	return &business{
		DB: conn,
	}
}

type testBusinessDBRepo struct {
	DB *sql.DB
}
func NewBusinessTestingDBRepo() repository.BusinessDBRepo {
	return &testBusinessDBRepo{
	}
}

func (m *business) CreateBusiness(ctx context.Context, tx *sql.Tx, business models.Business) (int, error){
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var newId int
	var err error

	queryFields := ""
	queryPlaceholders := ""
	var args []interface{}

	businessType := reflect.TypeOf(business)
	businessValue := reflect.ValueOf(business)

	for i := 0; i < businessType.NumField(); i++ {
		field := businessType.Field(i)
		value := businessValue.Field(i)
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
		INSERT INTO businesses
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