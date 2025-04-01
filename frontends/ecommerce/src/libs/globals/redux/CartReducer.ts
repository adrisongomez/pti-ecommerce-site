import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Product } from "../generated/requests";

type CartState = {
  data: { product: Product; quantity: number }[];
};

const INITIAL_STATE: CartState = {
  data: [],
};

const CartSlider = createSlice({
  name: "cart",
  initialState: INITIAL_STATE,
  reducers: {
    addCart(state, action: PayloadAction<Product>) {
      const { payload } = action;
      if (!state.data.find((v) => v.product.id === payload.id)) {
        state.data = [...state.data, { product: payload, quantity: 1 }];
      } else {
        state.data = state.data.map((v) =>
          v.product.id === payload.id
            ? { product: payload, quantity: v.quantity + 1 }
            : v,
        );
      }
    },
    removeCart(state, action: PayloadAction<number>) {
      state.data = state.data
        .map((p) =>
          p.product.id === action.payload
            ? { product: p.product, quantity: p.quantity - 1 }
            : p,
        )
        .filter((d) => d.quantity > 0);
    },
    clearCart(state) {
      state.data = [];
    },
  },
});

export default CartSlider;
