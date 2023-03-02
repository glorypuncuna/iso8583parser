package helper

import (
	"encoding/hex"
	"fmt"
	"mappertest/entity"
	"strconv"
)

func TypeFixed(mode int, num int, raw []string, length int) (string, []string) {
	var newVal string
	convert := CheckAlpha(strconv.Itoa(num))
	if convert && mode == 3 {
		length *= 2
	}

	de := raw[:length]
	fmt.Println(num, ": ", de)
	newRaw := raw[length:]
	newVal = ArrToString(de)
	if convert && mode == 3 {
		value, _ := hex.DecodeString(newVal)
		newVal = string(value)
	}

	return newVal, newRaw
}

// func TypeFixedModif(raw []string, length int) (string, []string, string) {
// 	mod := ArrToString(raw[:1])
// 	de := raw[1 : length+1]
// 	newRaw := raw[length+1:]
// 	return ArrToString(de), newRaw, mod
// }

func TypeVar(mode int, num int, raw []string, dataType int) (string, []string, int) {

	//only if the length is defined by hex numnber
	//lengthHex := ArrToString(raw[:dataType])
	// hex, err := strconv.ParseInt(lengthHex, 16, 64)
	// length := int(hex)

	length, err := strconv.Atoi(ArrToString(raw[:dataType]))
	if err != nil {
		fmt.Println("Length is not an integer")
	}

	convert := CheckAlpha(strconv.Itoa(num))
	if convert && mode == 3 {
		length *= 2
	}

	de := raw[dataType : length+dataType]
	newRaw := raw[length+dataType:]

	newVal := ArrToString(de)
	if convert && mode == 3 {
		value, _ := hex.DecodeString(newVal)
		newVal = string(value)
	}

	return newVal, newRaw, length
}

func AssignField(id int, label string) entity.Field {
	var field entity.Field
	field.Id = id
	field.Label = label
	return field
}
