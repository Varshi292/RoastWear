// src/components/Login/Register.test.js
import React from "react";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import Register from "./Register";
import { MemoryRouter } from "react-router-dom";

const mockNavigate = jest.fn();

jest.mock("react-router-dom", () => ({
  ...jest.requireActual("react-router-dom"),
  useNavigate: () => mockNavigate,
}));

describe("Register Component", () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  test("renders all form fields and register button", () => {
    render(
      <MemoryRouter>
        <Register />
      </MemoryRouter>
    );

    expect(screen.getByPlaceholderText("Username")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Email")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Password")).toBeInTheDocument();
    expect(
      screen.getByRole("button", { name: "Register" })
    ).toBeInTheDocument();
  });

  test("shows error if fields are empty on submit", () => {
    render(
      <MemoryRouter>
        <Register />
      </MemoryRouter>
    );

    fireEvent.click(screen.getByRole("button", { name: "Register" }));
    expect(
      screen.getByText("All fields are required!")
    ).toBeInTheDocument();
  });

  test("successful registration navigates to login", async () => {
    jest.useFakeTimers(); // control timeout
    global.fetch = jest.fn(() =>
      Promise.resolve({
        json: () =>
          Promise.resolve({
            success: true,
            message: "Registration successful!",
          }),
      })
    );

    render(
      <MemoryRouter>
        <Register />
      </MemoryRouter>
    );

    fireEvent.change(screen.getByPlaceholderText("Username"), {
      target: { value: "testuser" },
    });
    fireEvent.change(screen.getByPlaceholderText("Email"), {
      target: { value: "test@example.com" },
    });
    fireEvent.change(screen.getByPlaceholderText("Password"), {
      target: { value: "password123" },
    });

    fireEvent.click(screen.getByRole("button", { name: "Register" }));

    // Message check
    await screen.findByText("Registration successful!");

    // Fast forward time
    jest.advanceTimersByTime(2000);

    await waitFor(() => {
      expect(mockNavigate).toHaveBeenCalledWith("/login");
    });

    jest.useRealTimers();
  });

  test("shows error on failed API call", async () => {
    global.fetch = jest.fn(() => {
      throw new Error("API down");
    });

    render(
      <MemoryRouter>
        <Register />
      </MemoryRouter>
    );

    fireEvent.change(screen.getByPlaceholderText("Username"), {
      target: { value: "testuser" },
    });
    fireEvent.change(screen.getByPlaceholderText("Email"), {
      target: { value: "test@example.com" },
    });
    fireEvent.change(screen.getByPlaceholderText("Password"), {
      target: { value: "password123" },
    });

    fireEvent.click(screen.getByRole("button", { name: "Register" }));

    await screen.findByText("Registration failed. Please try again.");
  });
});
