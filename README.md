# Vending Machine System using Go
This is a simple vending machine system implemented in Go. The system is designed to be used by a vending machine to keep track of the items in the machine and to allow users to purchase items from the machine.

## Features
- Add items to the vending machine
- Remove items from the vending machine
- Purchase items from the vending machine
- View the items in the vending machine
- View the total amount of money in the vending machine
- Save the state of the vending machine to a file

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