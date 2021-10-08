package word

import (
	"strings"
	"unicode"
)

// Toupper 转换为大写，如：hello -> HELLO
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 转换为小写，如：HELLO -> hello
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase 下划线转换为大写驼峰，如：hello_world -> HelloWorld
func UnderscoreToUpperCamelCase(s string) string {
	// 下划线替换为空格
	s = strings.Replace(s, "_", " ", -1)
	// 每个单词首字母转换为大写标题格式
	s = strings.Title(s)
	// 空格替换为空
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 下划线转换为小写驼峰，如：hello_world -> helloWorld
func UnderscoreToLowerCamelCase(s string) string {
	// 转换为大写驼峰
	s = UnderscoreToUpperCamelCase(s)
	// 将首字母转换成小写
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUndersocre 驼峰转换为下划线，如：HelloWorld -> hello_world
func CamelCaseToUndersocre(s string) string {
	var output []rune
	// 遍历字符串
	for i, r := range s {
		// 首字母处理
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		// 遇到大写字母，先加下划线
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		// 转换为小写字母并添加
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
