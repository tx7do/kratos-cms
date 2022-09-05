package slug

import (
	"github.com/gosimple/slug"
)

// Generate 生成链接
func Generate(input string) string {
	return slug.MakeLang(input, "en")
}

// GenerateEnglish 生成英文链接
func GenerateEnglish(input string) string {
	return slug.MakeLang(input, "en")
}

// GenerateGerman 生成德文链接
func GenerateGerman(input string) string {
	return slug.MakeLang(input, "de")
}
