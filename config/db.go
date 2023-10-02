package config

import (
	"golang_basic_gin_juli_2023/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/golang_basic_gin_1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect database")
	}

	// Migrate the table
	DB.AutoMigrate(&models.Department{},
		&models.Position{}, &models.Employee{},
		&models.Inventory{}, &models.Archive{}, &models.EmployeeInventory{})

	// DB.Create(&models.Department{
	// 	Name: "Human Resources",
	// 	Code: "HR",
	// })
}
