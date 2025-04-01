import SimpleCenterLayout from "@/libs/globals/components/layouts/SimpleCenterLayout";
import { useAuthServicePostApiAuthLogin } from "@/libs/globals/generated/queries";
import { OpenAPI } from "@/libs/globals/generated/requests";
import { useAppDispatch } from "@/libs/globals/hooks/redux";
import { fetchUser } from "@/libs/globals/redux/AuthReducer";
import { writeCreds } from "@/libs/globals/utilities/auth";
import LoginForm from "@/libs/routes/auth/LoginForm";
import { createFileRoute } from "@tanstack/react-router";
import { useState } from "react";

export const Route = createFileRoute("/auth/login")({
  component: LoginPage,
  head: () => ({
    meta: [{ title: "Ecommerce | Login" }],
  }),
});

function LoginPage() {
  const [invalidCredentials, setInvalidCredentiasl] = useState(false);
  const nav = Route.useNavigate();
  const login = useAuthServicePostApiAuthLogin({
    mutationKey: ["LoginPage_Login"],
    onError() {
      setInvalidCredentiasl(true);
    },
  });
  const dispatch = useAppDispatch();
  return (
    <SimpleCenterLayout>
      <LoginForm
        showErrorCrendetialsMessage={invalidCredentials}
        onLogin={async (data) => {
          setInvalidCredentiasl(false);
          OpenAPI.USERNAME = data.email;
          OpenAPI.PASSWORD = data.password;
          const cred = await login.mutateAsync();
          writeCreds(cred);
          nav({
            to: "/",
          });
          OpenAPI.USERNAME = undefined;
          OpenAPI.PASSWORD = undefined;
          dispatch(fetchUser());
        }}
      />
    </SimpleCenterLayout>
  );
}
