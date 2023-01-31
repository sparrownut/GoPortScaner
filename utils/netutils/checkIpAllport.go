package netutils

import (
	"GoPortScaner/Global"
	"strconv"
	"time"
)

func CheckIpWithAllPort(ip string) string {
	threads := 0
	resCsvString := ""
	portStart := 0
	portEnd := 65536
	//println(Global.PORT)
	if Global.PORT != 0 {
		portStart = Global.PORT
		portEnd = Global.PORT + 1
	}
	//println(portStart)
	//println(portEnd)
	for i := portStart; i < portEnd; i++ {
		time.Sleep(time.Duration(time.Millisecond * 1))
	wait:
		if threads <= Global.SINHPORTMAXTHREADS {

			go func(host string, port string) {
				//if Global.DBG {
				//	println(fmt.Sprintf("执行%v:%v中 当前线程%v", host, port, threads))
				//}
				threads++
				defer func() {
					if r := recover(); r != nil {
						if Global.DBG {
							println("ERR")
						}
						threads--
					}
					//if Global.DBG {
					//	println(fmt.Sprintf("执行%v:%v完毕 当前线程%v", host, port, threads))
					//}
					threads--
				}()
				resStr := PortDataToCsvString(ScanOpenPort(host, port, Global.CHECKN))
				if resStr != "" {
					resCsvString += resStr + "\n"
				}
				for {
					select {
					case <-time.After(Global.PORTTIMEOUT):
						threads--
						return

					}
				}
			}(ip, strconv.Itoa(i))
			time.Sleep(time.Duration(time.Nanosecond * 10))
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
