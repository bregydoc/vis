package main

import (
	"encoding/json"
	"log"

	"context"

	"net/http"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/olahol/melody.v1"
)

type eventType string

const createNewDev eventType = "CREATE_DEV"
const createNewInvestor eventType = "CREATE_INVESTOR"
const createNewRGY eventType = "CREATE_RGY"

const buyShares eventType = "BUY_SHARES"

const devCreated eventType = "DEV_CREATED"
const investorRegistered eventType = "INVESTOR_REGISTERED"
const rgyDeployed eventType = "RGY_DEPLOYED"

const sharesAreSold eventType = "SHARES_ARE_SOLD"
const sharesAvailableToTransfer eventType = "SHARES_TRANSFER_AVAILABLE"

type visEvent struct {
	Type eventType   `json:"type"`
	Data interface{} `json:"data"`
}

type newDevParams struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type newInvestorParams struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type newRGYParams struct {
	DeveloperID string `json:"developer_id"`
	Name        string `json:"name"`
	TotalShares int64  `json:"total_shares"`
	CostCent    int64  `json:"cost_cent"`
}

type buySharesParams struct {
	InvestorID string `json:"investor_id"`
	RGYID      string `json:"rgy_id"`
	How        int64  `json:"how"`
}

func (state *RenvestgyState) launchBasicAPI(port string) {
	r := gin.Default()
	r.Use(cors.Default())

	m := melody.New()

	r.GET("/vis", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.GET("/state", func(c *gin.Context) {
		c.JSON(http.StatusOK, state)
	})

	go func() {
		for {
			select {
			case e := <-state.Events:
				if e.Type == "ShareSold" {
					message, err := json.Marshal(visEvent{
						Type: sharesAreSold,
						Data: e.Data,
					})
					if err != nil {
						log.Fatal(err)
					}
					m.Broadcast(message)
				} else if e.Type == "AvailableToTransferChange" {
					message, err := json.Marshal(visEvent{
						Type: sharesAvailableToTransfer,
						Data: e.Data,
					})
					if err != nil {
						log.Fatal(err)
					}
					m.Broadcast(message)
				}
			}
		}
	}()

	m.HandleConnect(func(s *melody.Session) {
		log.Println("[NEW CONNECTION]", s.Request.Host)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		event := new(visEvent)
		err := json.Unmarshal(msg, event)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("[NEW MESSAGE]", s.Request.Host, event.Type)
		switch event.Type {
		case createNewDev:
			paramsMap, isOk := event.Data.(map[string]interface{})
			if !isOk {
				log.Println("Invalid Params")
				return
			}

			params := new(newDevParams)
			if err = mapstructure.Decode(paramsMap, params); err != nil {
				log.Println("Invalid Params")
				return
			}

			dev, err := state.RegisterNewDeveloper(params.Name, common.HexToAddress(params.Address))
			if err != nil {
				log.Println(err)
			}

			message, err := json.Marshal(visEvent{
				Type: devCreated,
				Data: dev,
			})
			if err != nil {
				log.Println(err)
				return
			}

			m.Broadcast(message)
		case createNewInvestor:
			paramsMap, isOk := event.Data.(map[string]interface{})
			if !isOk {
				log.Println("Invalid Params")
				return
			}

			params := new(newInvestorParams)
			if err = mapstructure.Decode(paramsMap, params); err != nil {
				log.Println("Invalid Params")
				return
			}

			inv, err := state.RegisterNewInvestor(params.Name, common.HexToAddress(params.Address))
			if err != nil {
				log.Println(err)
				return
			}

			message, err := json.Marshal(visEvent{
				Type: investorRegistered,
				Data: inv,
			})
			if err != nil {
				log.Fatal(err)
			}

			m.Broadcast(message)

		case createNewRGY:
			paramsMap, isOk := event.Data.(map[string]interface{})
			if !isOk {
				log.Println("Invalid Params")
				return
			}

			params := new(newRGYParams)
			params.DeveloperID = paramsMap["developer_id"].(string)
			params.Name = paramsMap["name"].(string)
			params.TotalShares = int64(paramsMap["total_shares"].(float64))
			params.CostCent = int64(paramsMap["cost_cent"].(float64))

			rgyx, err := state.CreateNewRGYx(context.Background(), params.DeveloperID, params.Name, params.TotalShares, params.CostCent)
			if err != nil {
				log.Println(err)
				return
			}

			message, err := json.Marshal(visEvent{
				Type: rgyDeployed,
				Data: rgyx,
			})
			if err != nil {
				log.Fatal(err)
			}

			m.Broadcast(message)
		case buyShares:
			paramsMap, isOk := event.Data.(map[string]interface{})
			if !isOk {
				log.Println("Invalid Params")
				return
			}

			params := new(buySharesParams)

			params.InvestorID = paramsMap["investor_id"].(string)
			params.RGYID = paramsMap["rgy_id"].(string)
			params.How = int64(paramsMap["shares"].(float64))

			err = state.SellRGYToInvestor(context.Background(), params.RGYID, params.InvestorID, big.NewInt(params.How))
			if err != nil {
				log.Println(err)
				return
			}
		}
	})

	r.Run(port)
}
