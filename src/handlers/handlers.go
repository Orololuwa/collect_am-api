package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/config"
	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/Orololuwa/collect_am-api/src/mocks"
	"github.com/Orololuwa/collect_am-api/src/repository"
	dbrepo "github.com/Orololuwa/collect_am-api/src/repository/db-repo"
)

type Repository struct {
	App *config.AppConfig
	conn repository.DBInterface
	User repository.UserDBRepo
	Business repository.BusinessDBRepo
	Kyc repository.KycDBRepo
}

var Repo *Repository

// NewRepo function initializes the Repo
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		conn: db.Gorm,
		User: dbrepo.NewUserDBRepo(db),		
		Business: dbrepo.NewBusinessDBRepo(db),
		Kyc: dbrepo.NewKycDBRepo(db),
	}
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
