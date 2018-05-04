package sort

import (
	"fmt"
	"testing"
)

func Test_quick_sort(t *testing.T) {
	array := []int{5,4,3,7,2,0,1,30}
	fmt.Println("Initial array:", array)
	QuickSort(array)
	fmt.Println("QuickSort:", array)
}