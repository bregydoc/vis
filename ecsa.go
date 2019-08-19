package main

import (
	"context"

	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (state *RenvestgyState) getAuthFromPrivateKey(c context.Context) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(state.PrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := state.HTTPEthClient.PendingNonceAt(c, fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := state.HTTPEthClient.SuggestGasPrice(c)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(900000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}
