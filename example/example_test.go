package example_test

import (
	"testing"

	"github.com/cespare/asrt"
)

func TestAssert(t *testing.T) {
	asrt.Assert(t, false)
}

func TestEqual(t *testing.T) {
	asrt.Equal(t, 3, 4)
}

func TestDeepEqual(t *testing.T) {
	m1 := map[string]int{"foo": 1}
	m2 := map[string]int{"foo": 2}
	asrt.DeepEqual(t, m1, m2)
}
