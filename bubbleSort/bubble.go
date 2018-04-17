package bubbleSort

import "fmt"

func BubbleSort(array []int,c2 chan bool){
	var p=array
	fmt.Println("before sort")
	fmt.Println(p)
	var swapped=true
	for swapped{
		swapped=false

		for i:=2;i<len(p);i++{
			if p[i-1]>p[i]{
				temp:=p[i]
				p[i]=p[i-1]
				p[i-1]=temp
				swapped=true
			}
		}
	}
	fmt.Println("bubble array")
	fmt.Println(p)
	c2<-true

}
