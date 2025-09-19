import homepage from "./build/index.html";

Bun.serve({
  port: process.env.PORT,
  routes: {
    "/": homepage,
  },
});

console.log(
  `🚀 Development server running on http://localhost:${process.env.PORT}`,
);
