package solutions

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type WithCSVReaderAndAsynchronous struct{}

func (s WithCSVReaderAndAsynchronous) Resolve(usersFile, productsFile, resultFile string) error {
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

func (s *WithCSVReaderAndAsynchronous) readFiles(usersFile, productsFile string) (userMap, error) {
	var wg sync.WaitGroup

	safeUsersMap := safeUsers{
		value: make(userMap),
	}
	wg.Add(2)

	go s.readUsers(usersFile, &safeUsersMap, &wg)
	go s.readProducts(productsFile, &safeUsersMap, &wg)

	wg.Wait()

	return safeUsersMap.value, nil
}

func (s *WithCSVReaderAndAsynchronous) readUsers(usersFile string, usersMap *safeUsers, wg *sync.WaitGroup) error {
	defer wg.Done()

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

		usersMap.Lock()
		user, ok := usersMap.value[userID]
		if ok {
			user.ID = userID
			user.Name = userName
			usersMap.value[userID] = user
		} else {
			usersMap.value[userID] = User{
				ID:   userID,
				Name: userName,
			}
		}
		usersMap.Unlock()
	}

	return nil
}

func (s *WithCSVReaderAndAsynchronous) readProducts(productsFile string, usersMap *safeUsers, wg *sync.WaitGroup) error {
	defer wg.Done()

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

		usersMap.Lock()
		user, ok := usersMap.value[userID]
		if ok {
			user.Products = append(user.Products, product)
			usersMap.value[userID] = user
		} else {
			usersMap.value[userID] = User{
				ID: 		 userID,
				Products: []Product{product},
			}
		}
		usersMap.Unlock()
	}

	return nil
}
