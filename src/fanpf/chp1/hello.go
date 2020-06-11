/**
 * @Author: fanpengfei
 * @Description:
 * @File:  hello
 * @Version: 1.0.0
 * @Date: 2020/6/11 10:22
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("input your name:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(-1)
	}
	fmt.Println("input:", input[:len(input)-1])
}
