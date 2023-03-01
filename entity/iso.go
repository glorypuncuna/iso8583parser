package entity

import "mappertest/formatter"

type ISO struct {
	Header      string
	Mti         string
	MTIMess     formatter.MTIFormatter
	Bitmap      formatter.BitmapFormatter
	DataElement []Field
}
