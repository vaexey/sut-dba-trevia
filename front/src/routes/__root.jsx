import { createRootRoute, Outlet, useLocation } from "@tanstack/react-router";
import UtilButtons from "../components/UtilButtons";
import Navbar from "../components/Navbar";
import PageLogo from "../components/PageLogo";
import TextLogo from "../components/TextLogo";

export const Route = createRootRoute({
  component: () => {
    const location = useLocation();
    const hideNavbarPages = ["/", "/login", "/signup"];
    const hideUtilButtonsPages = ["/login", "/signup"];
    const isNavbarHidden = hideNavbarPages.includes(location.pathname);
    const areUtilButtonsHidden = hideUtilButtonsPages.includes(
      location.pathname,
    );

    return (
      <>
        {!isNavbarHidden && <Navbar />}
        {!areUtilButtonsHidden && <UtilButtons />}
        {isNavbarHidden && <TextLogo />}
        <PageLogo />
        <div
          className="outlet-wrapper"
          style={{
            height: isNavbarHidden
              ? "calc(100vh - 120px)"
              : "calc(100vh - 180px)",
          }}
        >
          <Outlet />
        </div>
      </>
    );
  },
});
