package storage

import (
	"os"
	"testing"
	"time"

	"logan-tolbert.com/gopher-connect/models"
)

func TestMain(testRunner *testing.M) {
	exitCode := testRunner.Run()
	os.Exit(exitCode)
}

func TestAddContact(test *testing.T) {
	// **Arrange** - Reset storage and create a test contact
	contacts = []models.Person{}

	testContact := models.Person{
		FirstName: "Logan",
		LastName:  "Tolbert",
		BirthDate: time.Date(1988, 7, 28, 0, 0, 0, 0, time.UTC),
	}

	// **Act** - Add the contact
	AddContact(testContact)

	// **Assert** - Verify the contact was added
	contactsList := ListContacts()
	if len(contactsList) != 1 {
		test.Fatalf("Expected 1 contact, got %d", len(contactsList))
	}

	// **Assert** - Verify the correct contact was added
	addedContact := contactsList[0]
	if addedContact.FirstName != testContact.FirstName || addedContact.LastName != testContact.LastName {
		test.Errorf("Expected %q %q, but got %q %q", testContact.FirstName, testContact.LastName, addedContact.FirstName, addedContact.LastName)
	}
}

func TestGetContact(test *testing.T) {
	// **Arrange** - Reset storage and create a test contact
	contacts = []models.Person{}

	testContact := models.Person{
		FirstName: "Logan",
		LastName:  "Tolbert",
		BirthDate: time.Date(1988, 7, 28, 0, 0, 0, 0, time.UTC),
	}

	// **Act** - Add the contact and retrieve it
	AddContact(testContact)
	result1 := GetContact("Logan", "Tolbert")
	result2 := GetContact("Unknown", "Person")

	// **Assert** - Verify the contact was added correctly
	if result1 == nil {
		test.Fatalf("Expected to retrieve 'Logan Tolbert', but got nil")
	}

	// Verify the contact
	if result1.FirstName != testContact.FirstName || result1.LastName != testContact.LastName {
		test.Errorf("Expected %q %q, but got %q %q", testContact.FirstName, testContact.LastName, result1.FirstName, result1.LastName)
	}

	// Verify nil result
	if result2 != nil {
		test.Errorf("Expected nil for non-existent contact, but got %v", result2)
	}
}

func TestListContacts(test *testing.T) {
	// **Arrange** - Reset storage and create test contacts
	contacts = []models.Person{}

	testContacts := []models.Person{
		{
			FirstName: "John",
			LastName:  "Doe",
			BirthDate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Phone:     "1234567890",
			Email:     "john.doe@example.com",
		},
		{
			FirstName: "Jane",
			LastName:  "Smith",
			BirthDate: time.Date(1985, 5, 20, 0, 0, 0, 0, time.UTC),
			Phone:     "9876543210",
			Email:     "jane.smith@example.com",
		},
	}

	// **Act** - Add contacts
	for _, contact := range testContacts {
		AddContact(contact)
	}

	// Retrieve contacts
	result := ListContacts()

	// **Assert** - Verify correct number of contacts
	if len(result) != len(testContacts) {
		test.Fatalf("Expected %d contacts, got %d", len(testContacts), len(result))
	}

	// Verify each contact matches
	for i, contact := range testContacts {
		if result[i].FirstName != contact.FirstName || result[i].LastName != contact.LastName {
			test.Errorf("Expected %q %q, but got %q %q", contact.FirstName, contact.LastName, result[i].FirstName, result[i].LastName)
		}
	}
}

func TestUpdatePhone(test *testing.T) {
	// **Arrange** - Reset storage and create test contact
	contacts = []models.Person{}
	testContact := models.Person{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
	}
	AddContact(testContact)

	// **Act** - Update phone number
	success := UpdatePhone("John", "Doe", "5556667777")
	invalidUpdate := UpdatePhone("John", "Doe", "invalid-phone")
	nonExistent := UpdatePhone("Jane", "Smith", "8889990000")

	// **Assert** - Verify update was successful
	result := GetContact("John", "Doe")
	if !success || result.Phone != "5556667777" {
		test.Errorf("Expected phone to be updated to '5556667777', but got %q", result.Phone)
	}

	// Verify invalid phone update fails
	if invalidUpdate {
		test.Error("Expected invalid phone update to fail, but it succeeded")
	}

	// Verify updating non-existent contact fails
	if nonExistent {
		test.Error("Expected update for non-existent contact to fail, but it succeeded")
	}
}

func TestUpdateEmail(test *testing.T) {
	// **Arrange** - Reset storage and create test contact
	contacts = []models.Person{}
	testContact := models.Person{
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "old.email@example.com",
	}
	AddContact(testContact)

	// **Act** - Update email
	success := UpdateEmail("Jane", "Smith", "new.email@example.com")
	invalidUpdate := UpdateEmail("Jane", "Smith", "invalid-email")
	nonExistent := UpdateEmail("John", "Doe", "test@example.com")

	// **Assert** - Verify update was successful
	result := GetContact("Jane", "Smith")
	if !success || result.Email != "new.email@example.com" {
		test.Errorf("Expected email to be updated to 'new.email@example.com', but got %q", result.Email)
	}

	// Verify invalid email update fails
	if invalidUpdate {
		test.Error("Expected invalid email update to fail, but it succeeded")
	}

	// Verify updating non-existent contact fails
	if nonExistent {
		test.Error("Expected update for non-existent contact to fail, but it succeeded")
	}
}

func TestDeleteContact(test *testing.T) {
	// **Arrange** - Reset storage and create test contact
	contacts = []models.Person{}
	testContact := models.Person{
		FirstName: "Alice",
		LastName:  "Brown",
	}
	AddContact(testContact)

	// **Act** - Delete the contact
	success := DeleteContact("Alice", "Brown")
	nonExistent := DeleteContact("Bob", "Marley")

	// **Assert** - Verify contact was deleted
	if !success {
		test.Error("Expected successful deletion, but delete failed")
	}

	// Ensure the contact is no longer retrievable
	result := GetContact("Alice", "Brown")
	if result != nil {
		test.Error("Expected contact to be deleted, but it still exists")
	}

	// Ensure deleting a non-existent contact returns false
	if nonExistent {
		test.Error("Expected deletion of non-existent contact to fail, but it succeeded")
	}
}
