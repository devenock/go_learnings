package goarrays

import "fmt"

func GoArrays() {
	// an empty array of 5 ints
	var marks [5]int
	// set the value at an index
	marks[3] = 8

	fmt.Println(marks[2])
	// get the length
	fmt.Printf("My array is %v items longer\n", len(marks))
	// ÷let the system count
	ages := [...]int{23, 40, 34, 56}
	fmt.Println("array is:", ages, " and length is: ", len(ages))
	// multi-dimensional array
	var twoD [2][3]int = [2][3]int{
		{1, 2, 3}, {4, 5, 6},
	}

	fmt.Println("My two dimansional array is:", twoD)
	fmt.Println("Student marks:", marks)
	// nums := []int{1, 2, 3, 4, 5}
	// getSum(nums)
}

func getSum(nums []int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}
