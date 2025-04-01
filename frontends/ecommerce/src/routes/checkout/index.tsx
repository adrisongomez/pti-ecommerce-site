import Button from "@/libs/globals/components/buttons/Button";
import Card from "@/libs/globals/components/cards/Card";
import Textfield from "@/libs/globals/components/fields/Textfield";
import MainLayout from "@/libs/globals/components/layouts/MainLayout";
import { PLACEHOLDER_IMAGE } from "@/libs/globals/constants";
import {
  useAddressServiceGetApiAddresses,
  useAddressServicePostApiAddresses,
  useOrderServicePostApiOrders,
} from "@/libs/globals/generated/queries";
import { AddressInput } from "@/libs/globals/generated/requests";
import { useAppDispatch, useAppSelector } from "@/libs/globals/hooks/redux";
import CartSlider from "@/libs/globals/redux/CartReducer";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { useFormik } from "formik";
import * as Yup from "Yup";

export const Route = createFileRoute("/checkout/")({
  component: CheckoutPage,
  head: () => ({
    meta: [{ title: "Ecommerce | Checkout" }],
  }),
});

function CheckoutPage() {
  const nav = useNavigate();
  const dispatch = useAppDispatch();
  const cart = useAppSelector((state) => state.cart.data);
  const user = useAppSelector((state) => state.auth.user);
  const createAddress = useAddressServicePostApiAddresses({
    mutationKey: ["CheckoutPage_CreateAddress_Mutation"],
    throwOnError: true,
  });
  const createOrder = useOrderServicePostApiOrders({
    mutationKey: ["CheckoutPage_CreateOrder_Mutation"],
    throwOnError: true,
  });
  const totalPrice =
    cart
      .flatMap((c) => parseInt(c.product.variants[0].price) * c.quantity)
      .reduce((acc, v) => acc + v, 0) * 100;
  const formik = useFormik<Omit<AddressInput, "userId">>({
    validationSchema: Yup.object({
      addressLine1: Yup.string().required(),
      addressLine2: Yup.string().optional(),
      state: Yup.string().required(),
      city: Yup.string().required(),
      country: Yup.string().required(),
    }),
    initialValues: {
      state: "",
      city: "",
      addressLine1: "",
      country: "",
    },
    onReset() {
      nav({ to: "/" });
    },
    async onSubmit(d) {
      if (!user?.id) {
        return;
      }
      const value = await createAddress.mutateAsync({
        requestBody: {
          country: d.country,
          zipCode: d.zipCode,
          city: d.city,
          addressLine2: d.addressLine2,
          addressLine1: d.addressLine1,
          userId: user?.id,
          state: d.state,
        },
      });
      await createOrder.mutateAsync({
        requestBody: {
          userId: user.id,
          totalPrice: totalPrice,
          lineItems: cart.map((c) => ({
            quantity: c.quantity,
            productId: c.product.id,
            price: parseInt(c.product.variants[0].price) * 100,
          })),
          addressId: value.id,
          email: user.email,
        },
      });
      dispatch(CartSlider.actions.clearCart());
      nav({ to: "/" });
    },
  });
  const addresses = useAddressServiceGetApiAddresses({}, [
    "CheckoutPage_ShippingAddress",
  ]);
  console.log(addresses);
  return (
    <MainLayout>
      <Card className="m-auto bg-(--bg-light) p-12">
        <form
          className="flex w-3xl flex-col gap-3"
          onSubmit={formik.handleSubmit}
          onReset={formik.handleReset}
        >
          <h1 className="m-auto mb-6 text-3xl font-bold">Checkout</h1>
          <div className="mb-6 flex flex-col gap-4 rounded-lg bg-(--bg-main) px-6 pt-6 pb-4">
            {cart.map((c) => (
              <div
                key={`product-item-${c.product.id}`}
                className="flex w-full items-center justify-between gap-6"
              >
                <img
                  className="size-12 rounded-md bg-slate-300 object-contain"
                  alt={c.product.title.trim()}
                  src={c.product.medias.at(0)?.url ?? PLACEHOLDER_IMAGE}
                />
                <h5 className="max-w-[40ch] flex-1 text-lg font-medium overflow-ellipsis">
                  {c.product.title}
                </h5>
                <div className="flex items-center gap-3">
                  <h5 className="text-lg font-medium">x{c.quantity}</h5>
                </div>
                <div>
                  <h5 className="text-lg font-medium">
                    $ {parseInt(c.product.variants[0].price) * c.quantity}
                  </h5>
                </div>
              </div>
            ))}
            <div className="flex items-center justify-end gap-6">
              <h5 className="text-xl font-bold">Total</h5>
              <h5 className="text-xl font-bold">$ {totalPrice / 100}</h5>
            </div>
          </div>
          <Textfield
            required
            name="addressLine1"
            value={formik.values.addressLine1}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            label="Address Line 1"
            error={formik.touched.addressLine1 && !!formik.errors.addressLine1}
            helperText={
              formik.touched.addressLine1
                ? formik.errors.addressLine1
                : undefined
            }
          />
          <Textfield
            label="Address Line 2"
            name="addressLine2"
            value={formik.values.addressLine2 ?? ""}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
          />
          <div className="flex gap-6">
            <Textfield
              fullWidth
              required
              label="City"
              value={formik.values.city}
              name="city"
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              error={formik.touched.city && !!formik.errors.city}
              helperText={formik.touched.city ? formik.errors.city : undefined}
            />
            <Textfield
              fullWidth
              required
              label="State"
              name="state"
              value={formik.values.state}
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              error={formik.touched.state && !!formik.errors.state}
              helperText={
                formik.touched.state ? formik.errors.state : undefined
              }
            />
          </div>
          <div className="flex gap-6">
            <Textfield
              fullWidth
              label="Zip Code"
              name="zipCode"
              value={formik.values.zipCode ?? ""}
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
            />
            <Textfield
              required
              fullWidth
              label="Country"
              name="country"
              value={formik.values.country ?? ""}
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              error={formik.touched.country && !!formik.errors.country}
              helperText={
                formik.touched.country ? formik.errors.country : undefined
              }
            />
          </div>
          <div className="mt-3 flex items-end justify-end gap-6">
            <Button
              loading={formik.isSubmitting}
              disabled={formik.isSubmitting || !formik.isValid}
              type="submit"
              variant="contained"
            >
              Submit
            </Button>
            <Button type="reset" variant="outlined">
              Cancel
            </Button>
          </div>
        </form>
      </Card>
    </MainLayout>
  );
}
