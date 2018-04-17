package main

import (
	"algo/insertSort"
	"fmt"
	"algo/bubbleSort"
	"time"
	"algo/mergeSort"
	"algo/quickSort"
)

func main(){
	c:=make(chan bool)
	c2:=make(chan bool)


	numArray:= []int{2,8,4,3,7,5}

	var emp []int
	var emp2 []int
	var emp3 []int
	var emp4 []int

	num := append(emp, numArray...)
	num2:=append(emp2,numArray...)
	num3:=append(emp3,numArray...)
	num4:=append(emp4,numArray...)

	startISort:=time.Now()
	go insertSort.InsertSort(num,c)
	<-c
	elapsed1:=time.Since(startISort)
	fmt.Println("time for Insertsort",elapsed1)


	startBsort:=time.Now()
	go bubbleSort.BubbleSort(num2,c2)
	<-c2
	elapsed2:=time.Since(startBsort)
	fmt.Println("time for bubble sort",elapsed2)

	startMsort:=time.Now()
	mergeSort.MergeFun(num3)

	elapsed3:=time.Since(startMsort)
	fmt.Println("time for Merge sort",elapsed3)

	startQsort:=time.Now()
	quickSort.QuickSort(num4,0,len(num4)-1)
	elapsed4:=time.Since(startQsort)
	fmt.Println("try timefor Merge sort",elapsed4)

	fmt.Println(numArray)

}
