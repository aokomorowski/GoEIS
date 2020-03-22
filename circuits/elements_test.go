package circuits

import "testing"

func TestResistor(t *testing.T) {
	testCases := []struct {
		desc        string
		resistance  float64
		shouldThrow bool
	}{
		{
			desc:        "R(200)",
			resistance:  200.0,
			shouldThrow: false,
		},
		{
			desc:        "non correct input",
			resistance:  -200.1,
			shouldThrow: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, gotErr := CreateResistor(tC.resistance)
			want := tC.resistance
			if gotErr != nil && !tC.shouldThrow {
				t.Errorf("Thrown an error, but shouldnt: %s", gotErr)
			}
			if got != want {
				t.Errorf("got %g, want %g", got, want)
			}
		})
	}
}

// func TestCreateCapacitor(t *testing.T) {
// 	testCases := []struct {
// 		desc     string
// 		capacity float64
// 	}{
// 		{
// 			desc:     "",
// 			capacity: 120.0,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {

// 		})
// 	}
// }
