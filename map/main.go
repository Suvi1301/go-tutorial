package main

import "fmt"

func main() {

	// declaring map with keys and values of type string.
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func notes() {
	// Ways of declaring maps:
	// 1.
	var map1 map[string]string

	// 2.
	map2 := map[int]string{}

	// 3.
	map3 := make(map[string]string)

	// Add more entries
	map2[1] = "One"

	// delete key value pair from map
	delete(map2, 1)

	fmt.Println(map1, map2, map3)
}
