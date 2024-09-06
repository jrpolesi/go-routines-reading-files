package solutions

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

type Product struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	UserID int     `json:"user_id"`
}

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

type userMap map[int]User

type safeUsers struct {
	value userMap
	sync.Mutex
}


func saveResultInJSONFile(resultFile string, usersMap userMap) error {
	jsonFile, err := os.Create(resultFile)
	if err != nil {
		errMsg := fmt.Sprintf("error creating result file: %s", err)
		return errors.New(errMsg)
	}

	jsonData, err := json.MarshalIndent(usersMap, "", "  ")
	if err != nil {
		errMsg := fmt.Sprintf("error marshalling data: %s", err)
		return errors.New(errMsg)
	}

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		errMsg := fmt.Sprintf("error saving data to json file: %s", err)
		return errors.New(errMsg)
	}

	return nil
}
