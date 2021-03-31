package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kyrgyz-bilim/entity"
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
	dbConfig := DBConfig{
		Host:     "172.19.0.1",
		Port:     5555,
		User:     "kyrgyzBilim",
		DBName:   "kyrgyzBilim",
		Password: "dbCXeDTR5vFhZkxRCq",
	}
	return &dbConfig
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
	db.AutoMigrate(&entity.User{})
}
