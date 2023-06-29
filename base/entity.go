package base

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
	"time"
)

type EntityGroup struct {
	fx.In
	Entities []interface{} `group:"entity"`
}

type ID uint64

func (my ID) GqlID() {}

type Primary struct {
	ID ID `gorm:"primary_key;AUTO_INCREMENT;comment:主键;"`
}

type General struct {
	State     int8      `gorm:"comment:状态;"`
	Remark    string    `gorm:"type:text;comment:备注"`
	CreatedAt time.Time `gorm:"comment:创建时间;"`
	UpdatedAt time.Time `gorm:"comment:更新时间;"`
}

type Entity struct {
	Primary
	General
}

type AuditorEntity struct {
	Entity
	CreatedBy *uint64 `gorm:"comment:创建人;"`
	UpdatedBy *uint64 `gorm:"comment:更新人;"`
}

type DeleteEntity struct {
	AuditorEntity
	DeletedBy *uint64        `gorm:"comment:删除人;"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:逻辑删除;"`
}
