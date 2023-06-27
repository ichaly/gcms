package data

import (
	"database/sql/driver"
	"github.com/ichaly/gcms/core"
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

func (Kind) GqlEnumDescription() string {
	return "The episodes in the Star Wars trilogy"
}

func (Kind) GqlEnumValues() core.EnumValueMapping {
	return core.EnumValueMapping{
		"MOMENT":   {Moment, "Star Wars Episode IV: A New Hope, released in 1977."},
		"QUESTION": {Question, "Star Wars Episode V: The Empire Strikes Back, released in 1980."},
		"ANSWER":   {Answer, "Star Wars Episode VI: Return of the Jedi, released in 1983."},
	}
}
