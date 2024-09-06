package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jrpolesi/go-routines-reading-files/solutions"
)

type Solution interface {
	Resolve(usersFile, productsFile, resultFile string) error
}

type Solutions struct {
	WithCSVReaderAndSynchronous              Solution
	WithCSVReaderAndAsynchronous             Solution
	WithCSVReaderAndAsynchronousWithChannel  Solution
	WithCSVReaderAndAsynchronousWithChannel2 Solution
}

func main() {
	usersFile, productsFile, resultFile, solution := readFlags()

	solutions := Solutions{
		WithCSVReaderAndSynchronous:              solutions.WithCSVReaderAndSynchronous{},
		WithCSVReaderAndAsynchronous:             solutions.WithCSVReaderAndAsynchronous{},
		WithCSVReaderAndAsynchronousWithChannel:  solutions.WithCSVReaderAndAsynchronousWithChannel{},
		WithCSVReaderAndAsynchronousWithChannel2: solutions.WithCSVReaderAndAsynchronousWithChannel2{},
	}

	startTime := time.Now()

	switch solution {
	case "1":
		solutions.WithCSVReaderAndSynchronous.Resolve(usersFile, productsFile, resultFile)
	case "2":
		solutions.WithCSVReaderAndAsynchronous.Resolve(usersFile, productsFile, resultFile)
	case "3":
		solutions.WithCSVReaderAndAsynchronousWithChannel.Resolve(usersFile, productsFile, resultFile)
	case "4":
		solutions.WithCSVReaderAndAsynchronousWithChannel2.Resolve(usersFile, productsFile, resultFile)
	default:
		log.Fatal("Invalid solution")
		os.Exit(1)
	}

	executionTime := time.Since(startTime)
	fmt.Printf("Execution time: %.2f seconds\n", executionTime.Seconds())

	fmt.Println("Data saved in", resultFile)
}

func readFlags() (string, string, string, string) {
	solution := flag.String("solution", "", "The solution to use")

	usersFile := flag.String("users", "", "The file containing the users")
	productsFile := flag.String("products", "", "The file containing the products")

	resultFile := flag.String("result", "", "The file to store the result")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	if len(os.Args) == 1 {
		fmt.Println("Usage: go-merge-files --users<file.csv> --products<file.csv> --result<file.csv>")
		os.Exit(1)
	}

	flag.Parse()

	return *usersFile, *productsFile, *resultFile, *solution
}
