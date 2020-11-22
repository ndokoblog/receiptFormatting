package receiptFormatting

import (
	"encoding/json"
	"strconv"
	"strings"
)

func NumFormat(s string) string {
	s = strings.TrimSuffix(s, ".00")
	v := len(s)
	if v <= 3 {
		return s
	}

	d := v % 3
	if d == 0 {
		d = 3
	}
	last := 0

	var parts []string

	for last < v {
		parts = append(parts, s[last:last+d])
		last = last + d
		d = 3
	}
	return strings.Join(parts, ".")
}

func Separate(source, separator string, pad int) string {
	v := len(source)
	if v <= pad {
		return source
	}

	var parts []string

	for last := 0; last < v; last += pad {
		end := last + pad
		if end > v {
			end = v
		}
		parts = append(parts, source[last:end])
	}
	return strings.Join(parts, separator)
}

func PadLeft(str, pad string, lenght int) string {
	for {
		if len(str) >= lenght {
			return str[0:lenght]
		}
		str = pad + str
	}
}

func PadRight(str, pad string, lenght int) string {
	for {
		if len(str) >= lenght {
			return str[0:lenght]
		}
		str = str + pad
	}
}

func ConvertToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func ConvertToMap(v interface{}) (m map[string]interface{}) {
	js, _ := json.Marshal(v)
	_ = json.Unmarshal(js, &m)
	return m
}

func Stringify(v interface{}) string {
	js, _ := json.Marshal(v)
	s := strings.ReplaceAll(string(js), `\`, ``)
	return s
}
