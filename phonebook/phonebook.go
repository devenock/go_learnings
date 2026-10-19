package phonebook

import (
	"encoding/json"
	"fmt"
	"os"
)

// IMPLEMENT ALL THE CRUD OPERATIONS

// 1. Create a contact
// 2. List all contacts
// 3. Get contact by ID
// 4. Update a contact by ID
// 5. Delete a contact by ID
// 6. Mark contact as saved

// define the contact struct
type Contact struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	IsSaved bool   `json:"is_saved"`
}

// declare the contacts slice
var contacts []Contact
var nextID = 1

func PhoneBook() {
	fmt.Println("Phonebook App!!")
}

// add a contact
func addContact(contact Contact) {
	contacts = append(contacts, contact)
}

// list all contacts
func listContacts() []Contact {
	return contacts
}

// get contact by ID
func getContact(id int) (Contact, error) {
	for _, contact := range contacts {
		if contact.ID == id {
			return contact, nil
		}
	}
	return Contact{}, fmt.Errorf("contact with id %d not found", id)
}

// update contact
func updateContact(id int, name string) error {
	for i := range contacts {
		if contacts[i].ID == id {
			contacts[i].Name = name
			return nil
		}
	}

	return fmt.Errorf("contact with id %d not found", id)
}

// delete contact
func deleteContact(id int) error {
	for i, contact := range contacts {
		if contact.ID == id {
			contacts = append(contacts[:i], contacts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("contact with id %d not found", id)
}

// mark as saved
func markAsSaved(id int) error {
	for i := range contacts {
		if contacts[i].ID == id {
			contacts[i].IsSaved = true
			return nil
		}
	}
	return fmt.Errorf("contact with id %v not found", id)
}

// save contacts (to a json file)
func saveContacts() error {
	data, err := json.MarshalIndent(contacts, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("contacts.json", data, 0644)
}

// load saved contacts
func loadContacts() error {
	data, err := os.ReadFile("contacts.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	err = json.Unmarshal(data, &contacts)
	if err != nil {
		return err
	}
	return nil
}
