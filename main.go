package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The function `getInput` takes a prompt string and a `bufio.Reader` as input, and returns a trimmed
// string and an error.
func getInput(p string,r *bufio.Reader)(string,error){
	fmt.Print(p)

	input,err:=r.ReadString('\n')

	return strings.TrimSpace(input), err
}

// The function creates a new bill by taking user input for the bill name and returns the created bill.
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name,_ := getInput("Create a new bill name: ",reader)

	b:=newBill(name)
	fmt.Println("Created the bill - ",b.name)

	return b
}

// The function `promptOptions` allows the user to choose options to add items, save the bill, or add a
// tip to a bill.
func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin)
	opt,_ := getInput("Choose option (a - add item, s - save bill, t - add tip): ",reader)

	switch opt{
	case "a":
		// add item - (name and price)
		name, _ := getInput("Item name: ",reader)
		price, _ := getInput("Item price: ",reader)

		p,err := strconv.ParseFloat(price,64)
		if err != nil{
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name,p)

		fmt.Println("Item added - ",name,price)
		promptOptions(b)
	

	case "t":
		tip, _ := getInput("Enter tip amount ($): ",reader)
		t,err := strconv.ParseFloat(tip,64)
		if err != nil{
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip added - ",tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the bill - ",b.name)
	default:
		fmt.Println("Please choose a valid option...")
		promptOptions(b)

	} 
}

// The main function creates a bill and prompts options for the user.
func main() {
	mybill := createBill()

	promptOptions(mybill)
}