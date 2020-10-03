package main

type Frequency float32

func (f Frequency) getValue() float32 {
	return float32(f)
}

type FrequencySeries []Frequency

type Impedance complex64
type ImpedanceSeries []Impedance

func (z Impedance) add(value Impedance) Impedance{
	return z + value
}
type Series struct {
	Frequencies FrequencySeries
	Impedances  ImpedanceSeries
}

func main() {
	resistor1 := createResistor(10)
	resistor2 := createResistor(20)
	elements := Elements{Element(resistor1),Element(resistor2)}
	addInSeries(elements)(10.0)
	return
}
