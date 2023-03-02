package helper

import (
	"strconv"
	"strings"
)

func ArrToString(raw []string) string {
	var convert string

	for _, c := range raw {
		convert += c
	}

	return convert
}

func ClearWhitespace(raw string) string {
	split := strings.Fields(raw)
	res := ArrToString(split)
	return res
}

func GetMapAlpha() map[int]int {
	dict := map[int]int{37: 0, 38: 0, 39: 0, 40: 0, 41: 0,
		42: 0, 43: 0, 44: 0, 45: 0, 46: 0,
		47: 0, 48: 0, 49: 0, 50: 0, 51: 0,
		54: 0, 55: 0, 56: 0, 57: 0, 58: 0,
		59: 0, 60: 0, 61: 0, 62: 0, 63: 0,
		91: 0, 92: 0, 93: 0, 94: 0, 95: 0,
		98: 0, 101: 0, 102: 0, 103: 0, 104: 0,
		105: 0, 106: 0, 107: 0, 108: 0, 109: 0, 110: 0,
		111: 0, 112: 0, 113: 0, 114: 0, 115: 0, 116: 0,
		117: 0, 118: 0, 119: 0, 120: 0, 121: 0, 122: 0,
		123: 0, 124: 0, 125: 0, 126: 0, 127: 0}

	return dict
}

// func ChangeLengthVar(raw []string, dt int, dataType int) []string {

// 	var lengthAndArr []string
// 	len, _ := strconv.Atoi(ArrToString(raw[:dataType]))
// 	length := len * 2
// 	newArray := ArrToString(raw[dataType:])

// 	newLength := strconv.Itoa(length)
// 	lengthAndString := newLength + newArray
// 	lengthAndArr = strings.Split(lengthAndString, "")

// 	return lengthAndArr
// }

func CheckAlpha(num string) bool {
	aMap := GetMapAlpha()
	car, _ := strconv.Atoi(num)
	_, ok := aMap[car]

	return ok
}
