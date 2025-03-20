import Breadcrumps from "@/libs/globals/components/breadcrumps/Breadcrumps";
import MainLayout from "@/libs/globals/components/layouts/MainLayout";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute(
  "/categories/$categoryHandle/products/$productId",
)({
  component: RouteComponent,
});

function RouteComponent() {
  // const params = Route.useParams();
  return (
    <MainLayout>
      <Breadcrumps
        crumps={[
          { label: "Home", to: "/" },
          { label: "Categories", to: "/categories" },
          { label: "Desk lamp", to: "/categories/$categoryHandle" },
        ]}
        currentCrump={"Product Name"}
      />
    </MainLayout>
  );
}
