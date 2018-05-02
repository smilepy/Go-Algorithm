package problem

import (
	"fmt"
)

func MaxProfit(array []int, n, m, p int,f [][]int) {
	for j := 1; j <= m; j++ {
		f[1][j] = 0
		t := -array[1] - p
		for i := 2; i <=n; i++ {
			f[i][j] = max(f[i-1][j], max(f[i][j-1], array[i]+t))
			if t < f[i-1][j-1]-array[i]-p {
				t = f[i-1][j-1] - array[i] - p
			}

		} //for
	} //for

	fmt.Printf("f[%d][%d]:%d \n", n,m,f[n][m])

}

func max(first, second int) int {
	if first < second {
		return second
	} else {
		return first
	}
}
