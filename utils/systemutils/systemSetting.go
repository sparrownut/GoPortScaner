package systemutils

import "runtime"

func SetCpuWithMax() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
