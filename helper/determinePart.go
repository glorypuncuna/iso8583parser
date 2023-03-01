package helper

import "mappertest/formatter"

func MTIFromISO(lengthHeader int, isoMessage []string) string {
	//MTI sesuai dengan header
	mtiLength := lengthHeader + 4

	//masukin mti nya
	mtiSlice := isoMessage[lengthHeader:mtiLength]
	mti := ArrToString(mtiSlice)
	return mti
}

func HeaderFromISO(lengthHeader int, isoMessage []string) string {
	headerSlice := isoMessage[:lengthHeader]
	header := ArrToString(headerSlice)
	return header
}

func BitmapFromISO(lengthHeader int, isoMessage []string) (formatter.BitmapFormatter, int) {
	startIndexBitmap := lengthHeader + 4
	bitmapLength := startIndexBitmap + 16
	secondBit := CheckSecondBitmap(isoMessage[startIndexBitmap])

	var bitmapSlice []string
	var bitmap formatter.BitmapFormatter

	if secondBit {
		bitmapLength = startIndexBitmap + 32
		bitmapSlice = isoMessage[startIndexBitmap:bitmapLength]
	} else {
		bitmapSlice = isoMessage[startIndexBitmap:bitmapLength]
	}

	bitmap.Field, bitmap.Bit = Splitter(ArrToString(bitmapSlice))

	return bitmap, bitmapLength
}
