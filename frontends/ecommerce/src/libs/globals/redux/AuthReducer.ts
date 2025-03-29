import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { AuthService, OpenAPI, User } from "../generated/requests";
import { getCreds, removeCreds } from "../utilities/auth";

export type AuthContextStatus = "logged" | "logout";

export type AuthContextType = {
  status: AuthContextStatus;
  loading: boolean;
  user: User | null;
};

export const AUTH_PROVIDER_INITIAL_STATE: AuthContextType = {
  status: "logout",
  loading: false,
  user: null,
};

export const fetchUser = createAsyncThunk("user/fetchUser", async () => {
  const creds = getCreds();
  if (!creds) {
    return null;
  }
  OpenAPI.TOKEN = creds.accessToken;
  return await AuthService.getApiAuthMe();
});

export const authSlice = createSlice({
  name: "auth",
  initialState: AUTH_PROVIDER_INITIAL_STATE,
  reducers: {
    logout(state) {
      removeCreds();
      state.user = null;
      state.status = "logout";
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchUser.pending, (state) => {
        state.loading = true;
      })
      .addCase(fetchUser.fulfilled, (state, action) => {
        state.loading = false;
        state.user = action.payload;
      })
      .addCase(fetchUser.rejected, (state) => {
        state.user = null;
        state.status = "logout";
        state.loading = false;
      });
  },
});

export default authSlice;
