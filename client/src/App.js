import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

import CreateRoom from "./components/CreateRoom";
import Room from "./components/Room";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" exact caseSensitive={false} element={<CreateRoom />}></Route>
        <Route path="/room/:roomID" caseSensitive={false} element={<Room />}></Route>
      </Routes>
    </Router>
  );
}

export default App;
