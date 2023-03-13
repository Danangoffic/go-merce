package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Danangoffic/go-merce/app/controllers/request"
	"github.com/Danangoffic/go-merce/app/helpers"
	"github.com/Danangoffic/go-merce/app/models"
	"github.com/Danangoffic/go-merce/app/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	db      *gorm.DB
	Service services.Services
}

// To init ProductController Class, pass the *gorm.DB interface
func NewProductController(db *gorm.DB, Service services.Services) *ProductController {
	return &ProductController{db: db, Service: Service}
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	var products []models.Product

	// Query semua produk
	category, _ := c.GetQuery("category")
	products, err := pc.Service.GetProducts(category)
	if err != nil {
		response := helpers.ResponseJSON(err.Error(), http.StatusNotFound, nil, err)
		c.JSON(http.StatusNotFound, response)
		return
	}

	// Jika tidak ada produk yang ditemukan
	if len(products) == 0 {
		response := helpers.ResponseJSON("Products not available", http.StatusNotFound, nil, errors.New("Products not available"))
		c.JSON(http.StatusNotFound, response)
		return
	}
	// Return produk sebagai response JSON
	response := helpers.ResponseJSON("success", http.StatusOK, products, nil)
	c.JSON(http.StatusOK, response)
	return
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
	// Mendapatkan ID dari parameter URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helpers.ResponseJSON("Invallid product id", http.StatusBadRequest, nil, errors.New("Invalid Product ID"))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var product models.Product

	// Query produk dengan ID yang sesuai
	if result := pc.db.First(&product, id); result.Error != nil {
		response := helpers.ResponseJSON("Product not found", http.StatusNotFound, nil, errors.New("Product not found"))
		c.JSON(http.StatusNotFound, response)
		return
	}

	// Return produk dengan ID tersebut sebagai response JSON
	response := helpers.ResponseJSON("success", http.StatusOK, product, nil)
	c.JSON(http.StatusOK, response)
	return
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var requestProduct request.StoreProduct

	// Binding data JSON ke model Product
	if err := c.ShouldBindJSON(&requestProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request body",
		})
		return
	}

	// Set nilai createdBy dengan user yang sedang login
	user := c.MustGet("createdById")
	currentUser := user.(*models.User)
	createdBy := models.User{ID: currentUser.ID}
	requestProduct.CreatedByID = strconv.Itoa(int(currentUser.ID))
	requestProduct.CreatedBy = createdBy
	// Tambah data ke database
	_, err := pc.Service.CreateProduct(requestProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to create product",
			"data":    nil,
		})
		return
	}

	// Response dengan data JSON
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusOK,
		"message": "Product created successfully",
		"data":    nil,
	})
}
