package base

import (
	"gorm.io/gorm"
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
	State     int8       `gorm:"index;comment:状态;"`
	Remark    string     `gorm:"type:text;comment:备注" json:",omitempty"`
	CreatedAt *time.Time `gorm:"comment:创建时间;" json:",omitempty"`
	UpdatedAt *time.Time `gorm:"comment:更新时间;" json:",omitempty"`
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
	if val, ok := getCurrentUserFromContext(tx); ok {
		my.CreatedBy = val
	}
	return nil
}

func (my *AuditorEntity) BeforeUpdate(tx *gorm.DB) error {
	if val, ok := getCurrentUserFromContext(tx); ok {
		my.UpdatedBy = val
	}
	return nil
}

func (my *AuditorEntity) BeforeDelete(tx *gorm.DB) error {
	if val, ok := getCurrentUserFromContext(tx); ok {
		my.UpdatedBy = val
	}
	return nil
}

func getCurrentUserFromContext(tx *gorm.DB) (*uint64, bool) {
	ctx := tx.Statement.Context
	if val, ok := ctx.Value(UserContextKey).(uint64); ok && val > 0 {
		return &val, ok
	}
	return nil, false
}
