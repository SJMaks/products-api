package handler

import (
	"net/http"
	"strconv"

	"github.com/SJMaks/products-api"
	"github.com/gin-gonic/gin"
)

type getMeasuresResponse struct {
	Data []products.Measure `json:"data"`
}

func (h *Handler) getMeasures(c *gin.Context) {
	measures, err := h.services.Measure.GetMeasures()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getMeasuresResponse{
		Data: measures,
	})
}

func (h *Handler) getMeasure(c *gin.Context) {
	input := c.Param("i")

	id, err := strconv.Atoi(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	measure, err := h.services.Measure.GetMeasure(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":   measure.Id,
		"name": measure.Name,
	})
}

func (h *Handler) createMeasure(c *gin.Context) {
	var input products.Measure

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Measure.CreateMeasure(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) updateMeasure(c *gin.Context) {
	var input products.Measure

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	param := c.Param("i")

	id, err := strconv.Atoi(param)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Measure.UpdateMeasure(input, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) deleteMeasure(c *gin.Context) {
	input := c.Param("i")

	id, err := strconv.Atoi(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Measure.DeleteMeasure(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
