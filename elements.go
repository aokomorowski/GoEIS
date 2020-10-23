package main

import (
	"math"
)

type Resistance float32
type Capacity float32
type Inductance float32

type Elements []Element

type Element struct {
	getImpedance func(f Frequency) Impedance
}
type Resistor Element
type Capacitor Element
type Inductor Element

func (r Resistance) getValue() float32 {
	return float32(r)
}

func (c Capacity) getValue() float32 {
	return float32(c)
}

func (l Inductance) getValue() float32 {
	return float32(l)
}

func returnImpedance(real float32, imaginary float32) Impedance {
	return Impedance(complex(real, imaginary))
}

func createResistor(r Resistance) Resistor {
	calculateImpedance := func(f Frequency) Impedance {
		return returnImpedance(r.getValue(), 0)
	}
	return Resistor{
		getImpedance: calculateImpedance,
	}
}

func createCapacitor(c Capacity) Capacitor {
	calculateImpedance := func(f Frequency) Impedance {
		imaginary := -1 / (2 * math.Pi * f.getValue() * c.getValue())
		return returnImpedance(0, imaginary)
	}
	return Capacitor{
		getImpedance: calculateImpedance,
	}
}

func createInductor(l Inductance) Inductor {
	calculateImpedance := func(f Frequency) Impedance {
		imaginary := 2 * math.Pi * f.getValue() * l.getValue()
		return returnImpedance(0, imaginary)
	}
	return Inductor{
		getImpedance: calculateImpedance,
	}
}

func sumInSeries(elements Elements) func(f Frequency) Impedance {
	return func(f Frequency) Impedance {
		z := Impedance(0)
		for _, element := range elements {
			z = z.add(element.getImpedance(f))
		}
		return z
	}
}

func sumInParallel(elements Elements) func(f Frequency) Impedance {
	return func(f Frequency) Impedance {
		z := Impedance(0)
		for _, element := range elements {
			z = z.add(1 / element.getImpedance(f))
		}
		return 1 / z
	}
}
