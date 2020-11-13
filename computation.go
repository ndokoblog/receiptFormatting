package receiptFormatting

import (
	"strconv"
	"strings"
)

func SumStringNumber(sources []string) string {
	sum := 0
	for _, val := range sources {
		num, _ := strconv.Atoi(strings.TrimSuffix(val, ".00"))
		sum += num
	}
	return strconv.Itoa(sum)
}
