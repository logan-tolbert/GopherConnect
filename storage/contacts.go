package storage

import "logan-tolbert.com/gopher-connect/models"

// In-memory data store
var contacts []models.Person

// AddContact adds a new contact to the in-memory storage.
func AddContact(person models.Person) {
	contacts = append(contacts, person)
}

// ListContacts returns a slice containing all stored contacts.
func ListContacts() []models.Person {
	return contacts
}

// GetContact searches for a contact by first and last name.
// It returns a pointer to the found contact or nil if no match is found.
func GetContact(firstName, lastName string) *models.Person {

	for i := range contacts {
		if contacts[i].FirstName == firstName && contacts[i].LastName == lastName {
			return &contacts[i]
		}
	}
	return nil
}

// UpdatePhone updates only the phone number of an existing contact.
// Returns true if successful, false if the contact was not found.
func UpdatePhone(firstName, lastName, newPhone string) bool {
	contact := GetContact(firstName, lastName)
	if contact == nil {
		return false
	}

	if !models.IsValidPhoneNumber(newPhone) {
		return false
	}

	contact.Phone = newPhone
	return true
}

// UpdateEmail updates only the email address of an existing contact.
// Returns true if successful, false if the contact was not found.
func UpdateEmail(firstName, lastName, newEmail string) bool {

	contact := GetContact(firstName, lastName)
	if contact == nil {
		return false
	}

	if !models.IsValidEmail(newEmail) {
		return false
	}

	contact.Email = newEmail
	return true
}

// DeleteContact removes a contact by first and last name.
// Returns true if the contact was deleted, false if not found.
func DeleteContact(firstName, lastName string) bool {
	contact := GetContact(firstName, lastName)
	if contact == nil {
		return false
	}

	for i := range contacts {
		if &contacts[i] == contact {
			contacts = append(contacts[:i], contacts[i+1:]...)
			return true
		}
	}

	return false
}
