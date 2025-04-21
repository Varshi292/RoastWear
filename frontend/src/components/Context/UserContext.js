import { createContext, useContext, useState, useEffect } from "react";

const UserContext = createContext();

export const UserProvider = ({ children }) => {
  const [userName, setUserName] = useState("");

  useEffect(() => {
    const stored = localStorage.getItem("userName");
    if (stored) {
      setUserName(stored);
    }
  }, []);

  const clearUser = () => {
    setUserName("");
    localStorage.removeItem("userName");
  };

  return (
    <UserContext.Provider value={{ userName, setUserName, clearUser }}>
      {children}
    </UserContext.Provider>
  );
};

export const useUser = () => useContext(UserContext);
