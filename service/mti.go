package service

import (
	"mappertest/helper"
	"mappertest/input"
)

type MTIService interface {
	ConvertMTI(inp input.MTIInput) []string
}

type mtiService struct {
}

func NewMTIService() *mtiService {
	return &mtiService{}
}

func (s *mtiService) ConvertMTI(inp input.MTIInput) []string {
	message := helper.ConvertMTISplit(inp.InputMTI)
	return message
}
