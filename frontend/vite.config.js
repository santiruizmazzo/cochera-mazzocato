import { defineConfig } from 'vite';

export default defineConfig({
  root: 'public',
  server: {
    host: true,
    port: import.meta.env.VITE_PORT
  },
  build: {
    outDir: '../build',
    emptyOutDir: true,
    minify: true
  },
});