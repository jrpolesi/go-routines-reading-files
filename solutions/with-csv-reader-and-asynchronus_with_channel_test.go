package solutions

import (
	"testing"
)

func BenchmarkWithCSVReaderAndAsynchronousWithChanel(b *testing.B) {
	s := WithCSVReaderAndAsynchronousWithChannel{}
	usersFile := "../users.csv"
	productsFile := "../products.csv"
	resultFile := "../result.json"

	for i := 0; i < b.N; i++ {
		s.Resolve(usersFile, productsFile, resultFile)
	}
}
