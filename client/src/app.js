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
