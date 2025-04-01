import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/profile/")({
  component: RouteComponent,
});
import { useMount } from "react-use";

function RouteComponent() {
  const nav = Route.useNavigate();
  useMount(() => {
    nav({ to: "/profile/$tabValue", params: { tabValue: "general" } });
  });
  return <div />;
}
