package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/xid"
)

// RegisterNewInvestor saves into our state the reference of a new Investor
func (state *RenvestgyState) RegisterNewInvestor(name string, publicKey common.Address) (*Investor, error) {
	if state.Investors == nil {
		state.Investors = make([]*Investor, 0)
	}

	newInvestor := &Investor{
		ID:              xid.New().String(),
		Name:            name,
		WalletPublicKey: publicKey,
	}
	state.Investors = append(state.Investors, newInvestor)

	err := state.saveState(".")
	if err != nil {
		return nil, err
	}

	return newInvestor, nil
}

// RegisterNewDeveloper register a new developer with or without its rgyx token
func (state *RenvestgyState) RegisterNewDeveloper(name string, publicKey common.Address, rgy ...*RGYx) (*Developer, error) {
	if state.Developers == nil {
		state.Developers = make([]*Developer, 0)
	}

	newDev := &Developer{
		ID:              xid.New().String(),
		Name:            name,
		WalletPublicKey: publicKey,
	}

	var linkedRGY *RGYx
	if len(rgy) > 0 {
		linkedRGY = rgy[0]
	}

	newDev.Rgy = linkedRGY
	state.Developers = append(state.Developers, newDev)

	err := state.saveState(".")
	if err != nil {
		return nil, err
	}

	return newDev, nil
}
