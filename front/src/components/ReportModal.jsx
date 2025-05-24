import React, { useState } from "react";
import "../styles/Modal.css";

const ReportAttractionModal = ({ onClose, onSubmit }) => {
  const [reportText, setReportText] = useState("");

  const handleSubmit = () => {
    if (reportText.trim() === "") {
      console.log("Report text cannot be empty.");
      return;
    }
    onSubmit(reportText);
    onClose();
  };

  return (
    <div className="modal">
      <div className="modal-content">
        <h2>Report</h2>
        <textarea
          placeholder="Enter your report here..."
          className="report-textbox"
          value={reportText}
          onChange={(e) => setReportText(e.target.value)}
          rows="10"
          style={{ width: "100%", marginBottom: "1rem" }}
        />
        <div className="modal-actions">
          <button onClick={handleSubmit} className="report-modal-button">
            Submit
          </button>
          <button onClick={onClose} className="report-modal-button">
            Cancel
          </button>
        </div>
      </div>
    </div>
  );
};

export default ReportAttractionModal;
