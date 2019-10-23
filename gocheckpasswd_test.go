package gocheckpasswd

import (
	"testing"
)

func Test12345(t *testing.T) {
	result := IsCommon("12345")
	if !result {
		t.Errorf("IsCommon(12345) was incorrect, got: %t, want: %t", result, false)
	}
}

func TestAlice(t *testing.T) {
	result := IsCommon("alice")
	if !result {
		t.Errorf("IsCommon(alice) was incorrect, got: %t, want: %t", result, false)
	}
}

func TestHorz71(t *testing.T) {
	result := IsCommon("horz71")
	if result {
		t.Errorf("IsCommon(horz712) was incorrect, got: %t, want: %t", result, true)
	}
}
