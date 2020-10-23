package main

import (
	"testing"
)

func TestCreateResistor(t *testing.T) {
	r := Resistance(30.0)
	got := createResistor(r).getImpedance(0.0)
	want := func() Impedance {
		return Impedance(complex(30.0, 0))
	}()

	if got != want {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestSumInSeries(t *testing.T) {
	resistor1 := createResistor(10)
	capacitor1 := createCapacitor(20)
	elements := Elements{Element(resistor1), Element(capacitor1)}
	got := sumInSeries(elements)(10.0)
	want := Impedance(10 - 0.0007957747i)

	if got != want {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestSumInParallel(t *testing.T) {
	resistor1 := createResistor(10)
	capacitor1 := createCapacitor(20)
	elements := Elements{Element(resistor1), Element(capacitor1)}
	got := sumInParallel(elements)(10.0)
	want := Impedance(6.3325736e-08 - 0.0007957747i)

	if got != want {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}
