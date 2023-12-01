import {
  ReactNode,
  createContext,
  useCallback,
  useContext,
  useMemo,
  useReducer,
} from "react";

interface AppState {
  token: string;
  login?: (token: string) => void;
  logout?: () => void;
}

const AuthContext = createContext<AppState>({ token: "" });

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [user, setUser] = useReducer(
    (state: AppState, next: AppState) => ({
      ...state,
      ...next,
    }),
    { token: "" }
  );

  // call this function when you want to authenticate the user
  const login = useCallback((token: string) => {
    setUser({ token });
  }, []);

  // call this function to sign out logged in user
  const logout = useCallback(() => {
    setUser({ token: "" });
  }, []);

  const value = useMemo(
    () => ({
      user,
      login,
      logout,
    }),
    [login, logout, user]
  );
  return (
    <AuthContext.Provider value={{ token: value.user.token, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);
