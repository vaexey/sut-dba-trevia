import { createLazyFileRoute, useNavigate } from "@tanstack/react-router";
import { useState, useEffect, useRef } from "react";

import funFactImg from "../assets/funFactBg.png";

import "../styles/Index.css";

export const Route = createLazyFileRoute("/")({
  component: Index,
});

function Index() {
  const [selectedButton, setSelectedButton] = useState("Wszystko");
  const [funFactBackground, setFunFactBackground] = useState("");
  const [isLoading, setIsLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState("");
  const [searchResults, setSearchResults] = useState([]);
  const [showDropdown, setShowDropdown] = useState(false);
  const inputBoxRef = useRef(null);
  const navigate = useNavigate();

  const handleDropdownClick = (name) => {
    setSearchQuery(name);
    setShowDropdown(false);
    inputBoxRef.current.blur();
  };

  var userSelected = false;

  const handleSearch = (event) => {
    event.preventDefault();

    const selectedLocation = searchResults.find(
      (result) => result.name === searchQuery,
    );

    if (!selectedLocation) {
      console.log("Please select a valid option from the dropdown.");
      return;
    }

    navigate({
      to: `/attractions?id=${selectedLocation.id}&category=${selectedButton}`,
    });
  };

  const handleButtonClick = (buttonName) => {
    setSelectedButton(buttonName);
  };

  useEffect(() => {
    const fetchBackgroundImage = async () => {
      setIsLoading(true);

      const response = funFactImg;

      setFunFactBackground(response);
      setIsLoading(false);
    };

    fetchBackgroundImage();
  }, []);

  useEffect(() => {
    if (searchQuery.trim() === "" || userSelected === true) {
      setSearchResults([]);
      return;
    }

    const fetchSearchResults = async () => {
      try {
        const response = await fetch(
          `/api/v1/locations/search?query=${searchQuery}`,
        );
        if (!response.ok) {
          throw new Error("Failed to fetch search results");
        }
        const data = await response.json();
        console.log("API Response:", data);
        setSearchResults(data);
      } catch (error) {
        console.error("Error fetching search results:", error);
        setSearchResults([]);
      }
    };

    const debounceTimeout = setTimeout(fetchSearchResults, 300);
    return () => clearTimeout(debounceTimeout);
  }, [searchQuery]);

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
          <input
            ref={inputBoxRef}
            type="text"
            name="search"
            placeholder="Search..."
            value={searchQuery}
            onChange={(e) => {
              setSearchQuery(e.target.value);
            }}
            onFocus={() => setShowDropdown(true)}
            onBlur={() => setTimeout(() => setShowDropdown(false), 200)}
            autoComplete="off"
          />
          <button type="submit">➤</button>
        </form>
        {showDropdown && searchResults.length > 0 && (
          <div className="dropdown-menu">
            {searchResults.map((result) => (
              <div
                key={result.id}
                className="dropdown-item"
                onClick={() => handleDropdownClick(result.name)}
              >
                {result.name}, {result.type}
              </div>
            ))}
          </div>
        )}
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
