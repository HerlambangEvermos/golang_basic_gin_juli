package routes

import (
	"golang_basic_gin_juli_2023/config"
	"golang_basic_gin_juli_2023/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetInventories(c *gin.Context) {
	inventory := []models.Inventory{}
	config.DB.Preload(clause.Associations).Find(&inventory)

	resInventories := []models.ResponseInventory{}

	for _, inv := range inventory {
		resInv := models.ResponseInventory{
			InventoryName:        inv.Name,
			InventoryDescription: inv.Description,
			Archive: models.ResponseArchive{
				ArchiveName:        inv.Archive.Name,
				ArchiveDescription: inv.Archive.Description,
			},
		}

		resInventories = append(resInventories, resInv)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get inventory",
		"data":    resInventories,
	})
}

func GetInventoriesByID(c *gin.Context) {
	id := c.Param("id")

	var inventory models.Inventory

	// data := config.DB.Preload("Department").First(&position, "id = ?", id)

	data := config.DB.Preload(clause.Associations).First(&inventory, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    inventory,
	})
}

func PostInventories(c *gin.Context) {
	// ambil data post dari json
	var reqInv models.RequestInventory
	c.BindJSON(&reqInv)

	inventory := models.Inventory{
		Name:        reqInv.InventoryName,
		Description: reqInv.InventoryDescription,
		Archive: models.Archive{
			Name:        reqInv.ArchiveName,
			Description: reqInv.InventoryDescription,
		},
	}

	// insert data to db
	config.DB.Create(&inventory)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully Create inventory",
		"data":    inventory,
	})
}

func PutInventories(c *gin.Context) {
	id := c.Param("id")

	data := config.DB.First(&models.Inventory{}, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Inventory not found",
		})

		return
	}

	var reqInv models.RequestInventory
	c.BindJSON(&reqInv)

	inventory := models.Inventory{
		Name:        reqInv.InventoryName,
		Description: reqInv.InventoryDescription,
	}
	config.DB.Model(&inventory).Where("id = ?", id).Updates(&inventory)

	archive := models.Archive{
		Name:        reqInv.ArchiveName,
		Description: reqInv.InventoryDescription,
	}
	config.DB.Model(&archive).Where("inventory_id = ?", id).Updates(&archive)

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Successfully",
		"data":    inventory,
	})

}

func DeleteInventories(c *gin.Context) {
	id := c.Param("id")

	var inventory models.Inventory

	data := config.DB.First(&inventory, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	config.DB.Delete(&inventory, id)

	c.JSON(200, gin.H{
		"message": "Delete Success",
	})
}
