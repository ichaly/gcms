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
	return "The episodes in the Star Wars trilogy"
}

func (Kind) GqlEnumValues() map[string]*graphql.EnumValueConfig {
	return map[string]*graphql.EnumValueConfig{
		"MOMENT":   {Value: Moment, Description: "Star Wars Episode IV: A New Hope, released in 1977."},
		"QUESTION": {Value: Question, Description: "Star Wars Episode V: The Empire Strikes Back, released in 1980."},
		"ANSWER":   {Value: Answer, Description: "Star Wars Episode VI: Return of the Jedi, released in 1983."},
	}
}
