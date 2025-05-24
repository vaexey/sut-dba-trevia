import { createLazyFileRoute, useRouterState } from "@tanstack/react-router";
import React, { useEffect, useState } from "react";
import ReportAttractionModal from "../components/ReportModal";
import UserComment from "../components/UserComment";
import "../styles/Attraction.css";

export const Route = createLazyFileRoute("/attraction")({
  component: AttractionPage,
});

function AttractionPage() {
  const search = useRouterState({ select: (s) => s.location.search });
  const params = new URLSearchParams(search);
  const id = params.get("id"); // Extract the "id" query parameter

  const [attractionData, setAttractionData] = useState(null);
  const [comments, setComments] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [showReportModal, setShowReportModal] = useState(false);
  const [showCommentReportModal, setShowCommentReportModal] = useState(false);
  const [reportedComment, setReportedComment] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem("token");
    setIsLoggedIn(!!token);

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
    fetchComments();
  }, [id]);

  const fetchComments = async () => {
    try {
      const token = localStorage.getItem("token");
      if (!token) {
        throw new Error("Authorization token is missing.");
      }

      const response = await fetch(`/api/v1/comments/${id}`, {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      const data = await response.json();
      console.log(data);
      setComments(data || []); // Store the fetched comments
    } catch (err) {
      console.error("Error fetching comments:", err);
    }
  };

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  const handleReport = () => {
    setShowReportModal(true);
  };

  const handleReportSubmit = (reportText) => {
    console.log("Report submitted for attraction:", attractionData.id);
    console.log("Report text:", reportText);

    const token = localStorage.getItem("token");
    const reportData = {
      AttractionId: attractionData.id,
      Content: reportText,
    };

    fetch("/api/v1/reports/attractions", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(reportData),
    })
      .then((res) => res.json())
      .then((data) => console.log(data))
      .catch((err) => console.error(err));
  };

  const handleCommentReport = (comment) => {
    setReportedComment(comment);
    setShowCommentReportModal(true);
  };

  const handleCommentReportSubmit = (reportText) => {
    console.log("Report submitted for comment:", reportedComment.id);
    console.log("Report text:", reportText);

    const token = localStorage.getItem("token");
    const reportData = {
      CommentId: reportedComment.id,
      Content: reportText,
    };

    fetch("/api/v1/reports/comments", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(reportData),
    })
      .then((res) => res.json())
      .then((data) => console.log(data))
      .catch((err) => console.error(err))
      .finally(() => {
        setShowCommentReportModal(false);
        setReportedComment(null);
      });
  };

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
        {isLoggedIn && (
          <button className="report-button" onClick={handleReport}>
            !
          </button>
        )}
      </div>
      <div className="comments-section">
        <h2>Comments</h2>
        {isLoggedIn ? (
          <>
            <UserComment
              onSubmit={fetchComments}
              attractionId={attractionData.id}
            />
            {comments.length > 0 ? (
              comments.map((comment, index) => (
                <div key={index} className="comment">
                  <div className="comment-header">
                    <span className="comment-author">{comment.username}</span>
                    <button
                      className="comment-report-button"
                      title="Report this comment"
                      onClick={() => handleCommentReport(comment)}
                    >
                      !
                    </button>
                  </div>
                  <p className="comment-text">{comment.comment}</p>
                </div>
              ))
            ) : (
              <p>No comments yet. Be the first to comment!</p>
            )}
          </>
        ) : (
          <p>Please log in to write and see user comments.</p>
        )}
      </div>
      {showReportModal && (
        <ReportAttractionModal
          onClose={() => setShowReportModal(false)}
          onSubmit={handleReportSubmit}
        />
      )}
      {showCommentReportModal && reportedComment && (
        <ReportAttractionModal
          onClose={() => setShowCommentReportModal(false)}
          onSubmit={handleCommentReportSubmit}
        />
      )}
    </div>
  );
}

export default AttractionPage;
