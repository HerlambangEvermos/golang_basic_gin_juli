package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Archive     Archive
	// Employess   []*Employee `gorm:"many2many:employee_inventories;"`
	Employees []EmployeeInventory `json:"employees"`
}

type RequestInventory struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"inventory_description"`
	ArchiveName          string `json:"archive_name"`
	ArchiveDescription   string `json:"archive_description"`
}

type ResponseInventory struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"inventory_description"`
	Archive              ResponseArchive
}

type ResponseInventoryEmployee struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"inventory_description"`
	EmployeeInventory    []RespEmployeeInventory
}
