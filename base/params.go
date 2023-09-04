package base

import (
	"fmt"
	"gorm.io/gorm"
)

type Params[T any] struct {
	Page   int
	Size   int
	Delete bool
	Data   T
	Sort   map[string]string
	Where  map[string]interface{}
}

func ParseOperator(str string, val interface{}) string {
	switch str {
	case "eq":
		return "="
	case "ne":
		return "<>"
	case "gt":
		return ">"
	case "lt":
		return "<"
	case "ge":
		return ">="
	case "le":
		return "<="
	case "like":
		return "LIKE"
	case "notLike":
		return "NOT LIKE"
	case "regex":
		return "~"
	case "notRegex":
		return "!~"
	case "isNull":
		if val.(bool) {
			return "IS NULL"
		} else {
			return "IS NOT NULL"
		}
	}
	return ""
}

func ParseWhere(data map[string]interface{}, tx *gorm.DB) {
	for key, val := range data {
		for k, v := range val.(map[string]interface{}) {
			tx = tx.Where(fmt.Sprintf("%s %s ?", key, ParseOperator(k, v)), v)
		}
	}
}

func ParseSort(data map[string]string, tx *gorm.DB) {
	for k, v := range data {
		switch v {
		case "ASC_NULLS_FIRST":
		case "ASC_NULLS_LAST":
		case "DESC_NULLS_FIRST":
		case "DESC_NULLS_LAST":
		case "ASC":
		case "DESC":
		default:
			tx = tx.Order(fmt.Sprintf("%s %s", k, v))
		}
	}
}
