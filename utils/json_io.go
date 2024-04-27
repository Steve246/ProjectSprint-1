package utils

import "encoding/json"

func ToJsonString(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	return string(jsonData), err
}

func FromJsonString(jsonString string, result interface{}) error {
	return json.Unmarshal([]byte(jsonString), result)
}
