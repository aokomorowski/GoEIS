package main

type Resistance float32

type Elements []Element

type Element struct {
	getImpedance func(f Frequency) Impedance
}
type Resistor Element


func (r Resistance) getValue() float32 {
	return float32(r)
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

func addInSeries(elements Elements) func(f Frequency) Impedance{
	return func(f Frequency) Impedance {
		z := Impedance(0)
		for _, element := range elements {
			z = z.add(element.getImpedance(f))
		}
		return z
	}
}

