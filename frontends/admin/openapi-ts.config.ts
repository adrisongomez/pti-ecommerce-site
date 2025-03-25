import { defineConfig } from "@hey-api/openapi-ts";

export default defineConfig({
  input: "../../backends/internal/gen/http/openapi3.json",
  output: "./src/generated",
  plugins: ["@hey-api/client-axios"],
});
