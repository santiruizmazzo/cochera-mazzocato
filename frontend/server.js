import { serve } from "bun";
import path from "path";

const server = serve({
  port: process.env.PORT,
  async fetch(req) {
    const url = new URL(req.url);

    // Default -> servir index.html
    let filePath = url.pathname === "/" ? "/index.html" : url.pathname;

    try {
      return new Response(Bun.file(`public${filePath}`));
    } catch {
      return new Response("Archivo no encontrado", { status: 404 });
    }
  },
});

console.log(`🚀 Servidor corriendo en http://localhost:${server.port}`);
