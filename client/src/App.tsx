import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { Login } from "./pages/login";
import Home from "./pages/home";
import { ProtectedRoute } from "./auth";
import { useQuery } from "react-query";
import { me } from "./services/auth";
import { Loader } from "lucide-react";
import { authStore } from "./auth";
import { useStore } from "zustand";

const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <ProtectedRoute>
        <Home />
      </ProtectedRoute>
    ),
  },
  {
    path: "/login",
    element: <Login />,
  },
]);

function App() {
  const setToken = useStore(authStore, (state) => state.setToken);

  const { isLoading: isLoadingUser } = useQuery("getUser", me, {
    onSuccess: (data) => {
      setToken(data?.accessToken ?? "");
    },
    retry: 0,
  });

  if (isLoadingUser) {
    return <Loader />;
  }

  return <RouterProvider router={router} />;
}

export default App;
