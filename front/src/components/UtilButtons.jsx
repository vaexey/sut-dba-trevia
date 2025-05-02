import { useNavigate } from "@tanstack/react-router";
import React, { useState } from "react";
import { FiLogIn, FiLogOut } from "react-icons/fi";
import "../styles/UtilButtons.css";

const UtilButtons = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const navigate = useNavigate();

  const handleLogin = () => {
    setIsLoggedIn(true);
    navigate({ to: "/login" });
  };
  const handleLogout = () => setIsLoggedIn(false);

  return (
    <div className="util-buttons">
      {isLoggedIn ? (
        <button onClick={handleLogout}>
          <FiLogOut /> Logout
        </button>
      ) : (
        <button onClick={handleLogin}>
          <FiLogIn /> Login
        </button>
      )}
    </div>
  );
};

export default UtilButtons;
