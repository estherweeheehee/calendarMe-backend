package controller

import (
	"net/http"

	"calendarme-backend/config"
	"calendarme-backend/models"

	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	notes := []models.Note{}
	config.DB.Find(&notes)
	c.JSON(http.StatusOK, &notes)
}



func FindNote(c *gin.Context) {
	// c.Header("Access-Control-Allow-Origin", "http://localhost:3001")
	// c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
    // c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	notes := []models.Note{}
	config.DB.Where("month = ?", c.Param("month")).Find(&notes)
	if len(notes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	} else {	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Found",
		"data":    &notes,
	})
	}
}

func CreateNote(c *gin.Context) {
	// c.Header("Access-Control-Allow-Origin", "http://localhost:3001")
	// c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
    // c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	note := models.Note{}
	c.BindJSON(&note)
	config.DB.Create(&note)
	c.JSON(http.StatusOK, &note)

}

func DeleteNote(c *gin.Context) {
	note := models.Note{}
	config.DB.Where("id = ?", c.Param("id")).Delete(&note)
	c.JSON(http.StatusOK, &note)
}

func UpdateNote(c *gin.Context) {
	note := models.Note{}
	config.DB.Where("id = ?", c.Param("id")).First(&note)
	c.BindJSON(&note)
	config.DB.Save(&note)
	c.JSON(http.StatusOK, &note)
}

func SearchTags(c *gin.Context) {
	notes := []models.Note{}
	config.DB.Where("tag = ?", c.Param("tag")).Find(&notes)
	if len(notes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	} else {	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Found",
		"data":    &notes,
	})
	}
}