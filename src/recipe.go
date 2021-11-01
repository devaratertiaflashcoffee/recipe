package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RecipeInsert ...
type RecipeInsert struct {
	ID          int    `json:"id"`
	CategoryIDs []int  `json:"category_ids"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AuthorID    int    `json:"author_id"`
}

// InsertRecipe ...
func InsertRecipe(c *gin.Context) {
	recipe := RecipeInsert{}

	if err := c.BindJSON(&recipe); err != nil {
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

	recipeDB := Recipe{
		Name:        recipe.Name,
		Description: recipe.Description,
		AuthorID:    recipe.AuthorID,
	}

	if err := tx.Table("recipe").Create(&recipeDB).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	}

	recipe.ID = recipeDB.ID

	for _, v := range recipe.CategoryIDs {
		if err := tx.Table("recipe_category_recipe").Create(&RecipeCategoryRecipe{RecipeCategoryID: v, RecipeID: recipe.ID}).Error; err != nil {
			tx.Rollback()
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, recipe)
		log.Println(recipe)
	}
}

// GetRecipe ...
func GetRecipe(c *gin.Context) {
	var recipes []Recipe
	var recipesResponse []RecipeResponse

	if err := DB.Table("recipe").Find(&recipes).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	}

	for _, v := range recipes {
		recipeResponse := RecipeResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
		}

		userData := User{}
		if err := DB.Table("users").Where("id = " + strconv.FormatInt(int64(v.AuthorID), 10)).Find(&userData).Error; err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
		}

		recipeResponse.Author = userData.Name

		categories := []RecipeCategory{}
		steps := []Step{}
		idStr := strconv.FormatInt(int64(v.ID), 10)
		DB.Table("recipe_category").
			Select("recipe_category.name").
			Joins("left join recipe_category_recipe on recipe_category_recipe.recipe_category_id = recipe_category.id").
			Where("recipe_category_recipe.recipe_id = " + idStr).
			Scan(&categories)

		if err := DB.Table("step").Where("recipe_id = " + strconv.FormatInt(int64(v.ID), 10)).Find(&steps).Error; err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
		}

		for _, v := range categories {
			recipeResponse.Category = append(recipeResponse.Category, v.Name)
		}

		for _, v1 := range steps {
			var ingredientResponse []IngredientResponse
			stepIDStr := strconv.FormatInt(int64(v1.ID), 10)
			DB.Table("step_ingredients").
				Select("ingredient.name, step_ingredients.amount, step_ingredients.unit").
				Joins("left join ingredient on ingredient.id = step_ingredients.ingredient_id").
				Where("step_ingredients.step_id = " + stepIDStr).
				Scan(&ingredientResponse)
			recipeResponse.Steps = append(recipeResponse.Steps, StepResponse{
				StepNumber:  v1.StepNumber,
				Description: v1.Description,
				Timer:       v1.Timer,
				Image:       v1.Image,
				Ingredients: ingredientResponse,
			})
		}
		recipesResponse = append(recipesResponse, recipeResponse)
	}

	c.JSON(http.StatusOK, recipesResponse)
	log.Println("OK")
}
