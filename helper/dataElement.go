package helper

import (
	"mappertest/entity"
)

//1 for lvar, 2 for llvar, and 3 for lllvar

func GetField(raw []string, bitmap []int, mode int, mti string) []entity.Field {
	var datel []entity.Field

	for _, c := range bitmap {

		switch c {
		case 2:
			field := AssignField(2, "Primary account number (PAN)")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 2)

			datel = append(datel, field)
		case 3:
			field := AssignField(3, "Processing Code")
			field.Length = 6
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 4:
			field := AssignField(4, "Amount Transaction")
			field.Length = 12
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 5:
			field := AssignField(5, "Amount, settlement")
			field.Length = 12
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 6:
			field := AssignField(6, "Amount, cardholder billing")
			field.Length = 12
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 7:
			field := AssignField(7, "Transmission date & time")
			field.Length = 10
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 8:
			field := AssignField(8, "Amount, cardholder billing fee")
			field.Length = 8
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 9:
			field := AssignField(9, "Conversion rate, settlement")
			field.Length = 8
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 10:
			field := AssignField(10, "Conversion rate, cardholder billing")
			field.Length = 8
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 11:
			field := AssignField(11, "System trace audit number (STAN)")
			field.Length = 6
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 12:
			field := AssignField(12, "Local transaction time (hhmmss)")
			field.Length = 6
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 13:
			field := AssignField(13, "Local transaction date (MMDD)")
			field.Length = 4
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 14:
			field := AssignField(14, "Expiration date (YYMM)")
			field.Length = 4
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 15:
			field := AssignField(15, "Settlement date")
			field.Length = 4
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 16:
			field := AssignField(16, "Currency conversion date")
			field.Length = 4
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 17:
			field := AssignField(17, "Capture date")
			field.Length = 4
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 18:
			field := AssignField(18, "Merchant type, or merchant category code")
			field.Length = 4
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 19:
			field := AssignField(19, "Acquiring institution (country code)")
			field.Length = 3
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 20:
			field := AssignField(20, "PAN extended (country code)")
			field.Length = 3
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 21:
			field := AssignField(21, "Forwarding institution (country code)")
			field.Length = 3
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 22:
			field := AssignField(22, "Point of service entry mode")
			field.Length = 3
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 23:
			field := AssignField(23, "Application PAN sequence number")
			field.Length = 3
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 24:
			field := AssignField(24, "Network international identifier (NII)")
			field.Length = 3
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 25:
			field := AssignField(25, "Point of service condition code")
			field.Length = 2
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 26:
			field := AssignField(26, "Point of service capture code")
			field.Length = 2
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 27:
			field := AssignField(27, "Authorizing identification response length")
			field.Length = 1
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 28:
			field := AssignField(28, "Amount, transaction fee")
			field.Length = 8
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 29:
			field := AssignField(29, "Amount, settlement fee")
			field.Length = 8
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 30:
			field := AssignField(30, "Amount, transaction processing fee")
			field.Length = 8
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 31:
			field := AssignField(31, "Amount, settlement processing fee")
			field.Length = 8
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 32:
			field := AssignField(32, "Acquiring institution identification code")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 2)

			datel = append(datel, field)
		case 33:
			field := AssignField(33, "Forwarding institution identification code")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 2)

			datel = append(datel, field)
		case 34:
			field := AssignField(34, "Primary account number, extended")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 2)

			datel = append(datel, field)
		case 35:
			field := AssignField(35, "Track 2 data")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 2)

			datel = append(datel, field)
		case 36:
			field := AssignField(36, "Track 3 data")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 37:
			field := AssignField(37, "Retrieval reference number")
			field.Length = 12
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 38:
			field := AssignField(38, "Authorization identification response")
			field.Length = 6
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 39:
			field := AssignField(39, "Response Code")
			field.Length = 2
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 40:
			field := AssignField(40, "Service restriction code")
			field.Length = 3
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 41:
			field := AssignField(41, "Card acceptor terminal identification")
			field.Length = 8
			if mode == 1 {
				field.Length = 16
			}
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 42:
			field := AssignField(42, "Card acceptor identification code")
			field.Length = 15
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 43:
			field := AssignField(43, "Card acceptor name or location")
			field.Length = 40
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 44:
			field := AssignField(44, "Additional response data")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 2)

			datel = append(datel, field)
		case 45:
			field := AssignField(45, "Track 1 data")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 2)

			datel = append(datel, field)
		case 46:
			field := AssignField(46, "Additional data (ISO)")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 47:
			field := AssignField(47, "Additional data (national)")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 48:
			field := AssignField(48, "Additional data (private)")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 49:
			field := AssignField(49, "Currency code, transaction")
			field.Length = 3
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 50:
			field := AssignField(50, "Currency code, settlement")
			field.Length = 3
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 51:
			field := AssignField(51, "Currency code, cardholder billing")
			field.Length = 3
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 52:
			field := AssignField(52, "Personal identification number data")
			field.Length = 64
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 53:
			field := AssignField(53, "Security related control information")
			field.Length = 2
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 54:
			field := AssignField(54, "Additional amounts")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 55:
			field := AssignField(55, "ICC data, EMV having multiple tags")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 56:
			field := AssignField(56, "Reserved (ISO)")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 57:
			field := AssignField(57, "Reserved (national)")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 58:
			field := AssignField(58, "Reserved (national)")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 59:
			field := AssignField(59, "Reserved (national)")

			field.Value, raw, field.Length = TypeVarCustom(mode, c, raw, 3, mti)

			datel = append(datel, field)
		case 60:
			field := AssignField(60, "Reserved (national)")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 61:
			field := AssignField(61, "Reserved (private)")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 62:
			field := AssignField(62, "Reserved (private)")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 4)

			datel = append(datel, field)
		case 63:
			field := AssignField(63, "Reserved (private)")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 64:
			field := AssignField(64, "Message authentication code (MAC)")
			field.Length = 64
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 65:
			field := AssignField(65, "Extended bitmap indicator")
			field.Length = 1
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 66:
			field := AssignField(66, "Settlement code")
			field.Length = 1
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 67:
			field := AssignField(67, "Extended payment code")
			field.Length = 2
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 68:
			field := AssignField(68, "Receiving institution country code")
			field.Length = 3
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 69:
			field := AssignField(69, "Settlement institution country code")
			field.Length = 3
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 70:
			field := AssignField(70, "Network management information code")
			field.Length = 3
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 71:
			field := AssignField(71, "Message number")
			field.Length = 4
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 72:
			field := AssignField(72, "Last message's number")
			field.Length = 4
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 73:
			field := AssignField(73, "Action date (YYMMDD)")
			field.Length = 6
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 74:
			field := AssignField(74, "Number of credits")
			field.Length = 10
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 75:
			field := AssignField(75, "Credits, reversal number")
			field.Length = 10
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 76:
			field := AssignField(76, "Number of debits")
			field.Length = 10
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 77:
			field := AssignField(77, "Debits, reversal number")
			field.Length = 10
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 79:
			field := AssignField(79, "Transfer, reversal number")
			field.Length = 10
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 80:
			field := AssignField(80, "Number of inquiries")
			field.Length = 10
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 81:
			field := AssignField(81, "Number of authorizations")
			field.Length = 10
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 82:
			field := AssignField(82, "Credits, processing fee amount")
			field.Length = 12
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 83:
			field := AssignField(83, "Credits, transaction fee amount")
			field.Length = 12
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 84:
			field := AssignField(84, "Debits, processing fee amount")
			field.Length = 12
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 85:
			field := AssignField(85, "Debits, transaction fee amount")
			field.Length = 12
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 86:
			field := AssignField(86, "Total amount of credits")
			field.Length = 16
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 87:
			field := AssignField(87, "Credits, reversal amount")
			field.Length = 16
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 88:
			field := AssignField(88, "Total amount of debits")
			field.Length = 16
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 89:
			field := AssignField(89, "Debits, reversal amount")
			field.Length = 16
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 90:
			field := AssignField(90, "Original data elements")
			field.Length = 42
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 91:
			field := AssignField(91, "File update code")
			field.Length = 1
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 92:
			field := AssignField(92, "File security code")
			field.Length = 2
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 93:
			field := AssignField(93, "Response indicator")
			field.Length = 5
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 94:
			field := AssignField(94, "Service indicator")
			field.Length = 7
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 95:
			field := AssignField(95, "Replacement amounts")
			field.Length = 42
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 96:
			field := AssignField(96, "Message security code")
			field.Length = 64
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 97:
			field := AssignField(97, "Net settlement amount")
			field.Length = 16
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 98:
			field := AssignField(98, "Payee")
			field.Length = 25
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		case 99:
			field := AssignField(99, "Settlement institution identification code")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 2)

			datel = append(datel, field)
		case 100:
			field := AssignField(100, "Receiving institution identification code")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 2)

			datel = append(datel, field)
		case 101:
			field := AssignField(101, "File name")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 2)

			datel = append(datel, field)
		case 102:
			field := AssignField(102, "Account identification 1")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 2)

			datel = append(datel, field)
		case 103:
			field := AssignField(103, "Account identification 1")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 2)

			datel = append(datel, field)
		case 104:
			field := AssignField(104, "Transaction description")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 105, 106, 107, 108, 109, 110, 111:
			field := AssignField(c, "Reserved for ISO use")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 112, 113, 114, 115, 116, 117, 118, 119:
			field := AssignField(c, "Reserved for national use")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 120, 121, 122, 123, 124, 125, 126, 127:
			field := AssignField(c, "Reserved for private use")

			field.Value, raw, field.Length = TypeVar(mode, c, raw, 3)

			datel = append(datel, field)
		case 128:
			field := AssignField(c, "Message authentication code")
			field.Length = 64
			field.Value, raw = TypeFixed(mode, c, raw, field.Length)

			datel = append(datel, field)
		}

	}
	return datel
}
