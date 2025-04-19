import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import { MemoryRouter, Route, Routes } from "react-router-dom";
import Login from "../Login"; // Adjust if necessary
import Register from "../Register"; // Adjust if necessary

// Mock the router for navigation
const renderWithRouter = (ui, { route = "/" } = {}) => {
  return render(
    <MemoryRouter initialEntries={[route]}>
      <Routes>
        <Route path="/" element={ui} />
        <Route path="/register" element={<Register />} />
      </Routes>
    </MemoryRouter>
  );
};

describe("Login Component", () => {
  test("renders login form with username, password, and submit button", () => {
    renderWithRouter(<Login />);

    expect(screen.getByPlaceholderText("Username")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Password")).toBeInTheDocument();
    expect(screen.getByText("Login")).toBeInTheDocument();
  });

  test("displays error message on failed login attempt", async () => {
    renderWithRouter(<Login />);

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
    renderWithRouter(<Login />);

    const registerLink = screen.getByText("Register here");
    expect(registerLink).toBeInTheDocument();

    fireEvent.click(registerLink);
    expect(window.location.pathname).toBe("/register");
  });
});

describe("Register Component", () => {
  test("renders register form with username, email, password, and submit button", () => {
    renderWithRouter(<Register />);

    expect(screen.getByPlaceholderText("Username")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Email")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Password")).toBeInTheDocument();
    expect(screen.getByText("Register")).toBeInTheDocument();
  });

  test("shows error message if any field is left empty", () => {
    renderWithRouter(<Register />);

    fireEvent.click(screen.getByText("Register"));
    const errorMessage = screen.getByText("All fields are required!");
    expect(errorMessage).toBeInTheDocument();
  });

  test("shows success message and navigates to login on successful registration", async () => {
    renderWithRouter(<Register />);

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
