import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/users/orders/$orderId/")({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Hello "/users/orders/$orderId/"!</div>;
}
