package stringcase

import (
	"fmt"
	"testing"
)

func TestToLowCamelCase(t *testing.T) {
	fmt.Println(ToLowCamelCase("snake_case"))
	fmt.Println(ToLowCamelCase("CamelCase"))
	fmt.Println(ToLowCamelCase("lowerCamelCase"))
}

func TestToPascalCase(t *testing.T) {
	fmt.Println(ToPascalCase("snake_case"))
	fmt.Println(ToPascalCase("CamelCase"))
	fmt.Println(ToPascalCase("lowerCamelCase"))
}

func TestToSnakeCase(t *testing.T) {
	fmt.Println(ToSnakeCase("snake_case"))
	fmt.Println(ToSnakeCase("CamelCase"))
	fmt.Println(ToSnakeCase("lowerCamelCase"))
}
