package math

const (
	//NONE is returned from a math function with no solution
	NONE uint8 = 0
	//ONE  is returned from a math function with one solution
	ONE uint8 = 1
	//INFINITE is returned from a math function with infinte solutions
	INFINITE uint8 = 2
)

//Linear is a linear function
type Linear struct {
	M, C float64
}

//Y calculates the value y of a linear function
func (linear *Linear) Y(x float64) float64 {
	return (linear.M * x) + linear.C
}

//Cross checks if the two linear functions cross at some point
func Cross(first *Linear, second *Linear) (float64, uint8) {
	if first.M == second.M {
		if first.C == second.C {
			return 0, INFINITE
		} //else
		return 0, NONE
	} //else

	return (second.C - first.C) / (first.M - second.M), ONE //calculate X
}
