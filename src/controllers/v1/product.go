package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/helpers"
)

func (m *V1) AddProduct(w http.ResponseWriter, r *http.Request){
	var body dtos.AddProduct
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}

	err = m.App.Validate.Struct(body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}
}