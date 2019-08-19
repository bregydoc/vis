package main

import (
	"context"

	"math/big"

	"github.com/bregydoc/vis/rgy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/xid"
)

// RegisterNewInvestor saves into our state the reference of a new Investor
func (state *RenvestgyState) RegisterNewInvestor(name string, publicKey common.Address) error {
	if state.Investors == nil {
		state.Investors = make([]*Investor, 0)
	}
	state.Investors = append(state.Investors, &Investor{
		ID:              xid.New().String(),
		Name:            name,
		WalletPublicKey: publicKey,
	})

	return state.saveState(".")
}

// RegisterNewDeveloper register a new developer with or without its rgyx token
func (state *RenvestgyState) RegisterNewDeveloper(name string, publicKey common.Address, rgy ...*RGYx) error {
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

	return state.saveState(".")
}

// CreateNewRGYx creates a new RGY Token, based on the developer
func (state *RenvestgyState) CreateNewRGYx(c context.Context, developerID string, name string, totalShares int64) error {
	dev, err := state.getDevByID(developerID)
	if err != nil {
		return err
	}

	auth, err := state.getAuthFromPrivateKey(c)
	if err != nil {
		return err
	}

	rgyAddress, _, _, err := rgy.DeployRgy(auth, state.HTTPEthClient, name, big.NewInt(totalShares))
	if err != nil {
		return err
	}

	// Not my fault
	auth.Nonce.Add(auth.Nonce, big.NewInt(1))

	newRGY := &RGYx{
		ID:              xid.New().String(),
		Developer:       dev.Name,
		DeveloperID:     dev.ID,
		Name:            name,
		ContractAddress: rgyAddress,
	}

	state.RGYs = append(state.RGYs, newRGY)

	for _, d := range state.Developers {
		if d.ID == dev.ID {
			d.Rgy = newRGY
		}
	}

	return state.saveState(".")
}
