package controllers

import (
	"GO-1/config"
	"GO-1/models"
	"GO-1/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// hashing password

	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)

	_, err := config.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(500, gin.H{"erorr:": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "user berhasil dibuat"})
}

func Login(c *gin.Context) {
	var input models.User // Gunakan nama 'input' agar beda dengan 'foundUser'

	// 1. Gunakan ShouldBindJSON untuk menangkap data dari Postman
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
		return
	}

	// 2. CEK: Pastikan data tidak kosong sebelum nembak DB
	if input.Username == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username dan Password wajib diisi"})
		return
	}

	var foundUser models.User
	// 3. CARI USER BERDASARKAN USERNAME
	filter := bson.M{"username": input.Username}

	err := config.UserCollection.FindOne(context.TODO(), filter).Decode(&foundUser)

	if err != nil {
		// Jika error, berarti username memang tidak ada di DB
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username tidak ditemukan"})
		return
	}

	// 4. BANDINGKAN PASSWORD
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}
	token, err := utils.GenerateToken(foundUser.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "gada token"})
		return
	}

	// 5. BERHASIL LOGIN
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Berhasil",
		"token":   token,
	})
}
