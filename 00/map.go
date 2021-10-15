package main

import (
	"fmt"
	"sort"
)

func main() {
	//Golang中map的底层实现是一个散列表，
	//golang的map是hashmap，是使用数组+链表的形式实现的，使用拉链法消除hash冲突
	//src/runtime/map.go
	//是一种 无序，不重复  的键值对集合

	//初始化map
	//var m map[string]string 只声明，不初始化，只有有个头节点，
	// 1. 用内置的make函数创建一个空map
	m1 := make(map[string]string)
	m1["China"] = "北京"
	m1["America"] = "华盛顿"
	m1["English"] = "伦敦"
	m1["日本"] = "dongjing"
	fmt.Println(m1)

	// 2. 用map字面量，还可以指定初始化值
	m2 := map[string]string{}
	m2["name"] = "呵呵哒"
	fmt.Println(m2)

	m3 := map[string]int{
		"age":   32,
		"alice": 33,
	}
	fmt.Println(m3)

	// 操作map
	m4 := make(map[int]string)

	// 1.添加键值对,再次赋值就是修改
	m4[0] = "hello"
	m4[1] = "呵呵"
	m4[999] = "长长99"
	//m4[0] = "哟哟"
	//fmt.Printf("%T-------len(m4) = %d\n",m4,len(m4))
	// fmt.Println(m4[666]) //key不存在，获取的就是value类型的零值

	//3. key是否存在
	// ok-idiom 判断map是否存在key/value
	// value, ok := m4[666]
	// fmt.Println(value, ok)

	// 2. 删除数据，
	m5 := make(map[string]string)
	m5["name"] = "田昌"
	m5["telphone"] = "17775701417"

	delete(m5,"telphone")
	delete(m5,"hehe") //key存在就删除，不存在删不了，就不删了，也不报错

	fmt.Println(m5)

	//map 的遍历
	m6 := make(map[int]string)
	m6[1] = "红孩儿"
	m6[2] = "小钻风"
	m6[3] = "白骨精"
	m6[6] = "金角大王"
	m6[22] = "白素贞"
	m6[777] = "小青"

	//无序打印
	for k,v := range m6{
		fmt.Println(k,v)
	}

	//有序打印
	keys := make([]int,0,len(m6))
	for k,_ := range m6{
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)+
	
	for i:=0;i<len(keys);i++{
		fmt.Println(keys[i],m6[keys[i]])
	}
	// for _,key := range keys{
	// 	fmt.Println(key,m6[key])
	// }
	

}
