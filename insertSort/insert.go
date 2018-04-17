package insertSort

import "fmt"

func InsertSort(array []int,c chan bool){

	var a=array
	fmt.Println("before sort")
	fmt.Println(a)
	for j:=2;j<len(a);j++{
		var key=a[j]
		i:=j-1
		for i>0 && a[i]>key{
			a[i+1]=a[i]
			i=i-1
			a[i+1]=key
		}
	}
	fmt.Println("insert sort")
	fmt.Println(a)
	c<-true
}
