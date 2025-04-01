import { configureStore } from "@reduxjs/toolkit";
import authSlice from "./AuthReducer";
import CartSlider from "./CartReducer";

export const store = configureStore({
  reducer: {
    auth: authSlice.reducer,
    cart: CartSlider.reducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
