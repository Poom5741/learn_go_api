package main

import (
	"github.com/Poom5741/learn_go_api/auth"
	"github.com/Poom5741/learn_go_api/todo"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&todo.Todo{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/tokenz", auth.AccessToken("==signature=="))
	protected := r.Group("", auth.Protect([]byte("==signature==")))
	handler := todo.NewTodoHandler(db)
	protected.POST("/todos", handler.NewTask)

	r.Run()
}
