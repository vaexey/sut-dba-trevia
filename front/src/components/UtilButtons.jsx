import { useNavigate } from "@tanstack/react-router";
import React, { useState, useEffect } from "react";
import { FiLogIn, FiLogOut } from "react-icons/fi";
import { FaPlus } from "react-icons/fa";
import AddAttractionModal from "./AddAttractionModal";
import "../styles/UtilButtons.css";

const UtilButtons = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [logoutMessage, setLogoutMessage] = useState("");
  const [showModal, setShowModal] = useState(false);
  const [formData, setFormData] = useState({
    name: "",
    description: "",
    funfact: "",
    photo: null,
    location: "",
    type: "Hotel",
  });
  const [searchResults, setSearchResults] = useState([]);
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
    window.location.reload();
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handlePhotoChange = (e) => {
    setFormData((prev) => ({ ...prev, photo: e.target.files[0] }));
  };

  const handleLocationSearch = async (query) => {
    if (!query.trim()) {
      setSearchResults([]);
      return;
    }

    try {
      const response = await fetch(`/api/v1/locations/search?query=${query}`);
      if (!response.ok) {
        throw new Error("Failed to fetch locations");
      }
      const data = await response.json();
      setSearchResults(data);
    } catch (error) {
      console.error("Error fetching locations:", error);
      setSearchResults([]);
    }
  };

  const handleAddAttraction = () => {
    console.log("Form Data Submitted:", formData);
    setShowModal(false);
  };

  return (
    <div className="util-buttons">
      {isLoggedIn ? (
        <>
          <button onClick={() => setShowModal(true)}>
            <FaPlus style={{ marginRight: "0" }} />
          </button>
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
