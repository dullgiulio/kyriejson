package json

import (
	"fmt"
)

// TODO: This should take a jsonCheckList as argument (what to do with the returned data)
func GetKey(key string) jsonCheck {
	return func(data jsonData) (jsonData, error) {
		m, ok := data.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Expected a map, but it wasn't")
		}
		d, ok := m[key]
		if !ok {
			return nil, fmt.Errorf("JSON key %s not found", key)
		}
		return d, nil
	}
}

func InArray(data jsonData) (jsonData, error) {
	arr, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("Expected an array but it wasn't")
	}
	return arr, nil
}

func Each(checks jsonCheckList) jsonCheck {
	return func(data jsonData) (jsonData, error) {
		arr, ok := data.([]interface{})
		if !ok {
			return nil, fmt.Errorf("Expected an array but it wasn't")
		}
		for i := range arr {
			if err := checks.run(arr[i]); err != nil {
				return nil, err
			}
		}
		return data, nil
	}
}
