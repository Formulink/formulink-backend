package utils

import "encoding/json"

func Parse(value []byte, err error) (result interface{}, e error) {
	if err != nil {
		return nil, err
	}
	var res interface{}
	err = json.Unmarshal(value, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
