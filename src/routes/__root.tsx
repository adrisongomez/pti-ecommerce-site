import MainLayout from "@/globals/components/layouts/MainLayout";
import { createRootRoute, HeadContent, Outlet } from "@tanstack/react-router";
import { lazy, Suspense } from "react";

const TanStackRouterDevtools =
  process.env.NODE_ENV === "production"
    ? () => null
    : lazy(() =>
        import("@tanstack/router-devtools").then((res) => ({
          default: res.TanStackRouterDevtools,
        })),
      );

export const Route = createRootRoute({
  head: () => ({
    meta: [{ title: "My Store Test" }],
  }),
  component: () => (
    <>
      <MainLayout>
        <HeadContent />
        <Outlet />
      </MainLayout>
      <Suspense>
        <TanStackRouterDevtools />
      </Suspense>
    </>
  ),
});
