package helper

import (
	"strconv"
)

func ConvertMTISplit(raw string) []string {
	var mes []string
	for i, rn := range raw {
		c := string(rn)
		num, _ := strconv.Atoi(c)
		dig := i + 1
		switch dig {
		case 1:
			mes = append(mes, MTIFirst(num))
		case 2:
			mes = append(mes, MTISecond(num))
		case 3:
			mes = append(mes, MTIThird(num))
		case 4:
			mes = append(mes, MTIFourth(num))
		}

	}
	return mes
}

func MTIFirst(bit int) string {
	switch bit {
	case 0:
		return "ISO 8583:1987"
	case 1:
		return "ISO 8583:1993"
	case 2:
		return "ISO 8583:2003"
	case 8:
		return "National use"
	case 9:
		return "Private use"
	default:
		return "Reserved by ISO"
	}
}

func MTISecond(bit int) string {
	switch bit {
	case 1:
		return "Authorization message"
	case 2:
		return "Financial messages"
	case 3:
		return "File actions message"
	case 4:
		return "Reversal and chargeback messages"
	case 5:
		return "Reconciliation message"
	case 6:
		return "Administrative message"
	case 7:
		return "Fee collection messages"
	case 8:
		return "Network management message"
	default:
		return "Reserved by ISO"
	}
}

func MTIThird(bit int) string {
	switch bit {
	case 0:
		return "Request"
	case 1:
		return "Request response"
	case 2:
		return "Advice"
	case 3:
		return "Advice response"
	case 4:
		return "Notification"
	case 5:
		return "Notification acknowledgement"
	case 6:
		return "Instruction"
	case 7:
		return "Instruction acknowledgement"
	default:
		return "Reserved for ISO use"
	}
}

func MTIFourth(bit int) string {
	switch bit {
	case 0:
		return "Acquirer"
	case 1:
		return "Acquirer repeat"
	case 2:
		return "Issuer"
	case 3:
		return "Issuer repeat"
	case 4:
		return "Other"
	default:
		return "Reserved by ISO"
	}
}
