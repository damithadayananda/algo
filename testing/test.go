package main

import (
	"fmt"
)

func aaa(p *[5]int) int{


	for i:=1;i<5;i++{
		p[i-1]=i
	}
	return p[3]

}

func bbb(q *[10]int) int{
	for i:=1;i<10;i++{
		q[i-1]=i
	}
	return q[3]
}
func print(a int,b int){
	fmt.Println(a*b)
}