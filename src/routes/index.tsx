import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  head: () => ({
    meta: [{ title: "Home Page" }],
  }),
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Hello "/"!</div>;
}
