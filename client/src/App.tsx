import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { Login } from "./pages/login";
import Home from "./pages/home";
import { ProtectedRoute } from "./auth/protected-route";

const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <>
        <ProtectedRoute>
          <Home />
        </ProtectedRoute>
      </>
    ),
  },
  {
    path: "/login",
    element: <Login />,
  },
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
