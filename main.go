package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tkanos/gonfig"
)

var year, inFile, outFile, configFile string

// ConfigType ..
type ConfigType struct {
	ProcessingDirection string            `ini:"processing_direction"`
	UserFinder          string            `ini:"user_finder"`
	SubNet              map[string]string `ini:"subnets"`
	IgnorList           map[string]string `ini:"ignor"`
}

func init() {
	flag.StringVar(&configFile, "c", "go-netflowsquidlog-filter.ini", "configuration file")
	// flag.StringVar(&mainSubnet, "ms", "192.168.0.0/24", "the main subnet where clients are located, for example 192.168.0.2, 192.168.0.3, 192.168.0.4,")
	flag.StringVar(&inFile, "in", "", "Temp log file for filtering")
	flag.StringVar(&outFile, "out", "/var/log/squid/access2.log", "Log file for further processing by the analyzer")
	flag.Parse()
	configuration := ConfigType{}
	err := gonfig.GetConf(configFile, &configuration)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}
}

func main() {
	file, err := os.Open(inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// TODO тут должна находится функция обработки строк лог-фала
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// Input - string in netflow format №5
//Start             End               Sif   SrcIPaddress    SrcP  DIf   DstIPaddress    DstP    P Fl Pkts       Octets
//Output - squid log format default

func parseNetFlowToSquidLine(strIn, year, collectorIP string) (string, error) {
	var protocol string
	strArray := strings.Fields(strIn)
	if len(strArray) <= 0 {
		return "", nil
	}
	unixStampStr := unixStampFromNetflowDateStr(strArray[0], year)
	startOfResponse := unixStampFromNetflowDate(strArray[0], year)
	endOfResponse := unixStampFromNetflowDate(strArray[1], year)
	delayStr := strconv.FormatInt((endOfResponse/1000 - startOfResponse/1000), 10)
	// user = "-"

	switch strArray[8] {
	case "6":
		protocol = "TCP_PACKET"
	case "17":
		protocol = "UDP_PACKET"
	default:
		protocol = "OTHER_PACKET"

	}
	//Start             End               Sif   SrcIPaddress    SrcP  DIf   DstIPaddress    DstP    P Fl Pkts       Octets
	//
	out := fmt.Sprintf("%v %6v %v %v:%v/200 %v HEAD %v:%v - FIRSTUP_PARENT/%v packet/netflow", unixStampStr, delayStr, strArray[3], protocol, strArray[4], strArray[len(strArray)-1], strArray[6], strArray[7], collectorIP)
	return out, nil
}

func unixStampFromNetflowDateStr(str, year string) string {
	str = year + str
	normalizedDate, err := time.Parse("20060102.15:04:05.000", str)
	if err != nil {
		return ""
	}

	timeUnix := normalizedDate.Unix()
	timeUnixStr := strconv.FormatInt(timeUnix, 10)
	timeUnixNanoStr := strconv.FormatInt(((normalizedDate.UnixNano() - timeUnix*1000000000) / 1000000), 10)
	if len(timeUnixNanoStr) == 1 {
		timeUnixNanoStr = timeUnixNanoStr + "00"
	} else if len(timeUnixNanoStr) == 2 {
		timeUnixNanoStr = timeUnixNanoStr + "0"
	}

	out := fmt.Sprintf("%v.%v", timeUnixStr, timeUnixNanoStr)
	return out
}

func unixStampFromNetflowDate(str, year string) int64 {
	str = year + str
	normalizedDate, err := time.Parse("20060102.15:04:05.000", str)
	if err != nil {
		return 0
	}

	timeUnixMili := normalizedDate.UnixNano() / 1000000
	return timeUnixMili
}
