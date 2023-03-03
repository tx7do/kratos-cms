package stringcase

import (
	"strings"
	"unicode"
)

// ToSnakeCase 把字符转换为 蛇形命名法（snake_case）
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

// ToPascalCase 把字符转换为 帕斯卡命名/大驼峰命名法（CamelCase）
func ToPascalCase(input string) string {
	return toCamelCase(input, true)
}

// ToLowCamelCase 把字符转换为 小驼峰命名法（lowerCamelCase）
func ToLowCamelCase(input string) string {
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
