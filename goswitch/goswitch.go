package goswitch

import "fmt"

func GoSwitch() {
	school := ""
	fmt.Printf("What is the name of your school?")
	fmt.Scan(&school)
	// switch school {
	// case "Maranda":
	// 	fmt.Printf("%v is a national school\n", school)
	// case "Starehe":
	// 	fmt.Printf("%v is a national school\n", school)
	// case "Oriwo":
	// 	fmt.Printf("%v is a national school\n", school)
	// case "Onyege":
	// 	fmt.Printf("%v is a district school\n", school)
	// case "Ober":
	// 	fmt.Printf("%v is an extra-county school\n", school)
	// case "Ngara":
	// 	fmt.Printf("%v is a district school\n", school)
	// default:
	// 	fmt.Printf("%v is a national school\n", school)
	// }

	// REFACTOR OF THE ABOVE CODE
	switch school {
	case "Maranda", "Starehe", "Oriwo":
		fmt.Printf("%v is a national school!\n", school)
	case "Onyege", "Ngara":
		fmt.Printf("%v is a district school!\n", school)
	case "Ober":
		fmt.Printf("%v is an extra-county school!\n", school)
	default:
		fmt.Printf("%v is an ordinary school!\n", school)
	}
}
