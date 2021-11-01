package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InsertAuthor ...
func InsertAuthor(c *gin.Context) {
	user := User{}

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	}

	if err := DB.Table("users").Create(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
		log.Println(user)
	}
}
