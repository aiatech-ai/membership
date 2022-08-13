package util

import "encoding/json"

func StructToJsonMap[T any](s T) (map[string]interface{}, error) {
	var inInterface map[string]interface{}
	inrec, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &inInterface)
	if err != nil {
		return nil, err
	}
	return inInterface, nil
}
