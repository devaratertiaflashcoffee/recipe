package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InsertRecipeCategory ...
func InsertRecipeCategory(c *gin.Context) {
	recipeCategory := RecipeCategory{}

	if err := c.BindJSON(&recipeCategory); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	}

	if err := DB.Table("recipe_category").Create(&recipeCategory).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, recipeCategory)
		log.Println(recipeCategory)
	}
}
