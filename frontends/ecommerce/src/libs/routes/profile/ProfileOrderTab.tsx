import Button from "@/libs/globals/components/buttons/Button";
import Table from "@/libs/globals/components/tables/Table";
import { useOrderServiceDeleteApiOrdersByOrderId } from "@/libs/globals/generated/queries";
import { useOrderServiceGetApiOrdersSuspense } from "@/libs/globals/generated/queries/suspense";
import { ReactNode } from "@tanstack/react-router";
import { format } from "date-fns";
import { FC } from "react";

const ProfileOrderTab: FC = () => {
  const { data, refetch } = useOrderServiceGetApiOrdersSuspense({}, [
    "ProfileOrderTab_Order_Query",
  ]);
  const cancel = useOrderServiceDeleteApiOrdersByOrderId({
    mutationKey: ["ProfileOrder_Cancel_Mutation"],
  });
  return (
    <div className="flex flex-col gap-6">
      <h3 className="mb-6 text-xl font-bold">Orders</h3>
      {data.data.length === 0 && (
        <div className="flex items-center justify-center">
          <p className="text-(--text-accent)">Not order has been submited</p>
        </div>
      )}
      {data.data.length !== 0 && (
        <Table<{
          id: ReactNode;
          address: ReactNode;
          actions: ReactNode;
          orderAt: ReactNode;
        }>
          heads={[
            { field: "id", label: "# Order" },
            { field: "address", label: "Address" },
            { field: "orderAt", label: "Ordered At" },
            { field: "actions", label: "" },
          ]}
          rows={data.data.map((order) => ({
            id: <p>{order.id}</p>,
            address: (
              <p>
                {[
                  order.address.addressLine1,
                  order.address.addressLine2,
                  order.address.city,
                  order.address.state,
                  order.address.country,
                ]
                  .filter(Boolean)
                  .join(", ")}
              </p>
            ),
            orderAt: <p>{format(order.createdAt, "MMMM dd, yyyy")}</p>,
            actions: (
              <>
                <Button
                  variant="text"
                  onClick={() => {
                    cancel.mutate(
                      {
                        orderId: order.id,
                      },
                      {
                        onSuccess: () => {
                          refetch();
                        },
                      },
                    );
                  }}
                >
                  Cancel
                </Button>
              </>
            ),
          }))}
        />
      )}
    </div>
  );
};

export default ProfileOrderTab;
