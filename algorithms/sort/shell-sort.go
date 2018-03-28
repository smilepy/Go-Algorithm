/*
  Time complexity: O(nlogn)
*/
package sort

import "fmt"

func ShellSort(array []int){
	checkArray(array)
	gap:=1
	for ;gap<len(array)/3;gap=gap*3+1{
	}

	for ;gap>0;gap=gap/3{
		fmt.Printf("gap:%d\n",gap)
		for i:=gap;i<len(array);i++{
			current:=array[i]
			j:=i-gap
			for ;j>=0&&array[j]>current;j-=gap{
				array[j+gap]=array[j]
			}
			array[j+gap]=current
		}
	}
}
