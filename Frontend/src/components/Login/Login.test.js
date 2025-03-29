import React from "react";
import {
  render,
  screen,
  fireEvent,
  waitFor,
} from "@testing-library/react";
import { MemoryRouter, Route, Routes } from "react-router-dom";
import Login from "./Login";
import Register from "./Register";

// 🔧 Utility to render with routing support
const renderWithRouter = (ui, { route = "/" } = {}) => {
  return render(
    <MemoryRouter initialEntries={[route]}>
      <Routes>
        <Route path="/" element={ui} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
      </Routes>
    </MemoryRouter>
  );
};

describe("Login Component", () => {
  beforeEach(() => {
    // Reset fetch before every test
    global.fetch = jest.fn();
  });

  test("renders login form with username, password, and submit button", () => {
    renderWithRouter(<Login />);
    expect(screen.getByPlaceholderText("Username")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Password")).toBeInTheDocument();
    expect(
      screen.getByRole("button", { name: /login/i })
    ).toBeInTheDocument();
  });

  test("displays error message on failed login attempt", async () => {
    // 🧪 Mock failed login
    global.fetch.mockResolvedValueOnce({
      json: async () => ({
        success: false,
        message: "Login failed. Please try again.",
      }),
    });

    renderWithRouter(<Login />);

    fireEvent.change(screen.getByPlaceholderText("Username"), {
      target: { value: "wrongUser" },
    });
    fireEvent.change(screen.getByPlaceholderText("Password"), {
      target: { value: "wrongPass" },
    });
    fireEvent.click(screen.getByRole("button", { name: /login/i }));

    const errorMessage = await screen.findByText(
      "Login failed. Please try again."
    );
    expect(errorMessage).toBeInTheDocument();
  });

  test("navigates to register page when 'Register here' is clicked", async () => {
    renderWithRouter(<Login />);
    const registerLink = screen.getByText("Register here");

    fireEvent.click(registerLink);

    await waitFor(() =>
      expect(
        screen.getByRole("button", { name: /register/i })
      ).toBeInTheDocument()
    );
  });
});

describe("Register Component", () => {
  beforeEach(() => {
    global.fetch = jest.fn();
  });

  test("renders register form with username, email, password, and submit button", () => {
    renderWithRouter(<Register />);
    expect(screen.getByPlaceholderText("Username")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Email")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Password")).toBeInTheDocument();
    expect(
      screen.getByRole("button", { name: /register/i })
    ).toBeInTheDocument();
  });

  test("shows error message if any field is left empty", () => {
    renderWithRouter(<Register />);
    fireEvent.click(screen.getByRole("button", { name: /register/i }));
    const errorMessage = screen.getByText("All fields are required!");
    expect(errorMessage).toBeInTheDocument();
  });

  test("shows success message and navigates to login on successful registration", async () => {
    // 🧪 Mock success response
    global.fetch.mockResolvedValueOnce({
      json: async () => ({
        success: true,
        message: "Registration successful!",
      }),
    });

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
    fireEvent.click(screen.getByRole("button", { name: /register/i }));

    const successMessage = await screen.findByText("Registration successful!");
    expect(successMessage).toBeInTheDocument();
  });
});
