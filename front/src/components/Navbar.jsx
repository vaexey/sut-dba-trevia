import React, { useState } from "react";
import "../styles/Navbar.css";

const Navbar = () => {
  const [selectedButton, setSelectedButton] = useState("Wszystko");

  const handleSearch = (event) => {
    event.preventDefault();
    const searchValue = event.target.search.value;
    console.log("Search value:", searchValue);
  };

  const handleButtonClick = (buttonName) => {
    setSelectedButton(buttonName);
  };

  return (
    <nav className="navbar">
      <form onSubmit={handleSearch}>
        <input type="text" name="search" placeholder="Search..." />
        <button type="submit">➤</button>
      </form>

      <div className="buttons">
        <button
          className={selectedButton === "Wszystko" ? "active" : "nonactive"}
          onClick={() => handleButtonClick("Wszystko")}
        >
          Wszystko
        </button>
        <button
          className={selectedButton === "Hotele" ? "active" : "nonactive"}
          onClick={() => handleButtonClick("Hotele")}
        >
          Hotele
        </button>
        <button
          className={selectedButton === "Restauracje" ? "active" : "nonactive"}
          onClick={() => handleButtonClick("Restauracje")}
        >
          Restauracje
        </button>
        <button
          className={selectedButton === "Inne" ? "active" : "nonactive"}
          onClick={() => handleButtonClick("Inne")}
        >
          Inne
        </button>
      </div>
    </nav>
  );
};

export default Navbar;
