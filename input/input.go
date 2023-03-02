package input

type BitmapHexInput struct {
	InputBitmapHex string `uri:"input" binding:"required"`
}

type BitmapBitInput struct {
	InputBitmapBit string `uri:"input" binding:"required"`
}

type ISOInput struct {
	InputISO          string `json:"iso"`
	InputLengthHeader int    `json:"lengthHeader"`
	ModeType          int    `json:"mode"`
}

type MTIInput struct {
	InputMTI string `uri:"input" binding:"required"`
}
