package search

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {
	m := len(list) / 2
	if len(list) <= 1 {
		if list[0] == x {
			return 0
		} else {
			return -1
		}
	}
	if list[m] == x {
		return m
	} else if list[m] < x {
		n := FindSorted(list[m:], x)
		if n >= 0 {
			return n + m
		} else {
			return -1
		}
	} else {
		n := FindSorted(list[:m], x)
		if n >= 0 {
			return n
		} else {
			return -1
		}
	}

}
