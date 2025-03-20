import MainLayout from "@/libs/globals/components/layouts/MainLayout";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/products/$productId")({
  component: RouteComponent,
});

function RouteComponent() {
  const data = Route.useParams();
  return <MainLayout>Hello "/products/{data.productId}"!</MainLayout>;
}
