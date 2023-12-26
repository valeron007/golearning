package main

import (
	"fmt"
)

func main() {
	s := "world"
	//buffer := bytes.Buffer{}
	length := len(s)
	var arr = make([]string, length)
	for i, c := range s {
		//fmt.Printf("%c", c)
		arr[i] = string(c)
		println(arr[i])
	}
	fmt.Println()
	fmt.Println(arr)
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res = res + string(s[i])
		/*
			fmt.Println(res)
			fmt.Println(strconv.FormatUint(uint64(s[i]), 10))
			buffer.WriteString(strconv.FormatUint(uint64(s[i]), 10))
		*/
	}
	//fmt.Println(s[len(s):])
	//t := string(buf)
	//var result string = buffer.String()
	//fmt.Println(t)

}
