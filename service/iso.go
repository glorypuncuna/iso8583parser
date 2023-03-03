package service

import (
	"encoding/hex"
	"fmt"
	"mappertest/entity"
	"mappertest/formatter"
	"mappertest/helper"
	"mappertest/input"
	"strings"
)

type ISOService interface {
	SplitISO(in input.ISOInput) entity.ISO
	ToHex(in input.ISOInput) entity.ISO
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

	fmt.Println("MTI ", iso.MTIMess, iso.Bitmap.Field)

	iso.DataElement = helper.GetField(dataEl, iso.Bitmap.Field, in.ModeType)

	return iso
}

func (s *isoService) ToHex(in input.ISOInput) entity.ISO {

	in.InputISO = helper.ClearWhitespace(in.InputISO)
	split := strings.Split(in.InputISO, "")

	mtiLengthInHex := 8
	length := in.InputLengthHeader + mtiLengthInHex
	header := split[:length]

	bitmapLength := 16
	endBitmap := length + bitmapLength

	secondBitmap := helper.CheckSecondBitmap(split[length])

	if secondBitmap {
		endBitmap += bitmapLength
	}

	newBitmap := helper.ArrToString(split[length:endBitmap])
	newHeader, _ := hex.DecodeString(helper.ArrToString(header))
	newDe, _ := hex.DecodeString(helper.ArrToString(split[endBitmap:]))

	noHexString := string(newHeader) + newBitmap + string(newDe)

	in.InputISO = noHexString
	in.InputLengthHeader /= 2

	newIso := s.SplitISO(in)

	return newIso
}
