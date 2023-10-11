package base

import (
	"context"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type userContextKeyType struct{}

var UserContextKey = userContextKeyType{}

type ID uint64

func (my ID) ID() {}

type Primary struct {
	ID ID `gorm:"primary_key;comment:主键;next:sonyflake;" json:",omitempty"`
}

type General struct {
	State     int8              `gorm:"index;comment:状态;"	`
	Remark    datatypes.JSONMap `gorm:"comment:备注" json:",omitempty"`
	CreatedAt *time.Time        `gorm:"comment:创建时间;" json:",omitempty"`
	UpdatedAt *time.Time        `gorm:"comment:更新时间;" json:",omitempty"`
}

type Entity struct {
	Primary `mapstructure:",squash"`
	General `mapstructure:",squash"`
}

type Deleted struct {
	DeletedAt *gorm.DeletedAt `gorm:"index;comment:逻辑删除;" json:",omitempty"`
}

type DeleteEntity struct {
	AuditorEntity `mapstructure:",squash"`
	Deleted       `mapstructure:",squash"`
}

type AuditorEntity struct {
	Entity    `mapstructure:",squash"`
	CreatedBy *uint64 `gorm:"comment:创建人;" json:",omitempty"`
	UpdatedBy *uint64 `gorm:"comment:更新人;" json:",omitempty"`
	DeletedBy *uint64 `gorm:"comment:删除人;" json:",omitempty"`
}

func (my *AuditorEntity) BeforeCreate(tx *gorm.DB) error {
	if val, ok := getCurrentUserFromContext(tx.Statement.Context); ok {
		my.CreatedBy = val
	}
	return nil
}

func (my *AuditorEntity) BeforeUpdate(tx *gorm.DB) error {
	if val, ok := getCurrentUserFromContext(tx.Statement.Context); ok {
		my.UpdatedBy = val
	}
	return nil
}

func (my *AuditorEntity) BeforeDelete(tx *gorm.DB) error {
	if val, ok := getCurrentUserFromContext(tx.Statement.Context); ok {
		my.UpdatedBy = val
	}
	return nil
}

func getCurrentUserFromContext(ctx context.Context) (*uint64, bool) {
	val, ok := ctx.Value(UserContextKey).(string)
	if !ok || val == "" {
		return nil, false
	}
	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return nil, false
	}
	return &num, ok
}
