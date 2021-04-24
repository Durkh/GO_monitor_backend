package Error

import "log"

func PSCPUError(err error) {
	if err != nil {
		log.Fatalf("PSUtil CPU Error: %s", err.Error())
	}
}

func PSMemError(err error) {
	if err != nil {
		log.Fatalf("PSUtil Memory Error: %s", err.Error())
	}
}
