package utils

import (
	"fmt"
	"math/rand"
	"os"
)

var names = []string{
	"John",
	"Jane",
	"Mary",
	"Peter",
	"Paul",
	"Mark",
	"Luke",
	"Matthew",
	"James",
	"Jude",
	"Simon",
	"Andrew",
	"Philip",
	"Bartholomew",
	"Thomas",
	"Thaddaeus",
	"Matthias",
	"Paul",
	"Stephen",
	"Timothy",
	"Titus",
	"Philemon",
	"James",
	"John",
	"Jude",
	"Peter",
	"Paul",
	"Mark",
	"Luke",
	"David",
	"Michael",
	"William",
	"Joseph",
	"Daniel",
	"Benjamin",
	"Samuel",
	"Joshua",
	"Jacob",
	"Andrew",
	"Gabriel",
	"Matthew",
	"Anthony",
	"Christopher",
	"Jackson",
	"Ethan",
	"James",
	"Alexander",
	"Sebastian",
	"Logan",
	"Jack",
	"Ryan",
	"Oliver",
	"Henry",
	"Owen",
	"Connor",
	"Nathan",
	"Isaac",
	"Elijah",
	"Luke",
	"Levi",
}

var productsNames = []string{
	"Product 1",
	"Product 2",
	"Product 3",
	"Product 4",
	"Product 5",
	"Product 6",
	"Product 7",
	"Product 8",
	"Product 9",
	"Product 10",
	"Product 11",
	"Product 12",
	"Product 13",
	"Product 14",
	"Product 15",
	"Product 16",
	"Product 17",
	"Product 18",
	"Product 19",
	"Product 20",
	"Product 21",
	"Product 22",
	"Product 23",
	"Product 24",
	"Product 25",
	"Product 26",
	"Product 27",
	"Product 28",
	"Product 29",
	"Product 30",
	"Product 31",
	"Product 32",
	"Product 33",
	"Product 34",
	"Product 35",
	"Product 36",
	"Product 37",
	"Product 38",
	"Product 39",
	"Product 40",
	"Product 41",
	"Product 42",
	"Product 43",
	"Product 44",
	"Product 45",
	"Product 46",
	"Product 47",
	"Product 48",
	"Product 49",
	"Product 50",
}

func GenerateCSVFiles() {
	file, err := os.OpenFile("users.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	file2, err := os.OpenFile("products.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	defer file2.Close()
	if err != nil {
		panic(err)
	}

	for i := 1; i < 500_000; i++ {
		row := fmt.Sprintf("%d,%s\n", i, getRandomName())
		file.Write([]byte(row))
	}

	for i := 1; i < 1_000_000; i++ {
		row := fmt.Sprintf("%d,%s,%d,%d\n", i, getRandomProductName(), rand.Intn(1_999), rand.Intn(499_999)+1)
		file2.Write([]byte(row))
	}

	fmt.Println("Success")
}

func getRandomName() string {
	return names[rand.Intn(len(names)-1)]
}

func getRandomProductName() string {
	return productsNames[rand.Intn(len(productsNames)-1)]
}
