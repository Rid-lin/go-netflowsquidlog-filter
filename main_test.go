package main

import "testing"

func TestCheckIP(t *testing.T) {

	testCases := []struct {
		name    string
		ip      string
		subnet  string
		isValid bool
	}{
		{
			name:    "valid",
			ip:      "192.168.0.1",
			subnet:  "192.168.0.0/24",
			isValid: true,
		},
		{
			name:    "Invalid IP",
			ip:      "192.168.1.1",
			subnet:  "192.168.0.0/24",
			isValid: false,
		},
		{
			name:    "Invalid Mask",
			ip:      "192.168.0.1",
			subnet:  "192.168.0.0/8",
			isValid: false,
		},
	}
	for _, tc := range testCases {
		result, err := checkIP(tc.subnet, tc.ip)
		if err != nil {
			t.Errorf("Test 'сheckIP' failed. Error:%v", err)
		}
		if result != tc.isValid {
			t.Errorf("Test 'сheckIP' OK")
		}
	}

}

// func TestLogFileFiltering(t *testing.T) {
// 	cfg := struct {
// 		ProcessingDirection string
// 		UserFinder          string
// 		SubNets             []string
// 		IgnorList           []string
// 	}{
// 		ProcessingDirection: "both",
// 		UserFinder:          "",
// 		SubNets: []string{
// 			"192.168.0.0/24",
// 			"192.168.13.0/24",
// 			"192.168.65.0/24",
// 		},
// 		IgnorList: []string{
// 			"UDP_PACKET",
// 			":53 ",
// 			":3128 ",
// 			":123 ",
// 			"SrcIPaddress",
// 		},
// 	}

// 	testCases := []struct {
// 		name    string
// 		line    string
// 		result  string
// 		isValid bool
// 	}{
// 		{
// 			name:    "valid",
// 			line:    "1593388784.885      0 192.168.65.175 UDP_PACKET:25570/200 88 HEAD 10.50.4.62:53 - FIRSTUP_PARENT/192.168.65.1 packet/netflow",
// 			result:  "1593388784.885      0 192.168.65.175 UDP_PACKET:25570/200 88 HEAD 10.50.4.62:53 - FIRSTUP_PARENT/192.168.65.1 packet/netflow",
// 			isValid: true,
// 		},
// 	}
// 	for _, tc := range testCases {
// 		result := cfg.logFileFiltering(tc.line)
// 		if result != tc.result {
// 			t.Errorf("Test 'сheckIP' OK")
// 		}
// 	}

// }
