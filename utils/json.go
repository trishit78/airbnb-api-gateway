package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	
)


var Validator *validator.Validate


func init(){
	Validator = NewValidator()
}

func NewValidator() *validator.Validate{
	return validator.New((validator.WithRequiredStructEnabled()))
}

func WriteJsonResponse(w http.ResponseWriter, status int ,data any) error {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}

func WriteJsonSuccessResponse(w http.ResponseWriter,status int, message string, data any) error{
	res:=map[string]any{}

	res["status"] = "success"
	res["message"] = message
	res["data"] = data
	return WriteJsonResponse(w,status,res)
}


func WriteJsonErrorResponse(w http.ResponseWriter,status int, message string, data any) error{
	res:=map[string]any{}

	res["status"] = "error"
	res["message"] = message
	res["data"] = data
	return WriteJsonResponse(w,status,res)
}

func ReadJsonBody(r *http.Request,res any) error{
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(res)
}