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

func (m *kyc) GetBusinessKyc(ctx context.Context, tx *sql.Tx, business_id int, b models.KYC) (*models.KYC, error) {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

	var kyc models.KYC

    // Prepare the base query
    query := `
        SELECT 
			id,
			certificate_of_registration,
			proof_of_address,
			bvn,
			created_at,
			updated_at
        FROM 
			kyc
        WHERE
			business_id = $1
    `

    var args []interface{}
	args = append(args, business_id)

    userType := reflect.TypeOf(b)
    userValue := reflect.ValueOf(b)

    for i := 0; i < userType.NumField(); i++ {
        field := userType.Field(i)
        value := userValue.Field(i)
		tagValue := field.Tag.Get("db")

        if value.IsZero() || tagValue == "" {
            continue
        }

        switch value.Interface().(type) {
        case int, int64:
            query += " AND " + tagValue + " = $" + strconv.Itoa(len(args)+1)
            args = append(args, value.Interface())
        case string:
            query += " AND " + tagValue + " = $" + strconv.Itoa(len(args)+1)
            args = append(args, value.Interface())
        // Add more cases as needed for other types
        }
    }

    // Execute the query
    var err error
    if tx != nil {
        err = tx.QueryRowContext(ctx, query, args...).Scan(
			&kyc.ID,
			&kyc.CertificateOfRegistration,
			&kyc.ProofOfAddress,
			&kyc.BVN,
            &kyc.CreatedAt,
            &kyc.UpdatedAt,
        )
    } else {
        err = m.DB.QueryRowContext(ctx, query, args...).Scan(
			&kyc.ID,
			&kyc.CertificateOfRegistration,
			&kyc.ProofOfAddress,
			&kyc.BVN,
            &kyc.CreatedAt,
            &kyc.UpdatedAt,
        )
    }

    if err != nil {
        return &kyc, err
    }

    return &kyc, nil
}