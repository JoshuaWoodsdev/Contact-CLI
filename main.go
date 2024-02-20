package main

import (
	"bufio"
	"fmt"
	"net/mail"
	"os"
	"strings"
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

// Core functions (core logic for program)
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
	fmt.Println("Contact has been added")
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
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nWelcome to the ContactBook App")

		fmt.Println("\nPlease choose an option:")
		fmt.Println("1: Add Contact")
		fmt.Println("2: Display Contacts")
		fmt.Println("3: Delete Contact")
		fmt.Println("4: Exit")

		fmt.Print("Your choice: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		switch choiceStr {
		case "1":
			fmt.Println("Enter Name:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Println("Enter Phone:")
			phone, _ := reader.ReadString('\n')
			phone = strings.TrimSpace(phone)

			fmt.Println("Enter Email:")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Println("Enter Note:")
			note, _ := reader.ReadString('\n')
			note = strings.TrimSpace(note)

			addContact(name, phone, email, note)
		case "2":
			displayContacts()
		case "3":
			deleteContact()
		case "4":
			fmt.Println("Exiting the application.")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
