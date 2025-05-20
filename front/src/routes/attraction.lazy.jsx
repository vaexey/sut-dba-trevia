import { createLazyFileRoute, useRouterState } from "@tanstack/react-router";
import React, { useEffect, useState } from "react";
import "../styles/Attraction.css";

export const Route = createLazyFileRoute("/attraction")({
  component: AttractionPage,
});

function AttractionPage() {
  const search = useRouterState({ select: (s) => s.location.search });
  const params = new URLSearchParams(search);
  const id = params.get("id"); // Extract the "id" query parameter

  const [attractionData, setAttractionData] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (!id) {
      setError("No attraction ID provided in query parameters.");
      setIsLoading(false);
      return;
    }

    const fetchAttractionData = async () => {
      try {
        setIsLoading(true);

        const response = await fetch(`/api/v1/attractions/${id}`);
        if (!response.ok) {
          throw new Error("Failed to fetch attraction data.");
        }

        const data = await response.json();

        // Convert binary photo data to Base64 if necessary
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

        setAttractionData(formattedData);
      } catch (err) {
        setError(err.message);
      } finally {
        setIsLoading(false);
      }
    };

    fetchAttractionData();
  }, [id]);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div className="attraction-page">
      <div className="attraction-info">
        <div className="attraction-textdata">
          <h1 className="attraction-name">{attractionData.name}</h1>
          <p className="attraction-rating">Rating: {attractionData.rating}</p>
          <p className="attraction-description">{attractionData.description}</p>
        </div>
        <div className="attraction-photo">
          <img src={attractionData.photo} alt={attractionData.name} />
        </div>
      </div>
      <div className="comments-section">
        <h2>Comments</h2>
        <p>Comments will be displayed here soon...</p>
      </div>
    </div>
  );
}

export default AttractionPage;
