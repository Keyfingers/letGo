package main

import (
	"fmt"
	"time"
)

func main() {
	// 给定的时间字符串
	timeStr := "2024-03-13T11:15:00.7611867-03:00"

	// 解析时间字符串为 time.Time 类型
	t, err := time.Parse(time.RFC3339Nano, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	// 将时间转换为 UTC
	utcTime := t.UTC()

	// 计算自 Unix 纪元以来的毫秒数
	timestamp := utcTime.UnixNano() / int64(time.Millisecond)

	// 打印时间戳
	fmt.Println("Timestamp:", timestamp)
}
