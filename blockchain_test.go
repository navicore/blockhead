package main

import "testing"

func TestNewBlockchain(t *testing.T) {
	bc := NewBlockchain()
	if len(bc.blocks) != 1 {
		t.Errorf("expecting 1 block, got %s", bc)
	}
}
