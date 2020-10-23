package main

import (
	"testing"
)

//
//func TestParseCircuit(t *testing.T)  {
//	got := ParseCircuit("R(RC)")
//}

func TestCalculateImpedanceCircuit(t *testing.T) {
	r1 := Element(createResistor(10.0))
	r2 := Element(createResistor(20.0))
	c1 := Element(createCapacitor(10.0))
	circuit := []Elements{Elements{r1}, Elements{r2, c1}}
	got := calculateImpedanceCircuit(circuit)(10)
	var want = func(f Frequency) Impedance {
		return r1.getImpedance(f) + sumInParallel(Elements{r2, c1})(f)
	}(10)

	if got != want {
		t.Errorf("got: %v, want %v", got, want)
	}
}
