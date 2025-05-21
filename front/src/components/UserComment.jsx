import React, { useState } from "react";
import "../styles/UserComment.css";

const UserComment = ({ onSubmit, attractionId }) => {
  const [commentText, setCommentText] = useState("");

  const handleSubmit = async () => {
    if (commentText.trim() === "") {
      console.log("Comment cannot be empty.");
      return;
    }

    const commentData = {
      attractionId, // Pass the attraction ID as a prop
      comment: commentText,
    };

    try {
      const token = localStorage.getItem("token"); // Retrieve the token from localStorage
      if (!token) {
        throw new Error("Authorization token is missing.");
      }

      const response = await fetch("/api/v1/comments", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`, // Include the Bearer token
        },
        body: JSON.stringify(commentData),
      });

      if (!response.ok) {
        throw new Error("Failed to submit comment.");
      }

      const result = await response.json();
      console.log("Comment successfully submitted:", result);

      // Clear the input after successful submission
      setCommentText("");

      // Call the onSubmit callback to refresh comments or handle UI updates
      onSubmit(commentText);
    } catch (error) {
      console.error("Error submitting comment:", error);
    }
  };

  return (
    <div className="user-comment">
      <textarea
        placeholder="Write your comment here..."
        value={commentText}
        onChange={(e) => setCommentText(e.target.value)}
        rows="4"
        className="comment-textbox"
      />
      <button onClick={handleSubmit} className="comment-submit-button">
        Submit
      </button>
    </div>
  );
};

export default UserComment;
