package main

import "fmt"

func main(){
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
	}

	// or
	//var colors map[string]string

	// or
	//colors := make(map[string]string)

	// add color
	colors["white"] = "#fffffff"

	// delete color
	//delete(colors,"white")

	printMap(colors)
	
	//fmt.Println(colors)
}

func printMap(c map[string]string)  {
	for color, hex := range c{
		fmt.Println("Hex code for ", color, "is\t", hex)
	}
}