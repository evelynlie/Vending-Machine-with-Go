package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coin struct {
	Denomination int
	Quantity     int
}

// Reads the coins file and returns a slice of Coin structs
func readCoinsFile() []Coin {
	// Open the file
	dat, err := os.ReadFile("coins.dat")

	// If there is an error, panic
	if err != nil {
		panic(err)
	}

	// Create a slice of Coin structs
	coins := make([]Coin, 0)

	// Loop through each line in the file
	for _, line := range strings.Split(string(dat), "\n") {
		if line == "" {
			break
		}

		// Split the line by the comma character
		parts := strings.Split(line, ",")

		// Check if the line has the correct number of parts
		if len(parts) != 2 {
			fmt.Println("Not enough parts in line:", line)
			break
		}

		// Parse the denomination and quantity as an int
		denomination, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error parsing denomination:", err)
			break
		}

		quantity, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error parsing quantity:", err)
			break
		}

		// Create Coin structs
		coin := Coin{
			Denomination: denomination,
			Quantity:     quantity,
		}

		coins = append(coins, coin)
	}

	return coins
}

// Print the balance of coins
func printCoins(coins []Coin) {
	fmt.Println("Balance Summary")
	fmt.Println("-------------")
	fmt.Println("Denom | Quantity | Value")
	fmt.Println("---------------------------")
	// Sort the coins in ascending denomination order
	sort.Slice(coins, func(i, j int) bool {
		return coins[i].Denomination < coins[j].Denomination
	})

	total := 0

	// Loop through each coin
	for _, coin := range coins {
		value := coin.Denomination * coin.Quantity
		fmt.Printf("%-5d | %-8d | %d\n", coin.Denomination, coin.Quantity, value)
		total += value
	}

	fmt.Println("---------------------------")
	fmt.Println("Total Value:", total)
	fmt.Println()
}