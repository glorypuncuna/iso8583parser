package helper

import (
	"fmt"
	"mappertest/entity"
	"strconv"
)

func TypeFixed(raw []string, length int) (string, []string) {
	de := raw[:length]
	newRaw := raw[length:]
	return ArrToString(de), newRaw
}

// func TypeFixedModif(raw []string, length int) (string, []string, string) {
// 	mod := ArrToString(raw[:1])
// 	de := raw[1 : length+1]
// 	newRaw := raw[length+1:]
// 	return ArrToString(de), newRaw, mod
// }

func TypeVar(raw []string, dataType int) (string, []string, int) {

	//only if the length is defined by hex numnber
	//lengthHex := ArrToString(raw[:dataType])
	// hex, err := strconv.ParseInt(lengthHex, 16, 64)
	// length := int(hex)

	length, err := strconv.Atoi(ArrToString(raw[:dataType]))

	fmt.Println(length)
	if err != nil {
		fmt.Println("Length is not an integer")
	}
	de := raw[dataType : length+dataType]
	newRaw := raw[length+dataType:]
	return ArrToString(de), newRaw, length
}

func AssignField(id int, label string) entity.Field {
	var field entity.Field
	field.Id = id
	field.Label = label
	return field
}
