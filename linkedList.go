package main

import "fmt"

type Node struct {
	Value Food
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

// Function to add a food item to the linked list
func (ll *LinkedList) Add(food Food) {
	newNode := &Node{Value: food}
	if ll.Head == nil {
		ll.Head = newNode
		return
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

// Function to remove a food item from the linked list by its ID
func (ll *LinkedList) Remove(id string) {
	if ll.Head == nil {
		return
	}
	if ll.Head.Value.ID == id {
		ll.Head = ll.Head.Next
		return
	}
	current := ll.Head
	for current.Next != nil {
		if current.Next.Value.ID == id {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// Function to display the food items in the linked list
func (ll *LinkedList) Display() {
	// Check if the list is empty
	if ll.Head == nil {
		fmt.Println("No food items available.")
		return
	}

	// Start with the head of the linked list
	current := ll.Head
	for current != nil {
		food := current.Value

		// Print the food details
		fmt.Printf("%-5s | %-30s | $%.2f\n", food.ID, food.Name, food.Price)

		// Move to the next node in the list
		current = current.Next
	}
}

// Find the length of the linked list
func (ll *LinkedList) Length() int {
	// Start with the head of the linked list
	current := ll.Head
	length := 0
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

// Find a food item by its ID
func (ll *LinkedList) Find(id string) *Food {
	// Start with the head of the linked list
	current := ll.Head
	for current != nil {
		if current.Value.ID == id {
			return &current.Value
		}
		current = current.Next
	}
	return nil
}
