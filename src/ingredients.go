package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IngredientInsert ...
type IngredientInsert struct {
	ID          int    `json:"id"`
	CategoryIDs []int  `json:"category_ids"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	IMG         string `json:"img"`
}

// InsertIngredient ...
func InsertIngredient(c *gin.Context) {
	ingredient := IngredientInsert{}

	if err := c.BindJSON(&ingredient); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	}

	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	}

	ingredientDB := Ingredient{
		Name:  ingredient.Name,
		Color: ingredient.Color,
		IMG:   ingredient.IMG,
	}

	if err := tx.Table("ingredient").Create(&ingredientDB).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	}

	ingredient.ID = ingredientDB.ID

	for _, v := range ingredient.CategoryIDs {
		if err := tx.Table("ingredient_category_ingredient").Create(&IngredientCategoryIngredient{IngredientCategoryID: v, IngredientID: ingredient.ID}).Error; err != nil {
			tx.Rollback()
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, ingredient)
		log.Println(ingredient)
	}
}

// InsertIngredientCategory ...
func InsertIngredientCategory(c *gin.Context) {
	ingredientCategory := IngredientCategory{}

	if err := c.BindJSON(&ingredientCategory); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	}

	if err := DB.Table("ingredient_category").Create(&ingredientCategory).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, ingredientCategory)
		log.Println(ingredientCategory)
	}
}
