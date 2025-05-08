import { createLazyFileRoute } from "@tanstack/react-router";
import { useState, useEffect } from "react";

import funFactImg from "../assets/funFactBg.png";

import "../styles/Index.css";

export const Route = createLazyFileRoute("/")({
  component: Index,
});

function Index() {
  const [selectedButton, setSelectedButton] = useState("Wszystko");
  const [funFactBackground, setFunFactBackground] = useState("");
  const [isLoading, setIsLoading] = useState(true);

  const handleSearch = (event) => {
    event.preventDefault();
    const searchValue = event.target.search.value;
    console.log("Search value:", searchValue);
  };

  const handleButtonClick = (buttonName) => {
    setSelectedButton(buttonName);
  };

  //todo: integrate with api
  useEffect(() => {
    const fetchBackgroundImage = async () => {
      setIsLoading(true);

      const response = funFactImg;

      setFunFactBackground(response);
      setIsLoading(false);
    };

    fetchBackgroundImage();
  }, []);

  return (
    <div className="home-area">
      <div className="home-search-area">
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
            className={
              selectedButton === "Restauracje" ? "active" : "nonactive"
            }
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
        <form onSubmit={handleSearch}>
          <input type="text" name="search" placeholder="Search..." />
          <button type="submit">➤</button>
        </form>
      </div>
      <div
        className="home-fun-fact-area"
        style={{
          backgroundImage: isLoading ? "none" : `url(${funFactBackground})`,
        }}
      >
        <div className="fun-fact-content">
          <p className="fun-fact-title">Eiffel Tower</p>
          <p className="fun-fact-location">Paris, France</p>
          <p className="fun-fact-description">
            Lorem ipsum dolor sit amet consectetur adipiscing elit. Amet
            consectetur adipiscing elit quisque faucibus ex sapien. Quisque
            faucibus ex sapien vitae pellentesque sem placerat. Vitae
            pellentesque sem placerat in id cursus mi.
          </p>
        </div>
      </div>
    </div>
  );
}
