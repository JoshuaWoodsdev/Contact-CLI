package storage

import (
	"contact-cli/contact"
	"encoding/json"
	"os"
)

// two functions one to save and one to load
func SaveContactsToFile(contacts []contact.Contact, filename string) error {
	data, err := json.Marshal(contacts)
	if err != nil {
		return err
	}
	//return something
	return os.WriteFile(filename, data, 0644)
}

func LoadContactsFromFile(filename string) ([]contact.Contact, error) {
	var contacts []contact.Contact
	//reading data from Readfile
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &contacts)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}
