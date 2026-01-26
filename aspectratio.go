package main

type floatVect struct{ X, Y float32 }

func getAspectRatio(a, b int32) floatVect {
	var GCD int32

	GCD = gcd(screenW, screenH)
	w := screenW / GCD
	h := screenH / GCD
	arw := float32(w) / float32(h) //take note '/' does not work the same as in C
	return floatVect{X: arw, Y: 1.0}
}

// gcd (Greatest Common Divisor) calculates the GCF of two numbers using the Euclidean algorithm.
func gcd(a, b int32) int32 {
	// Base case: if the second number (b) is 0, the GCD is the first number (a).
	if b == 0 {
		return a
	}
	// Recursive step: call gcd with the second number (b) and the remainder of a divided by b.
	return gcd(b, a%b)
}
