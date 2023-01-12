package slug

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoSimple(t *testing.T) {
	// 俄文
	text := slug.Make("Hellö Wörld хелло ворлд")
	assert.Equal(t, text, "hello-world-khello-vorld")
	fmt.Println(text)

	// 繁体中文
	someText := slug.Make("影師")
	assert.Equal(t, someText, "ying-shi")
	fmt.Println(someText)

	// 简体中文
	cnText := slug.Make("天下太平")
	assert.Equal(t, cnText, "tian-xia-tai-ping")
	fmt.Println(cnText)

	// 英文
	enText := slug.MakeLang("This & that", "en")
	assert.Equal(t, enText, "this-and-that")
	fmt.Println(enText)

	// 德文
	deText := slug.MakeLang("Diese & Dass", "de")
	assert.Equal(t, deText, "diese-und-dass")
	fmt.Println(deText)

	// 保持大小写
	slug.Lowercase = false
	deUppercaseText := slug.MakeLang("Diese & Dass", "de")
	assert.Equal(t, deUppercaseText, "Diese-und-Dass")
	fmt.Println(deUppercaseText)

	// 字符替换
	slug.CustomSub = map[string]string{
		"water": "sand",
	}
	textSub := slug.Make("water is hot")
	assert.Equal(t, textSub, "sand-is-hot")
	fmt.Println(textSub)
}

func TestGenerate(t *testing.T) {
	// 俄文
	text := Generate("Hellö Wörld хелло ворлд")
	assert.Equal(t, text, "hello-world-khello-vorld")
	fmt.Println(text)

	// 繁体中文
	someText := Generate("影師")
	assert.Equal(t, someText, "ying-shi")
	fmt.Println(someText)

	// 简体中文
	cnText := Generate("天下太平")
	assert.Equal(t, cnText, "tian-xia-tai-ping")
	fmt.Println(cnText)

	// 英文
	enText := GenerateEnglish("This & that")
	assert.Equal(t, enText, "this-and-that")
	fmt.Println(enText)

	// 德文
	deText := GenerateGerman("Diese & Dass")
	assert.Equal(t, deText, "diese-und-dass")
	fmt.Println(deText)

	slug.CustomSub = map[string]string{
		"water": "sand",
	}
	textSub := Generate("water is hot")
	assert.Equal(t, textSub, "sand-is-hot")
	fmt.Println(textSub)
}
