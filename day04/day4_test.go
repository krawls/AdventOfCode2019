package main

import "testing"

func TestValidateDigits(t *testing.T) {
	tables := []struct {
		i int
		o bool
	}{
		{12, true},
		{141235, true},
		{1969, true},
		{1007564, false},
		{54625, true},
	}

	for _, table := range tables {
		got := validateDigits(table.i)
		if got != table.o {
			t.Errorf("Validating %d as having at least 6 digits as %t was incorrect, got: %t, want: %t.", table.i, table.o, got, table.o)
		}
	}
}

func TestValidateRange(t *testing.T) {
	tables := []struct {
		i int
		l int
		u int
		o bool
	}{
		{12, 11, 23454, true},
		{141235, 12345, 141236, true},
		{1969, 1969, 1969, true},
		{100564, 100565, 546123, false},
	}

	for _, table := range tables {
		got := validateRange(table.i, table.l, table.u)
		if got != table.o {
			t.Errorf("Validating %d <= %d <= %d as %t was incorrect, got: %t, want: %t.", table.l, table.i, table.u, table.o, got, table.o)
		}
	}
}

func TestValidateAdjacent(t *testing.T) {
	tables := []struct {
		i int
		o bool
	}{
		{111111, true},
		{223450, true},
		{196984, false},
		{101564, false},
	}

	for _, table := range tables {
		got := validateAdjacent(table.i)
		if got != table.o {
			t.Errorf("Validating %d as having at least 2 adjacent equal digits as %t was incorrect, got: %t, want: %t.", table.i, table.o, got, table.o)
		}
	}
}

func TestValidateAdjacentExact(t *testing.T) {
	tables := []struct {
		i int
		o bool
	}{
		{112233, true},
		{123444, false},
		{111122, true},
		{555555, false},
		{788889, false},
	}

	for _, table := range tables {
		got := validateAdjacentExact(table.i)
		if got != table.o {
			t.Errorf("Validating %d as having at exactly 2 adjacent equal digits as %t was incorrect, got: %t, want: %t.", table.i, table.o, got, table.o)
		}
	}
}

func TestValidateNeverDecrease(t *testing.T) {
	tables := []struct {
		i int
		o bool
	}{
		{111123, true},
		{135679, true},
		{111111, true},
		{101564, false},
		{567898, false},
	}

	for _, table := range tables {
		got := validateNeverDecrease(table.i)
		if got != table.o {
			t.Errorf("Validating %d as having at digits that never decrease as %t was incorrect, got: %t, want: %t.", table.i, table.o, got, table.o)
		}
	}
}
