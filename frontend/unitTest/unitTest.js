import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import { BrowserRouter } from "react-router-dom";
import Login from "../src/components/Login/Login";
import Register from "../src/components/Login/Register";

describe("Login Component", () => {
  test("renders login form with username, password, and submit button", () => {
    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    );

    expect(screen.getByPlaceholderText("Username")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Password")).toBeInTheDocument();
    expect(screen.getByText("Login")).toBeInTheDocument();
  });

  test("displays error message on failed login attempt", async () => {
    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    );

    fireEvent.change(screen.getByPlaceholderText("Username"), {
      target: { value: "wrongUser" },
    });
    fireEvent.change(screen.getByPlaceholderText("Password"), {
      target: { value: "wrongPass" },
    });
    fireEvent.click(screen.getByText("Login"));

    const errorMessage = await screen.findByText(
      "Login failed. Please try again."
    );
    expect(errorMessage).toBeInTheDocument();
  });

  test("navigates to register page when 'Register here' is clicked", () => {
    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    );

    const registerLink = screen.getByText("Register here");
    expect(registerLink).toBeInTheDocument();

    fireEvent.click(registerLink);
    expect(window.location.pathname).toBe("/register");
  });
});

describe("Register Component", () => {
  test("renders register form with username, email, password, and submit button", () => {
    render(
      <BrowserRouter>
        <Register />
      </BrowserRouter>
    );

    expect(screen.getByPlaceholderText("Username")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Email")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Password")).toBeInTheDocument();
    expect(screen.getByText("Register")).toBeInTheDocument();
  });

  test("shows error message if any field is left empty", () => {
    render(
      <BrowserRouter>
        <Register />
      </BrowserRouter>
    );

    fireEvent.click(screen.getByText("Register"));
    const errorMessage = screen.getByText("All fields are required!");
    expect(errorMessage).toBeInTheDocument();
  });

  test("shows success message and navigates to login on successful registration", async () => {
    render(
      <BrowserRouter>
        <Register />
      </BrowserRouter>
    );

    fireEvent.change(screen.getByPlaceholderText("Username"), {
      target: { value: "testUser" },
    });
    fireEvent.change(screen.getByPlaceholderText("Email"), {
      target: { value: "test@example.com" },
    });
    fireEvent.change(screen.getByPlaceholderText("Password"), {
      target: { value: "password123" },
    });
    fireEvent.click(screen.getByText("Register"));

    const successMessage = await screen.findByText("Registration successful!");
    expect(successMessage).toBeInTheDocument();
    expect(window.location.pathname).toBe("/login");
  });
});
