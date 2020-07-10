package GoEIS

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	gotF, gotZ := readFile("./lol.dat", "gamry")
	var wantF FrequencySeries = []float32{0.0}
	var wantZ ImpedanceSeries = []complex64{complex(1, 2)}

	if reflect.DeepEqual(gotF, wantF) || reflect.DeepEqual(gotZ, wantZ) {
		t.Error("readFile unsuccessful")
	}
}
