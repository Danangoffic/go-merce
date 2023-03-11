package controllers

import (
	"net/http"
	"strconv"

	"github.com/Danangoffic/go-merce/app/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	db *gorm.DB
}

// To init ProductController Class, pass the *gorm.DB interface
func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{db: db}
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	var products []models.Product

	// Query semua produk
	if result := pc.db.Find(&products); result.Error != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "message": "Failed to get Products data", "data": nil})
		return
	}

	// Jika tidak ada produk yang ditemukan
	if len(products) == 0 {
		c.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "message": "No products found", "data": nil})
		return
	}

	// Return produk sebagai response JSON
	c.JSON(http.StatusOK,
		gin.H{"status": http.StatusOK, "message": "success", "data": products})
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
	// Mendapatkan ID dari parameter URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product

	// Query produk dengan ID yang sesuai
	if result := pc.db.First(&product, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Return produk dengan ID tersebut sebagai response JSON
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product

	// Binding data JSON ke model Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request body",
		})
		return
	}

	// Set nilai createdBy dengan user yang sedang login
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed! User not Found",
		})
		return
	}
	currentUser := user.(*models.User)
	createdBy := models.User{ID: currentUser.ID}
	product.CreatedBy = createdBy

	// Tambah data ke database
	if err := pc.db.Create(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to create product",
		})
		return
	}

	// Response dengan data JSON
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusOK,
		"message": "Product created successfully",
		"data":    product,
	})
}
