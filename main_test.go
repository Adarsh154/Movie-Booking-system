package main

import "testing"

var TestAvailableSeats *map[string]seats

func init() {
	TestAvailableSeats = initData()
}

func TestTotalUserCost(t *testing.T) {
	var tests = []struct {
		input  []string
		output float64
	}{{
		[]string{"B1", "B4"},
		644.0,
	},
		{
			[]string{"A1", "A2", "A3"},
			1104.0,
		},
		{
			[]string{"A1", "B1", "C1"},
			966.0,
		}}

	for _, test := range tests {
		if output := TotalUserCost(test.input); output != test.output {
			t.Errorf("expected %f but got %f", test.output, output)

		}
	}
}

func TestTotalRevenue(t *testing.T) {
	var tests = []struct {
		output float64
	}{{

		2360.0,
	}}

	for _, test := range tests {
		if output := TotalRevenue("admin123"); output != test.output {
			t.Errorf("expected %f but got %f", test.output, output)

		}
	}
}

func TestValidateSeats(t *testing.T) {
	var tests = []struct {
		input1 []string
		output bool
	}{{
		[]string{"B1", "B4"},
		true,
	},
		{
			[]string{"z1", "A2", "A3"},
			false,
		},
		{
			[]string{",", "A2", "A3"},
			false,
		},
		{
			[]string{"A1", "A2", "A3"},
			true,
		},
		{
			[]string{"A1", "A1", "A3"},
			false,
		}}

	for _, test := range tests {
		if output, _ := ValidateSeats(test.input1); output != test.output {
			t.Errorf("expected %t but got %t", test.output, output)

		}
	}
}

func TestBookTickets(t *testing.T) {
	type testStruct struct {
		bookingSeats      []string
		InpAvailableSeats *map[string]seats
		showNum           int
		bookingStatus     bool
	}

	tests := []testStruct{
		{[]string{}, TestAvailableSeats, 3, true},
		{[]string{"A1"}, TestAvailableSeats, 1, true},
		{[]string{"A2", "A1"}, TestAvailableSeats, 1, false}}

	for _, test := range tests {
		if statusOutput := BookTickets(test.bookingSeats, test.InpAvailableSeats, test.showNum); statusOutput != test.bookingStatus {
			t.Errorf("for input %v expected %t but got %t", test.bookingSeats, test.bookingStatus, statusOutput)

		}
	}
}
