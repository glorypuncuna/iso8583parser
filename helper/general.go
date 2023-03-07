package helper

import (
	"encoding/hex"
	"fmt"
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

func GetFieldWithCustomSpec() map[int]int {
	dict := map[int]int{
		4: 0, 59: 0,
	}
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

func CheckCustomSpec(num string) bool {
	aMap := GetFieldWithCustomSpec()
	car, _ := strconv.Atoi(num)
	_, ok := aMap[car]

	return ok
}

func getSpec(num string, value string, mti string) string {
	switch num {
	case "4":
		return "RP " + value + ",00"
	case "59":
		length := len(value)
		newArr := strings.Split(value, "")
		value = ""
		var result string
		if mti == "0200" {
			for i := 0; i < 16; i++ {
				if length == 0 {
					break
				}

				switch i {
				case 0:
					result, newArr = cutFixed(newArr, 4, true)
					value += result + "-"
					length -= 4
				case 1:
					result, newArr = cutFixed(newArr, 30, true)
					value += result + "-"
					length -= 18
				case 2:
					result, newArr = cutFixed(newArr, 10, false)
					value += result + "-"
					length -= 10
				case 3:
					result, newArr = cutFixed(newArr, 10, false)
					value += result + "-"
					length -= 10
				case 4:
					result, newArr = cutFixed(newArr, 30, true)
					value += result + "-"
					length -= 30
				case 5:
					result, newArr = cutFixed(newArr, 30, true)
					value += result + "-"
					length -= 30
				case 6:
					result, newArr = cutFixed(newArr, 2, false)
					value += result + "-"
					length -= 2
				case 7:
					result, newArr = cutFixed(newArr, 6, false)
					value += result + "-"
					length -= 6
				case 8:
					result, newArr = cutFixed(newArr, 6, false)
					value += result + "-"
					length -= 6
				case 9:
					result, newArr = cutFixed(newArr, 6, false)
					value += result + "-"
					length -= 6
				case 10:
					result, newArr = cutFixed(newArr, 6, false)
					value += result + "-"
					length -= 6
				case 11:
					result, newArr = cutFixed(newArr, 2, true)
					value += result + "-"
					length -= 2
				case 12:
					result, newArr = cutFixed(newArr, 4, false)
					value += result + "-"
					length -= 4
				case 13:
					result, newArr = cutFixed(newArr, 4, false)
					value += result + "-"
					length -= 4
				case 14:
					result, newArr = cutFixed(newArr, 4, false)
					value += result + "-"
					length -= 4
				case 15:
					result, newArr = cutFixed(newArr, 4, false)
					value += result + "-"
					length -= 4
				}
				fmt.Println("\nrseultny: ", result)
			}
		}

		if mti == "0210" {
			for i := 0; i < 3; i++ {
				if length == 0 {
					break
				}

				switch i {
				case 0:
					result, newArr = cutFixed(newArr, 4, false)
					value += result + "-"
					length -= 4
				case 1:
					result, newArr = cutFixed(newArr, 4, false)
					value += result + "-"
					length -= 4
				case 2:
					result, newArr = cutFixed(newArr, 2, true)
					value += result + "-"
					length -= 2
				}
			}
		}
		return value
	}
	return value
}

func cutFixed(arr []string, length int, convert bool) (string, []string) {
	val := ArrToString(arr[:length])
	if convert {
		miu, _ := hex.DecodeString(val)
		val = string(miu)
	}
	newArr := arr[length:]
	return val, newArr
}
