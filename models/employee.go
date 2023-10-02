package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string   `json:"name"`
	Address    string   `json:"address"`
	Email      string   `json:"email"`
	PositionID uint     `json:"position_id"`
	Position   Position `json:"position"`
	// Inventories []*Inventory `gorm:"many2many:employee_inventories;"`
	Inventories []EmployeeInventory `json:"inventories"`
}

type EmployeeResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

type GetEmployeeResponse struct {
	ID         uint             `json:"id"`
	Name       string           `json:"name"`
	Address    string           `json:"address"`
	Email      string           `json:"email"`
	PositionID uint             `json:"position_id"`
	Position   PositionResponse `json:"position"`
}
