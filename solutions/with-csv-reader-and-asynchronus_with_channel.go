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

type WithCSVReaderAndAsynchronousWithChannel struct{}

func (s WithCSVReaderAndAsynchronousWithChannel) Resolve(usersFile, productsFile, resultFile string) error {
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

func (s *WithCSVReaderAndAsynchronousWithChannel) readFiles(usersFile, productsFile string) (userMap, error) {
	var wg sync.WaitGroup

	usersChannels := make(chan User)
	productsChannels := make(chan Product)

	wg.Add(2)

	go s.readProducts(productsFile, productsChannels, &wg)
	go s.readUsers(usersFile, usersChannels, &wg)

	go func() {
		wg.Wait()
	}()

	safeUsersMap := safeUsers{
		value: make(userMap),
	}

	for newUser := range usersChannels {
		// safeUsersMap.Lock()
		user, ok := safeUsersMap.value[newUser.ID]
		if !ok {
			safeUsersMap.value[user.ID] = newUser
		} else {
			user.Name = newUser.Name
			safeUsersMap.value[user.ID] = user
		}
		// safeUsersMap.Unlock()
	}

	for product := range productsChannels {
		// safeUsersMap.Lock()
		user, ok := safeUsersMap.value[product.UserID]
		if !ok {
			user.ID = product.UserID
			user.Products = []Product{product}
			safeUsersMap.value[user.ID] = user
		} else {
			user.Products = append(user.Products, product)
			safeUsersMap.value[user.ID] = user
		}
		// safeUsersMap.Unlock()
	}

	return safeUsersMap.value, nil
}

func (s *WithCSVReaderAndAsynchronousWithChannel) readUsers(usersFile string, userChan chan<- User, wg *sync.WaitGroup) error {
	defer wg.Done()
	defer close(userChan)

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

		userChan <- User{
			Name: userName,
			ID:   userID,
		}
	}

	return nil
}

func (s *WithCSVReaderAndAsynchronousWithChannel) readProducts(productsFile string, productsChannel chan<- Product, wg *sync.WaitGroup) error {
	defer wg.Done()
	defer close(productsChannel)

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

		productsChannel <- Product{
			ID:     productID,
			Name:   productName,
			Price:  productPrice,
			UserID: userID,
		}
	}

	return nil
}
