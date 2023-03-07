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

	if !convert && mode == 3 && length%2 == 1 {
		length += 1
	}

	de := raw[:length]
	fmt.Println(num, ": ", de, " length :", length)
	newRaw := raw[length:]
	newVal = ArrToString(de)
	if convert && mode == 3 {
		value, _ := hex.DecodeString(newVal)
		newVal = string(value)
		fmt.Println(num, " isinya: ", newVal)
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
	if mode == 3 && dataType%2 == 1 {
		dataType += 1
	}

	length, err := strconv.Atoi(ArrToString(raw[:dataType]))
	if err != nil {
		fmt.Println("Length is not an integer")
	}

	//func covert use to check whether the field need to be decoded to hex or not
	convert := CheckAlpha(strconv.Itoa(num))
	if convert && mode == 3 {
		length *= 2
	}

	de := raw[dataType : length+dataType]
	fmt.Println(num, ": ", de, " length :", length)
	newRaw := raw[length+dataType:]

	newVal := ArrToString(de)
	if convert && mode == 3 {
		value, _ := hex.DecodeString(newVal)
		newVal = string(value)
		fmt.Println(num, " isinya: ", newVal)
	}

	return newVal, newRaw, length
}

func AssignField(id int, label string) entity.Field {
	var field entity.Field
	field.Id = id
	field.Label = label
	return field
}

func TypeVarCustom(mode int, num int, raw []string, dataType int, mti string) (string, []string, int) {
	//only if the length is defined by hex numnber
	//lengthHex := ArrToString(raw[:dataType])
	// hex, err := strconv.ParseInt(lengthHex, 16, 64)
	// length := int(hex)
	if mode == 3 && dataType%2 == 1 {
		dataType += 1
	}

	length, err := strconv.Atoi(ArrToString(raw[:dataType]))
	if err != nil {
		fmt.Println("Length is not an integer")
	}

	//func covert use to check whether the field need to be decoded to hex or not
	convert := CheckAlpha(strconv.Itoa(num))
	if convert && mode == 3 {
		length *= 2
	}

	de := raw[dataType : length+dataType]
	fmt.Println(num, ": ", de, " length :", length)
	newRaw := raw[length+dataType:]

	newVal := ArrToString(de)
	if convert && mode == 3 {
		value, _ := hex.DecodeString(newVal)
		newVal = string(value)
		fmt.Println(num, " isinya: ", newVal)
	}

	customSpec := CheckCustomSpec(strconv.Itoa(num))
	if customSpec {
		newVal = getSpec(strconv.Itoa(num), ArrToString(de), mti)
	}
	return newVal, newRaw, length
}

func TypeFixedCustom(mode int, num int, raw []string, length int, mti string) (string, []string) {
	var newVal string
	convert := CheckAlpha(strconv.Itoa(num))
	if convert && mode == 3 {
		length *= 2
	}

	if !convert && mode == 3 && length%2 == 1 {
		length += 1
	}

	de := raw[:length]
	fmt.Println(num, ": ", de, " length :", length)
	newRaw := raw[length:]
	newVal = ArrToString(de)
	if convert && mode == 3 {
		value, _ := hex.DecodeString(newVal)
		newVal = string(value)
		fmt.Println(num, " isinya: ", newVal)
	}

	customSpec := CheckCustomSpec(strconv.Itoa(num))
	if customSpec {
		newVal = getSpec(strconv.Itoa(num), ArrToString(de), mti)
	}
	return newVal, newRaw
}
