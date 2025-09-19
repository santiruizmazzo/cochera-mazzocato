import { defineConfig } from "vite";

export default defineConfig({
  root: "public",
  server: {
    host: true,
    port: import.meta.env.VITE_PORT || 5173,
  },
  build: {
    outDir: "../build",
    emptyOutDir: true,
    minify: true,
  },
});
