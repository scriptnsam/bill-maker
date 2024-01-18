package main

import (
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
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
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

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