package main

import (
	"GO-1/config"
	"GO-1/routes"
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin" // Hanya panggil satu kali
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Pastikan variabel ini memang dibutuhkan di file main.go,
// jika tidak, lebih baik dipindah ke folder/package lain (misalnya config/repository).
var taskCollection *mongo.Collection

func main() {
	// 1. Load variabel dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, menggunakan variabel environment sistem")
	}

	// 2. Ambil string dari .env
	frontendUrlsRaw := os.Getenv("FRONTEND_URLS")

	// Jika kosong, berikan fallback default
	if frontendUrlsRaw == "" {
		frontendUrlsRaw = "http://localhost:3000"
	}

	// 3. Pecah string menjadi array/slice string berdasarkan koma
	allowedOrigins := strings.Split(frontendUrlsRaw, ",")

	// Bersihkan spasi ekstra dari tiap URL
	for i := range allowedOrigins {
		allowedOrigins[i] = strings.TrimSpace(allowedOrigins[i])
	}

	// 4. Koneksi ke Database
	config.ConnectDB()

	// 5. Inisialisasi Gin Router
	r := gin.Default()

	// 6. Konfigurasi CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// 7. Daftarkan Routes DI SINI (SEBELUM r.Run)
	r.GET("/api/data", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Berhasil terhubung!"})
	})

	// Panggil fungsi route eksternal di sini
	routes.SetupTaskRouters(r)

	// 8. Ambil port dan Jalankan Server (r.Run HARUS PALING BAWAH)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server berjalan di port:", port)
	r.Run(":" + port)
}
