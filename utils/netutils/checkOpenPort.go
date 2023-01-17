package netutils

import (
	"GoPortScaner/Global"
	"fmt"
	"net"
	"strings"
	"time"
)

type PortDataSingle struct {
	Data string
	Open bool
	Time string
}
type PortData struct {
	Host         string
	Port         string
	CheckRes     []PortDataSingle
	FinalOpenRes bool
	//TimeSub int
}

func DataToCsvTitleGenerater() string {
	s1 := fmt.Sprintf("%v,%v,%v", "HOST", "PORT", "最终扫描输出")
	for i := 0; i < Global.CHECKN; i++ {
		s1 += fmt.Sprintf(",第%v次扫描结果", i+1)
	}
	s1 += "\n"
	return s1
}
func PortDataToCsvString(data PortData) string {
	// HOST PORT 扫描开放结果 历次扫描结果
	if !data.FinalOpenRes {
		return ""
	}
	s1 := fmt.Sprintf("%v,%v,%v", data.Host, data.Port, data.FinalOpenRes)
	for _, it := range data.CheckRes {
		itOpenString := "关闭"
		if it.Open {
			itOpenString = "开启"
		}
		s1 += "," + itOpenString + " " + it.Data + " " + it.Time
	}
	return s1
}

func ScanOpenPort(host string, port string, checkN int) PortData {
	isOpen := false
	var res PortData
	res.Host = host
	res.Port = port
	for i := 0; i < checkN; i++ { // 扫描checkN次
		dial, connecterr := net.DialTimeout("tcp", host+":"+port, time.Duration(5*time.Second)) // 超时3秒
		if connecterr == nil && dial != nil {
			//fmt.Printf("%v:%v开放", host, port)
			defer func() { // 关闭
				_ = dial.Close()
			}()
			isOpen = true
			res.FinalOpenRes = isOpen

			//res.TimeSub =
			// 开放试图读取端口数据
			_ = dial.SetReadDeadline(time.Now().Add(5 * time.Second))
			buf := [512]byte{}
			n, _ := dial.Read(buf[:])
			readPortStringData := string(buf[:n])
			readPortStringData = strings.ReplaceAll(readPortStringData, ",", "，") // 防止csv分割漏洞
			if readPortStringData != "" {
				res.CheckRes = append(res.CheckRes, PortDataSingle{ // 返回结果
					Open: true,
					Time: time.Now().String(),
					Data: readPortStringData,
				})
			} else {
				res.CheckRes = append(res.CheckRes, PortDataSingle{ // 返回结果
					Open: true,
					Time: time.Now().String(),
				})
			}
		} else {
			res.CheckRes = append(res.CheckRes, PortDataSingle{
				Open: false,
				Time: time.Now().String(),
			})
		}
	}
	return res
}
