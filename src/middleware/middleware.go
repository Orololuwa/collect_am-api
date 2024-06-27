package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Orololuwa/collect_am-api/src/config"
	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	dbrepo "github.com/Orololuwa/collect_am-api/src/repository/db-repo"
	"github.com/Orololuwa/collect_am-api/src/types"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Middleware struct {
	App      *config.AppConfig
	User     repository.UserDBRepo
	Business repository.BusinessDBRepo
}

func New(a *config.AppConfig, db *driver.DB) *Middleware {
	return &Middleware{
		App:      a,
		User:     dbrepo.NewUserDBRepo(db),
		Business: dbrepo.NewBusinessDBRepo(db),
	}
}

func NewTest(a *config.AppConfig) *Middleware {
	return &Middleware{
		App:      a,
		User:     dbrepo.NewUserTestingDBRepo(),
		Business: dbrepo.NewBusinessTestingDBRepo(),
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

func (m *Middleware) BusinessValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		businessId, err := strconv.Atoi(chi.URLParam(r, "businessId"))
		if err != nil {
			helpers.ClientError(w, err, http.StatusInternalServerError, "")
			return
		}

		// get the business' data from the db
		ctx := r.Context()
		business, err := m.Business.GetOneById(uint(businessId))
		if err != nil {
			if err.Error() == "record not found" {
				helpers.ClientError(w, errors.New("business not found"), http.StatusBadRequest, "")
			} else {
				helpers.ClientError(w, err, http.StatusBadRequest, "")
			}
			return
		}

		ctx = context.WithValue(ctx, "business", &business)
		r = r.WithContext(ctx)

		// get the user details and check if business belongs to user
		user, ok := r.Context().Value("user").(*models.User)
		if !ok || user == nil {
			helpers.ClientError(w, errors.New("unauthorized"), http.StatusUnauthorized, "")
			return
		}

		if user.ID != uint(business.UserID) {
			helpers.ClientError(w, errors.New("business doesn't belong to user"), http.StatusUnauthorized, "")
			return
		}

		// check if business is approved
		if !business.IsSetupComplete {
			helpers.ClientError(w, errors.New("business setup incomplete"), http.StatusUnauthorized, "")
			return
		}

		//check if business is disabled or thereabout

		next.ServeHTTP(w, r)
	})
}
