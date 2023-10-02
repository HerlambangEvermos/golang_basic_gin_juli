package models

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeInventory struct {
	gorm.Model
	EmployeeID  uint      `json:"employee_id"`
	Employee    Employee  `gorm:"foreignKey:EmployeeID;references:ID"`
	InventoryID uint      `json:"inventory_id"`
	Inventory   Inventory `gorm:"foreignKey:InventoryID;references:ID"`
	Description string    `json:"description"`
}

type RequestRental struct {
	EmployeeID  uint   `json:"employee_id"`
	InventoryID uint   `json:"inventory_id"`
	Description string `json:"description"`
}

type ResponseGetRental struct {
	Id            uint   `json:"id"`
	Description   string `json:"description"`
	EmployeeName  string `json:"employee_name"`
	InventoryName string `json:"inventory_name"`
	CreatedAt     time.Time
}

type RespEmployeeInventory struct {
	EmployeeID  uint   `json:"employee_id"`
	InventoryID uint   `json:"inventory_id"`
	Description string `json:"description"`
}
