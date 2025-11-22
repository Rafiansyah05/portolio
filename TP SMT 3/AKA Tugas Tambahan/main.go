package main
import "fmt"


func selectionSortIterative(TabN []int) {
	var n, i, j, minIndex int
	n = len(TabN)
	for i = 0; i < n-1; i++ {
		minIndex = i
		for j = i + 1; j < n; j++ {
			if TabN[j] < TabN[minIndex] {
				minIndex = j
			}
		}

		TabN[i] = TabN[minIndex]
		TabN[minIndex] = TabN[i]
		
	}
}

func selectionSortRecursive(TabN []int, start int){
	var n, i, minIndex int
	n = len(TabN)
	if start >= n-1 {
		return
	}

	minIndex := start
	for i := start + 1; i < n; i++ {
		if TabN[i] < TabN[minIndex] {
			minIndex = i
		}
	}
	TabN[start] = TabN[minIndex]
	TabN[minIndex] = TabN[start]

	selectionSortRecursive(TabN, start+1)
}


func main{
	
}