package sort

import (
	"fmt"
	"testing"
)

func Test_MergeSort(test *testing.T){
	array := []int{100,1000,50,-1000,1,3,5,4,10,12,30,31,24,22}
	fmt.Println("Initial array:", array)
	newArray:=MergeSort(array)
	fmt.Println("Sorted array:", newArray)
}
