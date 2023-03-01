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

func (h *isoHandler) TestDulu(c *gin.Context) {
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

	res := h.isoService.SplitISO(input)
	c.JSON(http.StatusOK, helper.APIResponse("Data has been proceed sucessfully",
		200, "success", res))
}
