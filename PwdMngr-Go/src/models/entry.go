package models

import "fmt"


// Entry represents a website-password entry
type Entry struct {
	Website  string
	Password string
}


// EntryOps performs various operations on an Entry based on the given operation
func (entry Entry) EntryOps(ops string, EntryDB map[string]string) {
	switch ops {
	case "CREATE":
		create(entry,EntryDB)
	case "READ":
		read(entry, EntryDB)
	case "UPDATE":
		update(entry, EntryDB)
	case "DELETE":
		del(entry.Website, EntryDB)
	}
}

// create adds an entry to the map with the given website and password
func create(entry Entry, EntryDB map[string]string) {

	// Check if website or password is empty
	if entry.Website == "" || entry.Password == "" {
		fmt.Printf("ERROR: Enter valid credentials\n")
	} else {
		// Add the entry's website as the key and the password as the value
		EntryDB[entry.Website] = entry.Password
	}
}

// read retrieves and prints either the password or website depending on the entry
func read(entry Entry, EntryDB map[string]string) {
	val := ""

	// Determine whether to retrieve password or website based on the entry
	if entry.Password == "" {
		val = EntryDB[entry.Website]
	} else {
		for key,value := range EntryDB{
			if value == entry.Password{
				val = key
				break
			}
		}
	}

	// Print the retrieved value or an error message if not found
	if val != "" {
		fmt.Printf("%s",val)
	} else {
		fmt.Print("ERROR: Target not found\n")
	}
}

// update updates the password for a given website, and if the updated password is empty, calls del()
func update(entry Entry, EntryDB map[string]string) {
	pwd := entry.Password
	web := entry.Website

	// Check if the website exists in the map
	if EntryDB[web] != "" {
		// If the password is empty, delete the entry for the given website
		if pwd == "" {
			del(web, EntryDB)
		}
		// Update the password for the given website
		EntryDB[web] = pwd
	} else {
		fmt.Printf("ERROR: Enter valid credentials\n")
	}
}

// del deletes an entry from the map for the given website
func del(web string, EntryDB map[string]string) {
	delete(EntryDB, web)
}

