package main

// User ...
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// RecipeCategory ...
type RecipeCategory struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Name     string `json:"name"`
}

// RecipeCategoryRecipe ...
type RecipeCategoryRecipe struct {
	RecipeCategoryID int `json:"recipe_category_id"`
	RecipeID         int `json:"recipe_id"`
}

// Recipe ...
type Recipe struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AuthorID    int    `json:"author_id"`
}

// RecipeResponse ...
type RecipeResponse struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Author      string         `json:"author"`
	Category    []string       `json:"category"`
	Steps       []StepResponse `json:"steps"`
}

// StepResponse ...
type StepResponse struct {
	StepNumber  int                  `json:"step_number"`
	Description string               `json:"description"`
	Timer       int                  `json:"timer"`
	Image       string               `json:"image"`
	Ingredients []IngredientResponse `json:"ingredients"`
}

// IngredientResponse ...
type IngredientResponse struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}

// IngredientCategory ...
type IngredientCategory struct {
	ID          int    `json:"id"`
	ParentID    int    `json:"parent_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// IngredientCategoryIngredient ...
type IngredientCategoryIngredient struct {
	IngredientCategoryID int `json:"ingredient_category_id"`
	IngredientID         int `json:"ingredient_id"`
}

// Ingredient ...
type Ingredient struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	IMG   string `json:"img"`
}

// Step ...
type Step struct {
	ID          int    `json:"id"`
	RecipeID    int    `json:"recipe_id"`
	StepNumber  int    `json:"step_number"`
	Description string `json:"description"`
	Timer       int    `json:"timer"`
	Image       string `json:"image"`
}

// StepIngredient ...
type StepIngredient struct {
	RecipeID     int    `json:"recipe_id"`
	IngredientID int    `json:"ingredient_id"`
	StepID       int    `json:"step_id"`
	Amount       int    `json:"amount"`
	Unit         string `json:"unit"`
}
