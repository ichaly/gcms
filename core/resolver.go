package core

import (
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

func QueryResolver[T any](in graphql.ResolveParams, db *gorm.DB) (interface{}, error) {
	var p Params[T]
	err := mapstructure.WeakDecode(in.Args, &p)
	if err != nil {
		return nil, err
	}
	tx := db.WithContext(in.Context).Model(new(T))
	if p.Where != nil {
		ParseWhere(p.Where, tx)
	}
	if p.Sort != nil {
		ParseSort(p.Sort, tx)
	}
	if p.Size > 1000 || p.Page < 0 {
		p.Size = 10
	}
	if p.Size > 0 {
		tx = tx.Limit(p.Size)
	}
	tx = tx.Offset(p.Size * p.Page)
	var res []T
	err = tx.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, err
}

func MutationResolver[T any](in graphql.ResolveParams, db *gorm.DB) (interface{}, error) {
	var p Params[T]
	err := mapstructure.WeakDecode(in.Args, &p)
	if err != nil {
		return nil, err
	}
	tx := db.WithContext(in.Context).Model(new(T))
	if p.Where != nil {
		ParseWhere(p.Where, tx)
	}
	if p.Delete {
		err = tx.Delete(&p.Data).Error
		return nil, err
	}
	err = tx.Save(&p.Data).Error
	if err != nil {
		return nil, err
	}
	return p.Data, nil
}