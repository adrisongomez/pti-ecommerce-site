import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";

export async function withMockAPI() {
  if (process.env.NODE_ENV !== "development") {
    const { worker } = await import("@/mocks/browser");
    return worker.start();
  }
}
withMockAPI().then(() =>
  createRoot(document.getElementById("root")!).render(
    <StrictMode>
      <App />
    </StrictMode>,
  ),
);
