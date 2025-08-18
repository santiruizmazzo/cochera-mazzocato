import fs from "fs";
import dotenv from "dotenv";

// Cargar variables del .env
dotenv.config();

// Crear el contenido de env.js para el navegador
const envContent = `window.ENV = {
  API_URL: "${process.env.API_URL || ""}",
};`;

// Guardar el archivo en /public
fs.writeFileSync("public/env.js", envContent);

console.log("Archivo public/env.js generado con variables de entorno.");
