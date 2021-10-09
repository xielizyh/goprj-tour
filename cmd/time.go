package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/xielizyh/goprj-tour/internal/timer"
)

// timeCmd 时间命令
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// nowTimeCmd 当前时间子命令
var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTime string
var duration string

// calculateTimeCmd 计算时间子命令
var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			} else if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			var err error
			// 如果解析不是标准格式，按照Unix时间戳处理
			location, _ := time.LoadLocation("Asia/Shanghai")
			if currentTimer, err = time.ParseInLocation(layout, calculateTime, location); err != nil {
				// Parse会尝试读取入参参数的时区信息，如果未指定，会使用UTC时间
				// if currentTimer, err = time.Parse(layout, calculateTime); err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime error: %v", err)
		}
		log.Printf("输出结果：%s, %d", t.Format(layout), t.Unix())
	},
}

func init() {
	// 注册命令
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	// 添加选项
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要计算的时间，有效单位为时间戳或已格式化后的时间`)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns", "us" (or "µ s"), "ms", "s", "m", "h"`)
}
