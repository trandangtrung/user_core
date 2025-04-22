package db

import (
	"demo/config"
	"demo/internal/entity"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB
}

var dbinstance *Database
var dbonce sync.Once

func GetDatabaseConnection() *Database {
	dbonce.Do(func() {
		dbinstance = &Database{
			Connection: dbConnect(),
		}
	})
	return dbinstance
}

func dbConnect() *gorm.DB {

	cfg := config.GetConfig().DbCfg
	dsn := cfg.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database, error: %v", err)
	}

	err = db.AutoMigrate(entity.Platform{}, entity.Role{}, entity.Token{},
		entity.UserPlatform{}, entity.UserRole{}, entity.User{})
	if err != nil {
		log.Fatalf("failed to migrate database, error: %v", err)
	}

	return db
}

func (d *Database) Close() error {
	sqlDB, err := d.Connection.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
