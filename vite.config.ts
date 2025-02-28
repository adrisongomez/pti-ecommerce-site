import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

// https://vite.dev/config/
export default defineConfig({
  base: "/programming-the-internet-tarea-2/", // this is just for enable github pages
  plugins: [react()],
});
