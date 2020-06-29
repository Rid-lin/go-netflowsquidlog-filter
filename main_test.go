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
