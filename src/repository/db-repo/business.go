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

func (m *business) GetUserBusiness(ctx context.Context, tx *sql.Tx, userId int, b models.Business) (*models.Business, error) {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    var business models.Business
	var kyc models.KYC

    // Prepare the base query
    query := `
        SELECT 
			b.id, 
			b.name, 
			b.email, 
			b.description, 
			b.sector, 
			b.is_corporate_affairs, 
			b.is_setup_complete, 
			b.logo, 
			b.created_at, 
			b.updated_at,
			k.id AS kyc_id,
			k.certificate_of_registration,
			k.proof_of_address,
			k.bvn,
			k.created_at AS kyc_created_at,
			k.updated_at AS kyc_updated_at
        FROM 
			businesses AS b
		LEFT JOIN
			kyc AS k ON b.id = k.business_id
        WHERE
			user_id = $1
    `

    var args []interface{}
	args = append(args, userId)

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
            &business.ID,
            &business.Name,
            &business.Email,
            &business.Description,
            &business.Sector,
			&business.IsCorporateAffair,
			&business.IsSetupComplete,
			&business.Logo,
            &business.CreatedAt,
            &business.UpdatedAt,
			&kyc.ID,
			&kyc.CertificateOfRegistration,
			&kyc.ProofOfAddress,
			&kyc.BVN,
            &kyc.CreatedAt,
            &kyc.UpdatedAt,
        )
    } else {
        err = m.DB.QueryRowContext(ctx, query, args...).Scan(
            &business.ID,
            &business.Name,
            &business.Email,
            &business.Description,
            &business.Sector,
			&business.IsCorporateAffair,
			&business.IsSetupComplete,
			&business.Logo,
            &business.CreatedAt,
            &business.UpdatedAt,
			&kyc.ID,
			&kyc.CertificateOfRegistration,
			&kyc.ProofOfAddress,
			&kyc.BVN,
            &kyc.CreatedAt,
            &kyc.UpdatedAt,
        )
    }

    if err == sql.ErrNoRows {
        return nil, nil // No rows found, return nil
    }

    if err != nil {
        return &business, err
    }

	business.KYC = kyc
    return &business, nil
}