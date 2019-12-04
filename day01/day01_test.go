package main

import "testing"

func TestCalcModuleFuelRequirementPart1(t *testing.T) {
	tables := []struct {
		m int
		r int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, table := range tables {
		got := calcModuleFuelRequirementPart1(table.m)
		if got != table.r {
			t.Errorf("Calculated fuel required (%d / 3 - 2) was incorrect, got: %d, want: %d.", table.m, got, table.r)
		}
	}
}

func TestCalcModuleFuelRequirementPart2(t *testing.T) {
	tables := []struct {
		m int
		r int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, table := range tables {
		got := calcModuleFuelRequirementPart2(table.m)
		if got != table.r {
			t.Errorf("Calculated fuel required for %d was incorrect, got: %d, want: %d.", table.m, got, table.r)
		}
	}
}
