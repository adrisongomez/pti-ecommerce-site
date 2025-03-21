import SimpleCenterLayout from "@/libs/globals/components/layouts/SimpleCenterLayout";
import LoginForm from "@/libs/routes/auth/LoginForm";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/auth/login")({
  component: LoginPage,
});

function LoginPage() {
  return (
    <SimpleCenterLayout>
      <LoginForm />
    </SimpleCenterLayout>
  );
}
