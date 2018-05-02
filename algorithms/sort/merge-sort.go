/*
归并排序是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。
Time complexity:O(N*logN)
*/
package sort

import "fmt"

func MergeSort(array []int) []int{
	length:=len(array)
	if length<2{
		return array
	}
	middle:=length/2
	fmt.Printf("length/2:%+v\n",middle)
	left:=array[:middle]
	right:=array[middle:]
	return merge(MergeSort(left),MergeSort(right))
}

func merge(left,right []int) []int{
	var result []int
	i:=0
	j:=0
	for ;i<len(left)&&j<len(right);{
		if left[i]<=right[j]{
			result=append(result,left[i])
			i++
		}else{
			result=append(result,right[j])
			j++
		}
	}
	for ;i<len(left);i++{
		result=append(result,left[i])
	}
	for ;j<len(right);j++{
		result=append(result,right[j])
	}

	return result
}
