/*
Time complexity: O(min(m,n)) , 其中m和n为两个字符串的长度
 */
package string

import (
	"testing"
	"fmt"
)

func Test_Concatenate(t *testing.T){
	first:="abcdef"
	second:="fedha"
	fmt.Printf("concatenate result:%s\n",Concatenate(first,second))

}
