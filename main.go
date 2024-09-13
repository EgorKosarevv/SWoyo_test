package main

import (
	"SWOYO/controllers"
	"SWOYO/database"
	"SWOYO/store"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Парсинг флага для выбора хранилища
	useDB := flag.Bool("d", false, "Use database storage")
	flag.Parse()

	// Инициализация хранилища
	var storage store.Store
	if *useDB {
		db, err := database.Connect()
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		storage = store.NewDBStore(db)
	} else {
		storage = store.NewMemoryStore()
	}

	// Инициализация контроллера
	controller := controllers.NewURLController(storage)

	// Инициализация роутера Gin
	r := gin.Default()

	// Определение маршрутов
	r.POST("/", controller.HandlePost)
	r.GET("/:shortURL", controller.HandleGet)

	log.Println("Server is running on port 8080")
	log.Fatal(r.Run(":8080"))
}
