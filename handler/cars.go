package handler

import (
	"go-rest-api/model"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var mu sync.Mutex

func ListProduct(db map[string]model.Car) gin.HandlerFunc {
	return func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()

		products := make([]model.Car, 0, len(db))
		for _, product := range db {
			products = append(products, product)
		}

		c.JSON(200, products)
	}
}

func GetProduct(db map[string]model.Car) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		mu.Lock()
		defer mu.Unlock()

		product, exists := db[id]
		if !exists {
			log.Printf("Product not found: %s", id)
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}

		c.JSON(200, product)
	}
}

func CreateProduct(db map[string]model.Car) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product model.Car
		if err := c.BindJSON(&product); err != nil {
			log.Printf("Product fields mismatch: %v", err)
			c.JSON(400, gin.H{"error": "Product not valid"})
			return
		}

		product.ID = uuid.New().String()

		mu.Lock()
		db[product.ID] = product
		mu.Unlock()

		c.JSON(201, product)
	}
}

func UpdateProduct(db map[string]model.Car) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var productReq model.Car
		if err := c.BindJSON(&productReq); err != nil {
			log.Printf("Product fields mismatch: %v", err)
			c.JSON(400, gin.H{"error": "Product not valid"})
			return
		}

		mu.Lock()
		defer mu.Unlock()

		product, exists := db[id]
		if !exists {
			log.Printf("Product not found: %s", id)
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}

		if productReq.Brand != "" {
			product.Brand = productReq.Brand
		}

		if productReq.HorsePower != 0 {
			product.HorsePower = productReq.HorsePower
		}

		db[id] = product
		c.JSON(200, product)
	}
}

func DeleteProduct(db map[string]model.Car) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		mu.Lock()
		defer mu.Unlock()

		_, exists := db[id]
		if !exists {
			log.Printf("Product not found: %s", id)
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}

		delete(db, id)
		c.JSON(204, gin.H{"message": "Product deleted"})
	}
}
