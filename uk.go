package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Model structs
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"uniqueIndex;not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Age       int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Photo struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Caption   string    `gorm:"not null"`
	PhotoURL  string    `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	PhotoID   uint      `gorm:"not null"`
	Message   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type SocialMedia struct {
	ID             uint      `gorm:"primaryKey"`
	Name           string    `gorm:"not null"`
	SocialMediaURL string    `gorm:"not null"`
	UserID         uint      `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

// JWT Signing Key (for simplicity, hardcoded here; in production, use environment variable)
var jwtSigningKey = []byte("your_jwt_signing_key")

// Middleware untuk validasi JWT
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// TODO: Validate JWT token (skipped for brevity)
		// Example using jwt-go library:
		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//     return jwtSigningKey, nil
		// })

		// if err != nil || !token.Valid {
		//     c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		//     c.Abort()
		//     return
		// }

		// Jika token valid, lanjutkan ke handler
		c.Next()
	}
}

// Setup database connection
func setupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("mygram.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate schema
	err = db.AutoMigrate(&User{}, &Photo{}, &Comment{}, &SocialMedia{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	return db
}

func main() {
	// Setup Gin router
	r := gin.Default()

	// Setup database
	db := setupDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Routes
	v1 := r.Group("/api")
	{
		// Registrasi Pengguna
		v1.POST("/register", func(c *gin.Context) {
			// TODO: Implement registration endpoint
			c.JSON(http.StatusOK, gin.H{"message": "Implement registration endpoint"})
		})

		// Login Pengguna
		v1.POST("/login", func(c *gin.Context) {
			// TODO: Implement login endpoint
			c.JSON(http.StatusOK, gin.H{"message": "Implement login endpoint"})
		})

		// Middleware untuk validasi JWT
		v1.Use(authMiddleware())

		// Tambah Foto
		v1.POST("/photos", func(c *gin.Context) {
			// TODO: Implement add photo endpoint
			c.JSON(http.StatusOK, gin.H{"message": "Implement add photo endpoint"})
		})

		// Tambah Komentar pada Foto
		v1.POST("/photos/:photo_id/comments", func(c *gin.Context) {
			// TODO: Implement add comment endpoint
			c.JSON(http.StatusOK, gin.H{"message": "Implement add comment endpoint"})
		})

		// Update Data Foto
		v1.PUT("/photos/:photo_id", func(c *gin.Context) {
			// TODO: Implement update photo endpoint
			c.JSON(http.StatusOK, gin.H{"message": "Implement update photo endpoint"})
		})

		// Hapus Data Foto
		v1.DELETE("/photos/:photo_id", func(c *gin.Context) {
			// TODO: Implement delete photo endpoint
			c.JSON(http.StatusOK, gin.H{"message": "Implement delete photo endpoint"})
		})
	}

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
