import { useState } from "react";
import { createRootRoute, Outlet } from "@tanstack/react-router";

export const Route = createRootRoute({
  component: () => {
    return (
      <>
        <div>
          <Outlet />
        </div>
      </>
    );
  },
});
