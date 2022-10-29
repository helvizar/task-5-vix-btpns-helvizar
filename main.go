package main

import (
	"fmt"
	"golang-project/controllers"
	"golang-project/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&models.Note{})
	if err != nil {
		panic(err.Error())
	}

	noteControllers := &controllers.NoteControllers{}

	router := httprouter.New()

	router.GET("/", noteControllers.Index)
	router.GET("/create", noteControllers.Create)
	router.POST("/create", noteControllers.Create)
	router.GET("/edit/:id", noteControllers.Edit)
	router.POST("/update/:id", noteControllers.Update)
	router.POST("/done/:id", noteControllers.Done)
	router.POST("/delete/:id", noteControllers.Delete)

	port := ":3000"
	fmt.Println("Aplikasi jalan di: http://localhost:3000")
	log.Fatal(http.ListenAndServe(port, router))

}
