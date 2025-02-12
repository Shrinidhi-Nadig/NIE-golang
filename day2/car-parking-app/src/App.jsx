import React from "react";
import { Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import Home from "./pages/home";
import CreateCar from "./pages/createcar";
import EditCar from "./pages/editcar";
import ViewCar from "./pages/viewcar";

const App = () => {
  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/create" element={<CreateCar />} />
        <Route path="/edit/:id" element={<EditCar />} />
        <Route path="/view/:id" element={<ViewCar />} />
      </Routes>
    </>
  );
};

export default App;