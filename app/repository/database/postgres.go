package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kyrgyz-bilim/entity"
	"os"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := &DBConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     5432,
		User:     os.Getenv("POSTGRES_USER"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}
	return dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(DbURL(BuildDBConfig())), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func SetupDB(db *gorm.DB) {
	err := db.AutoMigrate(&entity.User{})
	err = db.AutoMigrate(&entity.Course{})
	err = db.AutoMigrate(&entity.Section{})
	err = db.AutoMigrate(&entity.Topic{})
	err = db.AutoMigrate(&entity.SubTopic{})
	if err != nil {
		panic("Can't migrate database")
	}
}
