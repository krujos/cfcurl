package cfcurl

import "errors"

// Curl calls cf curl  and return the resulting json. This method will fail if
// the api is depricated
func Curl(path string) (map[string]interface{}, error) {
	return nil, errors.New("Not Implemented")
}

// CurlDepricated calls cf curl and return the resulting json, even if the api is depricated
func CurlDepricated(path string) (map[string]interface{}, error) {
	return nil, errors.New("Not implemented")
}
