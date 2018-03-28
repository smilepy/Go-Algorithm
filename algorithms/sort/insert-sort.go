/*
  Time complexity: O(N^2)
*/
package sort

func InsertSort(array []int) {
	checkArray(array)
	// The array is divided into two parts: sorted-part and unsorted-part. 
	// At the beginning, sorted-part contains the first element, and unsorted-part contains the rest.
	for i := 1; i < len(array); i++ {
		// For every element in unsorted-part
		j:=i-1
		current:=array[i]
		for ; j >= 0&& array[j]>current; j-- {
				array[j+1]=array[j]
		}
		// choose the right place in sorted-part to insert.
		array[j+1]=current
	}
}
