package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"encoding/json"
	"os"
	"path"

	"context"

	"github.com/bregydoc/vis/rgy"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// RenvestgyState represents a global state of renvestgy
type RenvestgyState struct {
	Version   string
	Timestamp time.Time

	WalletPulicKey common.Address
	PrivateKey     string

	Investors  []*Investor
	Developers []*Developer
	RGYs       []*RGYx

	HTTPEthClient *ethclient.Client `json:"-"`
	WSEthClient   *ethclient.Client `json:"-"`
}

// RenvestgyEvent is an renvestgy event
type RenvestgyEvent struct {
	Timestamp time.Time
	Type      string
	Info      string
}

// NewRenvestgyState returns a new instance of renvestgy state
func NewRenvestgyState(c context.Context, baseURL string, privateKey string, events chan RenvestgyEvent) (*RenvestgyState, error) {
	clientWS, err := ethclient.Dial("ws://" + baseURL)
	if err != nil {
		return nil, err
	}

	clientHTTP, err := ethclient.Dial("http://" + baseURL)
	if err != nil {
		return nil, err
	}

	s := &RenvestgyState{
		HTTPEthClient: clientHTTP,
		WSEthClient:   clientWS,
	}

	err = s.loadState(".")
	if err != nil {
		s.Timestamp = time.Now()
		s.PrivateKey = privateKey
		s.Version = "0.0.1"
		s.Investors = []*Investor{}
		s.Developers = []*Developer{}
		s.RGYs = []*RGYx{}

		err = s.saveState(".")
		if err != nil {
			return nil, err
		}
	}

	if len(s.RGYs) > 0 {
		for _, r := range s.RGYs {
			contract := r.ContractAddress
			query := ethereum.FilterQuery{
				Addresses: []common.Address{contract},
			}

			logs := make(chan types.Log)
			sub, err := s.WSEthClient.SubscribeFilterLogs(context.Background(), query, logs)
			if err != nil {
				log.Fatal(err)
			}

			go func() {
				for {
					select {
					case err := <-sub.Err():
						log.Fatal(err)
					case vLog := <-logs:

						contractAbi, err := abi.JSON(strings.NewReader(string(rgy.RgyABI)))
						if err != nil {
							log.Fatal(err)
						}

						event := new(rgy.RgyShareSold)
						err = contractAbi.Unpack(event, "ShareSold", vLog.Data)
						if err != nil {
							log.Fatal(err)
						}
						events <- RenvestgyEvent{
							Timestamp: time.Now(),
							Type:      "new sold",
							Info:      fmt.Sprintf("to: %s, how: %s", event.To.Hex(), event.How),
						}
						fmt.Println(event.To.Hex(), event.How)
					}
				}
			}()
		}
	}

	return s, nil

}

func (state *RenvestgyState) saveState(dir string) error {
	if dir == "" {
		dir = "."
	}

	stateFile := path.Join(dir, "state.json")

	file, err := os.OpenFile(stateFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	return enc.Encode(state)
}

func (state *RenvestgyState) loadState(dir string) error {
	if dir == "" {
		dir = "."
	}

	stateFile := path.Join(dir, "state.json")

	file, err := os.OpenFile(stateFile, os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	dec := json.NewDecoder(file)

	return dec.Decode(state)
}
