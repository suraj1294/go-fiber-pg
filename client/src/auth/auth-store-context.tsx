import { AuthStore, authStore } from "@/auth";
import { ReactNode, createContext } from "react";

export const AuthStoreContext = createContext<AuthStore>(authStore);

export const AuthStoreProvider = ({ children }: { children: ReactNode }) => {
  return (
    <AuthStoreContext.Provider value={authStore}>
      {children}
    </AuthStoreContext.Provider>
  );
};
