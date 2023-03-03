package handler

import (
	"mappertest/helper"
	"mappertest/input"
	"mappertest/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type isoHandler struct {
	isoService service.ISOService
}

func NewISOHandler(isoService service.ISOService) *isoHandler {
	return &isoHandler{isoService}
}

func (h *isoHandler) Mode1(c *gin.Context) {
	var input input.ISOInput
	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Cannot process your request. Bad request",
			422, "Unprocesable Entity", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.ModeType = 1
	res := h.isoService.SplitISO(input)
	c.JSON(http.StatusOK, helper.APIResponse("Data has been proceed sucessfully",
		200, "success", res))
}

func (h *isoHandler) Mode2(c *gin.Context) {
	var input input.ISOInput
	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Cannot process your request. Bad request",
			422, "Unprocesable Entity", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.ModeType = 2
	res := h.isoService.ToHex(input)
	c.JSON(http.StatusOK, helper.APIResponse("Data has been proceed sucessfully",
		200, "success", res))
}

func (h *isoHandler) Mode3(c *gin.Context) {

	var input input.ISOInput
	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Cannot process your request. Bad request",
			422, "Unprocesable Entity", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.InputISO = helper.ClearWhitespace(input.InputISO)
	input.ModeType = 3
	res := h.isoService.SplitISO(input)
	c.JSON(http.StatusOK, helper.APIResponse("Data has been proceed sucessfully",
		200, "success", res))
}
