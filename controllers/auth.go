package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Auth(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name, pss, ok := ctx.Request.BasicAuth()
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Username dan Password empty",
			})
			return
		}

		if name == "admin" && pss == "admin" {
			var id int
			err := db.QueryRow("SELECT id FROM users WHERE username = $1", name).Scan(&id)
			if err != nil {
				now := time.Now()
				_, err := db.Exec(`
					INSERT INTO users (username, password, created_at, created_by, modified_at, modified_by)
					VALUES ($1, $2, $3, $4, $5, $6)
				`, name, pss, now, name, now, name)
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error": "Error to save users",
					})
					return
				}
			}

			ctx.Set("username", name)
			ctx.Next()
			return
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
	}
}
