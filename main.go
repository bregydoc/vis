package main

import (
	"context"
)

func main() {
	ethNet := "127.0.0.1:7545"
	pKey := "75bb96751e7a5c4cad9e07e3c1e578fc467b0d4f879ae8056d2096d0402265d3"

	events := make(chan RenvestgyEvent, 1)

	ctx := context.Background()
	state, err := NewRenvestgyState(ctx, ethNet, pKey, events)
	if err != nil {
		panic(err)
	}

	err = state.getStateOfRGYs()
	if err != nil {
		panic(err)
	}
	// err = state.RegisterNewDeveloper("Energy Free 2020", common.HexToAddress("0x26f377dDF41B1E2Bd916e7Bdb52AC4214722C173"))
	// if err != nil {
	// 	panic(err)
	// }

	// devID := "bldf6q6ukpt47ddqjs7g"
	// err = state.CreateNewRGYx(ctx, devID, "RGY0", 100)
	// if err != nil {
	// 	panic(err)
	// }

}
