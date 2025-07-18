package main

import (
	"fmt"
	"time"
)

func main() {
	str := "2023-04-05"
	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t.Format(layout)) // 输出时间类型的变量
}
