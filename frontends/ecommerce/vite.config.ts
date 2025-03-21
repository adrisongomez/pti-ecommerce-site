import { defineConfig } from "vite";
import tsconfigPaths from "vite-tsconfig-paths";
import react from "@vitejs/plugin-react-swc";
import tailwindcss from "@tailwindcss/vite";
import { TanStackRouterVite } from "@tanstack/router-plugin/vite";

// https://vite.dev/config/
export default defineConfig({
  server: {
    port: 8080,
  },
  base:
    process.env.NODE_ENV !== "development"
      ? "/programming-the-internet-tarea-2/"
      : undefined, // this is just for enable github pages
  plugins: [
    react(),
    tailwindcss(),
    TanStackRouterVite({ target: "react", autoCodeSplitting: true }),
    tsconfigPaths(),
  ],
});
