package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/config"
	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/Orololuwa/collect_am-api/src/repository"
	dbrepo "github.com/Orololuwa/collect_am-api/src/repository/db-repo"
	"github.com/Orololuwa/collect_am-api/src/types"
	"github.com/go-playground/validator/v10"
)

type Middleware struct {
	App  *config.AppConfig
	User repository.UserDBRepo
}

func New(a *config.AppConfig, db *driver.DB) *Middleware {
	return &Middleware{
		App:  a,
		User: dbrepo.NewUserDBRepo(db),
	}
}

func NewTest(a *config.AppConfig) *Middleware {
	return &Middleware{
		App:  a,
		User: dbrepo.NewUserTestingDBRepo(),
	}
}

func (m *Middleware) ValidateReqBody(next http.Handler, requestBodyStruct interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(requestBodyStruct); err != nil {
			helpers.ClientError(w, err, http.StatusBadRequest, "failed to decode body")
			return
		}

		defer r.Body.Close()

		if err := m.App.Validate.Struct(requestBodyStruct); err != nil {
			errors := err.(validator.ValidationErrors)
			helpers.ClientError(w, err, http.StatusBadRequest, errors.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "validatedRequestBody", requestBodyStruct)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			helpers.ClientError(w, errors.New("missing token"), http.StatusUnauthorized, "")
			return
		}
		tokenString = tokenString[len("Bearer "):]

		token, err := helpers.VerifyJWTToken(tokenString)
		if err != nil {
			helpers.ClientError(w, errors.New("invalid or expired token"), http.StatusUnauthorized, "")
			return
		}

		claims, ok := token.Claims.(*types.JWTClaims)
		if ok {
			// get the user's data from the database and perform any verification necessary
			ctx := r.Context()
			user, err := m.User.GetOneByEmail(claims.Email)
			if err != nil {
				if err.Error() == "record not found" {
					helpers.ClientError(w, errors.New("user not found"), http.StatusBadRequest, "")
				} else {
					helpers.ClientError(w, err, http.StatusBadRequest, "")
				}
				return
			}

			ctx = context.WithValue(ctx, "user", &user)
			r = r.WithContext(ctx)

		} else {
			helpers.ClientError(w, errors.New("unknown claims type, cannot proceed"), http.StatusInternalServerError, "")
			return
		}

		next.ServeHTTP(w, r)
	})
}
