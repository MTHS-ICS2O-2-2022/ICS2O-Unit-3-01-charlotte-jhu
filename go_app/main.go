// Created by: Charlotte Jhu
// Created on: March 2023
//
// This program calculates the area of a trapezoid

package main

import "fmt"

func main() {
	// This function calculates the area of a trapezoid
	// variables
	var aBase float64
	var bBase float64
	var height float64

	// input
	fmt.Println("This program calculates the area of a trapezoid.")
	fmt.Println()
	fmt.Print("Enter the length of the first base (mm): ")
	fmt.Scanln(&aBase)
	fmt.Print("Enter the length of the second base (mm): ")
	fmt.Scanln(&bBase)
	fmt.Print("Enter the height (mm): ")
	fmt.Scanln(&height)

	// process
	area := (aBase + bBase) / 2 * height

	// output
	fmt.Println()
	fmt.Println("The area is ", area, " mm²")
	fmt.Println("\nDone.")
}
