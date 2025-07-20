package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"quiz_sanber_batch68/structs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PostCategories(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var category structs.Category
		if err := ctx.ShouldBindJSON(&category); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		username := ctx.GetString("username")
		now := time.Now()
		category.CreatedAt = &now
		category.CreatedBy = username
		category.ModifiedAt = &now
		category.ModifiedBy = username

		sqlStatement := `
		INSERT INTO categories (name, created_at, created_by, modified_at, modified_by)
		VALUES ($1, $2, $3, $4, $5)
		Returning id, created_at;
		`

		err := db.QueryRow(sqlStatement, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy).
			Scan(&category.ID, &category.CreatedAt)
		if err != nil {
			fmt.Printf("PostCategories DB error: %v\n", err)

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"categories": category,
		})
	}
}

func GetCategories(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var categories []structs.Category

		sqlStatement := `SELECT * FROM categories`

		rows, err := db.Query(sqlStatement)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			var category structs.Category

			err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.CreatedBy)
			if err != nil {
				panic(err)
			}
			categories = append(categories, category)
		}
		fmt.Println("category data :", categories)
		ctx.JSON(http.StatusOK, gin.H{
			"categories": categories,
		})
	}
}

func GetCategoryByID(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParams := ctx.Param("id")

		var category structs.Category

		sqlStatement := `SELECT * FROM categories WHERE id = $1`

		id, err := strconv.Atoi(idParams)

		if err != nil {
			panic(err)
		}

		err = db.QueryRow(sqlStatement, id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.CreatedBy)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"categories": category,
		})
	}
}

func DeleteCategoryByID(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParams := ctx.Param("id")

		sqlStatement := `DELETE FROM categories WHERE id = $1`

		id, err := strconv.Atoi(idParams)
		if err != nil {
			panic(err)
		}
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
				"error": "The category you want to delete does not exist",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": "data succesfully deleted"})
		}

	}
}
