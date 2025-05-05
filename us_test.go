package uniqueslice

import (
	"bufio"
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	r := strings.NewReader(`123
1234
123`)
	scanner := bufio.NewScanner(r)
	//var lines [][]byte
	var handles []Handle[[]byte, byte]
	for scanner.Scan() {
		//lines = append(lines, append(nil, scanner.Bytes()...))
		handles = append(handles, Make(scanner.Bytes()))
	}
	if err := scanner.Err(); err != nil {
		t.Errorf("scanner error: %v", err)
	}
	for i, h := range handles {
		t.Logf("handles[%d] = %v (%q)", i, h.Value(), h.Value())
	}
	if handles[0] != handles[2] {
		t.Error("handles[0] != handles[2]")
	}
	if handles[0] == handles[1] {
		t.Error("handles[0] == handles[1]")
	}
}
