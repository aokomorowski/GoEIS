package circuits

import "testing"

func TestMain(t *testing.T) {
	got := main()
	want := ""

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
