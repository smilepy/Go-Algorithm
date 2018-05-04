package sort

import (
	"testing"
	"fmt"
)

func Test_HeapSort(t *testing.T){
	array := []int{5, 3, 2, 6, 4, 1}
	fmt.Println("Initial array:", array)
	HeapSort(array)
	fmt.Println("InsertSort:", array)

	array=[]int{7,4,3,10,22,1,100,-6,9,31,77}
	fmt.Println("Initial array:", array)
	HeapSort(array)
	fmt.Println("InsertSort:", array)
}
