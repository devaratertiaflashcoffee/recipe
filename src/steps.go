package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StepInsert ...
type StepInsert struct {
	ID              int              `json:"id"`
	RecipeID        int              `json:"recipe_id"`
	StepNumber      int              `json:"step_number"`
	Description     string           `json:"description"`
	Timer           int              `json:"timer"`
	Image           string           `json:"image"`
	StepIngredients []StepIngredient `json:"step_ingredients"`
}

// InsertStep ...
func InsertStep(c *gin.Context) {
	step := StepInsert{}

	if err := c.BindJSON(&step); err != nil {
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

	stepDB := Step{
		RecipeID:    step.RecipeID,
		StepNumber:  step.StepNumber,
		Description: step.Description,
		Timer:       step.Timer,
		Image:       step.Image,
	}

	if err := tx.Table("step").Create(&stepDB).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	}

	step.ID = stepDB.ID

	for k, v := range step.StepIngredients {
		step.StepIngredients[k].RecipeID = step.RecipeID
		step.StepIngredients[k].StepID = step.ID
		ingredient := &StepIngredient{
			RecipeID:     step.RecipeID,
			IngredientID: v.IngredientID,
			StepID:       step.ID,
			Amount:       v.Amount,
			Unit:         v.Unit,
		}
		if err := tx.Table("step_ingredients").Create(ingredient).Error; err != nil {
			tx.Rollback()
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, step)
		log.Println(step)
	}
}
