import LoginForm from "@/libs/routes/auth/LoginForm";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/auth/login")({
  component: LoginPage,
});

function LoginPage() {
  return (
    <div id="fuck" className="flex w-full items-center justify-center">
      <LoginForm />
    </div>
  );
}
