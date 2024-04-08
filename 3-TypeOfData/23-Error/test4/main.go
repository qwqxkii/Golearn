package main

import (
	"fmt"
	"io/ioutil"
)

//练习题 4: 使用 fmt.Errorf() 携带原始错误信息
//题目：。。如果ioutil.ReadFile返回一个错误，使用fmt.Errorf()来创建一个新的错误，该错误应该携带原始错误信息，并说明无法读取配置文件。

// 假设你有一个ReadConfig函数，该函数从给定的文件路径读取配置
func ReadConfig(config string) error {
	//使用ioutil.ReadFile来读取文件内容
	_, err := ioutil.ReadFile(config)
	if err != nil {
		return fmt.Errorf("出现报错%v", err)
	}
	return nil

}

func main() {
	err := ReadConfig("1qwq.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("读取成功")
	}

}
