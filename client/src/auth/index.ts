import { useContext } from "react";
import { AuthStoreContext, AuthStoreProvider } from "./auth-store-context";

export const useAuthStore = () => {
  const context = useContext(AuthStoreContext);

  if (context) {
    return context;
  }

  throw new Error(`useAuthStore must be used within a AuthStoreProvider`);
};

export { AuthStoreContext, AuthStoreProvider };
export {
  type AuthState,
  type AuthStore,
  getAuthState,
  setAuthState,
  subscribeAuthState,
  authStore,
} from "./auth-store";
export { ProtectedRoute } from "./protected-route";
