package sort

import (
	"github.com/pkg/errors"
	"fmt"
	"Go-Algorithm/algorithms/sort"
)

var length int

func FindKBiggest(array []int,k int) (int,error) {
	length=len(array)
	if k<1 || k> length-1{
		return 0,errors.New(fmt.Sprintf("%d is unvaliad",k))
	}

	BuildMaxHeap(array)
	number :=0
	for i:= len(array) -1;i>=0;i--{
		sort.Swap(array,0,i)
		if number ==k{
			return array[i],nil
		}
		number++
		length--
		heapify(array,0)
	}

	return 0,errors.New(fmt.Sprintf("%d is unvaliad",k))
}

func BuildMaxHeap(array []int){
	length=len(array)
	for i:= length /2;i>=0;i--{
		heapify(array,i)
	}
}

func heapify(array []int,i int){
	largest:=i
	left:=2*i+1
	right:=2*i+2
	if left< length && array[left]>array[largest]{
		largest=left
	}
	if right< length && array[right]>array[largest] {
		largest=right
	}
	if largest!=i{
		sort.Swap(array,i,largest)
		heapify(array,largest)
	}
}