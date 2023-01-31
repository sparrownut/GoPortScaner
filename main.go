package main

import (
	"GoPortScaner/Global"
	"GoPortScaner/utils/netutils"
	"GoPortScaner/utils/output"
	"GoPortScaner/utils/systemutils"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
	"time"
)

func main() {
	app := &cli.App{
		Name:      "GoPortScaner",
		Usage:     "高性能端口扫描器 \n多次扫描 力保扫描准确性\n仅供授权的渗透测试使用 请遵守法律!", // 这里写协议
		UsageText: "lazy to write...",
		Version:   "0.2.9",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "InputFile", Aliases: []string{"F"}, Destination: &Global.INPUTFILE, Value: "list", Usage: "扫描输入文件", Required: true},
			&cli.BoolFlag{Name: "DBG", Aliases: []string{"D"}, Destination: &Global.DBG, Value: false, Usage: "DBG MOD", Required: false},
			&cli.IntFlag{Name: "checkN", Aliases: []string{"C"}, Destination: &Global.CHECKN, Value: 3, Usage: "同一端口检测次数", Required: false},
			&cli.IntFlag{Name: "port", Aliases: []string{"P"}, Destination: &Global.PORT, Usage: "目标端口", Required: false},
			&cli.IntFlag{Name: "PortScanThreads", Aliases: []string{"T"}, Destination: &Global.SINHPORTMAXTHREADS, Value: 4096, Usage: "单个目标扫描线程数", Required: false},

			//&cli.IntFlag{Name: "checkN", Aliases: []string{"C"}, Destination: &Global.CHECKN, Value: 3, Usage: "同一端口检测次数", Required: false},
		},
		HideHelpCommand: true,
		Action: func(c *cli.Context) error {
			err := do()
			if err != nil {

			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		//panic(err)
	}

	//fmt.Printf(os.Args[1])
}
func do() error {
	systemutils.SetCpuWithMax()
	//fmt.Printf(netutils.CheckIpWithAllPort("hk11.stuid-fish.co"))
	outfile, _ := os.OpenFile(fmt.Sprintf("scanoutput_%v.csv", time.Now().Format("2006-01-02-15-04-05")), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer func(file *os.File) {
		_ = file.Close()
	}(outfile)
	file, fileReaderr := os.ReadFile(Global.INPUTFILE)
	doList := []string{}
	if fileReaderr == nil {
		readList := strings.Split(string(file), "\n")
		fmt.Printf("原输入%v条\n", len(readList))
		for _, it := range readList {
			if strings.Contains(it, "/") {
				doList = append(doList, netutils.Cidr2IPs(it)...)
			} else {
				doList = append(doList, it)
			}
		}
		fmt.Printf("解析后总任务%v\n", len(doList))
		//for _, it := range doList {
		//	println(it)
		//}
		//for i := 0; i < 100; i++ {
		//	time.Sleep(time.Duration(time.Second))
		//	output.PrintProgressBar(i, "test")
		//}
		done := 0
		csvOutput := ""
		threads := 0
		threadsMax := 10
		if Global.PORT != 0 {
			threadsMax = 4096
		}
		for _, it := range doList {
			done++
			//println(done)
			//println(float32(done) / float32(len(doList)))
			output.PrintProgressBar(int(100*float32(done)/float32(len(doList))), it+"            ")
		waitToRetry:
			time.Sleep(time.Duration(time.Millisecond * 1))
			if threads <= threadsMax {
				go func(host string) {
					threads++
					defer func() {
						threads--
					}()
					csvOutput += netutils.CheckIpWithAllPort(host)
				}(it)
			} else {
				goto waitToRetry
			}
		}
	waitToEnd:
		if threads > 0 {
			goto waitToEnd
		}
		csvTitle := netutils.DataToCsvTitleGenerater()
		csvText := csvTitle + csvOutput
		_, err := outfile.WriteString(csvText)
		if err == nil {
			fmt.Printf("完成")
		}
		//println(csvOutput)
	} else {
		println("文件读取错误")
	}
	//output.PrintProgressBar(50)
	//fmt.Printf()
	return nil

}
