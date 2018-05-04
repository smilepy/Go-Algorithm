/*
Time complexity: O((k+1)*logN)
find the k-th largest value of L 实际上就是求原数组中的第(k+1)大的元素,因此此处选用堆排序比较合适.
 */
package sort

import (
	"testing"
	"fmt"
)

func Test_FindKBiggest(t *testing.T){
	array:=[]int{1,7,3,2,10,5,30,4}
	k:=1
	value,err:=FindKBiggest(array,k)
	if err!=nil{
		fmt.Printf("err:%s\n",err.Error())
	}else{
		fmt.Printf("the %d-th largest value of L is %d\n",k,value)
	}
}
