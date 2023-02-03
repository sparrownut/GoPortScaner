package netutils

import (
	"GoPortScaner/Global"
	"strconv"
	"time"
)

func GetPortWorkSum() int {
	portList := []int{}
	if len(Global.PORT) > 1 {
		portList = Global.PORT
	} else {
		for port := 1; port < 65536; port++ { // 初始化列表
			portList = append(portList, port)
		}
	}
	return len(portList)
}

func CheckIpWithAllPort(ip string) string {
	threads := 0
	resCsvString := ""
	portList := []int{}

	//println(Global.PORT)

	if len(Global.PORT) > 1 {
		portList = Global.PORT
	} else {
		for port := 1; port < 65536; port++ { // 初始化列表
			portList = append(portList, port)
		}
	}
	//println(portStart)
	//println(portEnd)
	for i := range portList {
		time.Sleep(time.Duration(time.Millisecond * 1))
	wait:
		if threads <= Global.SINHPORTMAXTHREADS {

			go func(host string, port string) {
				//if Global.DBG {
				//	println(fmt.Sprintf("执行%v:%v中 当前线程%v", host, port, threads))
				//}
				threads++
				defer func() {
					Global.DoneWork++
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
					case <-time.After(time.Duration(time.Duration(Global.PORTTIMEOUT) * time.Second)):
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
