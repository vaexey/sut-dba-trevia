import React, { useState } from "react";
import "../styles/Navbar.css";
import SearchComponent from "./SearchComponent";

const Navbar = () => {
  const [selectedButton, setSelectedButton] = useState("All");

  return (
    <nav className="navbar">
      <SearchComponent
        selectedButton={selectedButton}
        setSelectedButton={setSelectedButton}
      />
    </nav>
  );
};

export default Navbar;
