import { createLazyFileRoute, Link, useNavigate } from "@tanstack/react-router";
import { useState } from "react";
import "../styles/Auth.css";
import { signupUser } from "../api/signupApi";

export const Route = createLazyFileRoute("/signup")({
  component: SignUp,
});

function SignUp() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [displayName, setDisplayName] = useState("");
  const [error, setError] = useState("");
  const [successMessage, setSuccessMessage] = useState("");
  const navigate = useNavigate();

  const handleSignUp = async (e) => {
    e.preventDefault();
    setError("");
    setSuccessMessage("");
    try {
      const data = await signupUser(username, password, displayName);
      console.log("Sign-up successful:", data);

      setSuccessMessage(
        "Account created successfully! Redirecting to login...",
      );
      setTimeout(() => {
        navigate({ to: "/login" });
      }, 3000);
    } catch (err) {
      console.error("Sign-up error:", err.message);
      setError(err.message);
    }
  };

  return (
    <div className="auth-container">
      <h1 className="auth-title">Sign Up</h1>
      <form onSubmit={handleSignUp} className="auth-form">
        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          className="auth-input"
          required
        />
        <input
          type="text"
          placeholder="Display Name"
          value={displayName}
          onChange={(e) => setDisplayName(e.target.value)}
          className="auth-input"
          required
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="auth-input"
          required
        />
        <button type="submit" className="auth-button">
          Sign Up
        </button>
      </form>
      {error && <p className="auth-error">{error}</p>}
      {successMessage && <p className="auth-success">{successMessage}</p>}
      <p className="auth-redirect-text">
        Already have an account?{" "}
        <Link to="/login" className="auth-link">
          Login
        </Link>
      </p>
    </div>
  );
}
