package service

import (
	"fmt"
	"mappertest/entity"
	"mappertest/formatter"
	"mappertest/helper"
	"mappertest/input"
	"strings"
)

type ISOService interface {
	SplitISO(in input.ISOInput) entity.ISO
}

type isoService struct {
}

func NewISOService() *isoService {
	return &isoService{}
}

func (s *isoService) SplitISO(in input.ISOInput) entity.ISO {

	var iso entity.ISO
	split := strings.Split(in.InputISO, "")

	iso.Header = helper.HeaderFromISO(in.InputLengthHeader, split)
	iso.Mti = helper.MTIFromISO(in.InputLengthHeader, split)

	mtiArr := helper.ConvertMTISplit(iso.Mti)
	iso.MTIMess = formatter.FormatMTI(mtiArr[0], mtiArr[1],
		mtiArr[2], mtiArr[3])

	var startDataLen int
	iso.Bitmap, startDataLen = helper.BitmapFromISO(in.InputLengthHeader, split)

	dataEl := split[startDataLen:]
	fmt.Println(dataEl)

	iso.DataElement = helper.GetField(dataEl, iso.Bitmap.Field)

	return iso
}
