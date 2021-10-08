package current

import (
	"strings"
	"testing"
)

func TestNewPath(t *testing.T) {
	r := NewPath()
	s := r.Path()
	if !strings.HasSuffix(s, "/current") {
		t.Errorf("Wrong current path result: %v", s)
	}
	s = r.Join("../current/current_path.go")
	if !strings.HasSuffix(s, "/current/current_path.go") {
		t.Errorf("Wrong current path result: %v", s)
	}
}
