package solutions

import (
	"testing"
)

func BenchmarkWithCSVReaderAndAsynchronous(b *testing.B) {
	s := WithCSVReaderAndAsynchronous{}
	usersFile := "../users.csv"
	productsFile := "../products.csv"
	resultFile := "../result.json"

	for i := 0; i < b.N; i++ {
		s.Resolve(usersFile, productsFile, resultFile)
	}
}
