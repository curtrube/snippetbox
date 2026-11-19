package main

import (
	"testing"
	"time"

	"github.com/curtrube/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	// Create a slice of anonymous structs containing the test case name,
	// input to our humanDate() function (the tm field), and expected output
	// (the want field).
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2024 at 10:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CST",
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CST", -6*60*60)),
			want: "17 Mar 2024 at 16:15",
		},
	}

	for _, tt := range tests {
		// Use the t.Run() function to run a sub-test for each test case. The
		// first parameter to this is the name of the test (used to identify the
		// sub-test in any log output) and the second parameter is an anonymous
		// function containing the actual test for each case.
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			// Use the new assert.Equal() helper to compare the expected and
			// actual values.
			assert.Equal(t, hd, tt.want)
		})
	}
}
