package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/Orololuwa/collect_am-api/src/config"
	"github.com/theritikchoure/logx"
)

var app *config.AppConfig

func NewHelper(a *config.AppConfig){
	app = a
}

func ClientError(w http.ResponseWriter, err error, status int,  message string) {
	errorMessage := message
	if errorMessage == "" {
		errorMessage = err.Error()
	}

	logx.ColoringEnabled = true
	if app.GoEnv != "test" {
		logx.Log(err.Error(), logx.FGRED, logx.BGBLACK)
	}

	response := map[string]interface{}{"message": errorMessage, "error": err}
    errorResponse, errJson := json.Marshal(response)
    if errJson != nil {
        http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(errorResponse)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	logx.ColoringEnabled = true
	logx.Log(err.Error(), logx.FGRED, logx.BGBLACK)
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ClientResponseWriter(w http.ResponseWriter, data interface{}, status int, message string){
	response := map[string]interface{}{"message": message, "data": data}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(jsonResponse)
}

// func StringToBool(s string) bool {
//     return strings.EqualFold(s, "true")
// }

func StringToBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

// AssignIfExists mimics javascripts Object.assign
func AssignIfExists(src, dst map[string]interface{}, keys ...string) {
    for _, key := range keys {
        if value, ok := src[key]; ok {
            dst[key] = value
        }
    }
}