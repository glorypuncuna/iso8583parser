package helper

func ArrToString(raw []string) string {
	var convert string

	for _, c := range raw {
		convert += c
	}

	return convert
}
