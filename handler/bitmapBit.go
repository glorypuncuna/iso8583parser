package handler

import (
	"mappertest/formatter"
	"mappertest/helper"
	"mappertest/input"
	"mappertest/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bitmapBitHandler struct {
	bitmapBitService service.BitmapBitService
}

func NewBitmapBitHandler(bitmapBitService service.BitmapBitService) *bitmapBitHandler {
	return &bitmapBitHandler{bitmapBitService}
}

func (h *bitmapBitHandler) ResultField(c *gin.Context) {
	var input input.BitmapBitInput
	err := c.ShouldBindUri(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Cannot process your request. Bad request",
			422, "Unprocesable Entity", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	field, bit := h.bitmapBitService.DetermineField(input)
	res := formatter.BitmapFormatter{
		Field: field,
		Bit:   bit,
	}
	c.JSON(http.StatusOK, helper.APIResponse("Data has been proceed sucessfully",
		200, "success", res))
}
