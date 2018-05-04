/*
堆排序是选择排序的一种。
Time complexity:O(N*logN)
*/

package sort

var length int

func HeapSort(array []int){
	BuildMaxHeap(array)
	for i:=len(array)-1;i>0;i-- {
		Swap(array,0,i)
		length--
		heapify(array,0)
	}
}

func heapify(array []int, i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < length && array[left] > array[largest] {
		largest = left
	}
	if right < length && array[right] > array[largest] {
		largest = right
	}

	if largest != i {
		Swap(array, i, largest)
		heapify(array, largest)
	}
}

func BuildMaxHeap(array []int) {
	length = len(array)
	for i := length / 2; i >= 0; i-- {
		heapify(array,i)
	}
}
