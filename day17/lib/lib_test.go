package day17

import (
	"testing"
)

func TestSpinlock(t *testing.T) {
	s := NewSpinlock(3)
	index := int(0)
	for i := 0; i < 2017; i++ {
		index = s.Spin()
	}

	tests := map[int]uint{-3: 1512, -2: 1134, -1: 151, 0: 2017, 1: 638, 2: 1513, 3: 851}

	for offset, expected := range tests {
		i := index + offset
		result := s.Lookup(i)
		if result != expected {
			t.Errorf("Failed with index %d, got %d expected %d", i, result, expected)
		}
	}
}
