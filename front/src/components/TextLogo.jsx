import React from "react";
import { useNavigate } from "@tanstack/react-router";
import textLogo from "../assets/treviaTextLogo.png";
import "../styles/Logo.css";

const TextLogo = () => {
  const navigate = useNavigate();

  const redirectHome = () => {
    navigate({ to: "/" });
  };

  return (
    <div className="text-logo-area">
      <img
        src={textLogo}
        alt="Logo"
        className="text-logo-img"
        onClick={redirectHome}
      />
    </div>
  );
};

export default TextLogo;
