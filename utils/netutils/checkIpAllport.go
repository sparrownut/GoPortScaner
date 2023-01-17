package netutils

import (
	"GoPortScaner/Global"
	"fmt"
	"strconv"
	"time"
)

func CheckIpWithAllPort(ip string) string {
	threads := 0
	resCsvString := ""

	for i := 0; i < 65536; i++ {
		time.Sleep(time.Duration(time.Nanosecond * 10))
	wait:
		if threads <= 4096 {
			go func(host string, port string) {
				if Global.DBG {
					println(fmt.Sprintf("执行%v:%v中 当前线程%v", host, port, threads))
				}
				threads++
				defer func() {
					if Global.DBG {
						println(fmt.Sprintf("执行%v:%v完毕", host, port))
					}
					threads--
				}()
				resStr := PortDataToCsvString(ScanOpenPort(host, port, Global.CHECKN))
				if resStr != "" {
					resCsvString += resStr + "\n"
				}
			}(ip, strconv.Itoa(i))
		} else { //线程超标回头等待
			goto wait
		}
	}
waitToEnd:
	if threads > 0 {
		goto waitToEnd
	}
	return resCsvString
}
