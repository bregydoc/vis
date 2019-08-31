package main

import (
	"context"
	"log"
)

func main() {
	ethNet := "127.0.0.1:7545"
	pKey := "dfd5be1b32cbe936876dc9c5c4011f8ca6aa88106923c3bbd458258014898cf1"

	events := make(chan RenvestgyEvent, 1)

	ctx := context.Background()

	state, err := NewRenvestgyState(ctx, ethNet, pKey, events)
	if err != nil {
		panic(err)
	}

	log.Println("[STATE]", state.Version)

	// rgyID := "blipbl9rcvk4qlr92n6g"
	// err = state.SellRGYToInvestor(ctx, rgyID, common.HexToAddress("0xb4eF0F9D44c2104230cA1E2CEaD2228FdeF3016b"), big.NewInt(20))
	// if err != nil {
	// 	panic(err)
	// }

	// dev, err := state.RegisterNewDeveloper("EnergyVI", common.HexToAddress("0x0665a053AB38c02Aa00eAA9daEc15bFCa5f2aAB7"))
	// if err != nil {
	// 	panic(err)
	// }

	// devID := dev.ID
	// err = state.CreateNewRGYx(ctx, devID, "RGYX", 500, 1000)
	// if err != nil {
	// 	panic(err)
	// }

	// err = state.getStateOfRGYs()
	// if err != nil {
	// 	panic(err)
	// }

	// how, err := state.GetBalanceOf(ctx, "blipbl9rcvk4qlr92n6g", common.HexToAddress("0xb4eF0F9D44c2104230cA1E2CEaD2228FdeF3016b"))
	// if err != nil {
	// 	panic(err)
	// }

	// log.Println(how)

	state.launchBasicAPI(":5000")

}
