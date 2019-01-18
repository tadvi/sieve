package sieve

import "testing"

func TestSieve(t *testing.T) {

	p := Primes(20)

	want := []bool{false, false, true, true, false, true, false, true, false, false,
		false, true, false, true, false, false, false, true, false, true, false}

	for i := range want {
		if p[i] != want[i] {
			t.Errorf("prime got '%t', want '%t', at position %d", p[i], want[i], i)
		}
	}

}
