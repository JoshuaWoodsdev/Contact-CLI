package main

import (
	"bufio"
	"fmt"
	"net/mail"
	"os"
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

func displayContacts() {
	if len(contacts) == 0 {
		fmt.Println("No contacts are avaible")
		return
	}

	fmt.Println("Availble contacts")
	for _, contact := range contacts {
		fmt.Println("- Name: " + contact.Name + ", Email: " + contact.Email)
	}
}

// delete contact
func deleteContact() {
	displayContacts()

	var nameToDelete string
	fmt.Println("\nEnter the name of the contact to delete:")
	fmt.Scanln(&nameToDelete)

	found := false
	for i, contact := range contacts {
		if contact.Name == nameToDelete {
			contacts = append(contacts[:i], contacts[i+1:]...) // Corrected line
			fmt.Println("Contact deleted successfully.")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Contact not found.")
	}
}

func main() {
	//Controll flow if than runs in main
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Welcome to the ContactBook App")

		fmt.Println("\nPlease choose an option:")
		fmt.Println("1: Add Contact")
		fmt.Println("2: Display Contacts")
		fmt.Println("3: Delete Contact")
		fmt.Println("4: Exit")

		//create var to hold user choice
		var choice int
		fmt.Scanln(&choice)

		fmt.Println("Your choices")

		switch choice {
		case 1: //addContact func
		case 2: //view contact func do a loop to view the splice print to cli
		case 3: //delete
		case 4:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}

}
