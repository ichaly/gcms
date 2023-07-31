package core

import (
	"gorm.io/gorm"
	"time"
)

type userContextKeyType struct{}

var UserContextKey = userContextKeyType{}

type ID uint64

func (my ID) ID() {}

type Primary struct {
	ID ID `gorm:"primary_key;comment:主键;next:sonyflake;"`
}

type General struct {
	State     int8      `gorm:"comment:状态;"`
	Remark    string    `gorm:"type:text;comment:备注"`
	CreatedAt time.Time `gorm:"comment:创建时间;"`
	UpdatedAt time.Time `gorm:"comment:更新时间;"`
}

type Entity struct {
	Primary `mapstructure:",squash"`
	General `mapstructure:",squash"`
}

type AuditorEntity struct {
	Entity    `mapstructure:",squash"`
	CreatedBy *uint64 `gorm:"comment:创建人;"`
	UpdatedBy *uint64 `gorm:"comment:更新人;"`
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

type DeleteEntity struct {
	AuditorEntity `mapstructure:",squash"`
	DeletedBy     *uint64        `gorm:"comment:删除人;"`
	DeletedAt     gorm.DeletedAt `gorm:"index;comment:逻辑删除;"`
}

func (my *DeleteEntity) BeforeDelete(tx *gorm.DB) error {
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
