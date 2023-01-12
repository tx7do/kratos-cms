package entgo

import (
	"encoding/json"
	"entgo.io/ent/dialect/sql"
	"strings"
	"time"
	"unicode"
)

func QueryCommandToSelector(queryCmd, orderByCmd map[string]string) ([]func(*sql.Selector), []func(*sql.Selector)) {
	var whereCond []func(s *sql.Selector)
	var orderCond []func(*sql.Selector)

	if queryCmd != nil {
		queryStr, ok := queryCmd["query"]
		if ok {
			var queryMap map[string]string
			err := json.Unmarshal([]byte(queryStr), &queryMap)
			if err == nil {
				for k, v := range queryMap {
					key := ToSnakeCase(k)
					value := v
					whereCond = append(whereCond, func(s *sql.Selector) {
						s.Where(sql.EQ(s.C(key), value))
					})
				}
			}
		}
	}

	if orderByCmd != nil {
		orderByStr, ok := orderByCmd["orderBy"]
		if ok {
			var orderByMap map[string]string
			err := json.Unmarshal([]byte(orderByStr), &orderByMap)
			if err == nil {
				for k, v := range orderByMap {
					key := ToSnakeCase(k)
					switch strings.ToLower(v) {
					case "desc", "descending", "descend":
						orderCond = append(orderCond, func(s *sql.Selector) {
							s.OrderBy(sql.Desc(s.C(key)))
						})
					case "asc", "ascending", "ascend":
						orderCond = append(orderCond, func(s *sql.Selector) {
							s.OrderBy(sql.Asc(s.C(key)))
						})
					}
				}
			}
		}
	}

	return whereCond, orderCond
}

func ToSnakeCase(input string) string {
	if input == "" {
		return input
	}
	if len(input) == 1 {
		return strings.ToLower(input)
	}

	input = strings.Replace(input, " ", "", -1)

	source := []rune(input)
	dist := strings.Builder{}
	dist.Grow(len(input) + len(input)/3) // avoid reallocation memory, 33% ~ 50% is recommended
	skipNext := false
	for i := 0; i < len(source); i++ {
		cur := source[i]
		switch cur {
		case '-', '_':
			dist.WriteRune('_')
			skipNext = true
			continue
		}
		if unicode.IsLower(cur) || unicode.IsDigit(cur) {
			dist.WriteRune(cur)
			continue
		}

		if i == 0 {
			dist.WriteRune(unicode.ToLower(cur))
			continue
		}

		last := source[i-1]
		if (!unicode.IsLetter(last)) || unicode.IsLower(last) {
			if skipNext {
				skipNext = false
			} else {
				dist.WriteRune('_')
			}
			dist.WriteRune(unicode.ToLower(cur))
			continue
		}
		// last is upper case
		if i < len(source)-1 {
			next := source[i+1]
			if unicode.IsLower(next) {
				if skipNext {
					skipNext = false
				} else {
					dist.WriteRune('_')
				}
				dist.WriteRune(unicode.ToLower(cur))
				continue
			}
		}
		dist.WriteRune(unicode.ToLower(cur))
	}

	return dist.String()
}

func ToPascalCase(input string) string {
	return toCamelCase(input, true)
}

func ToCamelCase(input string) string {
	return toCamelCase(input, false)
}

func toCamelCase(s string, initCase bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	var uppercaseAcronym = map[string]string{}
	if a, ok := uppercaseAcronym[s]; ok {
		s = a
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := initCase
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return n.String()
}

const tsLayout = "2006-01-02 15:04:05"

// UnixMilliToStringPtr 毫秒时间戳转字符串
func UnixMilliToStringPtr(tm *int64) *string {
	if tm == nil {
		return nil
	}
	str := time.UnixMilli(*tm).Format(tsLayout)
	return &str
}

// UnixMilliStringToInt64Ptr 字符串转毫秒时间戳
func UnixMilliStringToInt64Ptr(tm *string) *int64 {
	if tm == nil {
		return nil
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	theTime, err := time.ParseInLocation(tsLayout, *tm, loc)
	if err != nil {
		return nil
	}
	unixTime := theTime.UnixMilli()
	return &unixTime
}
