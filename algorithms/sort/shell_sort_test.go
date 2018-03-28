package sort

import (
	"testing"
	"fmt"
)

func TestShellSort(t *testing.T) {
	array := []int{5, 3, 2, 6, 4, 1,7}
	fmt.Println("Initial array:", array)
	ShellSort(array)
	fmt.Println("SelectSort:", array)
}
