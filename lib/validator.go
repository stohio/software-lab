package softwarelab

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
// @param param: the name of the paramter that is being validated
// @param value: the value of the paramter that is being validated
// @param regex: the regular expression that value should be matched with
// @param w: if the value isn't validated or if it doesn't match regex, the response will indicate what went wrong
// @return: true if the parameter is present and matches the regular expression
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
// @param param: the name of the parameter that is being validated
// @param value: the vlaue of the parameter that is being validated
// @param w: If the parameter isn't validated then the response will indicate which paramater failed
// @return: true if the paramter is present, false otherwise
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

// ValidateAndUnmarshalJSON checks that the json in body is formatted correctly, and matches the interface
// DataStruct. If it does match then the json from body is copied over to DataStruct
// @param body: the http request body that should contain json
// @param DataStruct: a pointer to the structure the json in body should be copied over to
// @param w: If the unmarshaling fails then the error gotten will be sent with the response
// @return: true if the JSON in body was unmarshaled succsesfully
func ValidateAndUnmarshalJSON(body []byte, DataStruct interface{}, w http.ResponseWriter) bool {

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
