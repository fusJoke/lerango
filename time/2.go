package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()

	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)

	timeObj := time.Unix(timestamp1, 0)
	fmt.Println(timeObj)
	year   := timeObj.Year()
	month  := timeObj.Month()
	day    := timeObj.Day()
	hour   := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}