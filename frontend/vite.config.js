import { defineConfig } from "vite";

export default defineConfig({
  root: "public",
  server: {
    host: true,
    port: process.env.VITE_PORT || 5173,
  },
  build: {
    outDir: "../build",
    emptyOutDir: true,
    minify: true,
  },
  preview: {
    host: true,
    port: process.env.VITE_PORT || 5173,
  },
});
