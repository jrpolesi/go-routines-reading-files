package solutions

import (
	"testing"
)

func BenchmarkWithCSVReaderAndAsynchronousWithChanel2(b *testing.B) {
	s := WithCSVReaderAndAsynchronousWithChannel2{}
	usersFile := "../users.csv"
	productsFile := "../products.csv"
	resultFile := "../result.json"

	for i := 0; i < b.N; i++ {
		s.Resolve(usersFile, productsFile, resultFile)
	}
}
