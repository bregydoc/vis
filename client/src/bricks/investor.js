import { Box, Heading, Image, Text } from "grommet";

import React from "react";

export default props => {
  return (
    <Box
      pad="small"
      margin="small"
      border={{ color: "brand", size: "small" }}
      round={{ size: "xsmall" }}
    >
      <Box direction="row" width="auto">
        <Image
          src={
            "https://ui-avatars.com/api/?rounded=true&name=" +
            props.name.replace(" ", "+")
          }
          width="64"
          height="64"
        />
        <Box margin={{ horizontal: "medium", vertical: "xsmall" }}>
          <Text size="large" weight="bold">
            {props.name}
          </Text>
          <Text size="small" color="dark-6">
            {props.id}
          </Text>
        </Box>
        <Box
          margin={{ left: "auto", right: "xxsmall", vertical: "xsmall" }}
          justify="center"
          direction="column"
        >
          <Text size="medium" alignSelf="center" justify="center">
            Credits
          </Text>
          <Text size="medium" color="dark-2">
            $ 200.00
          </Text>
        </Box>
      </Box>
      <Box margin={{ vertical: "small" }} width="auto">
        <Box direction="row">
          <Text size="small" color="neutral-2" margin={{ right: "small" }}>
            Public key:
          </Text>
          <Text size="small">{props.wallet_public_key}</Text>
        </Box>
      </Box>
      {/* <Box>Hello World</Box> */}
    </Box>
  );
};
