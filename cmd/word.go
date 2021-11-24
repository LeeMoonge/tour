package cmd

import (
	"github.com/go-programming-tour-book/tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}

// 目前单词所支持的转换模式枚举值
const (
	ModeUpper = iota + 1 // 全部单词转大写
	ModeLower // 全部单词转小写
	ModeUnderScoreToUpperCamelCase // 下划线单词转为大写驼峰单词
	ModeUnderScoreToLowerCamelCase // 下划线单词转为小写驼峰单词
	ModeCamelCaseToUnderscore // 驼峰单词转为下划线单词
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部单词转为大写",
	"2：全部单词转为小写",
	"3：下划线单词转为大写驼峰单词",
	"4：下划线单词转为小写驼峰单词",
	"5：驼峰单词转为下划线单词",
}, "\n")



var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词转换格式",
	Long: desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderScoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderScoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该格式转换，请执行 help word 查看文档！")
		}

		log.Printf("输出结果：%s\n", content)
	},
}



