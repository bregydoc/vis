import { Box, Button, Grid, Heading, Image, Layer, Text } from "grommet";
import React, { useEffect, useState } from "react";

import BuyRGYx from "./bricks/buyRGYx";
import { ClassicState } from "./contexts";
import Dev from "./bricks/dev";
import Investor from "./bricks/investor";
import NewDevForm from "./bricks/newDev";
import NewRGY from "./bricks/newRGY";
import Rgy from "./bricks/rgy";
import { w3cwebsocket as W3CWebSocket } from "websocket";
import visLogo from "./images/vislogo.png";

const createNewDev = "CREATE_DEV";
const createNewInvestor = "CREATE_INVESTOR";
const createNewRGY = "CREATE_RGY";

const buyShares = "BUY_SHARES";

const devCreated = "DEV_CREATED";
const investorRegistered = "INVESTOR_REGISTERED";
const rgyDeployed = "RGY_DEPLOYED";

const sharesAreSold = "SHARES_ARE_SOLD";
const sharesAvailableToTransfer = "SHARES_TRANSFER_AVAILABLE";
const client = new W3CWebSocket("ws://127.0.0.1:5000/vis");

export default props => {
  const [showNewDev, setShowNewDev] = useState(false);
  const [showNewInvestor, setShowNewInvestor] = useState(false);
  const [showNewRGYGen, setShowNewRGYGen] = useState(false);
  const [showBuyRGYx, setShowBuyRGYx] = useState(false);

  const [newDevID, setNewDevID] = useState("");

  const { classicState, updateClassicState } = ClassicState.useContainer();

  const createNewRGYCallback = devID => {
    setShowNewRGYGen(true);
    setNewDevID(devID);
  };

  useEffect(() => {
    client.onopen = () => {
      console.log("WebSocket Client Connected");
    };
    client.onmessage = message => {
      const payload = JSON.parse(message.data);
      console.log(payload.type);
      updateClassicState();
      setShowNewDev(false);
      setShowNewRGYGen(false);
      setShowNewInvestor(false);
      setShowBuyRGYx(false);
    };
  }, []);

  const createNewDevCallback = data => {
    client.send(
      JSON.stringify({
        type: createNewDev,
        data: {
          name: data.name,
          address: data.address
        }
      })
    );
  };

  const createNewInvestorCallback = data => {
    console.log("inv", data);
    client.send(
      JSON.stringify({
        type: createNewInvestor,
        data: {
          name: data.name,
          address: data.address
        }
      })
    );
  };

  const createNewRGYWS = data => {
    client.send(
      JSON.stringify({
        type: createNewRGY,
        data: {
          developer_id: data.devID,
          name: data.name,
          total_shares: parseInt(data.shares),
          cost_cent: parseInt(data.cost)
        }
      })
    );
  };

  const buyRGYXToInvestor = data => {
    let ch = data.investor.split("(");
    const investorID = ch[1].slice(0, -1);

    ch = data.rgy.split("(");
    const rgyID = ch[1].slice(0, -1);

    console.log({
      investor_id: investorID,
      rgy_id: rgyID,
      shares: parseInt(data.shares)
    });

    client.send(
      JSON.stringify({
        type: buyShares,
        data: {
          investor_id: investorID,
          rgy_id: rgyID,
          shares: parseInt(data.shares)
        }
      })
    );
  };

  useEffect(() => {
    updateClassicState();
  }, []);

  return (
    <Grid
      rows={["1/3", "1/3", "1/3"]}
      columns={["1/2", "1/2"]}
      fill="vertical"
      areas={[
        { name: "investors", start: [0, 0], end: [0, 0] },
        { name: "devs", start: [0, 1], end: [0, 1] },
        { name: "rgys", start: [0, 2], end: [0, 2] },
        { name: "vis", start: [1, 0], end: [1, 2] }
      ]}
    >
      {showNewDev && (
        <Layer
          onEsc={() => setShowNewDev(false)}
          onClickOutside={() => setShowNewDev(false)}
        >
          <NewDevForm
            onClose={() => setShowNewDev(false)}
            onSubmit={createNewDevCallback}
          />
        </Layer>
      )}
      {showNewInvestor && (
        <Layer
          onEsc={() => setShowNewInvestor(false)}
          onClickOutside={() => setShowNewInvestor(false)}
        >
          <NewDevForm
            title="REGISTER INVESTOR"
            onClose={() => setShowNewInvestor(false)}
            onSubmit={createNewInvestorCallback}
          />
        </Layer>
      )}
      {showNewRGYGen && (
        <Layer
          onEsc={() => setShowNewRGYGen(false)}
          onClickOutside={() => setShowNewRGYGen(false)}
        >
          <NewRGY
            devID={newDevID}
            onClose={() => setShowNewRGYGen(false)}
            onSubmit={createNewRGYWS}
          />
        </Layer>
      )}
      {showBuyRGYx && (
        <Layer
          onEsc={() => setShowBuyRGYx(false)}
          onClickOutside={() => setShowBuyRGYx(false)}
        >
          <BuyRGYx
            investors={classicState.investors.map(i => `${i.name} (${i.id})`)}
            rgys={classicState.rgys.map(i => `${i.name} (${i.id})`)}
            onClose={() => setShowBuyRGYx(false)}
            onSubmit={buyRGYXToInvestor}
          />
        </Layer>
      )}

      <Box gridArea="devs" pad="small">
        <Box direction="row" justify="between">
          <Heading level="3" color="brand">
            Developers
          </Heading>
          <Box align="center" direction="column" justify="center">
            <Button
              label="Create New Dev"
              primary
              onClick={() => setShowNewDev(true)}
            />
          </Box>
        </Box>
        <div
          style={{
            display: "-webkit-inline-box",
            overflowX: "scroll",
            overflowY: "hidden"
          }}
        >
          {classicState.developers.map((dev, i) => {
            return (
              <Dev
                key={i}
                {...dev}
                onCreateRGY={devID => createNewRGYCallback(devID)}
              />
            );
          })}
        </div>
      </Box>
      <Box
        gridArea="rgys"
        pad="small"
        border={{ side: "top", size: "small", color: "light-5" }}
      >
        <Heading level="3" color="brand">
          RGYs
        </Heading>
        <div
          style={{
            display: "-webkit-inline-box",
            overflowX: "scroll",
            overflowY: "hidden"
          }}
        >
          {classicState.rgys.map((rgy, i) => {
            return <Rgy key={i} {...rgy} />;
          })}
        </div>
      </Box>

      <Box
        gridArea="investors"
        pad="small"
        border={{ side: "bottom", size: "small", color: "light-5" }}
      >
        <Box direction="row" justify="between">
          <Heading level="3" color="brand">
            Investors
          </Heading>
          <Box align="center" direction="column" justify="center">
            <Button
              label="Register New Investor"
              primary
              onClick={() => setShowNewInvestor(true)}
            />
          </Box>
        </Box>

        <div
          style={{
            display: "-webkit-inline-box",
            overflowX: "scroll",
            overflowY: "hidden"
          }}
        >
          {classicState.investors.map((inv, i) => {
            return <Investor key={i} {...inv} />;
          })}
        </div>
      </Box>
      <Box
        gridArea="vis"
        pad="small"
        border={{ side: "left", size: "small", color: "light-5" }}
      >
        <Box>
          <Box justify="between" direction="row" align="center">
            <Heading level="3" color="brand" margin={{ left: "small" }}>
              VIS Version: {classicState.version}
            </Heading>
            <Image
              src={visLogo}
              style={{ width: "auto", height: "48px" }}
              margin={{ right: "small" }}
            />
          </Box>
          <Box direction="row" margin={{ left: "small" }}>
            <Text size="small" color="neutral-2" margin={{ right: "small" }}>
              Public key:
            </Text>
            <Text size="small">{classicState.wallet_pulic_key}</Text>
          </Box>
          <Box direction="row" margin={{ left: "small" }}>
            <Text size="small" color="neutral-2" margin={{ right: "small" }}>
              Private key:
            </Text>
            <Text size="small">{classicState.private_key}</Text>
          </Box>
        </Box>
        <Box>
          <Button
            label={"Buy RGY to investor"}
            onClick={() => setShowBuyRGYx(true)}
          ></Button>
        </Box>
      </Box>
    </Grid>
  );
};
