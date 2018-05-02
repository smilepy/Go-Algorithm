/*
Time complexity: O((i+1)*(j+1)) , 其中j=i/2
动态规划，用f[i][j]表示前i天至多买卖j次(一买一卖记为一次)能得到的最大利润
*/
package problem

import "testing"

func Test_MaxProfit(t *testing.T){
	array:=[]int{0,0, 2, 1, 8, 4,9}
	maxTradingNum := (len(array)-1)/2
	p:=2
	//var f [6][3]int
	var f [][]int
	for i:=0;i<len(array);i++{
		var array []int
		for j:=0;j<maxTradingNum+1;j++{
			array=append(array,0)
		}
		f=append(f,array)
	}
	MaxProfit(array,len(array)-1,maxTradingNum,p,f)
}
