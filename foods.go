package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Food struct {
	ID          string
	Name        string
	Description string
	Price       float64
}

// Reads the foods file and returns a slice of Food structs
func readFoodsFile() LinkedList {
	// Open the file
	dat, err := os.ReadFile("foods.dat")

	// If there is an error, panic
	if err != nil {
		panic(err)
	}

	// Create a linked list of Food structs
	foods := LinkedList{}

	// Loop through each line in the file
	for _, line := range strings.Split(string(dat), "\n") {
		if line == "" {
			break
		}

		// Split the line by the pipe character
		parts := strings.Split(line, "|")

		// Check if the line has the correct number of parts
		if len(parts) != 4 {
			fmt.Println("Not enough parts in line:", line)
			break
		}

		// Parse the price as a float64
		price, err := strconv.ParseFloat(parts[3], 64)
		if err != nil {
			// Handle the error, for example:
			fmt.Println("Error parsing price:", err)
			break
		}

		// Create a Food struct
		food := Food{
			ID:          parts[0],
			Name:        parts[1],
			Description: parts[2],
			Price:       price,
		}

		foods.Add(food)
	}
	return foods
}

// Purchase Meal Function
func purchaseMeal(foods *LinkedList, coins *[]Coin) {
	fmt.Println("Purchase Meal")
	fmt.Println("-------------")
	fmt.Print("Enter the ID of the food item you want to purchase: ")

	// Read the ID of the food item to purchase
	var id string
	fmt.Scan(&id)

	// Find the food item in the linked list
	selectedFood := foods.Find(id)
	fmt.Println("You have selected \""+selectedFood.Name, "-", selectedFood.Description+"\". This will cost you $", selectedFood.Price, ".")
}

// Add Food Function
func addFood(foods *LinkedList) LinkedList {
	// ID is assigned by the system (e.g. F...) where ... is the next number in the sequence and has 4 digits
	var id string
	if (*foods).Head == nil {
		id = "F0001"
	} else {
		// Find the last food item in the slice
		current := (*foods).Head
		for current.Next != nil {
			current = current.Next
		}

		// Increment the ID of the last food item
		lastID, _ := strconv.Atoi(current.Value.ID[1:])

		// Create the new ID
		id = "F" + fmt.Sprintf("%04d", lastID+1)
	}

	fmt.Println("This new meal item will have the Item id of", id+".")

	// Read the food details from the user
	var name string
	fmt.Print("Enter the item name: ")
	fmt.Scan(&name)

	var description string
	fmt.Print("Enter the item description: ")
	fmt.Scan(&description)

	var price float64
	price = -1

	// Since the price must be at least a denomination of 0.05, we will keep asking the user for a valid price
	// Price must be a positive number, and must be divisible by 0.05
	for price <= 0 || int(price*100)%5 != 0 {
		fmt.Print("Enter the item price: ")
		fmt.Scan(&price)

		if price <= 0 {
			fmt.Println("Price must be a positive number.")
		} else if int(price*100)%5 != 0 {
			fmt.Println("Price must be divisible by 0.05.")
		}
	}

	// Create a new Food struct
	food := Food{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}

	fmt.Println("This item \""+food.Name, "-", food.Description+"\" has now been added to the food menu.")

	// Append the new food to the slice
	(*foods).Add(food)

	return *foods
}

// Remove Food Function
func removeFood(foods *LinkedList) LinkedList {
	fmt.Println()

	// Display the meal options
	displayMealOptions(*foods)

	var id string
	fmt.Print("Enter the ID of the food item you want to remove: ")

	// Read the ID of the food item to remove
	fmt.Scan(&id)

	selectedFood := foods.Find(id)

	// Remove the food item from the linked list
	(*foods).Remove(id)

	fmt.Println("\""+selectedFood.ID, "-", selectedFood.Name, "-", selectedFood.Description+"\"has been removed from the food menu.")

	return *foods
}

// Displays the meal options
func displayMealOptions(foods LinkedList) {
	fmt.Printf("Food Menu\n")
	fmt.Println("---------")
	fmt.Printf("%-5s | %-30s | %s\n", "ID", "Name", "Price")
	fmt.Println("------------------------------------------------")
	foods.Display()
}