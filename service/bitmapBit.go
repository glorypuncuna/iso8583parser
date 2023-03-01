package service

import (
	"mappertest/helper"
	"mappertest/input"
)

type BitmapBitService interface {
	DetermineField(input input.BitmapBitInput) ([]int, string)
}

type bitmapBitService struct {
}

func NewBitmapBitService() *bitmapBitService {
	return &bitmapBitService{}
}

func (s *bitmapBitService) DetermineField(in input.BitmapBitInput) ([]int, string) {
	field, bit := helper.Splitter(in.InputBitmapBit)
	return field, bit
}
