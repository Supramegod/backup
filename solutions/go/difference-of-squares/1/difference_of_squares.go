package differenceofsquares

func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
	panic("Please implement the SquareOfSum function")
}

func SumOfSquares(n int) int {
	squareSum := n * (n + 1) * (2*n + 1) / 6
	return squareSum
	panic("Please implement the SumOfSquares function")
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
