import Card from "@/libs/globals/components/cards/Card";
import MainLayout from "@/libs/globals/components/layouts/MainLayout";
import LoadingIndicator from "@/libs/globals/components/progress/LoadingIndicator";
import Tab from "@/libs/globals/components/tabs/Tab";
import ProfileGeneralTab from "@/libs/routes/profile/ProfileGeneralTab";
import ProfileOrderTab from "@/libs/routes/profile/ProfileOrderTab";
import { createFileRoute } from "@tanstack/react-router";
import { Suspense, useState } from "react";
import { Archive, Truck, User } from "react-feather";

export const Route = createFileRoute("/profile/$tabValue")({
  component: ProfilePage,
  head: () => ({
    meta: [{ title: "Ecommerce | Checkout" }],
  }),
});

type ProfilePageType = "general" | "address" | "orders";

function ProfilePage() {
  const params = Route.useParams();
  const nav = Route.useNavigate();
  const [currentTab, setCurrentTab] = useState<ProfilePageType>(
    mapToTabValue(params.tabValue),
  );

  return (
    <MainLayout>
      <Tab<ProfilePageType>
        options={[
          { value: "general", label: "General", icon: <User size={16} /> },
          { label: "Orders", value: "orders", icon: <Archive size={16} /> },
          { label: "Address", value: "address", icon: <Truck size={16} /> },
        ]}
        value={currentTab}
        variant="vertical"
        onChange={(_, v) => {
          setCurrentTab(v);
          nav({ to: "/profile/$tabValue", params: { tabValue: v } });
        }}
      />
      <div className="ml-9 h-40 flex-1">
        <Card className="max-w-5xl bg-(--bg-light) p-6">
          <Suspense
            fallback={
              <div className="flex items-center justify-center">
                <LoadingIndicator />
              </div>
            }
          >
            {currentTab === "general" && <ProfileGeneralTab />}
            {currentTab === "orders" && <ProfileOrderTab />}
          </Suspense>
        </Card>
      </div>
    </MainLayout>
  );
}

function mapToTabValue(s: string): ProfilePageType {
  switch (s) {
    case "orders": {
      return "orders";
    }
    case "address": {
      return "address";
    }
    case "general":
    default: {
      return "general";
    }
  }
}
