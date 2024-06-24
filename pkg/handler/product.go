package handler

import (
	"net/http"
	"strconv"

	"github.com/SJMaks/products-api"
	"github.com/gin-gonic/gin"
)

type getProductsResponse struct {
	Data []products.Product `json:"data"`
}

func (h *Handler) getProducts(c *gin.Context) {
	products, err := h.services.Product.GetProducts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getProductsResponse{
		Data: products,
	})
}

func (h *Handler) getProduct(c *gin.Context) {
	input := c.Param("i")

	id, err := strconv.Atoi(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	product, err := h.services.Product.GetProduct(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":         product.Id,
		"name":       product.Name,
		"quantity":   product.Quantity,
		"unit_coast": product.Unit_coast,
		"measure":    product.Measure,
	})
}

func (h *Handler) createProduct(c *gin.Context) {
	var input products.Product

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Product.CreateProduct(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateProduct(c *gin.Context) {
	var input products.Product

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	param := c.Param("i")

	id, err := strconv.Atoi(param)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Product.UpdateProduct(input, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) deleteProduct(c *gin.Context) {
	input := c.Param("i")

	id, err := strconv.Atoi(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Product.DeleteProduct(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
