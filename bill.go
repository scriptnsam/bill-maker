package main

import (
	"fmt"
	"os"
	"strings"
)

// The "bill" type represents a bill with a name, a map of items and their prices, and a tip amount.
// @property {string} name - The name property is a string that represents the name of the bill.
// @property items - The "items" property is a map that stores the name of each item as the key and its
// corresponding price as the value.
// @property {float64} tip - The "tip" property is a float64 type, which represents the amount of tip
// to be added to the bill.
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
// The function `newBill` creates a new bill with a given name and initializes an empty map for items
// and a tip of 0.
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
// The `format()` function is a method of the `bill` struct. It formats the bill information into a
// string that represents the breakdown of the bill.
func (b *bill) format() string {
	fs := fmt.Sprintf("Bill Breakdown for (%v) - : \n",strings.ToUpper(b.name))
	var total float64 = 0

	// list items
	for k,v := range b.items{
		fs += fmt.Sprintf("%-25v ...$%v \n",k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v \n","tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f \n","total:", total+b.tip)

	return fs

}

// update tip
// The `updateTip` function is a method of the `bill` struct. It takes a `tip` parameter of type
// `float64` and updates the `tip` field of the `bill` struct with the provided value. This function
// allows you to update the tip amount for a specific bill.
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
// The `addItem` function is a method of the `bill` struct. It takes two parameters, `name` of type
// `string` and `price` of type `float64`. This function is used to add an item to the bill by updating
// the `items` map of the `bill` struct. It assigns the `price` to the `name` key in the `items` map.
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// The `save()` function is a method of the `bill` struct. It is used to save the bill information to a
// file.
func (b *bill) save() {
	data := []byte(b.format())

	_, oserr := os.Stat("bills")
	if os.IsNotExist(oserr){
		// directory does not exist, so create it

		e := os.Mkdir("bills",0777)
		if e != nil{
			panic(e)
		}
	}

	err := os.WriteFile("bills/"+b.name+".txt",data,0644)

	if err != nil{
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}