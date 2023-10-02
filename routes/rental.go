package routes

import (
	"golang_basic_gin_juli_2023/config"
	"golang_basic_gin_juli_2023/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetRental(c *gin.Context) {
	EmployeeInventory := []models.EmployeeInventory{}

	// config.DB.Find(&EmployeeInventory)

	config.DB.Preload(clause.Associations).Find(&EmployeeInventory)

	responseGetRentals := []models.ResponseGetRental{}

	for _, ei := range EmployeeInventory {
		rgr := models.ResponseGetRental{
			Id:            ei.ID,
			Description:   ei.Description,
			EmployeeName:  ei.Employee.Name,
			InventoryName: ei.Inventory.Name,
			CreatedAt:     ei.CreatedAt,
		}
		responseGetRentals = append(responseGetRentals, rgr)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Data Rental",
		"data":    responseGetRentals,
	})
}

func RentalByEmployeeId(c *gin.Context) {
	var reqRental models.RequestRental

	if err := c.ShouldBindJSON(&reqRental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	rental := models.EmployeeInventory{
		EmployeeID:  reqRental.EmployeeID,
		InventoryID: reqRental.InventoryID,
		Description: reqRental.Description,
	}
	insert := config.DB.Create(&rental)
	if insert.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   insert.Error.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    rental,
		"message": "Insert Success",
	})

}

func GetRentalByInventoryID(c *gin.Context) {
	id := c.Param("id")

	Inventories := models.Inventory{}

	config.DB.Preload(clause.Associations).First(&Inventories, "id = ?", id)

	emInv := []models.RespEmployeeInventory{}
	for _, inv := range Inventories.Employees {
		emInv = append(emInv, models.RespEmployeeInventory{
			EmployeeID:  inv.EmployeeID,
			InventoryID: inv.InventoryID,
			Description: inv.Description,
		})
	}

	respInv := models.ResponseInventoryEmployee{
		InventoryName:        Inventories.Name,
		InventoryDescription: Inventories.Description,
		EmployeeInventory:    emInv,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Data Rental By Inventory",
		"data":    respInv,
	})
}
