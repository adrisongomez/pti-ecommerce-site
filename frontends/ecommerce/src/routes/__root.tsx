import { useAppDispatch } from "@/libs/globals/hooks/redux";
import { fetchUser } from "@/libs/globals/redux/AuthReducer";
import { getCreds } from "@/libs/globals/utilities/auth";
import { QueryClient } from "@tanstack/react-query";
import {
  createRootRouteWithContext,
  HeadContent,
  Outlet,
} from "@tanstack/react-router";
import { FC, lazy, Suspense } from "react";
import { useMount } from "react-use";

const ReactQueryDevtools =
  process.env.NODE_ENV === "production"
    ? () => null
    : lazy(() =>
        import("@tanstack/react-query-devtools").then((res) => ({
          default: res.ReactQueryDevtools,
        })),
      );
const TanStackRouterDevtools =
  process.env.NODE_ENV === "production"
    ? () => null
    : lazy(() =>
        import("@tanstack/router-devtools").then((res) => ({
          default: res.TanStackRouterDevtools,
        })),
      );

const RootComponent: FC = () => {
  const dispatch = useAppDispatch();
  useMount(() => {
    const creds = getCreds();
    if (creds) {
      dispatch(fetchUser());
    }
  });
  return (
    <>
      <HeadContent />
      <Outlet />
      <Suspense>
        <TanStackRouterDevtools />
        <ReactQueryDevtools />
      </Suspense>
    </>
  );
};
export const Route = createRootRouteWithContext<{
  queryClient: QueryClient;
}>()({
  head: () => ({
    meta: [{ title: "My Store Test" }],
  }),
  component: RootComponent,
});
