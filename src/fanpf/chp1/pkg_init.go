/**
 * @Author: fanpengfei
 * @Description:
 * @File:  pkg_init
 * @Version: 1.0.0
 * @Date: 2020/6/11 10:32
 */
/*
TODO:
1. init
2. g_
3. map
*/

package main

import "fmt"

var GInfo = map[string]int{"fanpf": 37, "rose": 35}

func ShowInfo() {
	for s, i := range GInfo {
		fmt.Println(s, ":", i)
	}
}

func SetInfo() {
	GInfo["fanpf"] = 40
}

func init() {
	fmt.Println("init...")
	ShowInfo()
	fmt.Println("init over...")
}

func main() {
	SetInfo()
	ShowInfo()
}
