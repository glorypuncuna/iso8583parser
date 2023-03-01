package handler

import (
	"mappertest/formatter"
	"mappertest/helper"
	"mappertest/input"
	"mappertest/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mTIHandler struct {
	mtiService service.MTIService
}

func NewMTIHandler(mtiService service.MTIService) *mTIHandler {
	return &mTIHandler{mtiService}
}

func (h *mTIHandler) MTIConverter(c *gin.Context) {
	var inp input.MTIInput
	err := c.ShouldBindUri(&inp)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Cannot process your request. Bad request",
			422, "Unprocesable Entity", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	res := h.mtiService.ConvertMTI(inp)
	mti := formatter.FormatMTI(res[0], res[1], res[2], res[3])
	c.JSON(http.StatusOK, helper.APIResponse("Data has been proceed sucessfully",
		200, "success", mti))
}
