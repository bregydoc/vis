package main

import (
	"github.com/ethereum/go-ethereum/common"
)

// RGYx represents a RGY n share
type RGYx struct {
	ID              string
	DeveloperID     string
	Developer       string
	Name            string
	ContractAddress common.Address
}

// Investor is any natural person who can save RGYs (RGYx)
type Investor struct {
	ID              string
	Name            string
	WalletPublicKey common.Address
}

// Developer is our energy developer partner
type Developer struct {
	ID              string
	Name            string
	WalletPublicKey common.Address
	Rgy             *RGYx
}
