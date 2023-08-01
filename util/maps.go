package util

import (
	"strings"
)

func QueryMap(data map[string]interface{}, path string) interface{} {
	arr := strings.SplitN(path, ".", 2)
	if len(arr) <= 1 {
		return data[path]
	}
	if val, ok := data[arr[0]].(map[string]interface{}); ok {
		return QueryMap(val, arr[1])
	}
	return nil
}

func EraseMap(data map[string]interface{}, path string) interface{} {
	arr := strings.SplitN(path, ".", 2)
	if len(arr) <= 1 {
		res := data[path]
		delete(data, path)
		return res
	}
	if val, ok := data[arr[0]].(map[string]interface{}); ok {
		return EraseMap(val, arr[1])
	}
	return nil
}
