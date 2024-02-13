package main

import (
	"fmt"
	"net/mail"
)

/*
Constants and Variables
*/
const appName = "Contact Book"

type Contact struct {
	Name  string
	Phone string
	Email string
	About string
}

// Create slice of Contact
var contacts []Contact

//Create functions and methods that will alter the struct data

// Utiltty functions (single tasks)
func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Core functions    (core logic for program)
func addContact(name, phone, email, note string) {
	if !valid(email) {
		fmt.Println("Invalid email format")
		return
	}

	//If clear delcares and creates a new contact type here define and desribe it
	newContact := Contact{
		Name:  name,
		Phone: phone,
		Email: email,
		About: note,
	}

	//add new contact to the slice
	contacts = append(contacts, newContact)
	fmt.Println("COntact has been added")
}

func main() {
	//Controll flow if than runs in main
	for {
		fmt.Println("Welcome to the ContactBook App")

		fmt.Println("\nPlease choose an option:")
		fmt.Println("1: Add Contact")
		fmt.Println("2: View Contacts")
		fmt.Println("3: Delete Contact")
		fmt.Println("4: Exit")

		//create var to hold user choice
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
		case 2:
		case 3:
		case 4:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}

	//user welcome prompt something that says welcome to the app
	/*write a for loop that wraps a... a case statment that has 3 otpions
		   add contant
		   view contacts
	       delete contact
	*/

	//also I think  things would be I would be setting up for user arguments. we are looking at
	/*
	  1 for the program itself,
	  2 the users choice
	  3 would be for "addcontact and delete"
	*/

}
