package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/config"
	"github.com/Orololuwa/collect_am-api/src/handlers"
	"github.com/Orololuwa/collect_am-api/src/helpers"
)

type V1 struct {
	App *config.AppConfig
	Handlers handlers.HandlerFunc
}

var v1 *V1

func NewController(a *config.AppConfig, h handlers.HandlerFunc) *V1 {
	v1Instance := &V1{
		App: a,
		Handlers: h,
	}
	v1= v1Instance

	return v1Instance
}

type jsonResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func (m *V1) Health(w http.ResponseWriter, r *http.Request){
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