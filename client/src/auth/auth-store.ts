import { createStore } from "zustand/vanilla";

export interface AuthState {
  token: string;
  setToken: (token: string) => void;
}

export const authStore = createStore<AuthState>((set) => ({
  token: "",
  setToken: (token: string) => set((state) => ({ ...state, token })),
}));
const { getState, setState, subscribe } = authStore;

export {
  getState as getAuthState,
  setState as setAuthState,
  subscribe as subscribeAuthState,
};

export type AuthStore = typeof authStore;

// export const useAuthStore = create(
//   subscribeWithSelector(() => ({ token: true }))
// );

//export const useAuthStore = (selector) => useStore(authStore, selector)
