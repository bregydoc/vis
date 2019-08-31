import React, { useState } from "react";

// import { w3cwebsocket as W3CWebSocket } from "websocket";
import Web3 from "web3";
import { abi } from "./abi";
import axios from "axios";
import { createContainer } from "unstated-next";

const API = "http://127.0.0.1:5000";
const ETHNET = "http://127.0.0.1:7545";

const web3 = new Web3(new Web3.providers.HttpProvider(ETHNET));

const useClassicState = () => {
  let [classicState, setClassicState] = useState({
    rgys: [],
    developers: [],
    investors: []
  });

  let updateClassicState = () => {
    axios.get(`${API}/state`).then(res => {
      if (res.status == 200) {
        const d = res.data;
        let sharesProms = [];
        let costsProms = [];
        let shareholders = {};
        console.log(res.data);
        res.data.rgys.map(rgy => {
          var rgyxContract = new web3.eth.Contract(abi, rgy.contract_address);
          rgyxContract.methods
            .getTotalShareholders()
            .call()
            .then(totalShareholders => {
              for (let i = 0; i < totalShareholders; i++) {
                if (typeof shareholders[rgy.id] === "undefined") {
                  shareholders[rgy.id] = [];
                }
                shareholders[rgy.id].push(
                  rgyxContract.methods.shareholders(i).call()
                );
              }
            });
          sharesProms.push(rgyxContract.methods.getAvailableShares().call());
          costsProms.push(rgyxContract.methods.cost().call());
        });

        Promise.all(sharesProms).then(results => {
          results.map((r, i) => {
            const rgy = res.data.rgys[i];
            d.rgys = d.rgys.map(rg =>
              rg.contract_address === rgy.contract_address
                ? { ...rg, shares: r }
                : rg
            );
          });
          Promise.all(costsProms)
            .then(results => {
              results.map((r, i) => {
                const rgy = res.data.rgys[i];
                d.rgys = d.rgys.map(rg =>
                  rg.contract_address === rgy.contract_address
                    ? { ...rg, cost: `$${(r / 100).toFixed(2)}` }
                    : rg
                );
              });
              // ! FIX IT LATER, BREGY!
              for (let k in shareholders) {
                Promise.all(shareholders[k]).then(results => {
                  results.map((r, i) => {
                    console.log(k, i, r);
                    d.rgys = d.rgys.map(rg => {
                      if (rg.id === k) {
                        if (typeof rg.investors === "undefined") {
                          rg.investors = [];
                        }
                        return {
                          ...rg,
                          investors: [...rg.investors, { address: r }]
                        };
                      } else {
                        return rg;
                      }
                    });
                  });
                  setClassicState(d);
                });
              }
              setClassicState(d);
            })

            .catch(err => {
              d.rgys = d.rgys.map(rg => ({ ...rg, cost: "undef" }));
              setClassicState(d);
            });
        });
      }
    });
  };
  return { classicState, updateClassicState };
};

let ClassicState = createContainer(useClassicState);

export { ClassicState };
