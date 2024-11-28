package testes

import "testing"

func TestCalculateTax(t *testing.T) {
	CalculateTax(500.00)
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.00)
	}
}
