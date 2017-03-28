package main

// TODO should this be in the lib directory instead of in remote directory?

import (
	"encoding/json"
	"net/http"
	"regexp"
)

// ParamError is a struct representing the what wen't wrong with a certain parameter and its value
type ParamError struct {
	Error string `json:"error,omitempty"`
	Param string `json:"param,omitempty"`
	Value string `json:"value,omitempty"`
}

// ValidateParamRegex checks that value is valid, and that it is matched by regex. It calls ValidateParam
// to check that value is valid, and if value doesn't match regex, then a response is sent saying the parameter
// is invalid
func ValidateParamRegex(param string, value *string, regex string, w http.ResponseWriter) bool {
	if !ValidateParam(param, value, w) {
		return false
	}

	if matched, _ := regexp.MatchString(regex, *value); matched {
		return true
	}

	response := ParamError{
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

// ValidateParam checks that value isn't null. If it is then it sends a response saying the parameter
// wasn't found
func ValidateParam(param string, value *string, w http.ResponseWriter) bool {

	if value != nil {
		return true
	}
	response := ParamError{
		Error: "Param not found",
		Param: param,
	}
	w.Header().Set("Content-Type", "applciation/json; charset=UTF-8")
	w.WriteHeader(400)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
	return false
}

// ValidateJSON checks that the json in body is formatted correctly, and matches the interface
// DataStruct. If it does match then the json from body is copied over to DataStruct
// TODO should ValidateJSON copy over the json from body to DataStruct, or should it just check,
// and there be another function to copy
func ValidateJSON(body []byte, DataStruct interface{}, w http.ResponseWriter) bool {

	if err := json.Unmarshal(body, DataStruct); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return false
	}
	return true
}
