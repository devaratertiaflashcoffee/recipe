package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB ...
var DB *gorm.DB

func main() {

	r := gin.Default()

	db, err := gorm.Open("sqlite3", "recipe.db")
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	defer db.Close()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"greet": "hello, world!",
		})
	})

	r.GET("/echo/:echo", func(c *gin.Context) {
		echo := c.Param("echo")
		c.JSON(http.StatusOK, gin.H{
			"echo": echo,
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			// c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"uploaded": len(files),
		})
	})

	r.POST("/user", InsertAuthor)

	r.POST("/recipe-category", InsertRecipeCategory)

	r.GET("/recipe", GetRecipe)

	r.POST("/recipe", InsertRecipe)

	r.POST("/ingredient", InsertIngredient)

	r.POST("/ingredient-category", InsertIngredientCategory)

	r.POST("/step", InsertStep)

	r.Run() // listen and serve on 0.0.0.0:8080
}
