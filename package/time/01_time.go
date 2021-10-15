package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		time包
		1 year = 365 day
		1 day = 24 hour
		1 hour = 60 minute
		1 minute = 60 second
		1 second = 1000 millisecond 毫秒ms
		1 millisecond = 1000 microsecond 微秒 us
		1 microsecond = 1000 nanosecond	 纳秒 ns
		1 nanosecond = 1000 picosecond 	皮秒 ps

	*/

	//1. 获取当前时间
	t1 := time.Now()
	//fmt.Printf("%T\n", t1)

	// 2. 获取指定时间
	// t1 = time.Date(1994, 7, 15, 12, 0, 0, 0, time.Local)

	// 3. 把time对象转换成string
	// s1 := t1.Format("2006-1-2 15:4:5")
	// fmt.Println(s1)

	// s1 := t1.String()
	// fmt.Println(s1)

	// 4. 把字符串转成time对象
	// s3 := "1999年10月9日"
	// t1,_ = time.Parse("2006年1月2日",s3)

	//5. 获取time对象的指定类容
	// year, month, day := t1.Date()
	// fmt.Println(year, month, day)

	// hour, min, sec := t1.Clock()
	// fmt.Println(hour, min, sec)
	// fmt.Println(t1.Year())
	// fmt.Println(t1.Month())
	// fmt.Println(t1.Day())
	// fmt.Println(t1.Hour())
	// fmt.Println(t1.Minute())
	// fmt.Println(t1.Second())
	// fmt.Println(t1.Nanosecond())

	// 5. 时间戳 ：指定日期，距离1970.01.01 0:0:0的差值， 秒，纳秒
	//t1 = time.Date(1970,1,1,0,0,0,0,time.UTC)
	// timeStamp := t1.Unix() //秒的差值
	// timeStamp := t1.UnixNano() //纳秒的差值

	// fmt.Println(timeStamp)

	// 6. 时间间隔
	t5 := t1.Add(time.Minute * 24)
	// fmt.Println(t5)
	dd := t5.Sub(t1)
	fmt.Println(dd)

	// time.Sleep(time.Second*3)

	// rand.Seed(time.Now().UnixNano())
	// randNum := rand.Intn(10) + 1
	// fmt.Println("Sleep", randNum, "Seconds")
	// time.Sleep(time.Duration(randNum) * time.Second)
	var fn1 = func() {
		fmt.Println("我是fn1")
	}
	timer := time.AfterFunc(time.Second*3, fn1)
	<-timer.C

	fmt.Println(t1)

}
