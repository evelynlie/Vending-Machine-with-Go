package main

import (
	"fmt"
	"os"
)

func displayMainMenu() {
	fmt.Println("Main Menu:")
	fmt.Println("   1. Display Meal Options")
	fmt.Println("   2. Purchase Meal")
	fmt.Println("   3. Save and Exit")
	fmt.Println("Administrator-Only Menu:")
	fmt.Println("   4. Add Food")
	fmt.Println("   5. Remove Food")
	fmt.Println("   6. Display Balance")
	fmt.Println("   7. Abort Program")
}

func saveFoodsToFile(list *LinkedList) {
	// Open the file for writing
	file, err := os.Create("foods.dat")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Loop through the linked list and write each food item to the file
	current := list.Head
	for current != nil {
		food := current.Value
		_, err := file.WriteString(fmt.Sprintf("%s|%s|%s|%.2f\n", food.ID, food.Name, food.Description, food.Price))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		current = current.Next
	}
}

func saveCoinsToFile(coins *[]Coin) {
	// Open the file for writing
	file, err := os.Create("coins.dat")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Loop through the coins slice and write each coin to the file
	for _, coin := range *coins {
		_, err := file.WriteString(fmt.Sprintf("%d|%d\n", coin.Denomination, coin.Quantity))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func saveAndExit(foods *LinkedList, coins *[]Coin) {
	// Save the foods to the foods file
	saveFoodsToFile(foods)

	// Save the coins to the coins file
	saveCoinsToFile(coins)

	fmt.Println("Data saved.")
	fmt.Println()
}

func main() {
	// Check if the user has provided the correct number of arguments
	if len(os.Args) != 3 {
		fmt.Println("Wrong number of arguments.\nUsage: ./vm <foodsfile> <coinsfile>")
		os.Exit(1)
	}

	// Read the foods file
	foods := readFoodsFile()

	// Read the coins file
	coins := readCoinsFile()

	// Display the main menu and abort the program if the user chooses to do so
	for {
		displayMainMenu()

		var choice int
		fmt.Print("Select your option (1-7): ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			fmt.Println()

			// Continue asking for input
			continue
		}

		switch choice {
		case 1:
			displayMealOptions(foods)

		case 2:
			// purchaseMeal(&foods, &coins)

		case 3:
			saveAndExit(&foods, &coins)

		case 4:
			addFood(&foods)

		case 5:
			removeFood(&foods)

		case 6:
			printCoins(coins)

		case 7:
			fmt.Println("Aborting program...")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice. Please try again.")
			fmt.Println()
		}
	}
}
