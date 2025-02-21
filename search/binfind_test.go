package search

import "fmt"

func ExampleFindSorted() {
	l1 := []int{1, 3, 4, 6, 10, 21, 38}

	fmt.Println(FindSorted(l1, 10))
	fmt.Println(FindSorted(l1, 11))

	// Output:
	// 4
	// -1
}
