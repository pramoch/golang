package main

import "fmt"

func main() {
	colors1 := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}
	delete(colors1, "red")
	fmt.Println(colors1)

	colors2 := make(map[string]string)
	colors2["white"] = "#ffffff"
	fmt.Println(colors2)

	colors3 := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	printMap(colors3)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex color for " + color + " is " + hex)
	}
}
