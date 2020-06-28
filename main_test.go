package main

import (
	"testing"
)

// flow-cat /home/rid-lin/flow/gw_root_75km/2020/2020-06/2020-06-23/ft-v05.2020-06-23.095336+0500 | flow-print -f 5
// Start             End               Sif   SrcIPaddress    SrcP  DIf   DstIPaddress    DstP    P Fl Pkts

var testingInData = `
0623.09:53:19.589 0623.09:53:20.649 8     192.168.65.155  50182 7     45.142.212.204  3128  6   0  3          144
0623.09:53:20.669 0623.09:53:20.669 8     192.168.65.195  55245 9     10.61.152.10    445   6   0  1          41
0623.09:53:20.669 0623.09:53:20.669 9     10.61.152.10    445   8     192.168.65.195  55245 6   0  1          52
0623.09:53:20.749 0623.09:53:20.749 8     192.168.65.143  38786 13    87.250.251.119  443   6   1  1          52
0623.09:52:34.249 0623.09:53:20.879 13    217.20.152.247  443   8     192.168.65.82   40192 6   2  47         41724`

var testingInLine = "0623.09:52:34.249 0623.09:53:20.879 13    217.20.152.247  443   8     192.168.65.82   40192 6   2  47         41724"
var testingOutLineOK = ""
var testingunixStampFromNetflowDateStr = "0623.09:52:34.249"
var testingunixStampFromNetflowDateStrOK = "1592905954.249"
var testingunixStampFromNetflowDate = "0623.09:52:34.249"
var testingunixStampFromNetflowDateOK int64 = 1592905954249

func TestParseNetFlowToSquidLine(t *testing.T) {
	result, err := parseNetFlowToSquidLine(testingInLine, "2020", "192.168.65.1")
	if err != nil {
		t.Errorf("Test OK failed: %s", err)
	}
	if result == testingOutLineOK {
		t.Errorf("Test OK failed, result not match")
	}
}

func TestUnixStampFromNetflowDateStr(t *testing.T) {
	result := unixStampFromNetflowDateStr(testingunixStampFromNetflowDateStr, "2020")
	if result != testingunixStampFromNetflowDateStrOK {
		t.Errorf("Test 'unixStampFromNetflowDate' failed, result not match")
	}
	result = unixStampFromNetflowDateStr(testingunixStampFromNetflowDateStr, "2019")
	if result == testingunixStampFromNetflowDateStrOK {
		t.Errorf("Test 'unixStampFromNetflowDate' with wrong data failed,  result not match")
	}
	result = unixStampFromNetflowDateStr(testingunixStampFromNetflowDateStr, "20")
	if result != "" {
		t.Errorf("Test 'unixStampFromNetflowDate' with wrong data failed, return wrong result")
	}
}

func TestUnixStampFromNetflowDate(t *testing.T) {
	result := unixStampFromNetflowDate(testingunixStampFromNetflowDate, "2020")
	if result != testingunixStampFromNetflowDateOK {
		t.Errorf("Test 'unixStampFromNetflowDate' failed, result not match")
	}
	result = unixStampFromNetflowDate(testingunixStampFromNetflowDate, "2019")
	if result == testingunixStampFromNetflowDateOK {
		t.Errorf("Test 'unixStampFromNetflowDate' with wrong data failed,  result not match")
	}
	result = unixStampFromNetflowDate(testingunixStampFromNetflowDate, "20")
	if result != 0 {
		t.Errorf("Test 'unixStampFromNetflowDate' with wrong data failed, return wrong result")
	}
}

// func TestFail(t *testing.T) {
// 	result, err := parseNetFlowToSquidLine(testingInLine)
// 	if err != nil {
// 		t.Errorf("Test OK failed: %s", err)
// 	}
// 	if result != testingOutLineOK {
// 		t.Errorf("Test OK failed, result not match")
// 	}
// }
