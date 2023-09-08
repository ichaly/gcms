package test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ichaly/gcms/apps/data"
	"github.com/ichaly/gcms/auth"
	"github.com/ichaly/gcms/base"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"sync"
	"testing"
	"time"
)

type Demo struct {
	Name        string `gorm:"size:200;comment:名字"`
	base.Entity `mapstructure:",squash"`
}

func (*Demo) TableName() string {
	return "sys_demo"
}

func TestConnect(t *testing.T) {
	cfg, err := base.NewConfig("../conf/dev.yml")
	if err != nil {
		t.Error(err)
	}
	d := &Demo{}
	db, err := base.NewConnect(cfg, []gorm.Plugin{base.NewSonyFlake()}, []interface{}{d})
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
				var ids []base.ID
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

func TestDemoUser(t *testing.T) {
	cfg, err := base.NewConfig("../conf/dev.yml")
	if err != nil {
		t.Error(err)
	}
	db, err := base.NewConnect(cfg, []gorm.Plugin{base.NewSonyFlake()}, []interface{}{})
	if err != nil {
		t.Error(err)
	}
	err = db.Model(&data.User{}).Save(&data.User{Username: "admin", Password: "123456"}).Error
	if err != nil {
		t.Error(err)
	}
}

func TestDemoClient(t *testing.T) {
	cfg, err := base.NewConfig("../conf/dev.yml")
	if err != nil {
		t.Error(err)
	}
	db, err := base.NewConnect(cfg, []gorm.Plugin{base.NewSonyFlake()}, []interface{}{})
	if err != nil {
		t.Error(err)
	}
	secret := strings.ReplaceAll(uuid.Must(uuid.NewRandom()).String(), "-", "")
	t.Log(secret)
	err = db.Model(&auth.Client{}).Save(&auth.Client{Secret: secret}).Error
	if err != nil {
		t.Error(err)
	}
}
