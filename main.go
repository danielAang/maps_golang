package main

import "fmt"

func main() {
	/* colors := make(map[string]string)
	colors["red"] = "#ff0000"
	fmt.Println(colors)
	delete(colors, "red")
	fmt.Println(colors) */
	/* var colors map[string]string */
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("key:", key, "=> value:", value)
	}
}
