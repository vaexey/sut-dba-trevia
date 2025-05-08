import { useNavigate } from "@tanstack/react-router";
import React, { useState, useEffect } from "react";
import { FiLogIn, FiLogOut } from "react-icons/fi";
import "../styles/UtilButtons.css";

const UtilButtons = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [logoutMessage, setLogoutMessage] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setIsLoggedIn(true);
    }
  }, []);

  const handleLogin = () => {
    setIsLoggedIn(true);
    navigate({ to: "/login" });
  };
  const handleLogout = () => {
    localStorage.removeItem("token");
    setIsLoggedIn(false);
    setLogoutMessage("You have been successfully logged out");

    setTimeout(() => {
      setLogoutMessage("");
    }, 3000);
  };

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
      {logoutMessage && <p className="logout-message">{logoutMessage}</p>}
    </div>
  );
};

export default UtilButtons;
