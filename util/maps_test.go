package util

import (
	"fmt"
	"testing"
)

var data = map[string]interface{}{
	"data": map[string]interface{}{
		"user": map[string]interface{}{
			"username": "ichaly",
			"password": "admin",
		},
	},
}

func TestQueryMap(t *testing.T) {
	res := QueryMap(data, "data.user.username")
	fmt.Println(res)
}

func TestEraseMap(t *testing.T) {
	fmt.Println(EraseMap(data, "data.user.password"))
	fmt.Println(EraseMap(data, "data.user.nickname"))
	fmt.Println(data)
}
