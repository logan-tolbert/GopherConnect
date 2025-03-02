package storage

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"

	"logan-tolbert.com/gopher-connect/models"
)

// loadData reads contact data from the JSON file.
//
// If the file exists, it returns the raw JSON data as a byte slice.
// If the file does not exist, it initializes the file with an empty JSON array (`[]`).
// All other errors encountered during file reading are returned.
//
// Returns:
//   - []byte: JSON-encoded contact data or an empty array if the file is missing.
//   - error: Any file read/write error, if applicable.
func loadData() ([]byte, error) {

	var storageDir = filepath.Join("storage", "data")
	var dbPath = filepath.Join(storageDir, "contactsDB.json")

	if err := os.MkdirAll(storageDir, os.ModePerm); err != nil {
		return nil, err
	}

	file, err := os.ReadFile(dbPath)
	if err != nil {
		if os.IsNotExist(err) {

			emptyJsonArr := []byte("[]")

			if err := os.WriteFile(dbPath, emptyJsonArr, fs.ModePerm); err != nil {
				return nil, err
			}
			return emptyJsonArr, nil
		}
		return nil, err
	}
	return file, nil
}

// saveData writes contact data to the JSON file.
//
// It ensures the storage directory exists before writing.
// The function converts the provided slice of contacts into JSON format
// and saves it to the database file, overwriting any existing data.
// Any errors encountered during JSON conversion or file writing are returned.
//
// Parameters:
//   - contacts []models.Person: The list of contacts to be saved.
//
// Returns:
//   - error: An error if the JSON conversion or file write operation fails.
func saveData(contacts []models.Person) error {

	var storageDir = filepath.Join("storage", "data")
	var dbPath = filepath.Join(storageDir, "contactsDB.json")

	if err := os.MkdirAll(storageDir, os.ModePerm); err != nil {
		return err
	}

	data, err := json.Marshal(contacts)

	if err != nil {
		return err
	}

	err = os.WriteFile(dbPath, data, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil

}
