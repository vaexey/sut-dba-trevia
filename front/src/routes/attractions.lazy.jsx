import {
  createLazyFileRoute,
  useRouterState,
  useNavigate,
} from "@tanstack/react-router";
import React, { useEffect, useState } from "react";
import "../styles/Attractions.css";

export const Route = createLazyFileRoute("/attractions")({
  component: AttractionsPage,
});

function AttractionsPage() {
  const search = useRouterState({ select: (s) => s.location.search });
  const { id, category } = search;

  const [attractionData, setAttractionData] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  const navigate = useNavigate();

  useEffect(() => {
    if (!id) {
      setError("No ID provided in query parameters.");
      setIsLoading(false);
      return;
    }

    const fetchAttractionData = async () => {
      try {
        setIsLoading(true);

        // Fetch data from both endpoints concurrently
        const [attractionsResponse, locationResponse] = await Promise.all([
          fetch(
            category !== "All"
              ? `/api/v1/attractions/location/${id}?category=${category.toLowerCase()}`
              : `/api/v1/attractions/location/${id}`,
          ),
          fetch(`/api/v1/locations/${id}`),
        ]);

        // Check if both responses are OK
        if (!attractionsResponse.ok || !locationResponse.ok) {
          throw new Error("Failed to fetch data from one or both endpoints.");
        }

        // Parse JSON responses
        const attractionsData = await attractionsResponse.json();
        const locationData = await locationResponse.json();

        // Convert binary photo data to Base64 for attractions
        const formattedAttractions = attractionsData.map((attraction) => ({
          ...attraction,
          photo: `data:image/jpeg;base64,${btoa(
            attraction.photo
              .replace(/\\x/g, "")
              .match(/.{1,2}/g)
              .map((byte) => String.fromCharCode(parseInt(byte, 16)))
              .join(""),
          )}`,
        }));

        // Combine location data and attractions data
        const combinedData = {
          location: {
            name: locationData.name,
            description: locationData.description,
          },
          attractions: formattedAttractions,
        };

        console.log("Combined Data:", combinedData);

        setAttractionData(combinedData);
      } catch (err) {
        setError(err.message);
      } finally {
        setIsLoading(false);
      }
    };

    fetchAttractionData();
  }, [id, category]);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div className="attractions-page">
      <div className="location-description">
        <h1>{attractionData.location.name}</h1>
        <p>{attractionData.location.description}</p>
      </div>
      <div className="attractions-grid-container">
        {attractionData.attractions.map((attraction) => (
          <div
            key={attraction.id}
            className="attractions-grid-item"
            onClick={(e) => {
              navigate({
                to: `/attraction?id=${attraction.id}`,
              });
            }}
          >
            <img
              alt={attraction.name}
              src={attraction.photo}
              className="attractions-photo"
            />
            <h3 className="attractions-name">{attraction.name}</h3>
            <p className="attractions-rating">Rating: {attraction.rating}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default AttractionsPage;
