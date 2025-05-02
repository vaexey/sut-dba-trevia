import { createRootRoute, Outlet, useLocation } from "@tanstack/react-router";
import UtilButtons from "../components/UtilButtons";
import Navbar from "../components/Navbar";

export const Route = createRootRoute({
  component: () => {
    const location = useLocation();
    const isLoginPage = location.pathname === "/login";
    const hideNavbarPages = ["/", "/login"];
    const isNavbarHidden = hideNavbarPages.includes(location.pathname);

    return (
      <>
        {!isNavbarHidden && <Navbar />}
        {!isLoginPage && <UtilButtons />}
        <div>
          <Outlet />
        </div>
      </>
    );
  },
});
