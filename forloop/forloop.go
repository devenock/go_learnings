package forloop

import "fmt"

func GoFor() {
	subjects := []string{"chemistry", "swahili", "biology", "maths", "physics"}

	fmt.Println("Here are your subjects:")
	// normal for loop
	// for i := 0; i < len(subjects); i++ {
	// 	fmt.Printf("%v. %v\n", i+1, subjects[i])
	// }

	// using range
	for index, subject := range subjects {

		// you can also skip and continue
		if len(subject) < 7 {
			continue
		}
		fmt.Printf("%v. %v\n", index+1, subject)
	}
}
