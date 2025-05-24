import { useNavigate } from "@tanstack/react-router";
import React, { useState, useEffect } from "react";
import { FiLogIn, FiLogOut } from "react-icons/fi";
import { FaPlus } from "react-icons/fa";
import AddAttractionModal from "./AddAttractionModal";
import "../styles/UtilButtons.css";

const UtilButtons = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [userRoleId, setUserRoleId] = useState(null);
  const [logoutMessage, setLogoutMessage] = useState("");
  const [showModal, setShowModal] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setIsLoggedIn(true);
      // Fetch user info to get roleId
      fetch("/api/v1/user", {
        headers: { Authorization: `Bearer ${token}` },
      })
        .then((res) => res.json())
        .then((data) => setUserRoleId(data.roleId))
        .catch(() => setUserRoleId(null));
    }
  }, []);

  const handleLogin = () => {
    setIsLoggedIn(true);
    navigate({ to: "/login" });
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    setIsLoggedIn(false);
    window.location.reload();
  };

  return (
    <div className="util-buttons">
      {isLoggedIn ? (
        <>
          <button onClick={() => setShowModal(true)}>
            <FaPlus style={{ marginRight: "0" }} />
          </button>
          {userRoleId === 3 && (
            <>
              <button onClick={() => navigate({ to: "/admin" })}>
                Admin Panel
              </button>
            </>
          )}
          {userRoleId === 2 && (
            <>
              <button onClick={() => navigate({ to: "/mod" })}>
                Moderation
              </button>
            </>
          )}
          <button onClick={handleLogout}>
            <FiLogOut /> Logout
          </button>
        </>
      ) : (
        <button onClick={handleLogin}>
          <FiLogIn /> Login
        </button>
      )}
      {logoutMessage && <p className="logout-message">{logoutMessage}</p>}

      {showModal && <AddAttractionModal onClose={() => setShowModal(false)} />}
    </div>
  );
};

export default UtilButtons;
