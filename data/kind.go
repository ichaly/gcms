package data

import (
	"database/sql/driver"
)

type Kind string

const (
	Moment   Kind = "MOMENT"
	Question Kind = "QUESTION"
	Answer   Kind = "ANSWER"
)

func (my *Kind) Scan(value interface{}) error {
	*my = Kind(value.(string))
	return nil
}

func (my Kind) Value() (driver.Value, error) {
	return string(my), nil
}

func (Kind) Description() string {
	return "Content type"
}

func (Kind) EnumValues() map[string]*struct {
	Value             interface{}
	Description       string
	DeprecationReason string
} {
	return map[string]*struct {
		Value             interface{}
		Description       string
		DeprecationReason string
	}{
		"MOMENT":   {Value: Moment, Description: "动态"},
		"ANSWER":   {Value: Answer, Description: "回答"},
		"QUESTION": {Value: Question, Description: "问题"},
	}
}
