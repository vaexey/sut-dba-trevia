import { createLazyFileRoute, Link } from "@tanstack/react-router";
import { useState } from "react";
import "../styles/Auth.css"; // Import the CSS file

export const Route = createLazyFileRoute("/signup")({
  component: SignUp,
});

function SignUp() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSignUp = (e) => {
    e.preventDefault();
    console.log("Sign-up attempt with:", { email, password });
    // todo: sign-up logic
  };

  return (
    <div className="auth-container">
      <h1 className="auth-title">Sign Up</h1>
      <form onSubmit={handleSignUp} className="auth-form">
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
          Sign Up
        </button>
      </form>
      <p className="auth-redirect-text">
        Already have an account?{" "}
        <Link to="/login" className="auth-link">
          Login
        </Link>
      </p>
    </div>
  );
}
