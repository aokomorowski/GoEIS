package GoEIS

import (
	"math"
	"testing"
)

func TestCreateResistor(t *testing.T) {
	r := Resistance(30.0)
	got := createResistor(r)
	want := func() Impedance {
		return Impedance(complex(30.0, 0))
	}

	if got() != want() {
		t.Errorf("Got: %v \n Want: %v", got(), want())
	}
}

func TestCreateCapacitor(t *testing.T) {
	c := Capacity(10.0)

	got := createCapacitor(c)
	want := func(f Frequency) Impedance {
		return Impedance(complex(0, -1/(2*math.Pi*c.getValue()*f.getValue())))
	}

	f := Frequency(100.0)
	if got(f) != want(f) {
		t.Errorf("Got: %v \n Want: %v", got(f), want(f))
	}
}
