package base

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewConnect(c *Config, p []gorm.Plugin, e []interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(
		buildDialect(c.Database),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	if err != nil {
		return nil, err
	}
	for _, v := range p {
		err = db.Use(v)
		if err != nil {
			return nil, err
		}
	}
	if c.App.Debug {
		err = db.AutoMigrate(e...)
		if err != nil {
			return nil, err
		}
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(90)
	sqlDb.SetConnMaxLifetime(5 * time.Minute)
	return db, nil
}

func buildDialect(ds *DataSource) gorm.Dialector {
	args := []interface{}{ds.Username, ds.Password, ds.Host, ds.Port, ds.Name}
	if ds.Dialect == "mysql" {
		return mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", args...,
		))
	} else {
		return postgres.Open(fmt.Sprintf(
			"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable TimeZone=Asia/Shanghai", args...,
		))
	}
}
