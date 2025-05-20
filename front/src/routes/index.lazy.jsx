import { createLazyFileRoute } from "@tanstack/react-router";
import { useState, useEffect } from "react";
import SearchComponent from "../components/SearchComponent";
import "../styles/Index.css";

export const Route = createLazyFileRoute("/")({
  component: Index,
});

function Index() {
  const [selectedButton, setSelectedButton] = useState("All");
  const [funFactData, setFunFactData] = useState(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchFunFactData = async () => {
      try {
        setIsLoading(true);

        const response = await fetch("/api/v1/attractions/funfact");
        if (!response.ok) {
          throw new Error("Failed to fetch fun fact data.");
        }

        const data = await response.json();

        const formattedData = {
          ...data,
          photo: `data:image/jpeg;base64,${btoa(
            data.photo
              .replace(/\\x/g, "")
              .match(/.{1,2}/g)
              .map((byte) => String.fromCharCode(parseInt(byte, 16)))
              .join(""),
          )}`,
        };

        setFunFactData(formattedData);
      } catch (error) {
        console.error("Error fetching fun fact data:", error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchFunFactData();
  }, []);

  return (
    <div className="home-area">
      <div className="home-search-area">
        <SearchComponent
          selectedButton={selectedButton}
          setSelectedButton={setSelectedButton}
        />
      </div>

      {isLoading ? (
        <p>Loading fun fact...</p>
      ) : funFactData ? (
        <div className="home-fun-fact-area">
          <div className="fun-fact-content">
            <p className="fun-fact-title">{funFactData.name}</p>
            <p className="fun-fact-description">{funFactData.funfact}</p>
          </div>
          <div className="fun-fact-photo">
            <img src={funFactData.photo} alt={funFactData.name} />
          </div>
        </div>
      ) : (
        <p>Failed to load fun fact.</p>
      )}
    </div>
  );
}

export default Index;
