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


func TestAddInSeries(t *testing.T) {
	resistor1 := createResistor(10)
	resistor2 := createResistor(20)
	elements := Elements{Element(resistor1),Element(resistor2)}
	got := addInSeries(elements)(10.0)
	want := Impedance(30.0)

	if got != want {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}