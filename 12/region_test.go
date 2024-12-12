package main

import (
	"testing"
)

func TestMapFindRegions(t *testing.T) {
	t.Run("ex 1", func(t *testing.T) {
		m := ContentToMap(ex1)
		regions := m.FindRegions()
		total := 0
		for _, r := range regions {
			total += r.FencePrice()
		}
		if total != 140 {
			t.Errorf("wrong fence price %v", total)
		}
	})
	t.Run("ex 1 bulk", func(t *testing.T) {
		m := ContentToMap(ex1)
		regions := m.FindRegions()
		total := 0
		for _, r := range regions {
			total += r.FencePriceBulk()
		}
		if total != 80 {
			t.Errorf("wrong fence price %v", total)
		}
	})
	t.Run("ex 3", func(t *testing.T) {
		m := ContentToMap(ex3)
		regions := m.FindRegions()
		total := 0
		for _, r := range regions {
			total += r.FencePrice()
		}
		if total != 1930 {
			t.Errorf("wrong fence price %v", total)
		}
	})
}
