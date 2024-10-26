package controllers

import (
	"bytes"
	"ecommerce/config"
	"ecommerce/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	router := gin.Default()
	router.POST("/products", CreateProduct)

	product := models.Product{Name: "Test Product", Price: 100}
	jsonValue, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// // Print the response code and body for debugging
	// t.Logf("Response code: %d", w.Code)
	// t.Logf("Response body: %s", w.Body.String())

	// // Check if the response code is as expected
	// assert.Equal(t, http.StatusOK, w.Code)

	// var response map[string]interface{}
	// err := json.Unmarshal(w.Body.Bytes(), &response)
	// if err != nil {
	// 	t.Fatalf("Failed to unmarshal response: %v", err)
	// }

	// // Check if the "data" field exists in the response
	// data, ok := response["data"].(map[string]interface{})
	// if !ok {
	// 	t.Fatalf("Expected 'data' field in response")
	// }

	// // Assert the product name
	// assert.Equal(t, product.Name, data["name"])
}

func TestGetProducts(t *testing.T) {
	router := gin.Default()
	router.GET("/products", GetProducts)

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotEmpty(t, response["data"])
}

func TestGetProduct(t *testing.T) {
	router := gin.Default()
	router.GET("/products/:id", GetProduct)

	product := models.Product{Name: "Test Product", Price: 100}
	config.DB.Create(&product)

	req, _ := http.NewRequest("GET", "/products/"+strconv.FormatUint(uint64(product.ID), 10), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, product.Name, response["data"].(map[string]interface{})["name"])
}

func TestUpdateProduct(t *testing.T) {
	router := gin.Default()
	router.PUT("/products/:id", UpdateProduct)

	product := models.Product{Name: "Test Product", Price: 100}
	config.DB.Create(&product)

	updatedProduct := models.Product{Name: "Updated Product", Price: 150}
	jsonValue, _ := json.Marshal(updatedProduct)
	req, _ := http.NewRequest("PUT", "/products/"+strconv.FormatUint(uint64(product.ID), 10), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, updatedProduct.Name, response["data"].(map[string]interface{})["name"])
}

func TestDeleteProduct(t *testing.T) {
	router := gin.Default()
	router.DELETE("/products/:id", DeleteProduct)

	product := models.Product{Name: "Test Product", Price: 100}
	config.DB.Create(&product)

	req, _ := http.NewRequest("DELETE", "/products/"+strconv.FormatUint(uint64(product.ID), 10), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.True(t, response["data"].(bool))
}
