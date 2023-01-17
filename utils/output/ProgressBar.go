package output

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrintProgressBar(progress int, str string) {
	if progress >= 0 && progress <= 100 {
		fmt.Printf("\r[%v%v]%v%% %v", strings.Repeat("=", progress), strings.Repeat(" ", 100-progress), strconv.Itoa(progress), str)
		err := os.Stdout.Sync()
		if err != nil {
			return
		}
	}
}
