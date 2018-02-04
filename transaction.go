package main

// Transaction data
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}
