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
	formattedTotal := fmt.Sprintf("%.2f", float64(total)/100)
	fmt.Println("Total Value: $", formattedTotal)
	fmt.Println()
}

func addCoin(coins *[]Coin, givenCoins *[]Coin) {
	// Loop through the given coins
	for _, givenCoin := range *givenCoins {
		// Find the coin in the coins slice
		found := false
		for i := range *coins {
			if (*coins)[i].Denomination == givenCoin.Denomination {
				(*coins)[i].Quantity += givenCoin.Quantity
				found = true
				break
			}
		}

		// If the coin is not found, add it to the coins slice
		if !found {
			*coins = append(*coins, givenCoin)
		}
	}
}

func splitIntoDenominations(change float64, coins *[]Coin) string {
	var result []string

	// Iterate over the coins in reverse order since the order is from lowest to highest
	for i := range *coins {
		coin := &(*coins)[i] // Use pointer to modify coin quantities

		// Calculate how many of this coin can be used
		coinValue := float64(coin.Denomination)

		for change >= coinValue && coin.Quantity > 0 {
			change -= coinValue
			coin.Quantity--

			// Format and append the denomination to the result
			if coin.Denomination >= 100 {
				result = append(result, fmt.Sprintf("$%d", coin.Denomination/100))
			} else {
				result = append(result, fmt.Sprintf("%dc", coin.Denomination))
			}
		}
	}

	// If change is still not zero, it means exact change could not be provided
	if change > 0 {
		return "Insufficient coins to provide exact change."
	}

	// Sort the result in descending order
	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	return fmt.Sprintf("Your change is %s", strings.Join(result, " "))
}
