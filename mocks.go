package main

import (
	"github.com/ethereum/go-ethereum/common"
)

// InvBregy are only a investor mockup
var InvBregy = &Investor{
	ID:              "0xbregy",
	Name:            "Bregy",
	WalletPublicKey: common.HexToAddress("0x41Cd526757c227e7375195c297C85EC757375E8B"),
}

// InvX are only a investor mockup
var InvX = &Investor{
	ID:              "0xinvx",
	Name:            "Inv X",
	WalletPublicKey: common.HexToAddress("0xB217C18CC17B4d2576283D5E31A2e1c81E56122d"),
}

// InvY are only a investor mockup
var InvY = &Investor{
	ID:              "0xinvy",
	Name:            "Inv Y",
	WalletPublicKey: common.HexToAddress("0x1eD873c9aD2406e16Ebb676dE6eBf809b40C9191"),
}

// InvZ are only a investor mockup
var InvZ = &Investor{
	ID:              "0xinvz",
	Name:            "Inv Z",
	WalletPublicKey: common.HexToAddress("0x1F6DfB38fEC01f8D4B315026FABf8AC460b175e9"),
}
