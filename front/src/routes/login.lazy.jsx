import { createLazyFileRoute, Link } from "@tanstack/react-router";
import { useState } from "react";
import "../styles/Auth.css"; // Import the CSS file

export const Route = createLazyFileRoute("/login")({
  component: Login,
});

function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = (e) => {
    e.preventDefault();
    console.log("Login attempt with:", { email, password });
    // todo: login logic
  };

  return (
    <div className="auth-container">
      <h1 className="auth-title">Login</h1>
      <form onSubmit={handleLogin} className="auth-form">
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
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
          Login
        </button>
      </form>
      <p className="auth-redirect-text">
        Don't have an account?{" "}
        <Link to="/signup" className="auth-link">
          Sign Up
        </Link>
      </p>
    </div>
  );
}
