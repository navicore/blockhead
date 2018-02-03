package main

import "testing"

func TestNewProofofwork(t *testing.T) {

	block := NewGenesisBlock()
	NewProofOfWork(block)
	// pow.Run()

}
