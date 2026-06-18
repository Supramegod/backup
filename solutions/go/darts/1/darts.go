package darts

func Score(x, y float64) int {
	dst := x*x + y*y
	switch {
	case dst <= 1:
		return 10
	case dst <= 25:
		return 5
	case dst <= 100:
		return 1
	default:
		return 0
	}
	panic("Please implement the Score function")
}
