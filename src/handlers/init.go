package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/config"
	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/Orololuwa/collect_am-api/src/mocks"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	dbrepo "github.com/Orololuwa/collect_am-api/src/repository/db-repo"
	"github.com/Orololuwa/collect_am-api/src/types"
)

type ErrorData struct {
	Message string
	Error error
	Status int
}

type HandlerFunc interface {
	SignUpV2(payload dtos.UserSignUp)(userId uint, errData *ErrorData)
	LoginUserV2(payload dtos.UserLoginBody)(data types.LoginSuccessResponse, errData *ErrorData)

	CreateBusiness(payload dtos.AddBusiness, options ...*Extras)(id uint, errData *ErrorData)
	GetBusinessV2(options ...*Extras)(data *models.Business, errData *ErrorData)
	UpdateBusinessV2(payload map[string]interface{}, options ...*Extras)(errData *ErrorData)
}

type Repository struct {
	App *config.AppConfig
	conn repository.DBInterface
	User repository.UserDBRepo
	Business repository.BusinessDBRepo
	Kyc repository.KycDBRepo
}

var Repo *Repository

// NewRepo function initializes the Repo
func NewRepo(a *config.AppConfig, db *driver.DB) HandlerFunc {
	handler := &Repository{
		App: a,
		conn: db.Gorm,
		User: dbrepo.NewUserDBRepo(db),		
		Business: dbrepo.NewBusinessDBRepo(db),
		Kyc: dbrepo.NewKycDBRepo(db),
	}

	Repo = handler

	return handler;
}

// NewRepo function initializes the Repo
func NewTestRepo(a *config.AppConfig) *Repository {
	mockDB := mocks.NewMockDB()

	return &Repository{
		App: a,
		conn: mockDB,
		User: dbrepo.NewUserTestingDBRepo(),
		Business: dbrepo.NewBusinessTestingDBRepo(),
		Kyc: dbrepo.NewKycTestingDBRepo(),
	}
}

func NewHandlers(r *Repository){
	Repo = r;
}

func Init(a *config.AppConfig, db *driver.DB) HandlerFunc {
	return &Repository{
		App: a,
		conn: db.Gorm,
		User: dbrepo.NewUserDBRepo(db),		
		Business: dbrepo.NewBusinessDBRepo(db),
		Kyc: dbrepo.NewKycDBRepo(db),
	}
}

type jsonResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func (m *Repository) Health(w http.ResponseWriter, r *http.Request){
	resp := jsonResponse{
		Message: "App Healthy",
		Data: nil,
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) ProtectedRoute(w http.ResponseWriter, r *http.Request){
	helpers.ClientResponseWriter(w, nil, http.StatusOK, "welcome to the protected route")
}

type Extras struct {
	User *models.User
}