package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var (
	db *gorm.DB
)

// Product model
type Product struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	// Setup database connection
	dsn := "host=localhost user=postgres password=your_password dbname=your_database port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate the Product struct to database
	err = db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Setup Gin router
	r := gin.Default()

	// Routes
	r.GET("/products", getProducts)
	r.POST("/products", createProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)

	// Run the server
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// Handler to get all products
func getProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(200, products)
}

// Handler to create a new product
func createProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&product)
	c.JSON(201, product)
}

// Handler to update a product by ID
func updateProduct(c *gin.Context) {
	var product Product
	id := c.Param("id")

	if err := db.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Save(&product)
	c.JSON(200, product)
}

// Handler to delete a product by ID
func deleteProduct(c *gin.Context) {
	var product Product
	id := c.Param("id")

	if err := db.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	db.Delete(&product)
	c.JSON(204, gin.H{})
}
