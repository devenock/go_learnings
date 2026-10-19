package ifelse

import "fmt"

func GoIfElse() {
	fmt.Println("Students Marks")
	// marks := make([]int, 10)
	marks := []int{5, 6, 9, 10, 7, 0, 40, 20, 1, 4}
	for _, mark := range marks {
		// fmt.Println(mark)
		if mark < 5 {
			fmt.Printf("%v\nComment: %d is below average\n", mark, mark)
		} else {
			fmt.Printf("%v\nComment: Congratulations! You passed.\n", mark)
		}
	}
}
