import React, { createContext, useContext, useEffect, useState } from "react";

import { ClassicState } from "./contexts";
import { Grommet } from "grommet";
import Vis from "./vis";

const theme = {
  global: {
    font: {
      family: "Roboto",
      size: "14px",
      height: "20px"
    },
    colors: {
      brand: "#4a65f5",
      "neutral-1": "#3debc3",
      "neutral-2": "#478ae7",
      "neutral-3": "#ff2c69"
    }
  }
};

export default () => {
  return (
    <ClassicState.Provider>
      <Grommet theme={theme} full={true}>
        <Vis />
      </Grommet>
    </ClassicState.Provider>
  );
};
