import React, { useState, useEffect, useRef } from "react";
import "../styles/Modal.css";

const AddAttractionModal = ({ onClose }) => {
  const [formData, setFormData] = useState({
    name: "",
    description: "",
    funfact: "",
    photo: null,
    location: "",
    type: "Hotel",
  });
  const [searchResults, setSearchResults] = useState([]);
  const [showDropdown, setShowDropdown] = useState(false);
  const [isFormValid, setIsFormValid] = useState(false); // State to track form validity
  const inputBoxRef = useRef(null);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handlePhotoChange = async (e) => {
    const file = e.target.files[0];
    const fileName = file?.name || "No file chosen";

    if (file) {
      try {
        // Read the file as binary data
        const arrayBuffer = await file.arrayBuffer();
        const byteArray = new Uint8Array(arrayBuffer);

        // Convert binary data to hexadecimal representation
        const hexString = Array.from(byteArray)
          .map((byte) => byte.toString(16).padStart(2, "0"))
          .join("");

        setFormData((prev) => ({
          ...prev,
          photo: hexString, // Store the hex string
          photoName: fileName,
        }));
      } catch (error) {
        console.error("Error processing the photo file:", error);
      }
    }
  };

  const handleDropdownClick = (name) => {
    setFormData((prev) => ({ ...prev, location: name }));
    setShowDropdown(false);
    inputBoxRef.current.blur();
  };

  useEffect(() => {
    if (formData.location.trim() === "") {
      setSearchResults([]);
      return;
    }

    const fetchSearchResults = async () => {
      try {
        const response = await fetch(
          `/api/v1/locations/search?query=${formData.location}`,
        );
        if (!response.ok) {
          throw new Error("Failed to fetch search results");
        }
        const data = await response.json();
        setSearchResults(data);
      } catch (error) {
        console.error("Error fetching search results:", error);
        setSearchResults([]);
      }
    };

    const debounceTimeout = setTimeout(fetchSearchResults, 300);
    return () => clearTimeout(debounceTimeout);
  }, [formData.location]);

  useEffect(() => {
    // Validate the form
    const isLocationValid = searchResults.some(
      (result) => result.name === formData.location,
    );
    const isFormComplete =
      formData.name.trim() &&
      formData.description.trim() &&
      formData.funfact.trim() &&
      formData.photo &&
      isLocationValid;

    setIsFormValid(isFormComplete);
  }, [formData, searchResults]);

  const handleAddAttraction = async () => {
    if (!isFormValid) {
      console.log(
        "Form is not valid. Please fill all fields and select a valid location.",
      );
      return; // Prevent submission if the form is invalid
    }

    // Find the selected location's ID
    const selectedLocation = searchResults.find(
      (result) => result.name === formData.location,
    );

    if (!selectedLocation) {
      console.log("Invalid location selected.");
      return;
    }

    // Prepare the request body
    const attractionData = {
      name: formData.name,
      description: formData.description,
      funfact: formData.funfact,
      photo: formData.photo, // Base64 string
      locationId: selectedLocation.id, // Use the location ID from the dropdown
      type: formData.type.toLowerCase(),
    };

    try {
      const token = localStorage.getItem("token"); // Retrieve the token from localStorage
      if (!token) {
        console.error("Authorization token is missing.");
        return;
      }

      const response = await fetch("/api/v1/attractions", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(attractionData),
      });

      if (!response.ok) {
        throw new Error("Failed to submit attraction data.");
      }

      const result = await response.json();
      console.log("Attraction successfully added:", result);

      // Close the modal after successful submission
      onClose();
    } catch (error) {
      console.error("Error submitting attraction data:", error);
    }
  };

  return (
    <div className="modal">
      <div className="modal-content">
        <h2>Add Attraction</h2>
        <form onSubmit={(e) => e.preventDefault()}>
          <label>
            Name:
            <input
              type="text"
              name="name"
              value={formData.name}
              onChange={handleInputChange}
              required
            />
          </label>
          <label>
            Description:
            <textarea
              name="description"
              value={formData.description}
              onChange={handleInputChange}
              required
            />
          </label>
          <label>
            Fun Fact:
            <textarea
              name="funfact"
              value={formData.funfact}
              onChange={handleInputChange}
              required
            />
          </label>
          <label>
            Photo:
            <div
              className="file-input-label"
              onClick={() => document.getElementById("file-input").click()}
            >
              <p className="choose-file">Choose file</p>
              <p className="file-name">
                {formData.photoName || "No file chosen"}
              </p>
            </div>
          </label>
          <input
            id="file-input"
            type="file"
            name="photo"
            accept="image/*"
            onChange={handlePhotoChange}
          />
          <label>
            Location:
            <input
              ref={inputBoxRef}
              type="text"
              name="location"
              placeholder="Search location..."
              value={formData.location}
              onChange={(e) => {
                handleInputChange(e);
                setShowDropdown(true);
              }}
              onFocus={() => setShowDropdown(true)}
              onBlur={() => setTimeout(() => setShowDropdown(false), 200)}
              autoComplete="off"
              required
            />
            {showDropdown && searchResults.length > 0 && (
              <div className="modal-dropdown-menu">
                {searchResults.map((result) => (
                  <div
                    key={result.id}
                    className="modal-dropdown-item"
                    onClick={() => handleDropdownClick(result.name)}
                  >
                    {result.name}, {result.type}
                  </div>
                ))}
              </div>
            )}
          </label>
          <label>
            Type:
            <select
              name="type"
              value={formData.type}
              onChange={handleInputChange}
            >
              <option value="Hotel">Hotel</option>
              <option value="Restaurant">Restaurant</option>
              <option value="Other">Other</option>
            </select>
          </label>
          <button type="button" onClick={handleAddAttraction}>
            Submit
          </button>
          <button type="button" onClick={onClose}>
            Cancel
          </button>
        </form>
      </div>
    </div>
  );
};

export default AddAttractionModal;
