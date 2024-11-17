# Vending Machine System using Go
This is a simple vending machine system implemented in Go. The system is designed to be used by a vending machine to keep track of the items in the machine and to allow users to purchase items from the machine.

## Features
- Add items to the vending machine
- Remove items from the vending machine
- Purchase items from the vending machine
- View the items in the vending machine
- View the total amount of money in the vending machine
- Save the state of the vending machine to a file

## Data Structures
The system utilizes two core data structures: **Linked List** for food items and **Slice** for coins. Each of these data structures plays an essential role in the functionality of the vending machine.
### 1. Food (Linked List)
The `Food` struct represents an item in the vending machine with the following fields:
- `ID` - The unique identifier of the food item
- `Name` - The name of the food item
- `Description` - The description of the food item
- `Price` (float) - The price of the food item

**Why Linked List?**
A Linked List is used to store food items due to its efficiency when performing operations like inserting and deleting items. In a vending machine, food items may need to be added or removed dynamically, and the Linked List allows for such operations to be done quickly without the need to shift elements as in a slice.

The Linked List is implemented using a `Node` struct, where each node contains a `Food` item and a pointer to the next node. The `LinkedList` struct has methods to:
- Add a food item to the list.
- Remove a food item by its ID.
- Find a food item by its ID.
- Display all food items in the vending machine.
- Read and write the food items data to a file.

### 2. Coin (Slice)
The `Coin` struct represents a coin in the vending machine with:
- `Denomination` (int) - Denomination of the coin in cents (e.g., 5 for 5¢, 100 for $1).
- `Quantity` (int) - Number of coins available for that denomination.

**Why Slice?**
A `Slice` is used to store coins because the number of different coin denominations is fixed, and slices provide efficient random access. This makes the system efficient when managing the available coins for change and making purchases.

The `Slice` of coins is represented as a list of `Coin` structs, and the following operations can be performed:
- Add a coin or its quantity to the list.
- Find a coin by denomination.
- Display the current stock of coins in the vending machine.
- Read and write the coins data to a file.

## Usage
1. To use the vending machine system, you can run the Makefile with the following command to build the executable file, which is called `vm`:
    ```bash
    make
    ```
2. You can then run the `vm` executable file along with a food data file and coins data file with the following command to start the vending machine system:
    ```bash
    ./vm foods.dat coins.dat
    ```

3. Once the vending machine system is running, you can use the following commands to interact with the system through the main menu.

## Commands
- `1` - Display items in vending machine
- `2` - Purchase item from vending machine
- `3` - Save the state of the vending machine to a file (food and coins into foods.dat and coins.dat respectively) and exit the program
- `4` - Add item to vending machine
- `5` - Remove item from vending machine
- `6` - Display total amount of money in vending machine
- `7` - Exit the program without saving the state of the vending machine

## Data Files
- The food data file should be in the following format:
    ```
    id|name|description|price
    ```
- The coins data file should be in the following format:
    ```
    denomination|quantity
    ```

**Note:** 
- The id of the food item should be in the format `F<id>`, where `<id>` is a unique integer identifier for the food item that is of length 4, starting from 0001.
- The denomination of the coin is in cents, starting from 5, which represents 5 cents, and goes on to 5000, which represents $50.