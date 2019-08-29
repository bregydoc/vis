import React, { useState } from "react";

// import { w3cwebsocket as W3CWebSocket } from "websocket";
import Web3 from "web3";
import { abi } from "./abi";
import axios from "axios";
import { createContainer } from "unstated-next";

const API = "http://127.0.0.1:5000";

const web3 = new Web3(new Web3.providers.HttpProvider("http://127.0.0.1:7545"));

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
        console.log(res.data);
        res.data.rgys.map(rgy => {
          var rgyxContract = new web3.eth.Contract(abi, rgy.contract_address);
          const l = rgyxContract.methods.shareholders.length;
          console.log(l);
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
                    ? { ...rg, cost: r / 100 }
                    : rg
                );
              });
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
