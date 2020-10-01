package GoEIS

type Frequency float32

func (f Frequency) getValue() float32 {
	return float32(f)
}

type FrequencySeries []Frequency

type Impedance complex64
type ImpedanceSeries []Impedance

func (z Impedance) add(value Impedance){
	z += value
}
type Series struct {
	Frequencies FrequencySeries
	Impedances  ImpedanceSeries
}

func main() {
	return
}
