package quickSort

import "fmt"

func QuickSort(a []int,lo,hi int){
	if lo>hi{
		return
	}
	p:=partition(a,lo,hi)
	QuickSort(a,lo,p-1)
	QuickSort(a,p+1,hi)

	fmt.Println(a)

}
func partition(a []int,lo,hi int)int{
	p:=a[hi]
	for j:=lo;j<hi;j++{
		if a[j]<p{
			a[j],a[hi]=a[hi],a[j]
			lo++
		}
	}
	a[lo],a[hi]=a[hi],a[lo]
	return lo
}
