package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Animal represents the structure used to store JSON data.
// We will be unmarshalling JSON data into this struct.
type Animal struct {
	Data map[string]string
}

func (a *Animal) UnmarshalJSON(data []byte) error {
	rawMap := make(map[string]any)

	// Unmarshal the data from JSON into rawMap
	err := json.Unmarshal(data, &rawMap)
	if err != nil {
		return err // If there is an error, return it
	}
	a.Data = make(map[string]string)

	// Loop over the temporary map (i.e., rawMap) and convert the values to strings.
	for key, value := range rawMap {
		// Use reflection to get a Value object of the value
		val := reflect.ValueOf(value)

		// If the value is a string, store it in the Data map of the Animal struct.
		// If the value is not a string, it is ignored.
		if val.Kind() == reflect.String {
			a.Data[key] = val.String()
		}
	}

	// If everything went well, return nil (i.e., no error)
	return nil

}

func main() {
	// JSON data with variable field names
	jsonData := []byte(`{
		"name": "Elephant",
		"diet": "Herbivore",
		"size": "Large",
		"a_random_field": "Random Value",
		"1": 1
	}`)

	var animal Animal // Initialize an Animal object to store the unmarshalled JSON data

	// Call the Unmarshal function to parse the JSON data into the Animal object
	err := json.Unmarshal(jsonData, &animal)

	// If there is an error during unmarshalling, print the error and stop the execution
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Print the unmarshalled JSON data
	fmt.Printf("Parsed Animal Data: %+v\n", animal.Data)
}
