package Global

import "time"

var (
	DBG                = false
	CHECKN             = 3
	INPUTFILE          = ""
	PORTTIMEOUT        = 5 * time.Second
	PORT               = 0
	SINHPORTMAXTHREADS = 4096
)
