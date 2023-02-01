package output

import (
	"fmt"
	"os"
	"strings"
)

func PrintProgressBar(progress float32, str string) {
	if int(progress) >= 0 && int(progress) <= 100 {
		fmt.Printf("\r[%v%v]%v%% %v", strings.Repeat("=", int(progress)), strings.Repeat(" ", 100-int(progress)), fmt.Sprintf("%.2f", progress), str)
		err := os.Stdout.Sync()
		if err != nil {
			return
		}
	}
}
