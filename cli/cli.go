package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"

	"logan-tolbert.com/gopher-connect/models"
	"logan-tolbert.com/gopher-connect/storage"
)

// Start initializes the CLI loop and displays the menu options.
func Start() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("===================================")
	fmt.Println("  Welcome to GopherConnect! ")
	fmt.Println("  Connecting contacts the Go way! ")
	fmt.Println("===================================")

	for {
		// Display Menu
		fmt.Println("\nChoose an option:")
		fmt.Println("1. List Contacts")
		fmt.Println("2. Add Contacts")
		fmt.Println("3. Get Contact By Name")
		fmt.Println("4. Update Phone Number")
		fmt.Println("5. Update Email Address")
		fmt.Println("6. Delete Contact")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Println("\n----------------------")
			listContacts()
		case "2":
			fmt.Println("\n----------------------")
			addContact(scanner)
		case "3":
			fmt.Println("\n----------------------")
			getContactByName(scanner)
		case "4":
			fmt.Println("\n----------------------")
			updatePhone(scanner)
		case "5":
			fmt.Println("\n----------------------")
			updateEmail(scanner)
		case "6":
			fmt.Println("\n----------------------")
			deleteContact(scanner)
		case "7":
			fmt.Println("\n----------------------")
			fmt.Println("Exiting GopherConnect. Goodbye!")
			return
		default:
			fmt.Println("\n----------------------")
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// isValidName checks if a name contains only alphabetic characters.
func isValidName(name string) bool {
	for _, r := range name {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// isEmptyInput checks if the input is empty.
func isEmptyInput(input string) bool {
	return strings.TrimSpace(input) == ""
}

// addContact prompts the user to enter contact details and adds a new contact.
func addContact(scanner *bufio.Scanner) {
	fmt.Print("Enter First Name: ")
	scanner.Scan()
	firstName := strings.TrimSpace(scanner.Text())
	if !isValidName(firstName) || isEmptyInput(firstName) {
		fmt.Println("Invalid first name. Please use only letters and do not leave it empty.")
		return
	}

	fmt.Print("Enter Last Name: ")
	scanner.Scan()
	lastName := strings.TrimSpace(scanner.Text())
	if !isValidName(lastName) || isEmptyInput(lastName) {
		fmt.Println("Invalid last name. Please use only letters and do not leave it empty.")
		return
	}

	fmt.Print("Enter Phone Number: ")
	scanner.Scan()
	phone := strings.TrimSpace(scanner.Text())
	if !models.IsValidPhoneNumber(phone) {
		fmt.Println("Invalid phone number format.")
		return
	}

	fmt.Print("Enter Email: ")
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())
	if !models.IsValidEmail(email) {
		fmt.Println("Invalid email format.")
		return
	}

	fmt.Print("Enter Birthdate (YYYY-MM-DD): ")
	scanner.Scan()
	birthdateStr := strings.TrimSpace(scanner.Text())

	birthdate, err := time.Parse("2006-01-02", birthdateStr)
	if err != nil {
		fmt.Println("Invalid birthdate format. Using default value.")
		birthdate = time.Time{}
	}

	contact := models.Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Email:     email,
		BirthDate: birthdate,
	}

	storage.AddContact(contact)
	fmt.Println("Contact added successfully!")
}

// listContacts fetches and displays all stored contacts.
func listContacts() {
	contacts := storage.ListContacts()
	if len(contacts) == 0 {
		fmt.Println("No contacts found. Would you like to add one? (y/n)")
		var response string
		fmt.Scanln(&response)
		if strings.ToLower(response) == "y" {
			addContact(bufio.NewScanner(os.Stdin))
		}
		return
	}
	fmt.Println("\nContacts:")
	for _, contact := range contacts {
		fmt.Println(contact)
	}
}

// getContactByName prompts the user to enter a first and last name to retrieve contact details.
func getContactByName(scanner *bufio.Scanner) {
	fmt.Print("Enter First Name: ")
	scanner.Scan()
	firstName := strings.TrimSpace(scanner.Text())
	if !isValidName(firstName) || isEmptyInput(firstName) {
		fmt.Println("Invalid first name. Please use only letters and do not leave it empty.")
		return
	}

	fmt.Print("Enter Last Name: ")
	scanner.Scan()
	lastName := strings.TrimSpace(scanner.Text())
	if !isValidName(lastName) || isEmptyInput(lastName) {
		fmt.Println("Invalid last name. Please use only letters and do not leave it empty.")
		return
	}

	contact := storage.GetContact(firstName, lastName)
	if contact == nil {
		fmt.Println("Contact not found.")
		return
	}

	fmt.Println("\n----------------------")
	fmt.Println(contact.String())
}

// updatePhone prompts the user to update the phone number of an existing contact.
func updatePhone(scanner *bufio.Scanner) {
	fmt.Print("Enter First Name: ")
	scanner.Scan()
	firstName := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter Last Name: ")
	scanner.Scan()
	lastName := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter New Phone Number: ")
	scanner.Scan()
	newPhone := strings.TrimSpace(scanner.Text())

	if !models.IsValidPhoneNumber(newPhone) {
		fmt.Println("Invalid phone number format. Update failed.")
		return
	}

	if storage.UpdatePhone(firstName, lastName, newPhone) {
		fmt.Println("Phone number updated successfully!")
	} else {
		fmt.Println("Contact not found. Update failed.")
	}
}

// updateEmail prompts the user to update the phone number of an existing contact
func updateEmail(scanner *bufio.Scanner) {
	fmt.Print("Enter First Name: ")
	scanner.Scan()
	firstName := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter Last Name: ")
	scanner.Scan()
	lastName := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter New Email: ")
	scanner.Scan()
	newEmail := strings.TrimSpace(scanner.Text())

	if !models.IsValidEmail(newEmail) {
		fmt.Println("Invalid email format. Update failed.")
		return
	}

	if storage.UpdateEmail(firstName, lastName, newEmail) {
		fmt.Println("Email updated successfully!")
	} else {
		fmt.Println("Contact not found. Update failed.")
	}
}

// deleteContact prompts for confirmation before deleting a contact.
func deleteContact(scanner *bufio.Scanner) {
	fmt.Print("Enter First Name: ")
	scanner.Scan()
	firstName := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter Last Name: ")
	scanner.Scan()
	lastName := strings.TrimSpace(scanner.Text())

	fmt.Printf("Are you sure you want to delete %s %s? (y/n): ", firstName, lastName)
	scanner.Scan()
	response := strings.TrimSpace(scanner.Text())
	if strings.ToLower(response) != "y" {
		fmt.Println("Deletion canceled.")
		return
	}

	if storage.DeleteContact(firstName, lastName) {
		fmt.Println("Contact deleted successfully!")
	} else {
		fmt.Println("Contact not found. Deletion failed.")
	}
}
