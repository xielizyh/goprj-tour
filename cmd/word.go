package cmd

import (
	"log"
	"strings"

	word "github.com/xielizyh/goprj-tour/internal/word"

	cobra "github.com/spf13/cobra"
)

// 单词转换模式
const (
	ModeUpper                      = iota + 1 // 全部转大写
	ModeLower                                 // 全部转小写
	ModeUnderscoreToUperCamelCase             // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下划线转小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)

var str string
var mode int8

// desc 长的帮助描述
var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1: 全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUndersocre(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}
		log.Printf("输出结果：%s", content)
	},
}

// init 每个包可以有多个 init 函数，执行顺序是按照文件名执行
func init() {
	// 对命令行参数str,mode的解析和绑定
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的格式")
}
