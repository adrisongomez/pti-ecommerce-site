import SimpleCenterLayout from "@/libs/globals/components/layouts/SimpleCenterLayout";
import { useAuthServicePostApiAuthSignup } from "@/libs/globals/generated/queries";
import { useAppDispatch, useAppSelector } from "@/libs/globals/hooks/redux";
import { fetchUser } from "@/libs/globals/redux/AuthReducer";
import { writeCreds } from "@/libs/globals/utilities/auth";
import SignUpForm from "@/libs/routes/auth/SignUpForm";
import { createFileRoute } from "@tanstack/react-router";
import { useMount } from "react-use";

export const Route = createFileRoute("/auth/sign-up")({
  component: RouteComponent,
});

function RouteComponent() {
  const nav = Route.useNavigate();
  const dispatch = useAppDispatch();
  const auth = useAppSelector((state) => state.auth);
  useMount(() => {
    if (auth.user) {
      nav({ to: "/" });
    }
  });
  const mutation = useAuthServicePostApiAuthSignup({});
  return (
    <SimpleCenterLayout>
      <SignUpForm
        onSignUpForm={async (d) => {
          const response = await mutation.mutateAsync({
            requestBody: {
              email: d.email,
              firstName: d.firstName,
              lastName: d.lastName,
              password: d.password,
            },
          });
          writeCreds(response);
          nav({
            to: "/",
          });
          dispatch(fetchUser());
        }}
      />
    </SimpleCenterLayout>
  );
}
