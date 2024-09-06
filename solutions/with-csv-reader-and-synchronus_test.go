package solutions

import (
	"testing"
)

func BenchmarkWithCSVReaderAndSynchronous(b *testing.B) {
	s := WithCSVReaderAndSynchronous{}
	usersFile := "../users.csv"
	productsFile := "../products.csv"
	resultFile := "../result.json"

	for i := 0; i < b.N; i++ {
		s.Resolve(usersFile, productsFile, resultFile)
	}
}
