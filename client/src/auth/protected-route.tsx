import { ReactNode } from "react";
import { useAuthStore } from ".";
import { Navigate } from "react-router-dom";
import { useStore } from "zustand";

export const ProtectedRoute = ({ children }: { children: ReactNode }) => {
  const store = useAuthStore();
  const token = useStore(store, (state) => state.token);

  // const { isLoading: isLoadingUser } = useQuery(
  //   "getRefreshToken",
  //   refreshToken,

  //   {
  //     enabled: !!token,
  //     onSuccess: (data) => {
  //       login?.(data?.accessToken ?? ""); //set refresh token
  //     },
  //     onError: (data) => {
  //       console.log(data);
  //       logout?.();
  //     },
  //     retry: 0,
  //     refetchOnWindowFocus: false,
  //     refetchInterval: 13 * 60 * 1000, // 13 minutes
  //   }
  // );

  // if (isLoadingUser) {
  //   return <Loader />;
  // }

  if (!token) {
    // user is not authenticated
    return <Navigate to="/login" />;
  }
  return children;
};
