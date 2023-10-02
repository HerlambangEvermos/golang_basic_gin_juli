package main

import (
	"golang_basic_gin_juli_2023/config"
	middlewares "golang_basic_gin_juli_2023/midlewares"
	"golang_basic_gin_juli_2023/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/home", getHome)

	// /api/v1/department

	api := r.Group("/api/v1")
	{

		user := api.Group("/user")
		{
			user.POST("/register", routes.RegisterUser)
			user.POST("/login", routes.GenerateToken)
		}

		department := api.Group("/departments").Use(middlewares.Auth())
		{
			department.GET("/", routes.GetDepartment)
			department.GET("/:id", routes.GetDepartmentById)
			department.POST("/", routes.PostDepartment)
			department.PUT("/:id", routes.PutDepartment)
			department.DELETE("/:id", routes.DeleteDepartment)

		}

		position := api.Group("/positions").Use(middlewares.Auth())
		{
			position.GET("/", routes.GetPositions)
			position.GET("/:id", routes.GetPositionsById)
			position.POST("/", routes.PostPositions)
			position.PUT("/:id", routes.PutPositions)
			position.DELETE("/:id", routes.DeletePositions)
		}

		employee := api.Group("/employees").Use(middlewares.Auth())
		{
			employee.GET("/", routes.GetEmployees)
			employee.GET("/:id", routes.GetEmployeesByID)
			employee.POST("/", routes.PostEmployees)
			employee.PUT("/:id", routes.PutEmployees)
			employee.DELETE("/:id", routes.DeleteEmployees)
		}

		inventories := api.Group("/inventories").Use(middlewares.Auth())
		{
			inventories.GET("/", routes.GetInventories)
			inventories.GET("/:id", routes.GetInventoriesByID)
			inventories.POST("/", routes.PostInventories)
			inventories.PUT("/:id", routes.PutInventories)
			inventories.DELETE("/:id", routes.DeleteInventories)
		}

		rental := api.Group("/rental").Use(middlewares.Auth())
		{
			rental.GET("/", routes.GetRental)
			rental.POST("/employee", routes.RentalByEmployeeId)
			rental.GET("/inventory/:id", routes.GetRentalByInventoryID)
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}

func getHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome home",
	})

}
