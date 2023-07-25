package core

import (
	"gorm.io/gorm"
	"time"
)

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

type DeleteEntity struct {
	AuditorEntity `mapstructure:",squash"`
	DeletedBy     *uint64        `gorm:"comment:删除人;"`
	DeletedAt     gorm.DeletedAt `gorm:"index;comment:逻辑删除;"`
}
