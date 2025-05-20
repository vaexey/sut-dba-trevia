import { createLazyFileRoute } from "@tanstack/react-router";
import { useState, useEffect } from "react";
import SearchComponent from "../components/SearchComponent";
import funFactImg from "../assets/funFactBg.png";
import "../styles/Index.css";

export const Route = createLazyFileRoute("/")({
  component: Index,
});

function Index() {
  const [selectedButton, setSelectedButton] = useState("All");
  const [funFactBackground, setFunFactBackground] = useState("");
  const [isLoading, setIsLoading] = useState(true);

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
        <SearchComponent
          selectedButton={selectedButton}
          setSelectedButton={setSelectedButton}
        />
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

export default Index;
