package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/bregydoc/vis/rgy"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/olekukonko/tablewriter"
	"github.com/rs/xid"
)

func (state *RenvestgyState) getRGYInsights(rgyContract common.Address) (*big.Int, *big.Int, error) {
	r, err := rgy.NewRgy(rgyContract, state.HTTPEthClient)
	if err != nil {
		return nil, nil, err
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

// CreateNewRGYx creates a new RGY Token, based on the developer
func (state *RenvestgyState) CreateNewRGYx(c context.Context, developerID string, name string, totalShares int64, costCent int64) (*RGYx, error) {
	dev, err := state.getDevByID(developerID)
	if err != nil {
		return nil, err
	}

	auth, err := state.getAuthFromPrivateKey(c)
	if err != nil {
		return nil, err
	}

	rgyAddress, _, _, err := rgy.DeployRgy(auth, state.HTTPEthClient, name, dev.WalletPublicKey, big.NewInt(totalShares), big.NewInt(costCent))
	if err != nil {
		return nil, err
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

	err = state.saveState(".")
	if err != nil {
		return nil, err
	}

	contract := newRGY.ContractAddress
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contract},
	}

	logs := make(chan types.Log)
	sub, err := state.WSEthClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Println("[STATE] RGY " + contract.String() + " suscribed to events")
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case vLog := <-logs:
				contractAbi, err := abi.JSON(strings.NewReader(string(rgy.RgyABI)))
				if err != nil {
					log.Fatal(err)
				}

				eventShareSold := new(rgy.RgyShareSold)
				eventAvailableToTransferChange := new(rgy.RgyAvailableToTransferChange)

				if err = contractAbi.Unpack(eventShareSold, "ShareSold", vLog.Data); err == nil {
					state.Events <- RenvestgyEvent{
						Timestamp: time.Now(),
						Type:      "ShareSold",
						Data:      eventShareSold,
					}
				} else if err = contractAbi.Unpack(eventAvailableToTransferChange, "AvailableToTransferChange", vLog.Data); err == nil {
					state.Events <- RenvestgyEvent{
						Timestamp: time.Now(),
						Type:      "AvailableToTransferChange",
						Data:      eventAvailableToTransferChange,
					}
				}
			}
		}
	}()

	return newRGY, nil
}

// SellRGYToInvestor transfer shares o
func (state *RenvestgyState) SellRGYToInvestor(c context.Context, rgyxID string, to common.Address, how *big.Int) error {
	auth, err := state.getAuthFromPrivateKey(c)
	if err != nil {
		return err
	}

	rgyType, err := state.getRGYXByID(rgyxID)
	if err != nil {
		return err
	}

	r, err := rgy.NewRgy(rgyType.ContractAddress, state.HTTPEthClient)
	if err != nil {
		return err
	}
	_, err = r.SellShares(auth, to, how)
	if err != nil {
		return err
	}

	return nil
}

// GetBalanceOf returns how rgys of rgyx type has an user
func (state *RenvestgyState) GetBalanceOf(c context.Context, rgyxID string, addr common.Address) (*big.Int, error) {
	rgyType, err := state.getRGYXByID(rgyxID)
	if err != nil {
		return nil, err
	}

	r, err := rgy.NewRgy(rgyType.ContractAddress, state.HTTPEthClient)
	if err != nil {
		return nil, err
	}
	
	return r.Balance(nil, addr)
}
