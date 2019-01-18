/* Sieve of Eratosthenes for primes */
package sieve

import "math"

func Primes(n int) []bool {
	p := make([]bool, n+1)
	for i := range p {
		p[i] = true
	}

	p[0] = false
	p[1] = false
	m := int(math.Sqrt(float64(n))) + 1
	for i := 2; i <= m; i++ {
		if p[i] {
			// if it is prime, mark all multiples of it as
			// non-primes
			for k := i * i; k <= n; k += i {
				p[k] = false
			}
		}
	}
	return p
}
