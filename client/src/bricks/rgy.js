import { Anchor, Box, Heading, Text } from "grommet";

import React from "react";

export default props => {
  let investors = [];
  if (props.investors !== undefined) {
    investors = props.investors;
  }
  return (
    <Box
      pad="medium"
      border={{ color: "brand", size: "small" }}
      margin="small"
      round={{ size: "xsmall" }}
      direction="column"
    >
      <Box direction="row" justify="between">
        <Box margin={{ horizontal: "small" }} justify="center" align="center">
          <Heading level={3}>{props.name}</Heading>
        </Box>
        <Box align="center" margin={{ left: "small" }} justify="center">
          <Text color="neutral-1">Cost</Text>
          <Text size="medium" color="neutral-1">
            {props.cost}
          </Text>
        </Box>
        <Box align="center" margin={{ left: "small" }} justify="center">
          <Text color="neutral-3">Shares</Text>
          <Text size="medium" color="neutral-3">
            {props.shares}
          </Text>
        </Box>
      </Box>
      <Anchor
        href={props.contract_address}
        label={"Contract: " + props.contract_address}
        margin={{ vertical: "xsmall" }}
        alignSelf="center"
      />
      <Box
        style={{ overflowY: "scroll", display: "inline" }}
        margin={{ top: "small" }}
      >
        {investors.map((inv, i) => {
          return (
            <Box key={i} direction="row" align="center" justify="center">
              <Text
                size="small"
                color="neutral-2"
                margin={{ right: "small", top: "xxsmall", bottom: "xxsmall" }}
                alignSelf="center"
              >
                Investor {i + 1}:
              </Text>
              <Text color="dark-1" size="small" alignSelf="center">
                {inv.address}
              </Text>
            </Box>
          );
        })}
      </Box>
    </Box>
  );
};
