package data

import "database/sql/driver"

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
