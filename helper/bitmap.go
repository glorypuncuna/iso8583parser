package helper

import "strconv"

func Splitter(raw string) ([]int, string) {
	counter := 1
	var fillFilled []int
	var bit string
	//loop through the runes of string
	for _, rn := range raw {
		//could not convert rune, hence change it to string first
		c := string(rn)
		//hexString to hex
		hex, _ := strconv.ParseUint(c, 16, 32)
		//convert to bit
		bits := AsBits(hex)
		//Determine whether a certain field exist or not
		for _, b := range bits {
			bit += strconv.Itoa(int(b))
			if b == 1 {
				fillFilled = append(fillFilled, counter)
			}
			counter++
		}
	}
	return fillFilled, bit
}

func AsBits(val uint64) []uint64 {
	bits := []uint64{}
	for i := 0; i < 4; i++ {
		bits = append([]uint64{val & 0x1}, bits...)
		val = val >> 1
	}
	return bits
}

func CheckSecondBitmap(c string) bool {
	hex, _ := strconv.ParseUint(c, 16, 32)
	bit := AsBits(hex)[0]
	// if int(bit) == 1 {
	// 	return true
	// }
	// return false
	return int(bit) == 1
}

// func GetBitmapValue(raw []string, filledField []int){

// 	for _, f := range filledField{
// 		switch f{
// 		case 1:
// 			continue
// 		case 2:
// 			field2 :=
// 		}
// 	}
// }
