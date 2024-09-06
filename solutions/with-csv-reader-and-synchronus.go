package solutions

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type WithCSVReaderAndSynchronous struct{}

func (s WithCSVReaderAndSynchronous) Resolve(usersFile, productsFile, resultFile string) error {
	mergedResult, err := s.readFiles(usersFile, productsFile)
	if err != nil {
		return err
	}

	err = saveResultInJSONFile(resultFile, mergedResult)
	if err != nil {
		return err
	}

	return nil
}

func (s *WithCSVReaderAndSynchronous) readFiles(usersFile, productsFile string) (userMap, error) {
	usersMap := make(userMap)

	err := s.readUsers(usersFile, &usersMap)
	if err != nil {
		return nil, err
	}

	err = s.readProducts(productsFile, &usersMap)
	if err != nil {
		return nil, err
	}

	return usersMap, nil
}

func (s *WithCSVReaderAndSynchronous) readUsers(usersFile string, usersMap *userMap) error {
	usersReader, err := os.Open(usersFile)
	if err != nil {
		errMsg := fmt.Sprintf("error opening users file: %s", err)
		return errors.New(errMsg)
	}
	defer usersReader.Close()

	csvReader := csv.NewReader(usersReader)

	rows, err := csvReader.ReadAll()
	if err != nil {
		errMsg := fmt.Sprintf("error reading users file: %s", err)
		return errors.New(errMsg)
	}

	for _, row := range rows {
		userID, err := strconv.Atoi(row[0])
		if err != nil {
			errMsg := fmt.Sprintf("error parsing user id %s: %s", row[0], err)
			return errors.New(errMsg)
		}
		userName := row[1]

		(*usersMap)[userID] = User{
			ID:  userID,
			Name: userName,
		}
	}

	return nil
}

func (s *WithCSVReaderAndSynchronous) readProducts(productsFile string, usersMap *userMap) error {
	productsReader, err := os.Open(productsFile)
	if err != nil {
		errMsg := fmt.Sprintf("error opening products file: %s", err)
		return errors.New(errMsg)
	}
	defer productsReader.Close()

	csvReader := csv.NewReader(productsReader)

	rows, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("error reading products file: ", err)
	}

	for _, row := range rows {
		productID, err := strconv.Atoi(row[0])
		if err != nil {
			errMsg := fmt.Sprintf("error parsing product id %s: %s", row[0], err)
			return errors.New(errMsg)
		}
		productName := row[1]
		productPrice, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			errMsg := fmt.Sprintf("error parsing product price for id %s: %s", row[0], err)
			return errors.New(errMsg)
		}
		userID, err := strconv.Atoi(row[3])
		if err != nil {
			errMsg := fmt.Sprintf("error parsing user id %s: %s", row[4], err)
			return errors.New(errMsg)
		}

		product := Product{
			ID:     productID,
			Name:   productName,
			Price:  productPrice,
			UserID: userID,
		}

		user, ok := (*usersMap)[userID]
		if ok {
			user.Products = append(user.Products, product)
			(*usersMap)[userID] = user
		} else {
			(*usersMap)[userID] = User{
				ID: 		 userID,
				Products: []Product{product},
			}
		}
	}

	return nil
}
