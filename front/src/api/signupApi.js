export async function signupUser(username, password, displayName) {
  const response = await fetch("/api/v1/sign-up", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, password, displayName }),
  });

  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Sign-up failed");
  }

  return response.json();
}
