package db

import (
	"database/sql"
	"e-money-svc/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"runtime"
	"time"
)

var (
	EmoneyDb *gorm.DB
)

func init() {
	cfg := config.GetConfig()
	EmoneyDb = OpenDBConnection(cfg.Db.EMoneySvc.Username, cfg.Db.EMoneySvc.Password, cfg.Db.EMoneySvc.Port, cfg.Db.EMoneySvc.Database, cfg.Db.EMoneySvc.Host)

}

func OpenDBConnection(user, pass, port, dbname, host string) *gorm.DB {

	fmt.Println("OpenDBConnection")
	dsn := user + ":" + pass + "@tcp(" + host + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	sqlDB, err := sql.Open("mysql", dsn)

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{Logger: logger.Default.LogMode((logger.Error))})

	if err != nil {
		if err.Error() == "Error 1049: Unknown database 'e_money_svc'" {

			return OpenDBConnection(user, pass, port, "", host)
		}

		panic(err)
		return nil
	}

	if dbname == "" {
		var dbNama = "e_money_svc"
		query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", dbNama)
		tx := db.Exec(query).Debug()

		if tx.Error != nil {
			panic(tx.Error.Error())
		}

		return OpenDBConnection(user, pass, port, dbNama, host)
	}

	dbo, _ := db.DB()

	dbo.SetMaxIdleConns(runtime.NumCPU())
	dbo.SetMaxOpenConns(runtime.NumCPU())
	dbo.SetConnMaxLifetime(time.Hour)

	return db
}
