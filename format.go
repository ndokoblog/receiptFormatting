package receiptFormatting

import (
	"encoding/json"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func NumFormat(s string) string {
	s = strings.TrimSuffix(s, ".00")
	s = strings.TrimLeft(s, "0")
	if s == "" {
		s = "0"
	}
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

func LocalMonth(s string, long bool) string {
	if long {
		if v, ok := localMonthLong[s]; ok {
			return v
		}
	} else {
		if v, ok := localMonthShort[s]; ok {
			return v
		}
	}
	return ""
}

func ConvertToInt(s string) int {
	f, _ := strconv.ParseFloat(s, 64)
	return int(f)
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

func Stringify2(v interface{}) string {
	s, _ := jsoniter.MarshalToString(v)
	return s
}
