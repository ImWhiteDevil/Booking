package main

import "fmt"

func main() {

	name := "soham"

	char := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// takes char,num,name as input argument
	trimArray(char, num, name)

}

// trimArray takes argument and trim it then prints it.
func trimArray(char []string, num []int, name string) {

	fmt.Printf("name :%s\n", name)

	// prints first 5 charaters of []char and appends to next line
	fmt.Printf("First 5 characters of given slice of string are : %s\n", char[:5])

	// prints first 5 charaters of []num and appends to next line
	fmt.Printf("First 5 characters of given slice of int are : %d\n", num[:5])

	// prints last 5 charaters of []char and appends to next line
	fmt.Printf("Last 5 characters of given slice of string are : %s\n", char[5:])

	// prints last 5 charaters of []num and appends to next line
	fmt.Printf("Last 5 characters of given slice of int are : %d\n", num[5:])

	// prints first 2 charaters of []char and appends to next line
	fmt.Printf("First 2 characters of given slice of string are : %s\n", char[:2])

	// prints first 2 charaters of []num and appends to next line
	fmt.Printf("First 2 characters of given slice of int are : %d\n", num[:2])

	// prints last 2 charaters of []char and appends to next line
	fmt.Printf("Last 2 characters of given slice of string are : %s\n", char[8:])

	// prints last 2 charaters of []num and appends to next line
	fmt.Printf("Last 2 characters of given slice of int are : %d\n", num[8:])

	// prints first  charater of string name and appends to next line
	fmt.Printf("First character of given string is : %s\n", name[:1])

	// prints without first  charater of string name and appends to next line
	fmt.Printf("Without First character of given string : %s\n", name[1:])

	// prints last  charater of string name and appends to next line
	fmt.Printf("Last character of given string is : %s\n", name[len(name)-1:])

	// prints without last charater of string name.
	fmt.Printf("Without Last character of given string : %s", name[:len(name)-1])
}
