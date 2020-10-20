package receiptFormatting

import (
	"strconv"
)

func SumStringNumber(sources []string) string {
	sum := 0
	for _, val := range sources {
		num, _ := strconv.Atoi(val)
		sum += num
	}
	return strconv.Itoa(sum)
}
