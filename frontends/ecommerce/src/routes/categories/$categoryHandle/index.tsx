import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/categories/$categoryHandle/")({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Hello "/categories/$categoryHandle/"!</div>;
}
