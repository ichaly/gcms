package data

import (
	"database/sql/driver"
	"github.com/graphql-go/graphql"
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

func (Kind) GqlDescription() string {
	return "Content type"
}

func (Kind) GqlEnumValues() map[string]*graphql.EnumValueConfig {
	return map[string]*graphql.EnumValueConfig{
		"MOMENT":   {Value: Moment, Description: "动态"},
		"ANSWER":   {Value: Answer, Description: "回答"},
		"QUESTION": {Value: Question, Description: "问题"},
	}
}
