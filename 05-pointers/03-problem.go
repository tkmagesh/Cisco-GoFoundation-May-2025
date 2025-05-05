/*
Build an interactive calculator. When the program is executed

Display the following menu
1. Add
2. Subtract
3. Multiply
4. Divide
5. Exit
Accept the user choice

if the choice == 5
	exit the application

if the choice is >= 1 and <= 4
	accept two operands (numbers)
	perform the corresponding operation
	print the result
	DISPLAY THE MENU AGAIN

else
	print "invalid choice"
	DISPLAY THE MENU AGAIN

*/

package main

import "fmt"

func main() {
	var userChoice, n1, n2, result int
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice :")
		fmt.Scanln(&userChoice)

		if userChoice == 5 {
			break
		}

		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}

		fmt.Println("Enter the operands :")
		fmt.Scanln(&n1, &n2)
		switch userChoice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		}
		fmt.Println("Result :", result)
	}
}
