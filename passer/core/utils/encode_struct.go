package utils

import "github.com/mitchellh/mapstructure"

func EncodeStructToMap(data interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := mapstructure.Decode(data, &result)
	return result, err
}
