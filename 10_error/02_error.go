package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	//error:内置的一个接口
	// type error interface {
	// 	Error() string
	// }

	// errors.New()
	// fmt.

	err1 := errors.New("自己创建玩的。。。")
	fmt.Println(err1)
	fmt.Printf("%T\n", err1)

	err2 := fmt.Errorf("错误信息码 ： %d", 100)
	fmt.Println(err2)
	fmt.Printf("%T\n", err2)

	err3 := checkAge(-11)
	if err3 != nil{
		log.Fatal(err3)
	}


}

//设计一个函数： 验证年龄是否合法，不合法就返回一个error
func checkAge(age int) error {
	if age < 0 {
		//return errors.New("年龄不合法")
		err := fmt.Errorf("你给定的年龄是：%d,不合法",age)
		return err
	}
	return nil
}
