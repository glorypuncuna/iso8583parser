package entity

import "mappertest/formatter"

type ISO struct {
	Header      string                    `json:"header"`
	Mti         string                    `json:"mti"`
	MTIMess     formatter.MTIFormatter    `json:"mtiMess"`
	Bitmap      formatter.BitmapFormatter `json:"bitmap"`
	DataElement []Field                   `json:"dataElement"`
}
