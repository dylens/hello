package main

import (
	//net "code.google.com/p/go.net/websocket"
	out "fmt"
	"math"
)

type name string

type (
	字符串 string
	s   字符串
)

func main() {
	//godoc -http=:6060
	var 玳 s
	玳 = "小玳"
	var y s = "yang"
	dye := 玳 + y
	out.Println(dye)
	out.Println(math.MaxInt32)
	var arr [10]int
	arr[0] = 42
	arr[1] = 12

	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [2][4]int{{1, 2, 3, 4}, {2, 3, 4, 5}}

	slice := []int{2, 3, 5}
	if arr[0]==42 {
		arr[1]=13
	}
}
