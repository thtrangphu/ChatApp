import React from "react";
import Chats from "./Chats.js";
import { Stack } from "@mui/material";

// const Cat = lazy(() => import("../../components/Cat"));
const GeneralApp = () => {
  return (
    <>
      {/* <Stack direction="row" sx={{ width: "100%" }}> */}
      <Chats />
      {/* </Stack> */}
    </>
  );
};

export default GeneralApp;
