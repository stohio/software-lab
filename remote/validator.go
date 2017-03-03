package main

import (
	"net/http"
	"regexp"
	"encoding/json"
)

type ParamError struct {
	Error	string `json:"error,omitempty"`
	Param	string	`json:"param,omitempty"`
	Value	string `json:"value,omitempty"`
}


func ValidateParamRegex(param string, value *string, regex string, w http.ResponseWriter) bool {
	if !ValidateParam(param, value, w) {
		return false
	}

	if matched, _ := regexp.MatchString(regex, *value); matched {
		return true
	} else {
		response := ParamError {
			Error: "Invalid param format",
			Param: param,
			Value: *value,
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(400)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
		return false
	}
}

func ValidateParam(param string, value *string, w http.ResponseWriter) bool {

	if value != nil {
		return true
			}
	response := ParamError {
		Error: "Param not found",
		Param: param,
	}
	w.Header().Set("Content-Type", "applciation/json; charset=UTF-8")
	w.WriteHeader(400)
	if err:= json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
	return false
}

func ValidateJson(body []byte, DataStruct interface{}, w http.ResponseWriter) bool {

	if err := json.Unmarshal(body, DataStruct); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err !=nil {
			panic(err)
		}
		return false
	}
	return true
}
