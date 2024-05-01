package dbrepo

import (
	"context"
	"database/sql"
	"reflect"
	"strconv"
	"time"

	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
)

type user struct {
	DB *sql.DB
}
func NewUserDBRepo(conn *sql.DB) repository.UserDBRepo {
	return &user{
		DB: conn,
	}
}

type testUserDBRepo struct {
	DB *sql.DB
}
func NewUserTestingDBRepo() repository.UserDBRepo {
	return &testUserDBRepo{
	}
}

func (m *user) CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error){
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var newId int

	query := `
			INSERT into users 
				(first_name, last_name, email, phone, password)
			values 
				($1, $2, $3, $4, $5)
			returning id`

	var err error;
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, 
			user.FirstName, 
			user.LastName, 
			user.Email, 
			user.Phone,
			user.Password,
		).Scan(&newId)
	}else{
		err = m.DB.QueryRowContext(ctx, query, 
			user.FirstName, 
			user.LastName, 
			user.Email, 
			user.Phone,
			user.Password,
		).Scan(&newId)
	}

	if err != nil {
		return 0, err
	}

	return newId, nil
}

func (m *user) GetAUser(ctx context.Context, tx *sql.Tx, u models.User) (*models.User, error) {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    var user models.User

    // Prepare the base query
    query := `
        SELECT id, first_name, last_name, email, phone, password, created_at, updated_at
        FROM users
        WHERE 1=1
    `

    var args []interface{}

    userType := reflect.TypeOf(u)
    userValue := reflect.ValueOf(u)

    for i := 0; i < userType.NumField(); i++ {
        field := userType.Field(i)
        value := userValue.Field(i)

        if value.IsZero() {
            continue
        }

        switch value.Interface().(type) {
        case int, int64:
            query += " AND " + field.Tag.Get("db") + " = $" + strconv.Itoa(len(args)+1)
            args = append(args, value.Interface())
        case string:
            query += " AND " + field.Tag.Get("db") + " = $" + strconv.Itoa(len(args)+1)
            args = append(args, value.Interface())
        // Add more cases as needed for other types
        }
    }

    // Execute the query
    var err error
    if tx != nil {
        err = tx.QueryRowContext(ctx, query, args...).Scan(
            &user.ID,
            &user.FirstName,
            &user.LastName,
            &user.Email,
            &user.Phone,
			&user.Password,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
    } else {
        err = m.DB.QueryRowContext(ctx, query, args...).Scan(
            &user.ID,
            &user.FirstName,
            &user.LastName,
            &user.Email,
            &user.Phone,
			&user.Password,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
    }

    if err == sql.ErrNoRows {
        return nil, nil // No rows found, return nil
    }

    if err != nil {
        return &user, err
    }

    return &user, nil
}


// func (m *user) GetAUser(ctx context.Context, tx *sql.Tx, u models.User) (*models.User, error){
// 	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
// 	defer cancel()

// 	var user models.User

// 	query := `
// 	SELECT id, first_name, last_name, email, phone, created_at, updated_at
// 	FROM users
// 	WHERE 1=1
// `

// 	var args []interface{}
// 	argIndex := 1 // Counter for argument placeholders

// 	if u.ID != 0 {
// 		query += " AND id=$" + strconv.Itoa(argIndex)
// 		args = append(args, u.ID)
// 		argIndex++
// 	}
// 	if u.Email != "" {
// 		query += " AND email=$" + strconv.Itoa(argIndex)
// 		args = append(args, u.Email)
// 		argIndex++
// 	}
// 	if u.FirstName != "" {
// 		query += " AND first_name=$" + strconv.Itoa(argIndex)
// 		args = append(args, u.FirstName)
// 		argIndex++
// 	}
// 	if u.LastName != "" {
// 		query += " AND last_name=$" + strconv.Itoa(argIndex)
// 		args = append(args, u.LastName)
// 		argIndex++
// 	}

// 	var err error
// 	if tx != nil {
// 		err = tx.QueryRowContext(ctx, query, args...).Scan(
// 			&user.ID,
// 			&user.FirstName,
// 			&user.LastName,
// 			&user.Email,
// 			&user.Phone,
// 			&user.CreatedAt,
// 			&user.UpdatedAt,
// 		)
// 	}else{
// 		err = m.DB.QueryRowContext(ctx, query, args...).Scan(
// 			&user.ID,
// 			&user.FirstName,
// 			&user.LastName,
// 			&user.Email,
// 			&user.Phone,
// 			&user.CreatedAt,
// 			&user.UpdatedAt,
// 		)
// 	}

// 	if err == sql.ErrNoRows {
// 		return nil, nil // No rows found, return nil
// 	}

// 	if err != nil {
// 		return &user, err
// 	}

// 	return &user, nil
// }

func (m *user) GetAllUser(ctx context.Context, tx *sql.Tx) ([]models.User, error){
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var users = make([]models.User, 0)

	query := `
		SELECT (id, first_name, last_name, email, phone, created_at, updated_at)
		from users
	`

	var rows *sql.Rows
	var err error

	if tx != nil {
		rows, err = tx.QueryContext(ctx, query)
	}else{
		rows, err = m.DB.QueryContext(ctx, query)
	}
	if err != nil {
		return users, err
	}

	for rows.Next(){
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Phone,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func (m *user) UpdateAUsersName(ctx context.Context, tx *sql.Tx, id int, firstName, lastName string)(error){
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
		UPDATE 
			users set (first_name, last_name) = ($1, $2)
		WHERE
			id = $3
	`

	var err error
	if tx != nil{
		_, err = tx.ExecContext(ctx, query, firstName, lastName, id)
	}else{
		_, err = m.DB.ExecContext(ctx, query, firstName, lastName, id)
	}

	if err != nil{
		return  err
	}

	return nil
}

func (m *user) DeleteUserByID(ctx context.Context, tx *sql.Tx, id int) error {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    query := "DELETE FROM users WHERE id = $1"

	var err error 

	if tx != nil {
		_, err = tx.ExecContext(ctx, query, id)
	}else{
		_, err = m.DB.ExecContext(ctx, query, id)
	}

    if err != nil {
        return err
    }

    return nil
}