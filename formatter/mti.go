package formatter

type MTIFormatter struct {
	ISOVersion      string `json:"isoVersion"`
	MessageClass    string `json:"messageClass"`
	MessageFunction string `json:"messageFunction"`
	MessageOrigin   string `json:"messageOrigin"`
}

func FormatMTI(bit1 string, bit2 string, bit3 string, bit4 string) MTIFormatter {
	mtiResult := MTIFormatter{
		ISOVersion:      bit1,
		MessageClass:    bit2,
		MessageFunction: bit3,
		MessageOrigin:   bit4,
	}
	return mtiResult
}
