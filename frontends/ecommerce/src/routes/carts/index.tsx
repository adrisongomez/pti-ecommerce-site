import Button from "@/libs/globals/components/buttons/Button";
import IconButton from "@/libs/globals/components/buttons/IconButton";
import Card from "@/libs/globals/components/cards/Card";
import MainLayout from "@/libs/globals/components/layouts/MainLayout";
import { PLACEHOLDER_IMAGE } from "@/libs/globals/constants";
import { useAppDispatch, useAppSelector } from "@/libs/globals/hooks/redux";
import CartSlider from "@/libs/globals/redux/CartReducer";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { MinusCircle, PlusCircle } from "react-feather";

export const Route = createFileRoute("/carts/")({
  component: RouteComponent,
  head: () => ({
    meta: [{ title: "Ecommerce | Cart" }],
  }),
});

function RouteComponent() {
  const nav = useNavigate();
  const cart = useAppSelector((state) => state.cart.data);
  const dispatch = useAppDispatch();
  return (
    <MainLayout>
      <Card className="m-auto flex w-3xl flex-col gap-3 bg-(--bg-light) p-6">
        <h3 className="mb-6 text-3xl font-bold">Cart</h3>
        {!cart.length && (
          <p className="text-center text-(--text-accent)">
            No selected item in the cart
          </p>
        )}
        {cart.map((c) => (
          <div
            key={`product-item-${c.product.id}`}
            className="flex items-center justify-between"
          >
            <div className="flex items-center gap-6">
              <img
                className="size-12 rounded-md bg-slate-300 object-contain"
                alt={c.product.title.trim()}
                src={c.product.medias.at(0)?.url ?? PLACEHOLDER_IMAGE}
              />
              <h5 className="max-w-[40ch] text-xl font-medium overflow-ellipsis">
                {c.product.title}
              </h5>
            </div>
            <div className="flex items-center gap-3">
              <IconButton
                onClick={() => {
                  dispatch(CartSlider.actions.addCart(c.product));
                }}
              >
                <PlusCircle />
              </IconButton>
              <h5 className="text-xl font-medium">{c.quantity}</h5>
              <IconButton
                onClick={() => {
                  dispatch(CartSlider.actions.removeCart(c.product.id));
                }}
              >
                <MinusCircle />
              </IconButton>
            </div>
          </div>
        ))}
        <div className="mt-6 flex items-end justify-end gap-6">
          <Button
            disabled={!cart.length}
            variant="contained"
            onClick={() => {
              nav({ to: "/checkout" });
            }}
          >
            Checkout
          </Button>
          <Button
            variant="outlined"
            onClick={() => {
              dispatch(CartSlider.actions.clearCart());
              nav({ to: "/" });
            }}
          >
            Clear
          </Button>
        </div>
      </Card>
    </MainLayout>
  );
}
