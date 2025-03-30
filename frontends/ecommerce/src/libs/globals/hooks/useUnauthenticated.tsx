import { useMount } from "react-use";
import { useAppSelector } from "./redux";
import { useNavigate } from "@tanstack/react-router";

export function useUnauthenticated() {
  const nav = useNavigate();
  const auth = useAppSelector((state) => state.auth);
  useMount(() => {
    if (auth.user) {
      nav({ to: "/" });
    }
  });
}
