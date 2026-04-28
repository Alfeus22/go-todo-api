package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func MyMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			fmt.Println("DEBUG: Header kosong")
			c.JSON(401, gin.H{"error": "Header tidak ditemukan"})
			c.Abort()
			return
		}

		if token != "RahasiaTod" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. ambil header authorization : Bearer <token>
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Akses di blokir"})
			ctx.Abort()
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// 2 validasi token
		token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// validasi kalau ethodnya beneran HS256
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Method signing salah : %v", t.Header["alg"])
			}
			return []byte("RahasiaDapurGalih22"), nil
		})

		// ... kode parse token kamu sebelumnya ...

		if token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok {
				// AMBIL user_id dari claims token
				userID := claims["user_id"].(string)

				// SIMPAN ke context Gin agar bisa dipanggil di controller mana pun
				ctx.Set("currentUser", userID)

				ctx.Next()
			}
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			ctx.Abort()
		}
	}
}
