package sorting

import (
	"fmt"
	"sort"
	"strings"
)

func GoSorting() {
	fmt.Println("Go Sorting!!")

	sortFind()

	sortFloat64s()
}

// Find method
// Find uses a binary search to find and return the smallest index i in [0,n]. If there is no such index i, it returns i = n.
// EXAMPLE
func sortFind() {
	a := []string{"Mike", "Molly", "Burn", "Bonkil"}
	for _, val := range []string{"Brian", "Burn"} {
		i, found := sort.Find(len(a), func(i int) int {
			return strings.Compare(val, a[i])
		})
		if found {
			fmt.Printf("found %s at index %d\n", val, i)
		} else {
			fmt.Printf("%s not found, would insert at %d\n", val, i)
		}
	}
}

// Float64s
// Float64s sorts a slice of float64 in increasing order. Not-a-number(NaN)values are ordered before other values.
// Note: in go 1.22, this function simply calls `slices.Sort`
func sortFloat64s() {
	a := []float64{1.2, 5.6, 8.9, 4.5, -3.9, 0.2}
	sort.Float64s(a)
	fmt.Println(a)
}
