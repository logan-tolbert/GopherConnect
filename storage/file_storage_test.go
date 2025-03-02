package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"logan-tolbert.com/gopher-connect/models"
)

// TestLoadData_FileExists verifies that loadData correctly loads
// contacts from an existing file containing valid JSON data.
func TestLoadData_FileExists(t *testing.T) {
	// **Arrange** - ensure directory and file exist, create test data, and write to file
	storageDir := filepath.Join("storage", "data")
	dbPath := filepath.Join(storageDir, "contactsDB.json")
	os.MkdirAll(storageDir, os.ModePerm)

	testData := []byte(`[{"FirstName":"John","LastName":"Doe","PhoneNumber":"555-555-555","Email":"johnDoe@email.com"}]`)

	os.WriteFile(dbPath, testData, os.ModeAppend)

	// **Act** - call loadData()
	data, err := loadData()

	// **Assert** - No error & data matches expected
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if string(data) != string(testData) {
		t.Errorf("Expected %s, got %s", string(testData), string(data))
	}

	// **Cleanup**
	os.Remove(dbPath)
}

// TestLoadData_FileDoesNotExist ensures that loadData creates a new contacts file
// and returns an empty JSON array if no file exists.
func TestLoadData_FileDoesNotExit(t *testing.T) {
	// **Arrange & Act** - call `loadData`
	data, err := loadData()

	// **Assert** - No error & should return `[]`
	if err != nil {
		t.Fatalf("Expected no error, go %v", err)
	}
	if string(data) != "[]" {
		t.Errorf("Expected empty JSON array [], got %s", string(data))
	}

	// **Cleanup** - Remove file created by `loadData`
	storageDir := filepath.Join("storage", "data")
	dbPath := filepath.Join(storageDir, "contactsDB.json")
	os.Remove(dbPath)
}

// TestLoadData_InvalidJSON checks that loadData returns an error
// if the contacts file contains invalid JSON.
func TestLoadData_InvalidJSON(t *testing.T) {
	// **Arrange** - ensure directory and file exist, create invalid JSON and write to file
	storageDir := filepath.Join("storage", "data")
	dbPath := filepath.Join(storageDir, "contactsDB.json")
	os.MkdirAll(storageDir, os.ModePerm)

	testData := []byte(`[{invalid JSON}]`)

	os.WriteFile(dbPath, testData, os.ModeAppend)

	// **Act** - call loadData()
	_, err := loadData()

	// **Assert** - should return an error
	if err == nil {
		t.Fatal("Expected and error due to invalid JSON, but got `nil`")
	}

	// **Cleanup**
	os.Remove(dbPath)
}

// TestSaveData_ValidContacts checks that saveData correctly
// writes valid contacts to the JSON file.
func TestSaveData_ValidContacts(t *testing.T) {
	// **Arrange** - Ensure directory and file exist
	storageDir := filepath.Join("storage", "data")
	dbPath := filepath.Join(storageDir, "contactsDB.json")
	os.MkdirAll(storageDir, os.ModePerm)

	testData := []models.Person{
		{
			FirstName: "John",
			LastName:  "Doe",
			BirthDate: time.Time{},
			Phone:     "555-555-555",
			Email:     "johnDoe@email.com",
		},
		{
			FirstName: "Jane",
			LastName:  "Smith",
			BirthDate: time.Time{},
			Phone:     "444-444-444",
			Email:     "janeSmith@email.com",
		},
	}

	// **Act** - Call saveData()
	err := saveData(testData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// **Assert** - Read the saved file
	data, err := os.ReadFile(dbPath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	expectedJSON, _ := json.Marshal(testData)

	if string(data) != string(expectedJSON) {
		t.Errorf("Expected %s, got %s", string(expectedJSON), string(data))
	}

	// **Cleanup**
	os.Remove(dbPath)
}

// TestSaveData_EmptyList ensures that saving
// an empty contact list results in an empty JSON array (`[]`).
func TestSaveData_EmptyList(t *testing.T) {
	// **Arrange** - Ensure directory and file exist
	storageDir := filepath.Join("storage", "data")
	dbPath := filepath.Join(storageDir, "contactsDB.json")
	os.MkdirAll(storageDir, os.ModePerm)

	testData := []models.Person{}

	// **Act** - Call saveData()
	err := saveData(testData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// **Assert** - Read the saved file
	data, err := os.ReadFile(dbPath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	expectedJSON, _ := json.Marshal(testData)

	if string(data) != string(expectedJSON) {
		t.Errorf("Expected %s, got %s", string(expectedJSON), string(data))
	}

	// **Cleanup**
	os.Remove(dbPath)
}
