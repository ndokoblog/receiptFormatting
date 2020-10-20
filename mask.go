package receiptFormatting

func AccountMask(source string) string {
	if len(source) == 15 {
		return source[:4] + " **** **** " + source[len(source)-3:]
	} else {
		return "-"
	}
}

func PhoneMask(source string) string {
	if len(source) >= 10 && len(source) <= 13 {
		return source[:2] + " **** **** " + source[len(source)-3:]
	} else {
		return "-"
	}
}
