package routes

import (
	"golang_basic_gin_juli_2023/config"
	"golang_basic_gin_juli_2023/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetEmployees(c *gin.Context) {
	employee := []models.Employee{}
	config.DB.Find(&employee)

	config.DB.Preload(clause.Associations).Find(&employee)

	// config.DB.Preload("EmployeeInventory.Inventory").First(&employee, 1)

	getEmployeeResponse := []models.GetEmployeeResponse{}

	for _, e := range employee {

		em := models.GetEmployeeResponse{
			ID:         e.ID,
			Name:       e.Name,
			Address:    e.Address,
			Email:      e.Email,
			PositionID: e.PositionID,
			Position: models.PositionResponse{
				ID:   e.Position.ID,
				Name: e.Position.Name,
				Code: e.Position.Code,
			},
		}

		getEmployeeResponse = append(getEmployeeResponse, em)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to employee",
		"data":    getEmployeeResponse,
	})
}

func GetEmployeesByID(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee

	data := config.DB.Preload(clause.Associations).First(&employee, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	em := models.GetEmployeeResponse{
		ID:         employee.ID,
		Name:       employee.Name,
		Address:    employee.Address,
		Email:      employee.Email,
		PositionID: employee.PositionID,
		Position: models.PositionResponse{
			ID:   employee.Position.ID,
			Name: employee.Position.Name,
			Code: employee.Position.Code,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    em,
	})

}

func PostEmployees(c *gin.Context) {
	var employee models.Employee
	c.BindJSON(&employee)

	config.DB.Create(&employee)

	c.JSON(http.StatusCreated, gin.H{
		"data":    employee,
		"message": "Insert Success",
	})
}

func PutEmployees(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	var reqEmployee models.Employee
	c.BindJSON(&reqEmployee)

	config.DB.Model(&employee).Where("id = ?", id).
		Updates(reqEmployee)

	c.JSON(200, gin.H{
		"message": "Updated",
		"data":    employee,
	})

}

func DeleteEmployees(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee

	data := config.DB.First(&employee, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	config.DB.Delete(&employee, id)

	c.JSON(200, gin.H{
		"message": "Delete Success",
	})
}
