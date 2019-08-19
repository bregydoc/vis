package main

import (
	"math/big"
	"os"

	"github.com/bregydoc/vis/rgy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/olekukonko/tablewriter"
)

func (state *RenvestgyState) getRGYInsights(rgyContract common.Address) (*big.Int, *big.Int, error) {
	r, err := rgy.NewRgy(rgyContract, state.HTTPEthClient)
	if err != nil {
		panic(err)
	}

	available, err := r.GetAvailableShares(nil)
	if err != nil {
		return nil, nil, err
	}

	genesis, err := r.GetGenesis(nil)
	if err != nil {
		return nil, nil, err
	}

	return available, genesis, nil
}

func (state *RenvestgyState) getStateOfRGYs() error {
	data := [][]string{}
	for _, r := range state.RGYs {
		av, gen, err := state.getRGYInsights(r.ContractAddress)
		if err != nil {
			return err
		}
		data = append(data, []string{r.Name, av.String(), gen.String()})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Available Shares", "Total Shares"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	return nil
}
