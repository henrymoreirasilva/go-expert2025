package testes

func CalculateTax(amount float32) float32 {
	amount = amount * 2
	if amount < 300.00 {
		return 0.2
	}
	if amount < 500.00 {
		return 0.5
	}
	if amount < 700.00 || amount == 1000.00 {
		return 0.7
	}
	return 1.0
}
