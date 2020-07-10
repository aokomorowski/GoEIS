package GoEIS

import (
	log "github.com/sirupsen/logrus"
)

func ContainsString(sl []string, v string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func readFile(filename, instrument string) (FrequencySeries, ImpedanceSeries) {

	supportedTypes := []string{
		"gamry", "autolab", "parstat", "zplot", "versastudio", "powersuite", "biologic", "chinstruments",
	}

	if ContainsString(supportedTypes, instrument) {
		log.Errorf("%s is not supported type of data")
	}

	var f FrequencySeries
	var Z ImpedanceSeries

	switch instrument {
	case "gamry":
		f, Z = readGamry(filename)
	case "autolab":
		f, Z = readAutolab(filename)
	case "parstat":
		f, Z = readParstat(filename)
	case "zplot":
		f, Z = readZplot(filename)
	case "versastudio":
		f, Z = readVersa(filename)
	case "powersuite":
		f, Z = readPowersuite(filename)
	case "biologic":
		f, Z = readBiologic(filename)
	case "chinstruments":
		f, Z = readChinstruments(filename)
	}

	return f, Z
}
