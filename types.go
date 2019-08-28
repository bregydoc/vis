package main

import (
	"github.com/ethereum/go-ethereum/common"
)

// RGYx represents a RGY n share
type RGYx struct {
	ID              string         `json:"id"`
	DeveloperID     string         `json:"developer_id"`
	Developer       string         `json:"developer"`
	Name            string         `json:"name"`
	ContractAddress common.Address `json:"contract_address"`
}

// Investor is any natural person who can save RGYs (RGYx)
type Investor struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	WalletPublicKey common.Address `json:"wallet_public_key"`
}

// Developer is our energy developer partner
type Developer struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	WalletPublicKey common.Address `json:"wallet_public_key"`
	Rgy             *RGYx          `json:"rgy"`
}
