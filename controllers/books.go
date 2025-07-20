package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"quiz_sanber_batch68/structs"

	"github.com/gin-gonic/gin"
)

func PostBooks(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newBook structs.Book

		if err := ctx.ShouldBindJSON(&newBook); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var maxCategoryID int
		err := db.QueryRow("SELECT COALESCE(MAX(category_id), 0) FROM books").Scan(&maxCategoryID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get max category_id"})
			return
		}
		newBook.CategoryID = maxCategoryID + 1

		username := ctx.GetString("username")

		now := time.Now()
		newBook.CreatedAt = &now
		newBook.CreatedBy = username
		newBook.ModifiedAt = &now
		newBook.ModifiedBy = username

		if newBook.TotalPage < 100 {
			newBook.Thickness = "tipis"
		} else {
			newBook.Thickness = "tebal"
		}

		if newBook.ReleaseYear <= 1980 && newBook.ReleaseYear >= 2024 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be in between 1980 to 2024"})
			return
		}

		sqlStatement := `
		INSERT INTO books 
			(title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id, created_at;
		`

		err = db.QueryRow(sqlStatement,
			newBook.Title,
			newBook.Description,
			newBook.ImageURL,
			newBook.ReleaseYear,
			newBook.Price,
			newBook.TotalPage,
			newBook.Thickness,
			newBook.CategoryID,
			newBook.CreatedAt,
			newBook.CreatedBy,
			newBook.ModifiedAt,
			newBook.ModifiedBy,
		).Scan(&newBook.ID, &newBook.CreatedAt)

		if err != nil {
			fmt.Println("Error inserting book:", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"books": newBook,
		})
	}
}

func GetBooks(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var getBook []structs.Book

		sqlStatement := `SELECT * from books`
		rows, err := db.Query(sqlStatement)

		if err != nil {
			panic(err)
		}

		defer rows.Close()

		for rows.Next() {
			var book = structs.Book{}

			err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)

			if err != nil {
				panic(err)
			}
			getBook = append(getBook, book)
		}

		fmt.Println("books data", getBook)
		ctx.JSON(http.StatusOK, gin.H{
			"books": getBook,
		})

	}
}

func GetBooksByID(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParams := ctx.Param("id")

		var book structs.Book

		sqlStatement := `SELECT * FROM books WHERE id = $1`

		id, err := strconv.Atoi(idParams)
		if err != nil {
			panic(err)
		}

		err = db.QueryRow(sqlStatement, id).Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"books": book,
		})

	}
}

func DeleteBooksByID(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParams := ctx.Param("id")
		var books structs.Book

		if err := ctx.ShouldBindJSON(&books); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"error": err,
			})
		}

		id, err := strconv.Atoi(idParams)
		if err != nil {
			panic(err)
		}

		sqlStatement := `DELETE FROM books WHERE id = $1`
		hasil, err := db.Exec(sqlStatement, id)
		if err != nil {
			panic(err)
		}
		rowsAffected, err := hasil.RowsAffected()
		if err != nil {
			panic(err)
		}

		if rowsAffected == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "The book you want to delete does not exist",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": "data has deleted"})
		}

	}
}
