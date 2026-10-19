package variables

import "fmt"

func GoVariables() {
	// In Go, variables are explicitly declared and used by the compiler.
	// we declare variables using the var keyword and it can declare 1 or more variables
	var color string = "blue"

	// You can also declare multiple variables at once.
	var age, marks int = 34, 80

	fmt.Println(color, age, marks)

	// Go can also infer the type of a variable based on the value
	var plan = "Bombard"
	fmt.Printf("The plan is of type: %T\n", plan)

	// variables declared without initialization are zero-valued
	var goal int
	fmt.Printf("The goal is of type: %T and the value is: %v\n", goal, goal)

	// you can also declare a variable using the shorthand syntax but it can only be used inside a function
	school := "Maranda"
	fmt.Println(school)
}
