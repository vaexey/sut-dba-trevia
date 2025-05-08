import React from "react";
import { useNavigate } from "@tanstack/react-router";
import logo from "../assets/legitTreviaLogo.png";
import "../styles/Logo.css";

const PageLogo = () => {
  const navigate = useNavigate();

  const redirectHome = () => {
    navigate({ to: "/" });
  };

  return (
    <div className="logo-area">
      <img src={logo} alt="Logo" className="logo-img" onClick={redirectHome} />
    </div>
  );
};

export default PageLogo;
