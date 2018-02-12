package math

//Operand type
type Operand func(a, b float64) float64

//Multiplication generic
func Multiplication(a, b float64) float64 {
	return a * b
}

//Addition generic
func Addition(a, b float64) float64 {
	return a + b
}

//Subtraction generic
func Subtraction(a, b float64) float64 {
	return a - b
}
