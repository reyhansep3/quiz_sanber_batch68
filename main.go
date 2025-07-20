package main

import (
	"database/sql"
	"fmt"
	"os"

	"quiz_sanber_batch68/controllers"
	database "quiz_sanber_batch68/databases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")

	if err != nil {
		panic("error load .env")
	}

	psqlInfo := fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("successfully connected to database")
	database.DBMigrate(db)

	router := gin.Default()
	router.POST("/api/books", controllers.Auth(db), controllers.PostBooks(db))
	router.GET("/api/books", controllers.Auth(db), controllers.GetBooks(db))
	router.GET("/api/books/:id", controllers.Auth(db), controllers.GetBooksByID(db))
	router.DELETE("/api/books/:id", controllers.Auth(db), controllers.DeleteBooksByID(db))

	router.POST("/api/categories", controllers.Auth(db), controllers.PostCategories(db))
	router.GET("/api/categories", controllers.Auth(db), controllers.GetCategories(db))
	router.GET("/api/categories/:id", controllers.Auth(db), controllers.GetCategoryByID(db))
	router.DELETE("/api/categories/:id", controllers.Auth(db), controllers.DeleteCategoryByID(db))

	router.Run(":8080")

}
