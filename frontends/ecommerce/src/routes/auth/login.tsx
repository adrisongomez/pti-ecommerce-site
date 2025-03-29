import SimpleCenterLayout from "@/libs/globals/components/layouts/SimpleCenterLayout";
import { useAuthServicePostApiAuthLogin } from "@/libs/globals/generated/queries";
import { OpenAPI } from "@/libs/globals/generated/requests";
import { useAppDispatch, useAppSelector } from "@/libs/globals/hooks/redux";
import { fetchUser } from "@/libs/globals/redux/AuthReducer";
import { writeCreds } from "@/libs/globals/utilities/auth";
import LoginForm from "@/libs/routes/auth/LoginForm";
import { createFileRoute } from "@tanstack/react-router";
import { useMount } from "react-use";

export const Route = createFileRoute("/auth/login")({
  component: LoginPage,
});

function LoginPage() {
  const nav = Route.useNavigate();
  const login = useAuthServicePostApiAuthLogin({
    mutationKey: ["LoginPage_Login"],
  });
  const dispatch = useAppDispatch();
  const auth = useAppSelector((state) => state.auth);
  useMount(() => {
    if (auth.user) {
      nav({ to: "/" });
    }
  });
  return (
    <SimpleCenterLayout>
      <LoginForm
        onLogin={async (data) => {
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
