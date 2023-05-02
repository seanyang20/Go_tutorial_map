package main

import "fmt"

func main() {
	// ---------------
	// Manipulating Maps
	// ---------------

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// 1st way of declaring empty map
	// var colors map[string]string

	// // 2nd way of declaring empty map
	// colors := make(map[string]string)

	// // add a key-value pair
	// colors["white"] = "#fffff"
	// // delete a key-value pair
	// delete(colors, "white")

	// fmt.Println(colors)

	// -------------
	// Iterating over Maps
	// --------------

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#fffff",
	}

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
