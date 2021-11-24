package cmd

import (
	"fmt"
	"github.com/go-programming-tour-book/tour/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime string
var duration string

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算" +
		"的时间，有效单位为时间戳或以格式化后的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "持续时间，有效时间" +
		"单位为`ns`,`us`(or`us`),`ms`,`s`,`m`,`h`")
}

var timeCmd = &cobra.Command{
	Use: "time",
	Short: "时间格式处理",
	Long: "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nowTimeCmd = &cobra.Command{
	Use: "now",
	Short: "获取当前时间",
	Long: "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("数据结果：%s，%d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
		// 可以使用预定义格式时间： nowTime.Format(time.RFC3339)
	},
}

var calculateTimeCmd = &cobra.Command{
	Use: "calc",
	Short: "计算所需时间",
	Long: "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		location, _ := time.LoadLocation("Asia/Shanghai")
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			fmt.Println("space: ", space)
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			//currentTimer, err = time.Parse(layout, calculateTime)
			// Parse解析时传入Local时区
			currentTimer, err = time.ParseInLocation(layout, calculateTime, location)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalcalateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err：%v", err)
		}

		log.Printf("输出结果：%s，%d", t.Format(layout), t.Unix())
	},
}