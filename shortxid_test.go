package shortxid

import (
	"strings"
	"testing"
)

func TestShortxid_NewID_Prepend(t *testing.T) {
	gen := NewGenerator(0, "GGG")

	if gen == nil {
		t.Fatal("failed to create new generator")
	}

	result := gen.NewID("III")

	if !strings.HasPrefix(result, "GGGIII") {
		t.Fatal("failed to generate with prefix")
	}

	result2 := gen.NewID("III")
	if result == result2 {
		t.Fatal("failed to generator unique id")
	}
}

func TestShortxid_NewID_Expected(t *testing.T) {
	gen := NewGenerator(12345, "")
	gen.TimeFunc = func() uint64 { return 1234567890 }

	result := gen.NewID("")
	if result != "1WrZ9RdWC1" {
		t.Fatal("new ID did not match expected: " + result)
	}
}
