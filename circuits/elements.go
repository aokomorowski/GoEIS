package circuits

// TODO: implement R, C, L, W, W_0, W_e, CPE
import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func CreateResistor(resistance float64) (float64, error) {
	var err error
	if resistance < 0 {
		err = fmt.Errorf("Given resistance is lesser than 0 : %g", resistance)
		log.Error(err)
	}
	return resistance, err
}

// func CreateCapacitor(conductivity float64) (float64 error) {
// 	return impedance, err
// }
