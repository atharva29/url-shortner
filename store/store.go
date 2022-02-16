package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// fileName stores the name of file where mapping is stored
const fileName = "store.txt"

// InitializeStore creates the store file if not present
func InitializeStore() {
	if _, err := os.Stat(fmt.Sprintf("./%s", fileName)); err == nil {
		return
	}
	file, _ := json.MarshalIndent(map[string]interface{}{}, "", " ")
	_ = ioutil.WriteFile(fileName, file, 0644)
}

// SaveUrlMapping saves the mapping of shortURL with respect to longURL
func SaveUrlMapping(shortUrl string, originalUrl string) error {
	data, err := readJSONFromFile()
	if err != nil {
		return err
	}
	data[shortUrl] = originalUrl
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(fileName, file, 0644)
	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortUrl, originalUrl)
	return nil
}

// RetrieveInitialUrl retrieves the longURL from text file
func RetrieveInitialUrl(shortUrl string) (string, error) {
	data, err := readJSONFromFile()
	if err != nil {
		return "", err
	}
	if _, ok := data[shortUrl]; !ok {
		return "", fmt.Errorf("key doesnt exists for %s", shortUrl)
	}
	return data[shortUrl], nil
}

// readJSONFromFile reads data from file
func readJSONFromFile() (map[string]string, error) {
	plan, _ := ioutil.ReadFile(fileName)
	var data map[string]string
	err := json.Unmarshal(plan, &data)
	if err != nil {
		return map[string]string{}, err
	}
	return data, nil
}
