import { Box, Button, Text } from "grommet";

import React from "react";

export default props => {
  return (
    <Box
      border={{ size: "small", color: "brand" }}
      pad="small"
      margin="small"
      round={{ size: "xsmall" }}
    >
      <Box margin={{ bottom: "small" }}>
        <Box>
          <Box align="center" justify="center">
            <Text size="large" color="dark-1" weight="bold">
              {props.name}
            </Text>
            <Text size="small" color="dark-6">
              {props.id}
            </Text>
          </Box>
          <Box direction="row">
            <Text size="small" color="neutral-2" margin={{ right: "small" }}>
              Public key:
            </Text>
            <Text size="small">{props.wallet_public_key}</Text>
          </Box>
        </Box>
      </Box>
      {props.rgy ? (
        <Box
          border={{ size: "small", color: "light-3" }}
          round={{ size: "xsmall" }}
          align="center"
          justify="center"
          pad="small"
        >
          <Text size="small">This Developer has a RGY</Text>
          <Text size="large" weight="bold" margin={{ vertical: "xxsmall" }}>
            {props.rgy.name}
          </Text>
          <Text size="small" color="dark-1">
            {props.rgy.contract_address}
          </Text>
        </Box>
      ) : (
        <Box align="center" justify="center" fill="vertical" direction="column">
          <Text size="small">You can to create a RGY</Text>
          <Box margin={{ top: "auto" }}>
            <Box fill="vertical" margin={{ top: "medium" }}>
              <Button
                label="Create RGY"
                onClick={() =>
                  props.onCreateRGY ? props.onCreateRGY(props.id) : null
                }
              ></Button>
            </Box>
          </Box>
        </Box>
      )}
    </Box>
  );
};
