// +build windows

package windows

import (
	"testing"
)

func TestInterfaceGenerator(t *testing.T) {
	g := &InterfaceGenerator{}

	_, err := g.Generate()
	if err != nil {
		t.Errorf("Generate() failed: %s", err)
	}
}
