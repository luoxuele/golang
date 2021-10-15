package main

import "fmt"

func main() {
    //Golang中map的底层实现是一个散列表，
    //是一种无序的键值对集合

    // var countryCapitalMap map[string]string
    // countryCapitalMap = make(map[string]string)
    countryCapitalMap := make(map[string]string)

    countryCapitalMap["China"] = "北京"
    countryCapitalMap["America"] = "华盛顿"
    countryCapitalMap["English"] = "伦敦"
    countryCapitalMap["日本"] = "dongjing"

    // fmt.Println(countryCapitalMap["China"])

    // fmt.Println(countryCapitalMap,len(countryCapitalMap))

    // fmt.Println(countryCapitalMap["越南"])
    country := "越南"
    capital,ok := countryCapitalMap[country]
    if(ok){
        fmt.Println(country,"的首都是",capital)
    }else{
        fmt.Println("没有找到",country,"的首都")
    }   

    //map的len就是键值对的个数
    delete(countryCapitalMap,"日本") //无返回值
    fmt.Println(len(countryCapitalMap))
    fmt.Println(countryCapitalMap["China"])
}
