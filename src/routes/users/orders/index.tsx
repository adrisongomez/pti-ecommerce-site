import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/users/orders/")({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Hello "/users/orders/"!</div>;
}
