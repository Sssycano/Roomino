package dao

import (
	"context"
	"fmt"
	"roomino/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func MySQLInit() {

	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 8889
	Dbname := "roomino"
	timeout := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("sqllink, error=" + err.Error())
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db
	err = db.AutoMigrate(model.Modlist...)
	if err != nil {
		panic("migrate, error=" + err.Error())
	}

}
func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
