import { Box, Button, Grid, Heading, Layer } from "grommet";
import React, { useEffect, useState } from "react";

import { ClassicState } from "./contexts";
import Dev from "./bricks/dev";
import NewDevForm from "./bricks/newDev";
import NewRGY from "./bricks/newRGY";
import Rgy from "./bricks/rgy";
import { w3cwebsocket as W3CWebSocket } from "websocket";

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
  const [showNewRGYGen, setShowNewRGYGen] = useState(false);
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

  useEffect(() => {
    updateClassicState();
  }, []);

  return (
    <Grid
      rows={["4/6", "5/6"]}
      columns={["1/2", "1/2"]}
      fill="vertical"
      areas={[
        { name: "devs", start: [0, 0], end: [0, 0] },
        { name: "rgys", start: [0, 1], end: [0, 1] },
        { name: "users", start: [1, 0], end: [1, 1] }
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
      <Box gridArea="rgys" pad="small">
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

      <Box gridArea="users" pad="small">
        Users
      </Box>
    </Grid>
  );
};
