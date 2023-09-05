package base

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
	"testing"
	"time"
)

type Demo struct {
	Name   string `gorm:"size:200;comment:名字"`
	Entity `mapstructure:",squash"`
}

func (*Demo) TableName() string {
	return "sys_demo"
}

func TestConnect(t *testing.T) {
	cfg, err := NewConfig("../conf/dev.yml")
	if err != nil {
		t.Error(err)
	}
	d := &Demo{}
	db, err := NewConnect(cfg, []gorm.Plugin{NewSonyFlake()}, []interface{}{d})
	if err != nil {
		t.Error(err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			err := db.Transaction(func(tx *gorm.DB) error {
				var demos []*Demo
				_tx := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&Demo{}).Where("state = ?", 0)
				_tx.Find(&demos)
				var ids []ID
				for _, d := range demos {
					ids = append(ids, d.ID)
				}
				time.Sleep(50 * time.Millisecond)
				return tx.Model(&Demo{}).Save(&Demo{Name: fmt.Sprintf("用户%d", n)}).Error
			})
			if err != nil {
				t.Error(err)
			}

			sqDb, _ := db.DB()
			t.Log("in use :", sqDb.Stats().InUse)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
