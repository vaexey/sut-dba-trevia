import { createLazyFileRoute } from "@tanstack/react-router";
import { useState, useEffect } from "react";
import { useNavigate } from "@tanstack/react-router";
import "../styles/AdminPage.css";

export const Route = createLazyFileRoute("/admin")({
  component: AdminPage,
});

function AdminPage() {
  const [commentReports, setCommentReports] = useState([]);
  const [attractionReports, setAttractionReports] = useState([]);
  const [loading, setLoading] = useState(true);
  const [roleId, setRoleId] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    setLoading(true);

    fetch("/api/v1/user", {
      headers: { Authorization: `Bearer ${token}` },
    })
      .then((res) => res.json())
      .then((data) => {
        setRoleId(data.roleId);
        if (data.roleId !== 3) {
          navigate({ to: "/" });
        }
      })
      .catch(() => navigate({ to: "/" }));

    Promise.all([
      fetch("/api/v1/reports/comments", {
        headers: { Authorization: `Bearer ${token}` },
      }).then((res) => res.json()),
      fetch("/api/v1/reports/attractions", {
        headers: { Authorization: `Bearer ${token}` },
      }).then((res) => res.json()),
    ])
      .then(([comments, attractions]) => {
        setCommentReports(comments);
        setAttractionReports(attractions);
      })
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div>Loading reports...</div>;

  return (
    <div className="admin-panel">
      <h1>Admin Panel</h1>
      <div className="admin-reports-columns">
        <div className="admin-report-column">
          <h2>Comment Reports</h2>
          {commentReports.length === 0 ? (
            <p>No comment reports.</p>
          ) : (
            <ul>
              {commentReports.map((report) => (
                <li key={report.id || report.Id}>
                  <strong>Comment ID:</strong>{" "}
                  {report.commentId || report.CommentId} <br />
                  <strong>Content:</strong> {report.content || report.Content}{" "}
                  <br />
                  <button
                    onClick={() =>
                      navigate({
                        to: `/attraction?id=${report.attractionId || report.AttractionId}`,
                      })
                    }
                  >
                    Go to Attraction
                  </button>
                </li>
              ))}
            </ul>
          )}
        </div>
        <div className="admin-report-column">
          <h2>Attraction Reports</h2>
          {attractionReports.length === 0 ? (
            <p>No attraction reports.</p>
          ) : (
            <ul>
              {attractionReports.map((report) => (
                <li key={report.id || report.Id}>
                  <strong>Attraction ID:</strong>{" "}
                  {report.attractionId || report.AttractionId} <br />
                  <strong>Content:</strong> {report.content || report.Content}{" "}
                  <br />
                  <button
                    onClick={() =>
                      navigate({
                        to: `/attraction?id=${report.attractionId || report.AttractionId}`,
                      })
                    }
                  >
                    Go to Attraction
                  </button>
                </li>
              ))}
            </ul>
          )}
        </div>
      </div>
    </div>
  );
}

export default AdminPage;
