package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type StorageService struct {
	URLMapping map[string]string
}

const fileName = "store.txt"

func InitializeStore() {
	if _, err := os.Stat(fmt.Sprintf("./%s", fileName)); err == nil {
		return
	}
	file, _ := json.MarshalIndent(map[string]interface{}{}, "", " ")
	_ = ioutil.WriteFile(fileName, file, 0644)
}

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

func readJSONFromFile() (map[string]string, error) {
	plan, _ := ioutil.ReadFile(fileName)
	var data map[string]string
	err := json.Unmarshal(plan, &data)
	if err != nil {
		return map[string]string{}, err
	}
	return data, nil
}
