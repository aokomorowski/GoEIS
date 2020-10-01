package GoEIS

import "math"


type Resistance float32
type Capacity float32
type Induction float32

type Elements []interface{}
type Resistor func() Impedance
type Capacitor func(f Frequency) Impedance
type Inductor func(f Frequency) Impedance

func (r Resistance) getValue() float32 {
	return float32(r)
}

func (c Capacity) getValue() float32 {
	return float32(c)
}

func (l Induction) getValue() float32 {
	return float32(l)
}

func returnImpedance(real float32, imaginary float32) Impedance {
	return Impedance(complex(real, imaginary))
}
func createResistor(r Resistance) Resistor {
	return func() Impedance {
		return returnImpedance(r.getValue(), 0)
	}
}

func createCapacitor(c Capacity) Capacitor {
	return func(f Frequency) Impedance {
		imaginary := -1 / (2 * math.Pi * f.getValue() * c.getValue())
		return returnImpedance(0, imaginary)
	}
}

func createInductor(l Induction) Inductor {
	return func(f Frequency) Impedance {
		imaginary := 2 * math.Pi * f.getValue() * l.getValue()
		return returnImpedance(0, imaginary)
	}
}

func serialAddition(e Elements) func(f Frequency) Impedance {
	return func (f Frequency) Impedance{
		z = Impedance(0)
		for _, element := range e {
			z.add(element(f))
		}
	}
}