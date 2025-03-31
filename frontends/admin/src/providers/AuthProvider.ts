import { AuthProvider } from "@refinedev/core";
import { getBaseApiClients } from "./utils";
import { authLogin } from "../generated/index";
import { getCreds, removeCreds, writeCreds } from "../utils/auth";

export default function Auth(baseUrl: string): AuthProvider {
  const client = getBaseApiClients(baseUrl);
  return {
    async login({ email, password }: { email: string; password: string }) {
      const basic = `${email}:${password}`;
      const creds = await authLogin({
        client,
        headers: {
          Authorization: `Basic ${btoa(basic)}`,
        },
        throwOnError: true,
      });
      writeCreds(creds.data);
      return {
        success: true,
        redirectTo: "/products",
      };
    },
    async logout() {
      removeCreds();
      return {
        success: true,
        successNotification: {
          message: "Logout",
        },
      };
    },
    async check() {
      const creds = getCreds();
      console.log("Auth creds", creds);
      return {
        authenticated: !!creds,
      };
    },
    async onError(err) {
      console.log("Auth err", err);
      return {
        redirectTo: "/",
        logout: true,
      };
    },
  };
}
