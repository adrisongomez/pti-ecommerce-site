import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute(
  "/categories/$categoryHandle/products/$productsId",
)({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Hello "/categories/$handle/products/$productsId"!</div>;
}
