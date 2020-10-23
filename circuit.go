package main

type Circuit []Elements

func calculateImpedanceCircuit(circuit Circuit) func(f Frequency) Impedance {
	var funcs []func(f Frequency) Impedance
	for _, item := range circuit {
		if len(item) > 1 {
			funcs = append(funcs, sumInParallel(item))
		} else {
			funcs = append(funcs, sumInSeries(item))
		}
	}
	return func(f Frequency) Impedance {
		z := Impedance(0)
		for _, item := range funcs {
			z += item(f)
		}
		return z
	}
}

// func ParseCircuit(circuit string)  {
// 	var re = regexp.MustCompile(`(?m)\((.*)\)`)
// 	return
// }
