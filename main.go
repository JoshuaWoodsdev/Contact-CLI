package main

/*
Constants and Variables
*/
const appName = "Contact Book"

/*Struct type
  type Contact struct {
   Name  string
   Phone string
   Email string
   About string
  }

  Create slice of Contact
   var contacts []Contact

  Use Bufio to read entire lines of user input
   call reader from bufio
   fmt.Println("command to the user to do something contact name: ")
   name = strings.TrimSpace(name) // use trim to clean the space from user input

   (func isValidEmail(email string) bool {
    at := strings.Index(email, "@")
    dot := strings.LastIndex(email, ".")

    // Check if "@" exists and is not the first character
    // Check if "." exists and is after the "@"
    // Ensure the domain part (after the ".") is not empty
    if at > 0 && dot > at+1 && dot < len(email)-1 {
        return true
    }
    return false
    }) checks email


*/
