import React, { useState, useRef, useEffect } from "react";
import { useNavigate } from "@tanstack/react-router";
import "../styles/SearchComponent.css";

function SearchComponent({ selectedButton, setSelectedButton }) {
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

  useEffect(() => {
    if (searchQuery.trim() === "") {
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
    <div className="search-component">
      <div className="buttons">
        <button
          className={selectedButton === "All" ? "active" : "nonactive"}
          onClick={() => setSelectedButton("All")}
        >
          Wszystko
        </button>
        <button
          className={selectedButton === "Hotel" ? "active" : "nonactive"}
          onClick={() => setSelectedButton("Hotel")}
        >
          Hotele
        </button>
        <button
          className={selectedButton === "Restaurant" ? "active" : "nonactive"}
          onClick={() => setSelectedButton("Restaurant")}
        >
          Restauracje
        </button>
        <button
          className={selectedButton === "Other" ? "active" : "nonactive"}
          onClick={() => setSelectedButton("Other")}
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
          className="text-search-box"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          onFocus={() => setShowDropdown(true)}
          onBlur={() => setTimeout(() => setShowDropdown(false), 200)}
          autoComplete="off"
        />
        <button type="submit" className="search-submit-button">
          ➤
        </button>
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
      </form>
    </div>
  );
}

export default SearchComponent;
